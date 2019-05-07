// package baseline contains functions to gather the current baseline and
// write them to GCS.
package baseline

import (
	"fmt"

	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/sklog"
	"go.skia.org/infra/go/tiling"
	"go.skia.org/infra/go/util"
	"go.skia.org/infra/golden/go/expstorage"
	"go.skia.org/infra/golden/go/tryjobstore"
	"go.skia.org/infra/golden/go/types"
)

// md5SumEmptyExp is the MD5 sum of an empty expectation.
// it is initialized in this file's init().
var md5SumEmptyExp = ""

func init() {
	var err error
	md5SumEmptyExp, err = util.MD5Sum(types.TestExp{})
	if err != nil {
		panic(fmt.Sprintf("Could not get the MD5 sum of an empty expectation: %s", err))
	}
}

// TODO(kjlubick): Add tests for GetBaselinesPerCommit.

// GetBaselinesPerCommit calculates the baselines for each commit in the tile.
// Of note, it only fills out the Positive matches - everything not seen is either untriaged
// or negative.
// If extraCommits is not empty they are assumed to be commits immediately following the commits
// in the given TileInfo and Baselines for these should be essentially copies of the last
// commit in the tile. This covers the case when the current tile is slightly behind all commits
// that have already been added to the repository.
func GetBaselinesPerCommit(exps types.TestExpBuilder, tileInfo TileInfo, extraCommits []*tiling.Commit) (map[string]*CommitableBaseline, error) {
	allCommits := tileInfo.AllCommits()
	if len(allCommits) == 0 {
		return map[string]*CommitableBaseline{}, nil
	}

	// Get the baselines for all commits for which we have data and that are not ignored.
	denseTile := tileInfo.GetTile(false)
	denseCommits := tileInfo.DataCommits()
	denseBaselines := make(map[string]types.TestExp, len(denseCommits))

	// Initialize the expectations for all data commits
	for _, commit := range denseCommits {
		denseBaselines[commit.Hash] = make(types.TestExp, len(denseTile.Traces))
	}

	// Sweep the tile and calculate the baselines.
	// For each trace we make a set of triaged, positive digests and for each
	// commit on the trace we add that set to the baseline.
	for _, trace := range denseTile.Traces {
		gTrace := trace.(*types.GoldenTrace)
		currDigests := map[string]types.Label{}
		testName := gTrace.Keys[types.PRIMARY_KEY_FIELD]
		for idx := 0; idx < len(denseCommits); idx++ {
			digest := gTrace.Digests[idx]

			// If the digest is not missing then add the digest to the running list of digests.
			if digest != types.MISSING_DIGEST {
				if _, ok := currDigests[digest]; !ok {
					if cl := exps.Classification(testName, digest); cl == types.POSITIVE {
						currDigests[digest] = cl
					}
				}
			}

			if len(currDigests) > 0 {
				denseBaselines[denseCommits[idx].Hash].AddDigests(testName, currDigests)
			}
		}
	}

	// Iterate over all commits. If the tile is sparse we substitute the expectations with the
	// expectations of the closest ancestor that has expecations. We also add the commits that
	// have landed already, but are not captured in the current tile.
	combined := allCommits
	if len(extraCommits) > 0 {
		combined = make([]*tiling.Commit, 0, len(allCommits)+len(extraCommits))
		combined = append(combined, allCommits...)
		combined = append(combined, extraCommits...)
	}

	ret := make(map[string]*CommitableBaseline, len(combined))
	var currBL *CommitableBaseline = nil
	for _, commit := range combined {
		bl, ok := denseBaselines[commit.Hash]
		if ok {
			md5Sum, err := util.MD5Sum(bl)
			if err != nil {
				return nil, skerr.Fmt("Error calculating MD5 sum: %s", err)
			}

			ret[commit.Hash] = &CommitableBaseline{
				StartCommit: commit,
				EndCommit:   commit,
				Total:       len(denseTile.Traces),
				Filled:      len(bl),
				Baseline:    bl,
				MD5:         md5Sum,
			}
			currBL = ret[commit.Hash]
		} else if currBL != nil {
			// Make a copy of the baseline of the previous commit and update the commit information.
			cpBL := *currBL
			cpBL.StartCommit = commit
			cpBL.EndCommit = commit
			ret[commit.Hash] = &cpBL
		} else {
			// Reaching this point means the first in the dense tile does not align with the first
			// commit of the sparse portion of the tile. This a sanity test and should only happen
			// in the presence of a programming error or data corruption.
			sklog.Errorf("Unable to get baseline for commit %s. It has not commits for immediate ancestors in tile.", commit.Hash)
		}
	}

	return ret, nil
}

// GetBaselineForIssue returns the baseline for the given issue. This baseline
// contains all triaged digests that are not in the master tile.
func GetBaselineForIssue(issueID int64, tryjobs []*tryjobstore.Tryjob, tryjobResults [][]*tryjobstore.TryjobResult, exp types.TestExpBuilder, commits []*tiling.Commit) (*CommitableBaseline, error) {
	var startCommit *tiling.Commit = commits[len(commits)-1]
	var endCommit *tiling.Commit = commits[len(commits)-1]

	b := types.TestExp{}
	for idx, tryjob := range tryjobs {
		for _, result := range tryjobResults[idx] {
			if result.Digest != types.MISSING_DIGEST && exp.Classification(result.TestName, result.Digest) == types.POSITIVE {
				b.AddDigest(result.TestName, result.Digest, types.POSITIVE)
			}

			_, c := tiling.FindCommit(commits, tryjob.MasterCommit)
			startCommit = minCommit(startCommit, c)
			endCommit = maxCommit(endCommit, c)
		}
	}

	md5Sum, err := util.MD5Sum(b)
	if err != nil {
		return nil, skerr.Fmt("Error calculating MD5 sum: %s", err)
	}

	// Note: CommitDelta, Total and Filled are not relevant for an issue baseline since there
	// are not traces and commits directly related to this.

	// TODO(stephana): Review whether StartCommit and EndCommit are useful here.

	ret := &CommitableBaseline{
		StartCommit: startCommit,
		EndCommit:   endCommit,
		Baseline:    b,
		Issue:       issueID,
		MD5:         md5Sum,
	}
	return ret, nil
}

// commitIssueBaseline commits the expectations for the given issue to the master baseline.
func commitIssueBaseline(issueID int64, user string, issueChanges types.TestExp, tryjobStore tryjobstore.TryjobStore, expStore expstorage.ExpectationsStore) error {
	if len(issueChanges) == 0 {
		return nil
	}

	syntheticUser := fmt.Sprintf("%s:%d", user, issueID)

	commitFn := func() error {
		if err := expStore.AddChange(issueChanges, syntheticUser); err != nil {
			return skerr.Fmt("Unable to add expectations for issue %d: %s", issueID, err)
		}
		return nil
	}

	return tryjobStore.CommitIssueExp(issueID, commitFn)
}

// minCommit returns newCommit if it appears before current (or current is nil).
func minCommit(current *tiling.Commit, newCommit *tiling.Commit) *tiling.Commit {
	if current == nil || newCommit == nil || newCommit.CommitTime < current.CommitTime {
		return newCommit
	}
	return current
}

// maxCommit returns newCommit if it appears after current (or current is nil).
func maxCommit(current *tiling.Commit, newCommit *tiling.Commit) *tiling.Commit {
	if current == nil || newCommit == nil || newCommit.CommitTime > current.CommitTime {
		return newCommit
	}
	return current
}
