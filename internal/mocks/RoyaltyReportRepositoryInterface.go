// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import billingpb "github.com/paysuper/paysuper-proto/go/billingpb"
import context "context"
import mock "github.com/stretchr/testify/mock"
import primitive "go.mongodb.org/mongo-driver/bson/primitive"

import time "time"

// RoyaltyReportRepositoryInterface is an autogenerated mock type for the RoyaltyReportRepositoryInterface type
type RoyaltyReportRepositoryInterface struct {
	mock.Mock
}

// FindByMerchantStatusDates provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6
func (_m *RoyaltyReportRepositoryInterface) FindByMerchantStatusDates(_a0 context.Context, _a1 string, _a2 []string, _a3 int64, _a4 int64, _a5 int64, _a6 int64) ([]*billingpb.RoyaltyReport, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6)

	var r0 []*billingpb.RoyaltyReport
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, int64, int64, int64, int64) []*billingpb.RoyaltyReport); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*billingpb.RoyaltyReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string, int64, int64, int64, int64) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCountByMerchantStatusDates provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *RoyaltyReportRepositoryInterface) FindCountByMerchantStatusDates(_a0 context.Context, _a1 string, _a2 []string, _a3 int64, _a4 int64) (int64, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, int64, int64) int64); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string, int64, int64) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *RoyaltyReportRepositoryInterface) GetAll(ctx context.Context) ([]*billingpb.RoyaltyReport, error) {
	ret := _m.Called(ctx)

	var r0 []*billingpb.RoyaltyReport
	if rf, ok := ret.Get(0).(func(context.Context) []*billingpb.RoyaltyReport); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*billingpb.RoyaltyReport)
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

// GetBalanceAmount provides a mock function with given fields: ctx, merchantId, currency
func (_m *RoyaltyReportRepositoryInterface) GetBalanceAmount(ctx context.Context, merchantId string, currency string) (float64, error) {
	ret := _m.Called(ctx, merchantId, currency)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) float64); ok {
		r0 = rf(ctx, merchantId, currency)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, merchantId, currency)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAcceptedExpireWithStatus provides a mock function with given fields: _a0, _a1, _a2
func (_m *RoyaltyReportRepositoryInterface) GetByAcceptedExpireWithStatus(_a0 context.Context, _a1 time.Time, _a2 string) ([]*billingpb.RoyaltyReport, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []*billingpb.RoyaltyReport
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, string) []*billingpb.RoyaltyReport); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*billingpb.RoyaltyReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *RoyaltyReportRepositoryInterface) GetById(ctx context.Context, id string) (*billingpb.RoyaltyReport, error) {
	ret := _m.Called(ctx, id)

	var r0 *billingpb.RoyaltyReport
	if rf, ok := ret.Get(0).(func(context.Context, string) *billingpb.RoyaltyReport); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.RoyaltyReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPayoutId provides a mock function with given fields: ctx, payoutId
func (_m *RoyaltyReportRepositoryInterface) GetByPayoutId(ctx context.Context, payoutId string) ([]*billingpb.RoyaltyReport, error) {
	ret := _m.Called(ctx, payoutId)

	var r0 []*billingpb.RoyaltyReport
	if rf, ok := ret.Get(0).(func(context.Context, string) []*billingpb.RoyaltyReport); ok {
		r0 = rf(ctx, payoutId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*billingpb.RoyaltyReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, payoutId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPeriod provides a mock function with given fields: _a0, _a1, _a2
func (_m *RoyaltyReportRepositoryInterface) GetByPeriod(_a0 context.Context, _a1 time.Time, _a2 time.Time) ([]*billingpb.RoyaltyReport, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []*billingpb.RoyaltyReport
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) []*billingpb.RoyaltyReport); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*billingpb.RoyaltyReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonPayoutReports provides a mock function with given fields: ctx, merchantId, currency
func (_m *RoyaltyReportRepositoryInterface) GetNonPayoutReports(ctx context.Context, merchantId string, currency string) ([]*billingpb.RoyaltyReport, error) {
	ret := _m.Called(ctx, merchantId, currency)

	var r0 []*billingpb.RoyaltyReport
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []*billingpb.RoyaltyReport); ok {
		r0 = rf(ctx, merchantId, currency)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*billingpb.RoyaltyReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, merchantId, currency)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportExists provides a mock function with given fields: ctx, merchantId, currency, from, to
func (_m *RoyaltyReportRepositoryInterface) GetReportExists(ctx context.Context, merchantId string, currency string, from time.Time, to time.Time) *billingpb.RoyaltyReport {
	ret := _m.Called(ctx, merchantId, currency, from, to)

	var r0 *billingpb.RoyaltyReport
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Time, time.Time) *billingpb.RoyaltyReport); ok {
		r0 = rf(ctx, merchantId, currency, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.RoyaltyReport)
		}
	}

	return r0
}

// GetRoyaltyHistoryById provides a mock function with given fields: ctx, id
func (_m *RoyaltyReportRepositoryInterface) GetRoyaltyHistoryById(ctx context.Context, id string) ([]*billingpb.RoyaltyReportChanges, error) {
	ret := _m.Called(ctx, id)

	var r0 []*billingpb.RoyaltyReportChanges
	if rf, ok := ret.Get(0).(func(context.Context, string) []*billingpb.RoyaltyReportChanges); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*billingpb.RoyaltyReportChanges)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, document, ip, source
func (_m *RoyaltyReportRepositoryInterface) Insert(ctx context.Context, document *billingpb.RoyaltyReport, ip string, source string) error {
	ret := _m.Called(ctx, document, ip, source)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.RoyaltyReport, string, string) error); ok {
		r0 = rf(ctx, document, ip, source)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, document, ip, source
func (_m *RoyaltyReportRepositoryInterface) Update(ctx context.Context, document *billingpb.RoyaltyReport, ip string, source string) error {
	ret := _m.Called(ctx, document, ip, source)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.RoyaltyReport, string, string) error); ok {
		r0 = rf(ctx, document, ip, source)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMany provides a mock function with given fields: ctx, query, set
func (_m *RoyaltyReportRepositoryInterface) UpdateMany(ctx context.Context, query primitive.M, set primitive.M) error {
	ret := _m.Called(ctx, query, set)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, primitive.M) error); ok {
		r0 = rf(ctx, query, set)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
