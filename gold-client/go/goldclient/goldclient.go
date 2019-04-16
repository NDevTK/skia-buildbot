package goldclient

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"go.skia.org/infra/go/fileutil"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/util"
	"go.skia.org/infra/golden/go/baseline"
	"go.skia.org/infra/golden/go/diff"
	"go.skia.org/infra/golden/go/jsonio"
	"go.skia.org/infra/golden/go/shared"
	"go.skia.org/infra/golden/go/types"
	"golang.org/x/sync/errgroup"
)

const (
	// jsonPrefix is the path prefix in the GCS bucket that holds JSON result files
	jsonPrefix = "dm-json-v1"

	// imagePrefix is the path prefix in the GCS bucket that holds images.
	imagePrefix = "dm-images-v1"

	// knownHashesPath is path on the Gold instance to retrieve the known image hashes that do
	// not need to be uploaded anymore.
	knownHashesPath = "json/hashes"

	// stateFile is the name of the file that holds the state in the work directory
	// between calls
	stateFile = "result-state.json"

	// jsonTempFile is the temporary file that is created to upload results via gsutil.
	jsonTempFile = "gsutil_dm.json"

	// goldHostTemplate constructs the URL of the Gold instance from the instance id
	goldHostTemplate = "https://%s-gold.skia.org"

	// bucketTemplate constructs the name of the ingestion bucket from the instance id
	bucketTemplate = "skia-gold-%s"
)

// md5Regexp is used to check whether strings are MD5 hashes.
var md5Regexp = regexp.MustCompile(`^[a-f0-9]{32}$`)

// GoldClient is the uniform interface to communicate with the Gold service.
type GoldClient interface {
	// SetSharedConfig populates the config with details that will be shared
	// with all tests. This is safe to be called more than once, although
	// new settings will overwrite the old ones. This will cause the
	// baseline and known hashes to be (re-)downloaded from Gold.
	SetSharedConfig(sharedConfig jsonio.GoldResults) error
	// Test adds a test result to the current testrun. If the GoldClient is configured to
	// return PASS/FAIL for each test, the returned boolean indicates whether the test passed
	// comparison with the expectations (this involves uploading JSON to the server).
	// This will upload the image if the hash of the pixels has not been seen before -
	// using auth.SetDryRun(true) can prevent that.
	//
	// An error is only returned if there was a technical problem in processing the test.
	Test(name string, imgFileName string) (bool, error)

	// Upload the JSON file for all Test() calls previously seen.
	// A no-op if configured for PASS/FAIL mode, since the JSON would have been uploaded
	// on the calls to Test().
	Finalize() error
}

// HTTPClient makes it easier to mock out goldclient's dependencies on
// http.Client by representing a smaller interface.
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

// cloudClient implements the GoldClient interface for the remote Gold service.
type cloudClient struct {
	// workDir is a temporary directory that has to exist between related calls
	workDir string

	// resultState keeps track of the all the information to generate and upload a valid result.
	resultState *resultState

	// ready caches the result of the isReady call so we avoid duplicate work.
	ready bool

	// these functions are overwritable by tests
	loadAndHashImage func(path string) ([]byte, string, error)
	now              func() time.Time

	// auth stores the authentication method to use.
	auth       AuthOpt
	httpClient HTTPClient
}

// GoldClientConfig is a config structure to configure GoldClient instances
type GoldClientConfig struct {
	// WorkDir is a temporary directory that caches data for one run with multiple calls to GoldClient
	WorkDir string

	// InstanceID is the id of the backend Gold instance
	InstanceID string

	// PassFailStep indicates whether each call to Test(...) should return a pass/fail value.
	PassFailStep bool

	// OverrideGoldURL is optional and allows to override the GoldURL for testing.
	OverrideGoldURL string
}

// resultState is an internal container for all information to upload results
// to Gold, including the jsonio.GoldResult structure itself.
type resultState struct {
	// SharedConfig is all the data that is common test to test, for example, the
	// keys about this machine (e.g. GPU, OS).
	SharedConfig    *jsonio.GoldResults
	PerTestPassFail bool
	InstanceID      string
	GoldURL         string
	Bucket          string
	KnownHashes     util.StringSet
	Expectations    types.TestExp
}

// NewCloudClient returns an implementation of the GoldClient that relies on the Gold service.
// If a new instance is created for each call to Test, the arguments of the first call are
// preserved. They are cached in a JSON file in the work directory.
func NewCloudClient(authOpt AuthOpt, config GoldClientConfig) (*cloudClient, error) {
	// Make sure the workdir was given and exists.
	if config.WorkDir == "" {
		return nil, skerr.Fmt("No 'workDir' provided to NewCloudClient")
	}

	workDir, err := fileutil.EnsureDirExists(config.WorkDir)
	if err != nil {
		return nil, skerr.Fmt("Error setting up workdir: %s", err)
	}

	if config.InstanceID == "" {
		return nil, skerr.Fmt("Can't have empty config")
	}

	ret := cloudClient{
		workDir:          workDir,
		auth:             authOpt,
		loadAndHashImage: loadAndHashImage,
		now:              defaultNow,

		resultState: newResultState(nil, &config),
	}
	if err := ret.setHttpClient(); err != nil {
		return nil, skerr.Fmt("Error setting http client: %s", err)
	}

	// write it to disk
	if err := saveJSONFile(ret.getResultStatePath(), ret.resultState); err != nil {
		return nil, skerr.Fmt("Could not write the state to disk: %s", err)
	}

	return &ret, nil
}

// LoadCloudClient returns a GoldClient that has previously been stored to disk
// in the path given by workDir.
func LoadCloudClient(authOpt AuthOpt, workDir string) (*cloudClient, error) {
	// Make sure the workdir was given and exists.
	if workDir == "" {
		return nil, skerr.Fmt("No 'workDir' provided to LoadCloudClient")
	}
	ret := cloudClient{
		workDir:          workDir,
		auth:             authOpt,
		loadAndHashImage: loadAndHashImage,
		now:              defaultNow,
	}
	var err error
	ret.resultState, err = loadStateFromJson(ret.getResultStatePath())
	if err != nil {
		return nil, skerr.Fmt("Could not load disk from state: %s", err)
	}
	if err = ret.setHttpClient(); err != nil {
		return nil, skerr.Fmt("Error setting http client: %s", err)
	}

	return &ret, nil
}

// SetSharedConfig implements the GoldClient interface.
func (c *cloudClient) SetSharedConfig(sharedConfig jsonio.GoldResults) error {
	existingConfig := GoldClientConfig{
		WorkDir: c.workDir,
	}
	if c.resultState != nil {
		existingConfig.InstanceID = c.resultState.InstanceID
		existingConfig.PassFailStep = c.resultState.PerTestPassFail
		existingConfig.OverrideGoldURL = c.resultState.GoldURL
	}
	c.resultState = newResultState(&sharedConfig, &existingConfig)

	// The GitHash may have changed (or been set for the first time),
	// So we can now load the baseline. We can also download the hashes
	// at this time, although we could have done it at any time before since
	// that does not depend on the GitHash we have.
	if err := c.downloadHashesAndBaselineFromGold(); err != nil {
		return skerr.Fmt("Error downloading from Gold: %s", err)
	}

	return saveJSONFile(c.getResultStatePath(), c.resultState)
}

// Test implements the GoldClient interface.
func (c *cloudClient) Test(name string, imgFileName string) (bool, error) {
	return c.addTest(name, imgFileName)
}

// addTest adds a test to results. If perTestPassFail is true it will also upload the result.
// Returns true if the test was added (and maybe uploaded) successfully.
func (c *cloudClient) addTest(name string, imgFileName string) (bool, error) {
	if err := c.isReady(); err != nil {
		return false, skerr.Fmt("Unable to process test result. Cloud Gold Client not ready: %s", err)
	}

	// Get an uploader. This is either based on an authenticated client or on gsutils.
	uploader, err := c.auth.GetGoldUploader()
	if err != nil {
		return false, skerr.Fmt("Error retrieving uploader: %s", err)
	}

	// Load the PNG from disk and hash it.
	imgBytes, imgHash, err := c.loadAndHashImage(imgFileName)
	if err != nil {
		return false, err
	}
	fmt.Printf("Given image with hash %s for test %s\n", imgHash, name)
	for expectHash, expectLabel := range c.resultState.Expectations[name] {
		fmt.Printf("Expectation for test: %s (%s)\n", expectHash, expectLabel.String())
	}

	var egroup errgroup.Group
	// Check against known hashes and upload if needed.
	if !c.resultState.KnownHashes[imgHash] {
		egroup.Go(func() error {
			gcsImagePath := c.resultState.getGCSImagePath(imgHash)
			if err := uploader.UploadBytes(imgBytes, imgFileName, prefixGCS(gcsImagePath)); err != nil {
				return skerr.Fmt("Error uploading image %s to %s. Got: %s", imgFileName, gcsImagePath, err)
			}
			return nil
		})
	}

	// Add the result of this test.
	c.addResult(name, imgHash)

	// At this point the result should be correct for uploading.
	if _, err := c.resultState.SharedConfig.Validate(false); err != nil {
		return false, err
	}

	// If we do per test pass/fail then upload the result and compare it to the baseline.
	ret := true
	if c.resultState.PerTestPassFail {
		egroup.Go(func() error {
			return c.uploadResultJSON(uploader)
		})

		ret = c.resultState.Expectations[name][imgHash] == types.POSITIVE
	}

	if err := egroup.Wait(); err != nil {
		return false, err
	}
	return ret, nil
}

func (c *cloudClient) Finalize() error {
	if err := c.isReady(); err != nil {
		return skerr.Fmt("Cannot finalize - client not ready: %s", err)
	}
	uploader, err := c.auth.GetGoldUploader()
	if err != nil {
		return skerr.Fmt("Error retrieving uploader: %s", err)
	}
	return c.uploadResultJSON(uploader)
}

func (c *cloudClient) uploadResultJSON(uploader GoldUploader) error {
	localFileName := filepath.Join(c.workDir, jsonTempFile)
	resultFilePath := c.resultState.getResultFilePath(c.now())
	if err := uploader.UploadJSON(c.resultState.SharedConfig, localFileName, resultFilePath); err != nil {
		return skerr.Fmt("Error uploading JSON file to GCS path %s: %s", resultFilePath, err)
	}
	return nil
}

// setHttpClient sets authenticated httpClient, if authentication was configured via SetAuthConfig.
// It also retrieves a token of the configured source to make sure it works.
func (c *cloudClient) setHttpClient() error {
	// If no auth option was set, we return an unauthenticated client.
	client, err := c.auth.GetHTTPClient()
	if err != nil {
		return err
	}
	c.httpClient = client
	return nil
}

// saveAuthOpt assumes that auth has been set. It saves it to the work directory for retrieval
// during later calls.
func (c *cloudClient) saveAuthOpt() error {
	outFile := filepath.Join(c.workDir, authFile)
	return saveJSONFile(outFile, c.auth)
}

// isReady returns true if the instance is ready to accept test results (all necessary info has been
// configured)
func (c *cloudClient) isReady() error {
	if c.ready {
		return nil
	}

	// if resultState hasn't been set yet, then we are simply not ready.
	if c.resultState == nil {
		return skerr.Fmt("No result state object available")
	}

	// Check whether we have some means of uploading results
	if c.auth == nil {
		return skerr.Fmt("No authentication information provided.")
	}
	if err := c.auth.Validate(); err != nil {
		return skerr.Fmt("Invalid auth: %s", err)
	}

	// Check if the GoldResults instance is complete once results are added.
	if _, err := c.resultState.SharedConfig.Validate(true); err != nil {
		return skerr.Fmt("Gold results fields invalid: %s", err)
	}

	c.ready = true
	return nil
}

// getResultStatePath returns the path of the temporary file where the state is cached as JSON
func (c *cloudClient) getResultStatePath() string {
	return filepath.Join(c.workDir, stateFile)
}

// addResult adds the given test to the overall results.
func (c *cloudClient) addResult(name, imgHash string) {
	newResult := &jsonio.Result{
		Digest: imgHash,
		Key:    map[string]string{types.PRIMARY_KEY_FIELD: name},

		// TODO(stephana): check if the backend still relies on this.
		Options: map[string]string{"ext": "png"},
	}

	// TODO(kjlubick): Maybe make the corpus field an option.
	if _, ok := c.resultState.SharedConfig.Key[types.CORPUS_FIELD]; !ok {
		newResult.Key[types.CORPUS_FIELD] = c.resultState.InstanceID
	}
	c.resultState.SharedConfig.Results = append(c.resultState.SharedConfig.Results, newResult)
}

// downloadHashesAndBaselineFromGold downloads the hashes and baselines
// and stores them to resultState.
func (c *cloudClient) downloadHashesAndBaselineFromGold() error {
	// What hashes have we seen already (to avoid uploading them again).
	if err := c.resultState.loadKnownHashes(c.httpClient); err != nil {
		return err
	}

	// Fetch the baseline (may be empty but should not fail).
	if err := c.resultState.loadExpectations(c.httpClient); err != nil {
		return err
	}
	return nil
}

// loadAndHashImage loads an image from disk and hashes the internal Pixel buffer. It returns
// the bytes of the encoded image and the MD5 hash as hex encoded string.
func loadAndHashImage(fileName string) ([]byte, string, error) {
	// Load the image
	reader, err := os.Open(fileName)
	if err != nil {
		return nil, "", err
	}
	defer util.Close(reader)

	imgBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, "", skerr.Fmt("Error loading file %s: %s", fileName, err)
	}

	img, err := png.Decode(bytes.NewBuffer(imgBytes))
	if err != nil {
		return nil, "", skerr.Fmt("Error decoding PNG in file %s: %s", fileName, err)
	}
	nrgbaImg := diff.GetNRGBA(img)
	md5Hash := fmt.Sprintf("%x", md5.Sum(nrgbaImg.Pix))
	return imgBytes, md5Hash, nil
}

// defaultNow returns what time it is now in UTC
func defaultNow() time.Time {
	return time.Now().UTC()
}

// newResultState creates a new instance of resultState
func newResultState(sharedConfig *jsonio.GoldResults, config *GoldClientConfig) *resultState {

	// TODO(stephana): Move deriving the URLs and the bucket to a central place in the backend
	// or get rid of the bucket entirely and expose an upload URL (requires authentication)

	goldURL := config.OverrideGoldURL
	if goldURL == "" {
		goldURL = getHostURL(config.InstanceID)
	}

	ret := &resultState{
		SharedConfig:    sharedConfig,
		PerTestPassFail: config.PassFailStep,
		InstanceID:      config.InstanceID,
		GoldURL:         goldURL,
		Bucket:          getBucket(config.InstanceID),
	}

	return ret
}

// loadStateFromJson loads a serialization of a resultState instance that was previously written
// via the save method.
func loadStateFromJson(fileName string) (*resultState, error) {
	ret := &resultState{}
	exists, err := loadJSONFile(fileName, ret)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}
	return ret, nil
}

// loadKnownHashes loads the list of known hashes from the Gold instance.
func (r *resultState) loadKnownHashes(httpClient HTTPClient) error {
	r.KnownHashes = util.StringSet{}

	// Fetch the known hashes via http
	hashesURL := fmt.Sprintf("%s/%s", r.GoldURL, knownHashesPath)
	resp, err := httpClient.Get(hashesURL)
	if err != nil {
		return skerr.Fmt("Error retrieving known hashes file: %s", err)
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil
	}

	// Retrieve the body and parse the list of known hashes.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return skerr.Fmt("Error reading body of HTTP response: %s", err)
	}
	if err := resp.Body.Close(); err != nil {
		return skerr.Fmt("Error closing HTTP response: %s", err)
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(body))
	for scanner.Scan() {
		// Ignore empty lines and lines that are not valid MD5 hashes
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) > 0 && md5Regexp.Match(line) {
			r.KnownHashes[string(line)] = true
		}
	}
	if err := scanner.Err(); err != nil {
		return skerr.Fmt("Error scanning response of HTTP request: %s", err)
	}
	return nil
}

// loadExpectations fetches the expectations from Gold to compare to tests.
func (r *resultState) loadExpectations(httpClient HTTPClient) error {
	urlPath := strings.Replace(shared.EXPECTATIONS_ROUTE, "{commit_hash}", r.SharedConfig.GitHash, 1)
	if r.SharedConfig.Issue > 0 {
		urlPath = fmt.Sprintf("%s?issue=%d", urlPath, r.SharedConfig.Issue)
	}
	url := fmt.Sprintf("%s/%s", r.GoldURL, strings.TrimLeft(urlPath, "/"))

	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return skerr.Fmt("Error reading body of request to %s: %s", url, err)
	}
	if err := resp.Body.Close(); err != nil {
		return skerr.Fmt("Error closing response from request to %s: %s", url, err)
	}

	exp := &baseline.CommitableBaseLine{}

	if err := json.Unmarshal(jsonBytes, exp); err != nil {
		fmt.Printf("Fetched from %s\n", url)
		if len(jsonBytes) > 200 {
			fmt.Printf(`Invalid JSON: "%s..."`, string(jsonBytes[0:200]))
		} else {
			fmt.Printf(`Invalid JSON: "%s"`, string(jsonBytes))
		}
		return skerr.Fmt("Error parsing JSON; this sometimes means auth issues: %s", err)
	}

	r.Expectations = exp.Baseline
	return nil
}

// getResultFilePath returns that path in GCS where the result file should be stored.
//
// The path follows the path described here:
//    https://github.com/google/skia-buildbot/blob/master/golden/docs/INGESTION.md
// The file name of the path also contains a timestamp to make it unique since all
// calls within the same test run are written to the same output path.
func (r *resultState) getResultFilePath(now time.Time) string {
	year, month, day := now.Date()
	hour := now.Hour()

	// Assemble a path that looks like this:
	// <path_prefix>/YYYY/MM/DD/HH/<git_hash>/<build_id>/<time_stamp>/<per_run_file_name>.json
	// The first segments up to 'HH' are required so the Gold ingester can scan these prefixes for
	// new files. The later segments are necessary to make the path unique within the runs of one
	// hour and increase readability of the paths for troubleshooting.
	// It is vital that the times segments of the path are based on UTC location.
	fileName := fmt.Sprintf("dm-%d.json", now.UnixNano())
	segments := []interface{}{
		jsonPrefix,
		year,
		month,
		day,
		hour,
		r.SharedConfig.GitHash,
		r.SharedConfig.BuildBucketID,
		now.Unix(),
		fileName}
	path := fmt.Sprintf("%s/%04d/%02d/%02d/%02d/%s/%d/%d/%s", segments...)

	if r.SharedConfig.Issue > 0 {
		path = "trybot/" + path
	}
	return fmt.Sprintf("%s/%s", r.Bucket, path)
}

// getGCSImagePath returns the path in GCS where the image with the given hash should be stored.
func (r *resultState) getGCSImagePath(imgHash string) string {
	return fmt.Sprintf("%s/%s/%s.png", r.Bucket, imagePrefix, imgHash)
}

// loadJSONFile loads and parses the JSON in 'fileName'. If the file doesn't exist it returns
// (false, nil). If the first return value is true, 'data' contains the parse JSON data.
func loadJSONFile(fileName string, data interface{}) (bool, error) {
	if !fileutil.FileExists(fileName) {
		return false, nil
	}

	err := util.WithReadFile(fileName, func(r io.Reader) error {
		return json.NewDecoder(r).Decode(data)
	})
	if err != nil {
		return false, skerr.Fmt("Error reading/parsing JSON file: %s", err)
	}

	return true, nil
}

// saveJSONFile stores the given 'data' in a file with the given name
func saveJSONFile(fileName string, data interface{}) error {
	err := util.WithWriteFile(fileName, func(w io.Writer) error {
		return json.NewEncoder(w).Encode(data)
	})
	if err != nil {
		return skerr.Fmt("Error writing/serializing to JSON: %s", err)
	}
	return nil
}

const (
	// Skia's naming conventions are old and don't follow the patterns that
	// newer clients do. One day, it might be nice to align the skia names
	// to match the rest.
	bucketSkiaLegacy     = "skia-infra-gm"
	hostSkiaLegacy       = "https://gold.skia.org"
	instanceIDSkiaLegacy = "skia-legacy"
)

// getBucket returns the bucket name for a given instance id.
// This is usually a formulaic transform, but there are some special cases.
func getBucket(instanceID string) string {
	if instanceID == instanceIDSkiaLegacy {
		return bucketSkiaLegacy
	}
	return fmt.Sprintf(bucketTemplate, instanceID)
}

// getHostURL returns the hostname for a given instance id.
// This is usually a formulaic transform, but there are some special cases.
func getHostURL(instanceID string) string {
	if instanceID == instanceIDSkiaLegacy {
		return hostSkiaLegacy
	}
	return fmt.Sprintf(goldHostTemplate, instanceID)
}

// Make sure cloudClient fulfils the GoldClient interface
var _ GoldClient = (*cloudClient)(nil)
