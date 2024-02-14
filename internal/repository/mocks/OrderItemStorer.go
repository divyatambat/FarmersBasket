// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	repository "github.com/divyatambat/FarmersBasket/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// OrderItemStorer is an autogenerated mock type for the OrderItemStorer type
type OrderItemStorer struct {
	mock.Mock
}

// GetOrderItemsByOrderID provides a mock function with given fields: ctx, orderID
func (_m *OrderItemStorer) GetOrderItemsByOrderID(ctx context.Context, orderID int64) ([]repository.OrderItem, error) {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderItemsByOrderID")
	}

	var r0 []repository.OrderItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]repository.OrderItem, error)); ok {
		return rf(ctx, orderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []repository.OrderItem); ok {
		r0 = rf(ctx, orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.OrderItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreOrderItems provides a mock function with given fields: ctx, orderItems
func (_m *OrderItemStorer) StoreOrderItems(ctx context.Context, orderItems []repository.OrderItem) error {
	ret := _m.Called(ctx, orderItems)

	if len(ret) == 0 {
		panic("no return value specified for StoreOrderItems")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []repository.OrderItem) error); ok {
		r0 = rf(ctx, orderItems)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrderItemStorer creates a new instance of OrderItemStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderItemStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderItemStorer {
	mock := &OrderItemStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
