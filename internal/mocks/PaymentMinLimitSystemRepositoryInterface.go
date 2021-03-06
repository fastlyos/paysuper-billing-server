// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import billingpb "github.com/paysuper/paysuper-proto/go/billingpb"
import context "context"
import mock "github.com/stretchr/testify/mock"

// PaymentMinLimitSystemRepositoryInterface is an autogenerated mock type for the PaymentMinLimitSystemRepositoryInterface type
type PaymentMinLimitSystemRepositoryInterface struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: _a0
func (_m *PaymentMinLimitSystemRepositoryInterface) GetAll(_a0 context.Context) ([]*billingpb.PaymentMinLimitSystem, error) {
	ret := _m.Called(_a0)

	var r0 []*billingpb.PaymentMinLimitSystem
	if rf, ok := ret.Get(0).(func(context.Context) []*billingpb.PaymentMinLimitSystem); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*billingpb.PaymentMinLimitSystem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByCurrency provides a mock function with given fields: _a0, _a1
func (_m *PaymentMinLimitSystemRepositoryInterface) GetByCurrency(_a0 context.Context, _a1 string) (*billingpb.PaymentMinLimitSystem, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *billingpb.PaymentMinLimitSystem
	if rf, ok := ret.Get(0).(func(context.Context, string) *billingpb.PaymentMinLimitSystem); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.PaymentMinLimitSystem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MultipleInsert provides a mock function with given fields: _a0, _a1
func (_m *PaymentMinLimitSystemRepositoryInterface) MultipleInsert(_a0 context.Context, _a1 []*billingpb.PaymentMinLimitSystem) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*billingpb.PaymentMinLimitSystem) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *PaymentMinLimitSystemRepositoryInterface) Upsert(_a0 context.Context, _a1 *billingpb.PaymentMinLimitSystem) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.PaymentMinLimitSystem) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
