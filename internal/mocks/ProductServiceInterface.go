// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import billingpb "github.com/paysuper/paysuper-proto/go/billingpb"
import context "context"
import mock "github.com/stretchr/testify/mock"

// ProductServiceInterface is an autogenerated mock type for the ProductServiceInterface type
type ProductServiceInterface struct {
	mock.Mock
}

// CountByProjectSku provides a mock function with given fields: _a0, _a1, _a2
func (_m *ProductServiceInterface) CountByProjectSku(_a0 context.Context, _a1 string, _a2 string) (int64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: _a0, _a1
func (_m *ProductServiceInterface) GetById(_a0 context.Context, _a1 string) (*billingpb.Product, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *billingpb.Product
	if rf, ok := ret.Get(0).(func(context.Context, string) *billingpb.Product); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.Product)
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

// List provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7
func (_m *ProductServiceInterface) List(_a0 context.Context, _a1 string, _a2 string, _a3 string, _a4 string, _a5 int64, _a6 int64, _a7 int32) (int64, []*billingpb.Product) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, int64, int64, int32) int64); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 []*billingpb.Product
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, int64, int64, int32) []*billingpb.Product); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*billingpb.Product)
		}
	}

	return r0, r1
}

// Upsert provides a mock function with given fields: ctx, product
func (_m *ProductServiceInterface) Upsert(ctx context.Context, product *billingpb.Product) error {
	ret := _m.Called(ctx, product)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.Product) error); ok {
		r0 = rf(ctx, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
