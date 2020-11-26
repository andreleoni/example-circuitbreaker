package externalcontract

import (
	"context"
	"encoding/json"
	"time"
)

// ExternalContract implements our shipment service interface
type ExternalContract interface {
	GetServiceName(ctx context.Context) GetServiceNameOutput

	GetExternalServiceTrackingCode(ctx context.Context, identifier GetExternalServiceTrackingCodeInput) GetExternalServiceTrackingCodeOutput

	AvailableService(ctx context.Context, availableServiceInput AvailableServiceInput) AvailableServiceOutput

	GetShippingStatuses(ctx context.Context, getShippingStatusesInput GetShippingStatusesInput) (GetShippingStatusesOutput, error)

	CreateTrackingLabel(ctx context.Context, createTrackingLabelInput CreateTrackingLabelInput) (CreateTrackingLabelOutput, error)

	Cancel(ctx context.Context, cancelOrderInput CancelOrderInput) error

	GetTrackingLabel(ctx context.Context, getTrackingLabelInput GetTrackingLabelInput) (GetTrackingLabelResponse, error)
}

type GetServiceNameOutput struct {
	Name string
}

type GetExternalServiceTrackingCodeInput struct {
	identifier Identifier
}

type GetExternalServiceTrackingCodeOutput struct {
	TrackingCode string
}

type AvailableServiceInput struct {
	From uint64
	To   uint64
}

type AvailableServiceOutput struct {
	Result bool
}

type GetShippingStatusesInput struct {
	Identifiers map[string]Identifier
}

type GetShippingStatusesOutput struct {
	ShippingStatuses map[string]ShippingStatus
}

type CreateTrackingLabelInput struct {
	AllShippingParams string
}

type CreateTrackingLabelOutput struct {
	LabelByte  []byte `json:"label"`
	Identifier string `json:"identifier"`
}

type CancelOrderInput struct {
	Identifier map[string]interface{}
}

type GetTrackingLabelInput struct {
	Identifier Identifier
}

type GetTrackingLabelResponse struct {
	LabelByte []byte `json:"label"`
}

// Base structs

// Identifier
type Identifier struct {
	ExternalID            string      `json:"external_id,omitempty"`
	ExternalSellerID      json.Number `json:"external_seller_id,omitempty"`
	TrackingCode          string      `json:"tracking_code,omitempty"`
	PlpID                 json.Number `json:"plp_id,omitempty"`
	TrackingURL           string      `json:"tracking_url,omitempty"`
	Provider              string      `json:"provider,omitempty"`
	EstimatedDeliveryDate string      `json:"estimated_delivery_date,omitempty"`
	Service               string      `json:"service,omitempty"`
	Code                  string      `json:"code,omitempty"`
	ShipmentID            string      `json:"shipment_id,omitempty"`
	Status                string      `json:"status,omitempty"`

	// Read only fields
	Number    string `gorm:"-" json:"number,omitempty"`
	OrderUUID string `gorm:"-" json:"order_uuid,omitempty"`
}

type ShippingStatus struct {
	Status     string
	Price      float64
	Identifier Identifier
	Events     []Event
}

type Event struct {
	Title       string
	Description string
	Status      string
	Kind        string
	OccurredAt  time.Time
}
