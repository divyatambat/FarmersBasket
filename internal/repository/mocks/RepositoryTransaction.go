// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	repository "github.com/divyatambat/FarmersBasket/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryTransaction is an autogenerated mock type for the RepositoryTransaction type
type RepositoryTransaction struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx
func (_m *RepositoryTransaction) BeginTx(ctx context.Context) (repository.Transaction, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BeginTx")
	}

	var r0 repository.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (repository.Transaction, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) repository.Transaction); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleTransaction provides a mock function with given fields: ctx, tx, incomingErr
func (_m *RepositoryTransaction) HandleTransaction(ctx context.Context, tx repository.Transaction, incomingErr error) error {
	ret := _m.Called(ctx, tx, incomingErr)

	if len(ret) == 0 {
		panic("no return value specified for HandleTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction, error) error); ok {
		r0 = rf(ctx, tx, incomingErr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepositoryTransaction creates a new instance of RepositoryTransaction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoryTransaction(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryTransaction {
	mock := &RepositoryTransaction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
