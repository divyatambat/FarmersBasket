// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/divyatambat/FarmersBasket/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetProductByID provides a mock function with given fields: ctx, productID
func (_m *Service) GetProductByID(ctx context.Context, productID int64) (dto.Product, error) {
	ret := _m.Called(ctx, productID)

	if len(ret) == 0 {
		panic("no return value specified for GetProductByID")
	}

	var r0 dto.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (dto.Product, error)); ok {
		return rf(ctx, productID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) dto.Product); ok {
		r0 = rf(ctx, productID)
	} else {
		r0 = ret.Get(0).(dto.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProducts provides a mock function with given fields: ctx
func (_m *Service) ListProducts(ctx context.Context) ([]dto.Product, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListProducts")
	}

	var r0 []dto.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]dto.Product, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []dto.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProductQuantity provides a mock function with given fields: ctx, productsQuantityMap
func (_m *Service) UpdateProductQuantity(ctx context.Context, productsQuantityMap map[int64]int64) error {
	ret := _m.Called(ctx, productsQuantityMap)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProductQuantity")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[int64]int64) error); ok {
		r0 = rf(ctx, productsQuantityMap)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
