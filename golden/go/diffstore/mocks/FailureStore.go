// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	diff "go.skia.org/infra/golden/go/diff"

	mock "github.com/stretchr/testify/mock"

	types "go.skia.org/infra/golden/go/types"
)

// FailureStore is an autogenerated mock type for the FailureStore type
type FailureStore struct {
	mock.Mock
}

// AddDigestFailure provides a mock function with given fields: failure
func (_m *FailureStore) AddDigestFailure(failure *diff.DigestFailure) error {
	ret := _m.Called(failure)

	var r0 error
	if rf, ok := ret.Get(0).(func(*diff.DigestFailure) error); ok {
		r0 = rf(failure)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PurgeDigestFailures provides a mock function with given fields: digests
func (_m *FailureStore) PurgeDigestFailures(digests types.DigestSlice) error {
	ret := _m.Called(digests)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.DigestSlice) error); ok {
		r0 = rf(digests)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnavailableDigests provides a mock function with given fields:
func (_m *FailureStore) UnavailableDigests() (map[types.Digest]*diff.DigestFailure, error) {
	ret := _m.Called()

	var r0 map[types.Digest]*diff.DigestFailure
	if rf, ok := ret.Get(0).(func() map[types.Digest]*diff.DigestFailure); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[types.Digest]*diff.DigestFailure)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
