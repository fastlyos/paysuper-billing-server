// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: grpc/grpc.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	grpc/grpc.proto

It has these top-level messages:
	EmptyRequest
	EmptyResponse
	PaymentCreateRequest
	PaymentCreateResponse
	PaymentFormJsonDataRequest
	PaymentFormJsonDataProject
	PaymentFormJsonDataResponse
	PaymentNotifyRequest
	PaymentNotifyResponse
	ConvertRateRequest
	ConvertRateResponse
	OnboardingBanking
	OnboardingRequest
	FindByIdRequest
	MerchantListingRequest
	Merchants
	MerchantChangeStatusRequest
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import billing "github.com/paysuper/paysuper-billing-server/pkg/proto/billing"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = billing.Merchant{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for BillingService service

type BillingService interface {
	OrderCreateProcess(ctx context.Context, in *billing.OrderCreateRequest, opts ...client.CallOption) (*billing.Order, error)
	PaymentFormJsonDataProcess(ctx context.Context, in *PaymentFormJsonDataRequest, opts ...client.CallOption) (*PaymentFormJsonDataResponse, error)
	PaymentCreateProcess(ctx context.Context, in *PaymentCreateRequest, opts ...client.CallOption) (*PaymentCreateResponse, error)
	PaymentCallbackProcess(ctx context.Context, in *PaymentNotifyRequest, opts ...client.CallOption) (*PaymentNotifyResponse, error)
	RebuildCache(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EmptyResponse, error)
	UpdateOrder(ctx context.Context, in *billing.Order, opts ...client.CallOption) (*EmptyResponse, error)
	UpdateMerchant(ctx context.Context, in *billing.Merchant, opts ...client.CallOption) (*EmptyResponse, error)
	GetConvertRate(ctx context.Context, in *ConvertRateRequest, opts ...client.CallOption) (*ConvertRateResponse, error)
	GetMerchantById(ctx context.Context, in *FindByIdRequest, opts ...client.CallOption) (*billing.Merchant, error)
	GetMerchantByExternalId(ctx context.Context, in *FindByIdRequest, opts ...client.CallOption) (*billing.Merchant, error)
	ListMerchants(ctx context.Context, in *MerchantListingRequest, opts ...client.CallOption) (*Merchants, error)
	ChangeMerchant(ctx context.Context, in *OnboardingRequest, opts ...client.CallOption) (*billing.Merchant, error)
	ChangeMerchantStatus(ctx context.Context, in *MerchantChangeStatusRequest, opts ...client.CallOption) (*billing.Merchant, error)
}

type billingService struct {
	c    client.Client
	name string
}

func NewBillingService(name string, c client.Client) BillingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "grpc"
	}
	return &billingService{
		c:    c,
		name: name,
	}
}

func (c *billingService) OrderCreateProcess(ctx context.Context, in *billing.OrderCreateRequest, opts ...client.CallOption) (*billing.Order, error) {
	req := c.c.NewRequest(c.name, "BillingService.OrderCreateProcess", in)
	out := new(billing.Order)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) PaymentFormJsonDataProcess(ctx context.Context, in *PaymentFormJsonDataRequest, opts ...client.CallOption) (*PaymentFormJsonDataResponse, error) {
	req := c.c.NewRequest(c.name, "BillingService.PaymentFormJsonDataProcess", in)
	out := new(PaymentFormJsonDataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) PaymentCreateProcess(ctx context.Context, in *PaymentCreateRequest, opts ...client.CallOption) (*PaymentCreateResponse, error) {
	req := c.c.NewRequest(c.name, "BillingService.PaymentCreateProcess", in)
	out := new(PaymentCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) PaymentCallbackProcess(ctx context.Context, in *PaymentNotifyRequest, opts ...client.CallOption) (*PaymentNotifyResponse, error) {
	req := c.c.NewRequest(c.name, "BillingService.PaymentCallbackProcess", in)
	out := new(PaymentNotifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) RebuildCache(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "BillingService.RebuildCache", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) UpdateOrder(ctx context.Context, in *billing.Order, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "BillingService.UpdateOrder", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) UpdateMerchant(ctx context.Context, in *billing.Merchant, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "BillingService.UpdateMerchant", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) GetConvertRate(ctx context.Context, in *ConvertRateRequest, opts ...client.CallOption) (*ConvertRateResponse, error) {
	req := c.c.NewRequest(c.name, "BillingService.GetConvertRate", in)
	out := new(ConvertRateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) GetMerchantById(ctx context.Context, in *FindByIdRequest, opts ...client.CallOption) (*billing.Merchant, error) {
	req := c.c.NewRequest(c.name, "BillingService.GetMerchantById", in)
	out := new(billing.Merchant)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) GetMerchantByExternalId(ctx context.Context, in *FindByIdRequest, opts ...client.CallOption) (*billing.Merchant, error) {
	req := c.c.NewRequest(c.name, "BillingService.GetMerchantByExternalId", in)
	out := new(billing.Merchant)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) ListMerchants(ctx context.Context, in *MerchantListingRequest, opts ...client.CallOption) (*Merchants, error) {
	req := c.c.NewRequest(c.name, "BillingService.ListMerchants", in)
	out := new(Merchants)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) ChangeMerchant(ctx context.Context, in *OnboardingRequest, opts ...client.CallOption) (*billing.Merchant, error) {
	req := c.c.NewRequest(c.name, "BillingService.ChangeMerchant", in)
	out := new(billing.Merchant)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) ChangeMerchantStatus(ctx context.Context, in *MerchantChangeStatusRequest, opts ...client.CallOption) (*billing.Merchant, error) {
	req := c.c.NewRequest(c.name, "BillingService.ChangeMerchantStatus", in)
	out := new(billing.Merchant)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BillingService service

type BillingServiceHandler interface {
	OrderCreateProcess(context.Context, *billing.OrderCreateRequest, *billing.Order) error
	PaymentFormJsonDataProcess(context.Context, *PaymentFormJsonDataRequest, *PaymentFormJsonDataResponse) error
	PaymentCreateProcess(context.Context, *PaymentCreateRequest, *PaymentCreateResponse) error
	PaymentCallbackProcess(context.Context, *PaymentNotifyRequest, *PaymentNotifyResponse) error
	RebuildCache(context.Context, *EmptyRequest, *EmptyResponse) error
	UpdateOrder(context.Context, *billing.Order, *EmptyResponse) error
	UpdateMerchant(context.Context, *billing.Merchant, *EmptyResponse) error
	GetConvertRate(context.Context, *ConvertRateRequest, *ConvertRateResponse) error
	GetMerchantById(context.Context, *FindByIdRequest, *billing.Merchant) error
	GetMerchantByExternalId(context.Context, *FindByIdRequest, *billing.Merchant) error
	ListMerchants(context.Context, *MerchantListingRequest, *Merchants) error
	ChangeMerchant(context.Context, *OnboardingRequest, *billing.Merchant) error
	ChangeMerchantStatus(context.Context, *MerchantChangeStatusRequest, *billing.Merchant) error
}

func RegisterBillingServiceHandler(s server.Server, hdlr BillingServiceHandler, opts ...server.HandlerOption) error {
	type billingService interface {
		OrderCreateProcess(ctx context.Context, in *billing.OrderCreateRequest, out *billing.Order) error
		PaymentFormJsonDataProcess(ctx context.Context, in *PaymentFormJsonDataRequest, out *PaymentFormJsonDataResponse) error
		PaymentCreateProcess(ctx context.Context, in *PaymentCreateRequest, out *PaymentCreateResponse) error
		PaymentCallbackProcess(ctx context.Context, in *PaymentNotifyRequest, out *PaymentNotifyResponse) error
		RebuildCache(ctx context.Context, in *EmptyRequest, out *EmptyResponse) error
		UpdateOrder(ctx context.Context, in *billing.Order, out *EmptyResponse) error
		UpdateMerchant(ctx context.Context, in *billing.Merchant, out *EmptyResponse) error
		GetConvertRate(ctx context.Context, in *ConvertRateRequest, out *ConvertRateResponse) error
		GetMerchantById(ctx context.Context, in *FindByIdRequest, out *billing.Merchant) error
		GetMerchantByExternalId(ctx context.Context, in *FindByIdRequest, out *billing.Merchant) error
		ListMerchants(ctx context.Context, in *MerchantListingRequest, out *Merchants) error
		ChangeMerchant(ctx context.Context, in *OnboardingRequest, out *billing.Merchant) error
		ChangeMerchantStatus(ctx context.Context, in *MerchantChangeStatusRequest, out *billing.Merchant) error
	}
	type BillingService struct {
		billingService
	}
	h := &billingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BillingService{h}, opts...))
}

type billingServiceHandler struct {
	BillingServiceHandler
}

func (h *billingServiceHandler) OrderCreateProcess(ctx context.Context, in *billing.OrderCreateRequest, out *billing.Order) error {
	return h.BillingServiceHandler.OrderCreateProcess(ctx, in, out)
}

func (h *billingServiceHandler) PaymentFormJsonDataProcess(ctx context.Context, in *PaymentFormJsonDataRequest, out *PaymentFormJsonDataResponse) error {
	return h.BillingServiceHandler.PaymentFormJsonDataProcess(ctx, in, out)
}

func (h *billingServiceHandler) PaymentCreateProcess(ctx context.Context, in *PaymentCreateRequest, out *PaymentCreateResponse) error {
	return h.BillingServiceHandler.PaymentCreateProcess(ctx, in, out)
}

func (h *billingServiceHandler) PaymentCallbackProcess(ctx context.Context, in *PaymentNotifyRequest, out *PaymentNotifyResponse) error {
	return h.BillingServiceHandler.PaymentCallbackProcess(ctx, in, out)
}

func (h *billingServiceHandler) RebuildCache(ctx context.Context, in *EmptyRequest, out *EmptyResponse) error {
	return h.BillingServiceHandler.RebuildCache(ctx, in, out)
}

func (h *billingServiceHandler) UpdateOrder(ctx context.Context, in *billing.Order, out *EmptyResponse) error {
	return h.BillingServiceHandler.UpdateOrder(ctx, in, out)
}

func (h *billingServiceHandler) UpdateMerchant(ctx context.Context, in *billing.Merchant, out *EmptyResponse) error {
	return h.BillingServiceHandler.UpdateMerchant(ctx, in, out)
}

func (h *billingServiceHandler) GetConvertRate(ctx context.Context, in *ConvertRateRequest, out *ConvertRateResponse) error {
	return h.BillingServiceHandler.GetConvertRate(ctx, in, out)
}

func (h *billingServiceHandler) GetMerchantById(ctx context.Context, in *FindByIdRequest, out *billing.Merchant) error {
	return h.BillingServiceHandler.GetMerchantById(ctx, in, out)
}

func (h *billingServiceHandler) GetMerchantByExternalId(ctx context.Context, in *FindByIdRequest, out *billing.Merchant) error {
	return h.BillingServiceHandler.GetMerchantByExternalId(ctx, in, out)
}

func (h *billingServiceHandler) ListMerchants(ctx context.Context, in *MerchantListingRequest, out *Merchants) error {
	return h.BillingServiceHandler.ListMerchants(ctx, in, out)
}

func (h *billingServiceHandler) ChangeMerchant(ctx context.Context, in *OnboardingRequest, out *billing.Merchant) error {
	return h.BillingServiceHandler.ChangeMerchant(ctx, in, out)
}

func (h *billingServiceHandler) ChangeMerchantStatus(ctx context.Context, in *MerchantChangeStatusRequest, out *billing.Merchant) error {
	return h.BillingServiceHandler.ChangeMerchantStatus(ctx, in, out)
}
