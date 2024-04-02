package ozon

import (
	"context"
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Passes struct {
	client *core.Client
}

type ListPassesParams struct {
	// Cursor for the next data sample
	Cursor string `json:"curson"`

	// Filters
	Filter ListPassesFilter `json:"filter"`

	// Limit on number of entries in a reply. Default value is 1000. Maximum value is 1000
	Limit int32 `json:"limit"`
}

type ListPassesFilter struct {
	// Filter by pass identifier
	ArrivalPassIds []string `json:"arrival_pass_ids"`

	// Filter by purpose of arrival:
	//
	// FBS_DELIVERY—delivery.
	// FBS_RETURN—take out returns.
	// If the parameter isn't specified, both purposes are considered.
	//
	// The specified purpose must be in the list of reasons for passes
	ArrivalReason string `json:"arrival_reason"`

	// Filter by drop-off points identifier
	DropoffPointIds []int64 `json:"dropoff_point_ids"`

	// true to get only active pass requests
	OnlyActivePasses bool `json:"only_active_passes"`

	// Filter by seller's warehouses identifier
	WarehouseIds []int64 `json:"warehouse_ids"`
}

type ListPassesResponse struct {
	core.CommonResponse

	// List of passes
	ArrivalPasses []ListPassesArrivalPass `json:"arrival_passes"`

	// Cursor for the next data sample. If the parameter is empty, there is no more data
	Cursor string `json:"cursor"`
}

type ListPassesArrivalPass struct {
	// Pass identifier
	ArrivalPassId int64 `json:"arrival_pass_id"`

	// Arrival purpose
	ArrivalReasons []string `json:"arrival_reasons"`

	// Date and time of arrival in UTC format
	ArrivalTime time.Time `json:"arrival_time"`

	// Driver full name
	DriverName string `json:"driver_name"`

	// Driver phone number
	DriverPhone string `json:"driver_phone"`

	// Drop-off point identifier
	DropoffPointId int64 `json:"dropoff_point_id"`

	// true if the request is active
	IsActive bool `json:"is_active"`

	// Car license plate
	VehicleLicensePlate string `json:"vehicle_license_plate"`

	// Car model
	VehicleModel string `json:"vehicle_model"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

func (c Passes) List(ctx context.Context, params *ListPassesParams) (*ListPassesResponse, error) {
	url := "/v1/pass/list"

	resp := &ListPassesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CreateCarriageParams struct {
	// List of passes
	ArrivalPasses []CarriageArrivalPass `json:"arrival_passes"`

	// Freight identifier
	CarriageId int64 `json:"carriage_id"`
}

type CarriageArrivalPass struct {
	// Driver full name
	DriverName string `json:"driver_name"`

	// Driver phone number
	DriverPhone string `json:"driver_phone"`

	// Car license plate
	VehicleLicensePlate string `json:"vehicle_license_plate"`

	// Car model
	VehicleModel string `json:"vehicle_model"`

	// true if you will export returns. Default is false
	WithReturns bool `json:"with_returns"`
}

type CreateCarriageResponse struct {
	core.CommonResponse

	// Pass identifiers
	ArrivalPassIds []string `json:"arrival_pass_ids"`
}

func (c Passes) CreateCarriage(ctx context.Context, params *CreateCarriageParams) (*CreateCarriageResponse, error) {
	url := "/v1/carriage/pass/create"

	resp := &CreateCarriageResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateCarriageParams struct {
	ArrivalPasses []UpdateCarriageArrivalPass `json:"arrival_passes"`

	CarriageId int64 `json:"carriage_id"`
}

type UpdateCarriageArrivalPass struct {
	CarriageArrivalPass

	Id int64 `json:"id"`
}

type UpdateCarriageResponse struct {
	core.CommonResponse
}

func (c Passes) UpdateCarriage(ctx context.Context, params *UpdateCarriageParams) (*UpdateCarriageResponse, error) {
	url := "/v1/carriage/pass/update"

	resp := &UpdateCarriageResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
