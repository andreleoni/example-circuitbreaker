package service1

import (
	"context"
	"example/pkg/circuitbreaker"
	contracts "example/pkg/externalcontract"
)

type Service1 struct {
	breaker circuitbreaker.BreakerIface
}

func New(breaker circuitbreaker.BreakerIface) contracts.ExternalContract {
	return &Service1{breaker}
}

func (Service1) GetServiceName(ctx context.Context) contracts.GetServiceNameOutput {
	return contracts.GetServiceNameOutput{}
}

func (Service1) GetExternalServiceTrackingCode(
	ctx context.Context, identifier contracts.GetExternalServiceTrackingCodeInput) contracts.GetExternalServiceTrackingCodeOutput {

	return contracts.GetExternalServiceTrackingCodeOutput{}
}

func (Service1) AvailableService(
	ctx context.Context, availableServiceInput contracts.AvailableServiceInput) contracts.AvailableServiceOutput {

	return contracts.AvailableServiceOutput{}
}

func (Service1) GetShippingStatuses(
	ctx context.Context, getShippingStatusesInput contracts.GetShippingStatusesInput) (contracts.GetShippingStatusesOutput, error) {

	return contracts.GetShippingStatusesOutput{}, nil
}

func (Service1) CreateTrackingLabel(
	ctx context.Context, createTrackingLabelInput contracts.CreateTrackingLabelInput) (contracts.CreateTrackingLabelOutput, error) {

	return contracts.CreateTrackingLabelOutput{}, nil
}

func (Service1) Cancel(ctx context.Context, cancelOrderInput contracts.CancelOrderInput) error {
	return nil
}

func (Service1) GetTrackingLabel(
	ctx context.Context, getTrackingLabelInput contracts.GetTrackingLabelInput) (contracts.GetTrackingLabelResponse, error) {

	return contracts.GetTrackingLabelResponse{}, nil
}
