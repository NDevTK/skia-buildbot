package gitstore

import (
	"context"
	"fmt"
	"sort"
	"testing"

	assert "github.com/stretchr/testify/require"
	"go.skia.org/infra/go/git/gitinfo"
	"go.skia.org/infra/go/testutils"
	"go.skia.org/infra/go/timer"
	"go.skia.org/infra/go/vcsinfo"
	vcs_testutils "go.skia.org/infra/go/vcsinfo/testutils"
)

// This test requires a checkout of a repo (can be really any repo) in a directory named 'skia'
// in the same directory as this test.
func TestLargeGitStore(t *testing.T) {
	testutils.ManualTest(t)
	testGitStore(t, skiaRepoURL, skiaRepoDir, true)
}

func TestNormURL(t *testing.T) {
	testutils.SmallTest(t)
	httpURL := "https://github.com/skia-dev/textfiles.git"
	gitURL := "ssh://git@github.com/skia-dev/textfiles"
	gitURLWithExt := "ssh://git@github.com:skia-dev/textfiles.git"
	normHTTP, err := NormalizeURL(httpURL)
	assert.NoError(t, err)
	normGit, err := NormalizeURL(gitURL)
	assert.NoError(t, err)
	normGitWithExt, err := NormalizeURL(gitURLWithExt)
	assert.NoError(t, err)

	// Make sure they all match.
	assert.Equal(t, "github.com/skia-dev/textfiles", normHTTP)
	assert.Equal(t, normHTTP, normGit)
	assert.Equal(t, normHTTP, normGitWithExt)
}

func TestGitStoreLocalRepo(t *testing.T) {
	testutils.LargeTest(t)

	repoDir, cleanup := vcs_testutils.InitTempRepo()
	defer cleanup()
	testGitStore(t, localRepoURL, repoDir, true)
}

func testGitStore(t *testing.T, repoURL, repoDir string, freshLoad bool) {
	// Get all commits that have been added to the gitstore.
	_, longCommits, gitStore := setupAndLoadGitStore(t, repoURL, repoDir, freshLoad)

	// Sort long commits they way they are sorted by BigTable (by timestamp/hash)
	sort.Slice(longCommits, func(i, j int) bool {
		tsI := longCommits[i].Timestamp.Unix()
		tsJ := longCommits[j].Timestamp.Unix()
		return (tsI < tsJ) || ((tsI == tsJ) && (longCommits[i].Hash < longCommits[j].Hash))
	})
	indexCommits := make([]*vcsinfo.IndexCommit, len(longCommits))
	for idx, commit := range longCommits {
		indexCommits[idx] = &vcsinfo.IndexCommit{
			Index:     idx,
			Hash:      commit.Hash,
			Timestamp: commit.Timestamp,
		}
	}

	// Find all the commits in the repository independent of branches.
	foundIndexCommits, foundLongCommits := getFromRange(t, gitStore, 0, len(longCommits), "")
	assert.Equal(t, len(indexCommits), len(foundIndexCommits))
	assert.Equal(t, len(longCommits), len(foundLongCommits))

	// Make sure they match what we found.
	for idx, expected := range longCommits {
		foundLongCommits[idx].Branches = expected.Branches
		assert.Equal(t, expected, foundLongCommits[idx])
	}

	// Verify that the branches from the GitStore match what's in the checkout.
	branchNames, branchCommits := getBranchCommits(t, repoDir)
	for branchIdx, branchName := range branchNames {
		expHashes := branchCommits[branchIdx]
		foundIndexCommits, foundLongCommits := getFromRange(t, gitStore, 0, len(longCommits), branchName)
		assert.Equal(t, len(expHashes), len(foundIndexCommits))
		assert.Equal(t, len(expHashes), len(foundLongCommits))
		expIdx := len(expHashes) - 1
		for idx := len(foundIndexCommits) - 1; idx >= 0; idx-- {
			expHash := expHashes[expIdx]
			assert.Equal(t, foundIndexCommits[idx].Hash, foundLongCommits[idx].Hash)
			assert.Equal(t, expHash, foundIndexCommits[idx].Hash)
			expIdx--
		}
	}
}

func getBranchCommits(t *testing.T, repoDir string) ([]string, [][]string) {
	ctx := context.TODO()
	vcs, err := gitinfo.NewGitInfo(ctx, repoDir, false, true)
	assert.NoError(t, err)

	branches, err := vcs.GetBranches(ctx)
	assert.NoError(t, err)

	branchNames := make([]string, 0, len(branches))
	branchCommits := make([][]string, 0, len(branches))
	for _, branch := range branches {
		// if strings.Contains(branch.Name, "m62") {
		// 	continue
		// }
		branchNames = append(branchNames, branch.Name)
		indexCommits, err := gitinfo.GetBranchCommits(ctx, repoDir, branch.Name)
		assert.NoError(t, err)
		commitHashes := make([]string, len(indexCommits))
		for idx, idxCommit := range indexCommits {
			commitHashes[idx] = idxCommit.Hash
		}
		branchCommits = append(branchCommits, commitHashes)
	}

	return branchNames, branchCommits
}

func getFromRange(t *testing.T, gitStore GitStore, startIdx, endIdx int, branchName string) ([]*vcsinfo.IndexCommit, []*vcsinfo.LongCommit) {
	ctx := context.TODO()

	tQuery := timer.New(fmt.Sprintf("RangeN %d - %d commits from branch %q", startIdx, endIdx, branchName))
	foundIndexCommits, err := gitStore.RangeN(ctx, startIdx, endIdx, branchName)
	assert.NoError(t, err)
	tQuery.Stop()

	hashes := make([]string, 0, len(foundIndexCommits))
	for _, commit := range foundIndexCommits {
		hashes = append(hashes, commit.Hash)
	}

	tLongCommits := timer.New(fmt.Sprintf("Get %d LongCommits from branch %q", len(hashes), branchName))
	foundLongCommits, err := gitStore.Get(ctx, hashes)
	assert.NoError(t, err)
	assert.Equal(t, len(foundIndexCommits), len(foundLongCommits))
	tLongCommits.Stop()

	return foundIndexCommits, foundLongCommits
}
