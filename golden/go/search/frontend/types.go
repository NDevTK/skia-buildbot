// Package frontend contains structs that represent how the frontend expects output from the search
// package.
package frontend

import (
	"time"

	"go.skia.org/infra/go/paramtools"
	"go.skia.org/infra/golden/go/expectations"
	"go.skia.org/infra/golden/go/search/common"
	"go.skia.org/infra/golden/go/tiling"
	"go.skia.org/infra/golden/go/types"
	"go.skia.org/infra/golden/go/web/frontend"
)

// SearchResponse is the structure returned by the Search(...) function of SearchAPI and intended
// to be returned as JSON in an HTTP response.
type SearchResponse struct {
	Results []*SearchResult `json:"digests" go2ts:"ignorenil"`
	// Offset is the offset of the digest into the total list of digests.
	Offset int `json:"offset"`
	// Size is the total number of Digests that match the current query.
	Size    int               `json:"size"`
	Commits []frontend.Commit `json:"commits"`
	// BulkTriageData contains *all* digests that match the query as keys. The value for each key is
	// an expectations.Label value giving the label of the closest triaged digest to the key digest
	// or empty string if there is no "closest digest". Note the similarity to the
	// frontend.TriageRequest type.
	BulkTriageData frontend.TriageRequestData `json:"bulk_triage_data"`
}

// TriageHistory represents who last triaged a certain digest for a certain test.
type TriageHistory struct {
	User string    `json:"user"`
	TS   time.Time `json:"ts"`
}

// SearchResult is a single digest produced by one or more traces for a given test.
type SearchResult struct {
	// Digest is the primary digest to which the rest of the data in this struct belongs.
	Digest types.Digest `json:"digest"`
	// Test is the name of the test that produced the primary digest. This is needed because
	// we might have a case where, for example, a blank 100x100 image is correct for one test,
	// but not for another test and we need to distinguish between the two cases.
	Test types.TestName `json:"test"`
	// Status is positive, negative, or untriaged. This is also known as the expectation for the
	// primary digest (for Test).
	Status expectations.Label `json:"status"`
	// TriageHistory is a history of all the times the primary digest has been retriaged for the
	// given Test.
	TriageHistory []TriageHistory `json:"triage_history"`
	// ParamSet is all the keys and options of all traces that produce the primary digest and
	// match the given search constraints. It is for frontend UI presentation only; essentially a
	// word cloud of what drew the primary digest.
	ParamSet paramtools.ParamSet `json:"paramset"`
	// TODO(kjlubick) make use of these instead of the combined ParamSet.
	TracesKeys    paramtools.ParamSet `json:"-"`
	TracesOptions paramtools.ParamSet `json:"-"`
	// TraceGroup represents all traces that produced this digest at least once in the sliding window
	// of commits.
	TraceGroup TraceGroup `json:"traces"`
	// RefDiffs are comparisons of the primary digest to other digests in Test. As an example, the
	// closest digest (closest being defined as least different) also triaged positive is usually
	// in here (unless there are no other positive digests).
	// TODO(kjlubick) map is confusing because it's just 2 things. Use struct instead.
	RefDiffs map[common.RefClosest]*SRDiffDigest `json:"refDiffs"`
	// ClosestRef labels the reference from RefDiffs that is the absolute closest to the primary
	// digest.
	ClosestRef common.RefClosest `json:"closestRef"` // "pos" or "neg"
}

// SRDiffDigest captures the diff information between a primary digest and the digest given here.
// The primary digest is generally shown on the left in the frontend UI, and the data here
// represents a digest on the right that the primary digest is being compared to.
type SRDiffDigest struct {
	// NumDiffPixels is the absolute number of pixels that are different.
	NumDiffPixels int `json:"numDiffPixels"`

	// CombinedMetric is a value in [0, 10] that represents how large the diff is between two
	// images. It is based off the MaxRGBADiffs and PixelDiffPercent.
	CombinedMetric float32 `json:"combinedMetric"`

	// PixelDiffPercent is the percentage of pixels that are different.
	PixelDiffPercent float32 `json:"pixelDiffPercent"`

	// MaxRGBADiffs contains the maximum difference of each channel.
	MaxRGBADiffs [4]int `json:"maxRGBADiffs"`

	// One of CombinedMetric, PixelDiffPercent, or NumDiffPixels depending on the requested
	// metric name (see query.go). Used internally in search.
	QueryMetric float32 `json:"-"`

	// DimDiffer is true if the dimensions between the two images are different.
	DimDiffer bool `json:"dimDiffer"`

	// Digest identifies which image we are comparing the primary digest to. Put another way, what
	// is the image on the right side of the comparison.
	Digest types.Digest `json:"digest"`
	// Status represents the expectation.Label for this digest.
	Status expectations.Label `json:"status"`
	// ParamSet is all of the params of all traces that produce this digest (the digest on the right).
	// It is for frontend UI presentation only; essentially a word cloud of what drew the primary
	// digest.
	ParamSet paramtools.ParamSet `json:"paramset"`
	// TODO(kjlubick) make use of these instead of the combined ParamSet.
	TracesKeys    paramtools.ParamSet `json:"-"`
	TracesOptions paramtools.ParamSet `json:"-"`
}

// DigestDetails contains details about a digest.
type DigestDetails struct {
	Result  SearchResult      `json:"digest"`
	Commits []frontend.Commit `json:"commits"`
}

// Trace describes a single trace, used in TraceGroup.
type Trace struct {
	// The id of the trace. Keep the json as label to be compatible with dots-sk.
	ID tiling.TraceID `json:"label"`
	// RawTrace is meant to be used to hold the raw trace (that is, the tiling.Trace which has not yet
	// been converted for frontend display) until all the raw traces for a given
	// TraceGroup can be converted to the frontend representation. The conversion process needs to be
	// done once all the RawTraces are available so the digest indices can be in agreement for a given
	// TraceGroup. It is not meant to be exposed to the frontend in its raw form.
	RawTrace *tiling.Trace `json:"-"`
	// DigestIndices represents the index of the digest that was part of the trace. -1 means we did
	// not get a digest at this commit. There is one entry per commit. DigestIndices[0] is the oldest
	// commit in the trace, DigestIndices[N-1] is the most recent. The index here matches up with
	// the Digests in the parent TraceGroup.
	DigestIndices []int `json:"data"`
	// Params are the key/value pairs that describe this trace.
	Params map[string]string `json:"params"`
	// TODO(kjlubick) Use these split values instead of the combined one.
	Keys    map[string]string `json:"-"`
	Options map[string]string `json:"-"`
	// CommentIndices are indices into the TraceComments slice on the final result. For example,
	// a 1 means the second TraceComment in the top level TraceComments applies to this trace.
	CommentIndices []int `json:"comment_indices"`
}

// TraceGroup is info about a group of traces. The concrete use of TraceGroup is to represent all
// traces that draw a given digest (known below as the "primary digest") for a given test.
type TraceGroup struct {
	// Traces represents all traces in the TraceGroup. All of these traces have the primary digest.
	Traces []Trace `json:"traces"`
	// Digests represents the triage status of the primary digest and the first N-1 digests that
	// appear in Traces, starting at head on the first trace. N is search.maxDistinctDigestsToPresent.
	Digests []DigestStatus `json:"digests"`
	// TotalDigests is the count of all unique digests in the set of Traces. This number can
	// exceed search.maxDistinctDigestsToPresent.
	TotalDigests int `json:"total_digests"`
}

// DigestStatus is a digest and its status, used in TraceGroup.
type DigestStatus struct {
	Digest types.Digest       `json:"digest"`
	Status expectations.Label `json:"status"`
}

// DigestComparison contains the result of comparing two digests.
type DigestComparison struct {
	Left  SearchResult  `json:"left"`  // The left hand digest and its params.
	Right *SRDiffDigest `json:"right"` // The right hand digest, its params and the diff result.
}

// UntriagedDigestList represents multiple digests that are untriaged for a given query.
type UntriagedDigestList struct {
	Digests []types.Digest `json:"digests"`

	// Corpora is filed with the strings representing a corpus that has one or more Digests belong
	// to it. In other words, it summarizes where the Digests come from.
	Corpora []string `json:"corpora"`

	// TS is the time that this data was created. It might be served from a cache, so this time will
	// not necessarily be "now".
	TS time.Time `json:"ts"`
}
