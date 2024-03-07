package midpoint

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"go.skia.org/infra/go/depot_tools/deps_parser"
	"go.skia.org/infra/go/gitiles"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/sklog"
)

const (
	GitilesEmptyResponseErr = "Gitiles returned 0 commits, which should not happen."
	chromiumSrcGit          = "https://chromium.googlesource.com/chromium/src.git"
)

// A Commit represents a commit of a given repository.
// TODO(jeffyoon@) - Reorganize this into a types folder.
type Commit struct {
	// GitHash is the Git SHA1 hash to build for the project.
	GitHash string

	// RepositoryUrl is the url to the repository, ie/ https://chromium.googlesource.com/chromium/src
	RepositoryUrl string
}

func NewCommit(h string) *Commit {
	return &Commit{
		GitHash:       h,
		RepositoryUrl: chromiumSrcGit,
	}
}

// A CombinedCommit represents one main base commit with any dependencies that require
// overrides as part of the build request.
// For example, if Commit is chromium/src@1, Dependency may be V8@2 which is passed
// along to Buildbucket as a deps_revision_overrides.
type CombinedCommit struct {
	// Main is the main base commit, usually a Chromium commit.
	Main *Commit
	// ModifiedDeps is a list of commits to provide as overrides, ie/ V8.
	ModifiedDeps []*Commit
}

// TODO(jeffyoon@) - move this to a deps folder, likely with the types restructure above.
// DepsToMap translates all deps into a map.
func (cc *CombinedCommit) DepsToMap() map[string]string {
	resp := make(map[string]string, 0)
	for _, c := range cc.ModifiedDeps {
		resp[c.RepositoryUrl] = c.GitHash
	}
	return resp
}

// GetMainGitHash returns the git hash of main.
func (cc *CombinedCommit) GetMainGitHash() string {
	if cc.Main == nil {
		return ""
	}

	return cc.Main.GitHash
}

// Key returns all git hashes combined to use for map indexing
func (cc *CombinedCommit) Key() string {
	if cc.Main == nil {
		return ""
	}
	key := cc.Main.GitHash
	for _, d := range cc.ModifiedDeps {
		if d != nil {
			key = fmt.Sprintf("%s+%s", key, d.GitHash)
		}
	}
	return key
}

func NewCombinedCommit(main *Commit, deps ...*Commit) *CombinedCommit {
	return &CombinedCommit{
		Main:         main,
		ModifiedDeps: deps,
	}
}

// CommitRange provides information about the left and right commits used to determine
// the next commit to bisect against.
type CommitRange struct {
	Left  *CombinedCommit
	Right *CombinedCommit
}

// HasLeftGitHash checks if left main git hash is set.
func (cr *CommitRange) HasLeftGitHash() bool {
	return cr.Left.Main.GitHash != ""
}

// HasRightGitHash checks if left main git hash is set
func (cr *CommitRange) HasRightGitHash() bool {
	return cr.Right.Main.GitHash != ""
}

// MidpointHandler encapsulates all logic to determine the next potential candidate for Bisection.
type MidpointHandler struct {
	// repos is a map of repository url to a GitilesRepo object.
	repos map[string]gitiles.GitilesRepo

	c *http.Client
}

// New returns a new MidpointHandler.
func New(ctx context.Context, c *http.Client) *MidpointHandler {
	return &MidpointHandler{
		repos: make(map[string]gitiles.GitilesRepo, 0),
		c:     c,
	}
}

// WithRepo returns a MidpointHandler with the repository url mapped to a GitilesRepo object.
func (m *MidpointHandler) WithRepo(url string, r gitiles.GitilesRepo) *MidpointHandler {
	m.repos[url] = r
	return m
}

// getOrCreateRepo fetches the gitiles.GitilesRepo object for the repository url.
// If not present, it'll create an authenticated Repo client.
func (m *MidpointHandler) getOrCreateRepo(url string) gitiles.GitilesRepo {
	gr, ok := m.repos[url]
	if !ok {
		gr = gitiles.NewRepo(url, m.c)
		m.repos[url] = gr
	}
	return gr
}

// findMidpoint identiifes the median commit given a start and ending git hash.
func (m *MidpointHandler) findMidpoint(ctx context.Context, url, startGitHash, endGitHash string) (string, error) {
	if startGitHash == endGitHash {
		return "", skerr.Fmt("Both git hashes are the same; Start: %s, End: %s", startGitHash, endGitHash)
	}

	gc := m.getOrCreateRepo(url)

	// Find the midpoint between the provided commit hashes. Take the lower bound
	// if the list is an odd count. If the gitiles response is == endGitHash, it
	// this means both start and end are adjacent, and DEPS needs to be unravelled
	// to find the potential culprit.
	// LogLinear will return in reverse chronological order, inclusive of the end git hash.
	lc, err := gc.LogLinear(ctx, startGitHash, endGitHash)
	if err != nil {
		return "", err
	}

	// The list can only be empty if the start and end commits are the same.
	if len(lc) == 0 {
		return "", skerr.Fmt("%s. Start %s and end %s hashes may be reversed.", GitilesEmptyResponseErr, startGitHash, endGitHash)
	}

	// Two adjacent commits returns one commit equivalent to the end git hash.
	if len(lc) == 1 && lc[0].ShortCommit.Hash == endGitHash {
		return startGitHash, nil
	}

	// Pop off the first element, since it's the end hash.
	// Golang divide will return lower bound when odd.
	lc = lc[1:]

	// Sort to chronological order before taking the midpoint. This means for even
	// lists, we opt to the higher index (ie/ in [1,2,3,4] len == 4 and midpoint
	// becomes index 2 (which = 3))
	slices.Reverse(lc)
	mlc := lc[len(lc)/2]

	return mlc.ShortCommit.Hash, nil
}

// fetchGitDeps calls Gitiles to read the DEPS content and parses out only the git-based dependencies.
func (m *MidpointHandler) fetchGitDeps(ctx context.Context, gc gitiles.GitilesRepo, gitHash string) (map[string]string, error) {
	denormalized := make(map[string]string, 0)

	content, err := gc.ReadFileAtRef(ctx, "DEPS", gitHash)
	if err != nil {
		return denormalized, err
	}

	entries, err := deps_parser.ParseDeps(string(content))
	if err != nil {
		return denormalized, err
	}

	// We have no good way of determining whether the current DEP is a .git based
	// DEP or a CIPD based dep using the existing deps_parser, so we filter by
	// whether the Id has ".com" to differentiate. This likely needs refinement.
	for id, depsEntry := range entries {
		if !strings.Contains(id, ".com") {
			continue
		}
		// We want it in https://{DepsEntry.Id} format, without the .git
		u, _ := url.JoinPath("https://", id)
		denormalized[u] = depsEntry.Version
	}

	return denormalized, nil
}

// findRolledDep searches for the dependency that may have been rolled.
func (m *MidpointHandler) findRolledDep(startDeps, endDeps map[string]string) string {
	for k, v := range startDeps {
		// If the dep doesn't exist, it couldn't have been rolled. Skip.
		if _, ok := endDeps[k]; !ok {
			continue
		}
		if v != endDeps[k] {
			return k
		}
	}

	return ""
}

// determineRolledDep coordinates the search to find which dep may have been rolled for adjacent commits.
func (m *MidpointHandler) determineRolledDep(ctx context.Context, url, startGitHash, endGitHash string) (*CombinedCommit, *Commit, *Commit, error) {
	gc := m.getOrCreateRepo(url)

	// Fetch deps for each git hash for the project
	startDeps, err := m.fetchGitDeps(ctx, gc, startGitHash)
	if err != nil {
		return nil, nil, nil, err
	}

	endDeps, err := m.fetchGitDeps(ctx, gc, endGitHash)
	if err != nil {
		return nil, nil, nil, err
	}

	// Find the first URL.
	diffUrl := m.findRolledDep(startDeps, endDeps)

	// DEPS are the same.
	if diffUrl == "" {
		return nil, nil, nil, nil
	}

	dStart := startDeps[diffUrl]
	left := &Commit{
		RepositoryUrl: diffUrl,
		GitHash:       dStart,
	}
	dEnd := endDeps[diffUrl]
	right := &Commit{
		RepositoryUrl: diffUrl,
		GitHash:       dEnd,
	}

	dNext, err := m.findMidpoint(ctx, diffUrl, dStart, dEnd)
	if err != nil {
		return nil, nil, nil, err
	}
	next := &CombinedCommit{
		Main: &Commit{
			RepositoryUrl: url,
			// Start and End githash only diffs in DEPS, pick the lower bound
			GitHash: startGitHash,
		},
		ModifiedDeps: []*Commit{
			{
				RepositoryUrl: diffUrl,
				GitHash:       dNext,
			},
		},
	}

	return next, left, right, nil
}

// FindDepsCommit finds the commit in the DEPS for the given repo.
//
// It returns a Commit that can be used to search for middle commit in the DEPS and then construct
// a CombinedCommit to build Chrome with modified DEPS.
func (m *MidpointHandler) FindDepsCommit(ctx context.Context, c *Commit, repoUrl string) (*Commit, error) {
	gc := m.getOrCreateRepo(c.RepositoryUrl)
	deps, err := m.fetchGitDeps(ctx, gc, c.GitHash)
	if err != nil {
		return nil, skerr.Wrap(err)
	}

	h, ok := deps[repoUrl]
	if !ok {
		return nil, skerr.Fmt("%s doesn't exist in DEPS", repoUrl)
	}

	return &Commit{
		RepositoryUrl: repoUrl,
		GitHash:       h,
	}, nil
}

// FindMidCommit finds the middle commit from the two given commits.
//
// It uses gitiles API to find the middle commit, and it also handles DEPS rolls when two commits
// are adjacent. If two commits are adjacent and no DEPS roll, then the first commit is returned;
// If two commits are adjacent and there is a DEPS roll on the second commit, then it will search
// for rolled repositories and find the middle commit between the roll.
//
// Note the returned Commit can be a different repo because it looks at DEPS, but it only looks at
// one level. If the DEPS of DEPS has rolls, it will not continue to search.
func (m *MidpointHandler) FindMidCommit(ctx context.Context, startCommit, endCommit *Commit) (*Commit, error) {
	if startCommit.RepositoryUrl != endCommit.RepositoryUrl {
		return nil, skerr.Fmt("two commits are from different repos")
	}

	baseUrl := startCommit.RepositoryUrl
	startGitHash, endGitHash := startCommit.GitHash, endCommit.GitHash
	nextCommitHash, err := m.findMidpoint(ctx, baseUrl, startGitHash, endGitHash)
	if err != nil {
		return nil, err
	}

	// If startGitHash and endGitHash are not adjacent, return the found commit right away.
	//
	// We use HasPrefix because nextCommitHash will always be the full SHA git hash,
	// but the provided startGitHash may be a short SHA.
	if !strings.HasPrefix(nextCommitHash, startGitHash) {
		return &Commit{
			RepositoryUrl: baseUrl,
			GitHash:       nextCommitHash,
		}, nil
	}

	// The nextCommit == startHash. This means start and end are adjacent commits.
	// Assume a DEPS roll, so we'll find the next candidate by parsing DEPS rolls.
	sklog.Debugf("Start hash %s and end hash %s are adjacent to each other. Assuming a DEPS roll.", startGitHash, endGitHash)

	next, _, _, err := m.determineRolledDep(ctx, baseUrl, startGitHash, endGitHash)
	if err != nil {
		return nil, err
	}

	// If endGitHash doesn't have DEPS rolls, return the first commit.
	if next == nil {
		return startCommit, nil
	}

	return next.ModifiedDeps[0], nil
}
