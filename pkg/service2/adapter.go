package service2

import (
	"context"
	"example/pkg/circuitbreaker"
	contracts "example/pkg/externalcontract"
)

type Service2 struct {
	breaker circuitbreaker.BreakerIface
}

func New(breaker circuitbreaker.BreakerIface) contracts.ExternalContract {
	return &Service2{breaker}
}

func (Service2) GetServiceName(ctx context.Context) contracts.GetServiceNameOutput {
	return contracts.GetServiceNameOutput{}
}

func (Service2) GetExternalServiceTrackingCode(
	ctx context.Context, identifier contracts.GetExternalServiceTrackingCodeInput) contracts.GetExternalServiceTrackingCodeOutput {

	return contracts.GetExternalServiceTrackingCodeOutput{}
}

func (Service2) AvailableService(
	ctx context.Context, availableServiceInput contracts.AvailableServiceInput) contracts.AvailableServiceOutput {

	return contracts.AvailableServiceOutput{}
}

func (Service2) GetShippingStatuses(
	ctx context.Context, getShippingStatusesInput contracts.GetShippingStatusesInput) (contracts.GetShippingStatusesOutput, error) {

	return contracts.GetShippingStatusesOutput{}, nil
}

func (Service2) CreateTrackingLabel(
	ctx context.Context, createTrackingLabelInput contracts.CreateTrackingLabelInput) (contracts.CreateTrackingLabelOutput, error) {

	return contracts.CreateTrackingLabelOutput{}, nil
}

func (Service2) Cancel(ctx context.Context, cancelOrderInput contracts.CancelOrderInput) error {
	return nil
}

func (Service2) GetTrackingLabel(
	ctx context.Context, getTrackingLabelInput contracts.GetTrackingLabelInput) (contracts.GetTrackingLabelResponse, error) {

	return contracts.GetTrackingLabelResponse{}, nil
}
