// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	frontend "go.skia.org/infra/golden/go/web/frontend"

	paramtools "go.skia.org/infra/go/paramtools"

	query "go.skia.org/infra/golden/go/search/query"

	search2 "go.skia.org/infra/golden/go/search2"

	time "time"

	types "go.skia.org/infra/golden/go/types"
)

// API is an autogenerated mock type for the API type
type API struct {
	mock.Mock
}

// ChangelistLastUpdated provides a mock function with given fields: ctx, qCLID
func (_m *API) ChangelistLastUpdated(ctx context.Context, qCLID string) (time.Time, error) {
	ret := _m.Called(ctx, qCLID)

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(context.Context, string) time.Time); ok {
		r0 = rf(ctx, qCLID)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, qCLID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountDigestsByTest provides a mock function with given fields: ctx, q
func (_m *API) CountDigestsByTest(ctx context.Context, q frontend.ListTestsQuery) (frontend.ListTestsResponse, error) {
	ret := _m.Called(ctx, q)

	var r0 frontend.ListTestsResponse
	if rf, ok := ret.Get(0).(func(context.Context, frontend.ListTestsQuery) frontend.ListTestsResponse); ok {
		r0 = rf(ctx, q)
	} else {
		r0 = ret.Get(0).(frontend.ListTestsResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, frontend.ListTestsQuery) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlamesForUntriagedDigests provides a mock function with given fields: ctx, corpus
func (_m *API) GetBlamesForUntriagedDigests(ctx context.Context, corpus string) (search2.BlameSummaryV1, error) {
	ret := _m.Called(ctx, corpus)

	var r0 search2.BlameSummaryV1
	if rf, ok := ret.Get(0).(func(context.Context, string) search2.BlameSummaryV1); ok {
		r0 = rf(ctx, corpus)
	} else {
		r0 = ret.Get(0).(search2.BlameSummaryV1)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, corpus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChangelistParamset provides a mock function with given fields: ctx, crs, clID
func (_m *API) GetChangelistParamset(ctx context.Context, crs string, clID string) (paramtools.ReadOnlyParamSet, error) {
	ret := _m.Called(ctx, crs, clID)

	var r0 paramtools.ReadOnlyParamSet
	if rf, ok := ret.Get(0).(func(context.Context, string, string) paramtools.ReadOnlyParamSet); ok {
		r0 = rf(ctx, crs, clID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(paramtools.ReadOnlyParamSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, crs, clID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCluster provides a mock function with given fields: ctx, opts
func (_m *API) GetCluster(ctx context.Context, opts search2.ClusterOptions) (frontend.ClusterDiffResult, error) {
	ret := _m.Called(ctx, opts)

	var r0 frontend.ClusterDiffResult
	if rf, ok := ret.Get(0).(func(context.Context, search2.ClusterOptions) frontend.ClusterDiffResult); ok {
		r0 = rf(ctx, opts)
	} else {
		r0 = ret.Get(0).(frontend.ClusterDiffResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, search2.ClusterOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommitsInWindow provides a mock function with given fields: ctx
func (_m *API) GetCommitsInWindow(ctx context.Context) ([]frontend.Commit, error) {
	ret := _m.Called(ctx)

	var r0 []frontend.Commit
	if rf, ok := ret.Get(0).(func(context.Context) []frontend.Commit); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]frontend.Commit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDigestDetails provides a mock function with given fields: ctx, grouping, digest, clID, crs
func (_m *API) GetDigestDetails(ctx context.Context, grouping paramtools.Params, digest types.Digest, clID string, crs string) (frontend.DigestDetails, error) {
	ret := _m.Called(ctx, grouping, digest, clID, crs)

	var r0 frontend.DigestDetails
	if rf, ok := ret.Get(0).(func(context.Context, paramtools.Params, types.Digest, string, string) frontend.DigestDetails); ok {
		r0 = rf(ctx, grouping, digest, clID, crs)
	} else {
		r0 = ret.Get(0).(frontend.DigestDetails)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, paramtools.Params, types.Digest, string, string) error); ok {
		r1 = rf(ctx, grouping, digest, clID, crs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDigestsDiff provides a mock function with given fields: ctx, grouping, left, right, clID, crs
func (_m *API) GetDigestsDiff(ctx context.Context, grouping paramtools.Params, left types.Digest, right types.Digest, clID string, crs string) (frontend.DigestComparison, error) {
	ret := _m.Called(ctx, grouping, left, right, clID, crs)

	var r0 frontend.DigestComparison
	if rf, ok := ret.Get(0).(func(context.Context, paramtools.Params, types.Digest, types.Digest, string, string) frontend.DigestComparison); ok {
		r0 = rf(ctx, grouping, left, right, clID, crs)
	} else {
		r0 = ret.Get(0).(frontend.DigestComparison)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, paramtools.Params, types.Digest, types.Digest, string, string) error); ok {
		r1 = rf(ctx, grouping, left, right, clID, crs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDigestsForGrouping provides a mock function with given fields: ctx, grouping
func (_m *API) GetDigestsForGrouping(ctx context.Context, grouping paramtools.Params) (frontend.DigestListResponse, error) {
	ret := _m.Called(ctx, grouping)

	var r0 frontend.DigestListResponse
	if rf, ok := ret.Get(0).(func(context.Context, paramtools.Params) frontend.DigestListResponse); ok {
		r0 = rf(ctx, grouping)
	} else {
		r0 = ret.Get(0).(frontend.DigestListResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, paramtools.Params) error); ok {
		r1 = rf(ctx, grouping)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPrimaryBranchParamset provides a mock function with given fields: ctx
func (_m *API) GetPrimaryBranchParamset(ctx context.Context) (paramtools.ReadOnlyParamSet, error) {
	ret := _m.Called(ctx)

	var r0 paramtools.ReadOnlyParamSet
	if rf, ok := ret.Get(0).(func(context.Context) paramtools.ReadOnlyParamSet); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(paramtools.ReadOnlyParamSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAndUntriagedSummaryForCL provides a mock function with given fields: ctx, qCLID
func (_m *API) NewAndUntriagedSummaryForCL(ctx context.Context, qCLID string) (search2.NewAndUntriagedSummary, error) {
	ret := _m.Called(ctx, qCLID)

	var r0 search2.NewAndUntriagedSummary
	if rf, ok := ret.Get(0).(func(context.Context, string) search2.NewAndUntriagedSummary); ok {
		r0 = rf(ctx, qCLID)
	} else {
		r0 = ret.Get(0).(search2.NewAndUntriagedSummary)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, qCLID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: _a0, _a1
func (_m *API) Search(_a0 context.Context, _a1 *query.Search) (*frontend.SearchResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *frontend.SearchResponse
	if rf, ok := ret.Get(0).(func(context.Context, *query.Search) *frontend.SearchResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*frontend.SearchResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *query.Search) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
