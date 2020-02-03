// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import billingpb "github.com/paysuper/paysuper-proto/go/billingpb"
import context "context"
import mock "github.com/stretchr/testify/mock"
import pkg "github.com/paysuper/paysuper-billing-server/internal/pkg"

// MoneyBackCostMerchantRepositoryInterface is an autogenerated mock type for the MoneyBackCostMerchantRepositoryInterface type
type MoneyBackCostMerchantRepositoryInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *MoneyBackCostMerchantRepositoryInterface) Delete(_a0 context.Context, _a1 *billingpb.MoneyBackCostMerchant) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.MoneyBackCostMerchant) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8
func (_m *MoneyBackCostMerchantRepositoryInterface) Find(_a0 context.Context, _a1 string, _a2 string, _a3 string, _a4 string, _a5 string, _a6 string, _a7 string, _a8 int32) ([]*pkg.MoneyBackCostMerchantSet, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8)

	var r0 []*pkg.MoneyBackCostMerchantSet
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string, string, string, int32) []*pkg.MoneyBackCostMerchantSet); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pkg.MoneyBackCostMerchantSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, string, string, string, int32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllForMerchant provides a mock function with given fields: _a0, _a1
func (_m *MoneyBackCostMerchantRepositoryInterface) GetAllForMerchant(_a0 context.Context, _a1 string) (*billingpb.MoneyBackCostMerchantList, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *billingpb.MoneyBackCostMerchantList
	if rf, ok := ret.Get(0).(func(context.Context, string) *billingpb.MoneyBackCostMerchantList); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.MoneyBackCostMerchantList)
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

// GetById provides a mock function with given fields: _a0, _a1
func (_m *MoneyBackCostMerchantRepositoryInterface) GetById(_a0 context.Context, _a1 string) (*billingpb.MoneyBackCostMerchant, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *billingpb.MoneyBackCostMerchant
	if rf, ok := ret.Get(0).(func(context.Context, string) *billingpb.MoneyBackCostMerchant); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.MoneyBackCostMerchant)
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

// Insert provides a mock function with given fields: _a0, _a1
func (_m *MoneyBackCostMerchantRepositoryInterface) Insert(_a0 context.Context, _a1 *billingpb.MoneyBackCostMerchant) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.MoneyBackCostMerchant) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MultipleInsert provides a mock function with given fields: _a0, _a1
func (_m *MoneyBackCostMerchantRepositoryInterface) MultipleInsert(_a0 context.Context, _a1 []*billingpb.MoneyBackCostMerchant) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*billingpb.MoneyBackCostMerchant) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *MoneyBackCostMerchantRepositoryInterface) Update(_a0 context.Context, _a1 *billingpb.MoneyBackCostMerchant) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.MoneyBackCostMerchant) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
