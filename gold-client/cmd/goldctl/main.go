package main

// goldctl is a CLI for working with the Gold service.

import (
	"fmt"
	"os"

	"go.skia.org/infra/golden/go/search"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"go.skia.org/infra/gold-client/go/goldclient"
	"go.skia.org/infra/golden/go/jsonio"
)

var (
	// Root command: goldctl itself.
	rootCmd *cobra.Command

	// Flags used throughout all commands.
	flagVerbose bool

	// Flags used by validate
	flagFile string

	// Flags used by imgtest:init and imgtest:add command
	flagCommit       string
	flagKeysFile     string
	flagIssueID      string
	flagPatchsetID   string
	flagJobID        string
	flagInstandID    string
	flagWorkDir      string
	flagPassFailStep bool
	flagFailureFile  string

	// Flags used by imgtest:add
	flagTestName string
	flagPNGFile  string
)

func init() {
	// Set up the root command.
	rootCmd = &cobra.Command{
		Use: "goldctl",
		Long: `
goldctl interacts with the Gold service.
It can be used directly or in a scripted environment. `,
	}
	rootCmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "Verbose prints out extra information")

	// validate command
	validateCmd := &cobra.Command{
		Use:     "validate",
		Aliases: []string{"va"},
		Short:   "Validate JSON",
		Long: `
Validate JSON input whether it complies with the format required for Gold
ingestion.`,
		Run: runValidateCmd,
	}
	validateCmd.Flags().StringVarP(&flagFile, "file", "f", "", "Input file to use instead of stdin")
	validateCmd.Args = cobra.NoArgs

	// auth command
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate against GCP",
		Long: `
Authenticate against GCP - TODO: How to specify the service account file ? `,
		Run: runAuthCommand,
	}

	// imgtest command and it's sub commands
	imgTestCmd := &cobra.Command{
		Use:   "imgtest",
		Short: "Collect  and upload test results as images",
		Long: `
Collect and upload test results to the Gold backend.`,
	}

	// cmd: imgtest init
	imgTestInitCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a  testing environment",
		Long: `
Start a testing session during which tests are added. This initializes the environment.
It gathers wether the 'add' command returns a pass/fail value and the common
keys shared by all tests that are added via 'add'.`,
		Run: runImgTestInitCommand,
	}
	addEnvFlags(imgTestInitCmd, false)

	imgTestAddCmd := &cobra.Command{
		Use:   "add",
		Short: "Adds a test image to the results.",
		Long: `
Add images generated by the tests to the test results. This requires two arguments:
			 - The test name
			 - The path to the resulting PNG.
`,
		Run:  runImgTestAddCommand,
		Args: cobra.NoArgs,
	}
	addEnvFlags(imgTestAddCmd, false) // Eventually the argument should be true.
	imgTestAddCmd.Flags().StringVarP(&flagTestName, "test-name", "", "", "Unique name of the test, must not contain spaces.")
	imgTestAddCmd.Flags().StringVarP(&flagPNGFile, "png-file", "", "", "Path to the PNG file that contains the test results.")

	imgTestFinalizeCmd := &cobra.Command{
		Use:   "finalize",
		Short: "Finish adding tests and process results.",
		Long: `
All tests have been added. Upload images and generate and upload the JSON file that captures
test results.`,
		Run: runImgTestFinalizeCommand,
	}

	imgTestPassFailCmd := &cobra.Command{
		Use:   "passfail",
		Short: "Checks whether the results match expectations",
		Long: `
Check against Gold or local baseline whether the results match the expectations`,
		Run: runImgTestPassFailCommand,
	}

	// assemble the imgtest command.
	imgTestCmd.AddCommand(
		imgTestInitCmd,
		imgTestAddCmd,
		imgTestFinalizeCmd,
		imgTestPassFailCmd,
	)

	// Wire up the commands as children of the root command.
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(imgTestCmd)
}

func main() {
	// Execute the root command.
	if err := rootCmd.Execute(); err != nil {
		logErrAndExit(rootCmd, err)
	}
}

func runImgTestAddCommand(cmd *cobra.Command, args []string) {
	// TODO(stephana): Remove after this stub lands.
	os.Exit(0)

	keyMap, err := readKeysFile(flagKeysFile)
	ifErrLogExit(cmd, err)

	validation := search.Validation{}
	issueID := validation.Int64Value("issue", flagIssueID, 0)
	patchsetID := validation.Int64Value("pachset", flagPatchsetID, 0)
	jobID := validation.Int64Value("jobid", flagJobID, 0)
	ifErrLogExit(cmd, validation.Errors())

	gr := &jsonio.GoldResults{
		GitHash:       flagCommit,
		Key:           keyMap,
		Issue:         issueID,
		Patchset:      patchsetID,
		BuildBucketID: jobID,
	}

	up, err := goldclient.NewUploadResults(gr, flagInstandID, flagPassFailStep)
	ifErrLogExit(cmd, err)
	logInfof(cmd, "CONFIG: \n%s", spew.Sdump(up))

	goldClient, err := goldclient.NewCloudClient(up)
	ifErrLogExit(cmd, err)
	fmt.Printf("\n\n\nBerofr\n\n\n")

	pass, err := goldClient.Test(flagTestName, flagPNGFile)
	fmt.Printf("\n\n\nAFTER\n\n\n")
	ifErrLogExit(cmd, err)

	if !pass {
		os.Exit(1)
	}
}

func runAuthCommand(cmd *cobra.Command, args []string)            { notImplemented(cmd) }
func runImgTestInitCommand(cmd *cobra.Command, args []string)     { notImplemented(cmd) }
func runImgTestFinalizeCommand(cmd *cobra.Command, args []string) { notImplemented(cmd) }
func runImgTestPassFailCommand(cmd *cobra.Command, args []string) { notImplemented(cmd) }

func readKeysFile(fileName string) (map[string]string, error) {
	return nil, nil
}

func notImplemented(cmd *cobra.Command) {
	logErr(cmd, fmt.Errorf("Command not implemented yet."))
	os.Exit(1)
}

// func ifErrLogExit(err error) {
// 	if err != nil {
// 		sklog.Errorf("Error: %s", err)
// 		os.Exit(1)
// 	}
// }

// TODO(stephana): REMOVE !!
// global flags
// --instance
// --work-dir
// --passfail

// --commit
// --keys-file
// --issue
// --patchset
// --jobid
// --failure-file

// per test flags
// --test-name
// --png-file

func addEnvFlags(cmd *cobra.Command, optional bool) {
	cmd.Flags().StringVarP(&flagInstandID, "instance", "", "", "ID of the Gold instance.")
	cmd.Flags().StringVarP(&flagWorkDir, "work-dir", "", "", "Temporary work directory")
	cmd.Flags().BoolVarP(&flagPassFailStep, "passfail", "", false, "Whether the 'add' call returns a pass/fail for each test.")

	cmd.Flags().StringVarP(&flagCommit, "commit", "", "", "Git commit hash")
	cmd.Flags().StringVarP(&flagKeysFile, "keys-file", "", "", "File containing key/value pairs commmon to all tests")
	cmd.Flags().StringVarP(&flagIssueID, "issue", "", "", "Gerrit issue if this is trybot run. ")
	cmd.Flags().StringVarP(&flagPatchsetID, "patchset", "", "", "Gerrit patchset number if this is a trybot run. ")
	cmd.Flags().StringVarP(&flagJobID, "jobid", "", "", "Job ID if this is a tryjob run. Current the BuildBucket id.")
	cmd.Flags().StringVarP(&flagFailureFile, "failure-file", "", "", "Path to the file where to write failure information")

	if !optional {
		cmd.MarkFlagRequired("instance")
		cmd.MarkFlagRequired("work-dir")
		cmd.MarkFlagRequired("passfail")
		cmd.MarkFlagRequired("commit")
		cmd.MarkFlagRequired("keys-file")
	}
}

// runValidateCmd implements the validation logic.
func runValidateCmd(cmd *cobra.Command, args []string) {
	f, closeFn, err := getFileOrStdin(flagFile)
	if err != nil {
		logErrfAndExit(cmd, "Error opeing input: %s", err)
	}

	goldResult, errMessages, err := jsonio.ParseGoldResults(f)
	if err != nil {
		if len(errMessages) == 0 {
			logErrfAndExit(cmd, "Error parsing JSON: %s", err)
		}

		logErr(cmd, "JSON validation failed:\n")
		for _, msg := range errMessages {
			logErrf(cmd, "   %s\n", msg)
		}
		os.Exit(1)
	}
	ifErrLogExit(cmd, closeFn())
	logVerbose(cmd, fmt.Sprintf("Result:\n%s\n", spew.Sdump(goldResult)))
	logVerbose(cmd, "JSON validation succeeded.\n")
}

// getFileOrStdin returns an file to read from based on the whether file flag was set.
func getFileOrStdin(inputFile string) (*os.File, func() error, error) {
	if inputFile == "" {
		return os.Stdin, func() error { return nil }, nil
	}

	f, err := os.Open(inputFile)
	if err != nil {
		return nil, nil, err
	}
	return f, f.Close, nil
}

// logErrf logs a formatted error based on the output settings of the command.
func logErrf(cmd *cobra.Command, format string, args ...interface{}) {
	_, _ = fmt.Fprintf(cmd.OutOrStderr(), format, args...)
}

// logErr logs an error based on the output settings of the command.
func logErr(cmd *cobra.Command, args ...interface{}) {
	_, _ = fmt.Fprint(cmd.OutOrStderr(), args...)
}

// logErrAndExit logs a formatted error and exits with a non-zero exit code.
func logErrAndExit(cmd *cobra.Command, err error) {
	logErr(cmd, err)
	os.Exit(1)
}

// logErrfAndExit logs an error and exits with a non-zero exit code.
func logErrfAndExit(cmd *cobra.Command, format string, err error) {
	logErrf(cmd, format, err)
	os.Exit(1)
}

// ifErrLogExit logs an error if the proviced error is not nil and exits
// with a non-zero exit code.
func ifErrLogExit(cmd *cobra.Command, err error) {
	if err != nil {
		logErr(cmd, err)
		os.Exit(1)
	}
}

// logInfo logs the given arguments based on the output settings of the command.
func logInfo(cmd *cobra.Command, args ...interface{}) {
	_, _ = fmt.Fprint(cmd.OutOrStdout(), args...)
}

// logInfo logs the given arguments based on the output settings of the command.
func logInfof(cmd *cobra.Command, format string, args ...interface{}) {
	_, _ = fmt.Fprintf(cmd.OutOrStdout(), format, args...)
}

// logVerbose logs the given arguments if the verbose flag is true.
func logVerbose(cmd *cobra.Command, args ...interface{}) {
	if flagVerbose {
		logInfo(cmd, args...)
	}
}
