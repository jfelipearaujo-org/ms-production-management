// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	order_entity "github.com/jfelipearaujo-org/ms-production-management/internal/entity/order_entity"
	mock "github.com/stretchr/testify/mock"
)

// MockGetOrderProductionByIdService is an autogenerated mock type for the GetOrderProductionByIdService type
type MockGetOrderProductionByIdService[T interface{}] struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, request
func (_m *MockGetOrderProductionByIdService[T]) Handle(ctx context.Context, request T) (order_entity.Order, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 order_entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, T) (order_entity.Order, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, T) order_entity.Order); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(order_entity.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, T) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockGetOrderProductionByIdService creates a new instance of MockGetOrderProductionByIdService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetOrderProductionByIdService[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetOrderProductionByIdService[T] {
	mock := &MockGetOrderProductionByIdService[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
