package ozon

import (
	"context"
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Warehouses struct {
	client *core.Client
}

type GetListOfWarehousesResponse struct {
	core.CommonResponse

	Result []GetListOfWarehousesResult `json:"result"`
}

type GetListOfWarehousesResult struct {
	// Trusted acceptance attribute. `true` if trusted acceptance is enabled in the warehouse
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`

	// Indication that the warehouse works under the rFBS scheme:
	//   - true — the warehouse works under the rFBS scheme;
	//   - false — the warehouse does not work under the rFBS scheme.
	IsRFBS bool `json:"is_rfbs"`

	// Warehouse name
	Name string `json:"name"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`

	// Possibility to print an acceptance certificate in advance. `true` if printing in advance is possible
	CanPrintActInAdvance bool `json:"can_print_act_in_advance"`

	// FBS first mile
	FirstMileType GetListOfWarehousesResultFirstMile `json:"first_mile_type"`

	// Indication if there is a limit on the minimum number of orders. `true` if there is such a limit
	HasPostingsLimit bool `json:"has_postings_limit"`

	// Indication that the warehouse is not working due to quarantine
	IsKarantin bool `json:"is_karantin"`

	// Indication that the warehouse accepts bulky products
	IsKGT bool `json:"is_kgt"`

	// true if the warehouse handles economy products
	IsEconomy bool `json:"is_economy"`

	// Indication that warehouse schedule can be changed
	IsTimetableEditable bool `json:"is_timetable_editable"`

	// Minimum limit value: the number of orders that can be brought in one shipment
	MinPostingsLimit int32 `json:"min_postings_limit"`

	// Limit value. -1 if there is no limit
	PostingsLimit int32 `json:"postings_limit"`

	// Number of warehouse working days
	MinWorkingDays int64 `json:"min_working_days"`

	// Warehouse status
	Status string `json:"status"`

	// Warehouse working days
	WorkingDays []WorkingDay `json:"working_days"`
}

type GetListOfWarehousesResultFirstMile struct {
	// DropOff point identifier
	DropoffPointId string `json:"dropoff_point_id"`

	// DropOff timeslot identifier
	DropoffTimeslotId int64 `json:"dropoff_timeslot_id"`

	// Indication that the warehouse settings are being updated
	FirstMileIsChanging bool `json:"first_mile_is_changing"`

	// First mile type:
	//
	// Enum: "DropOff" "Pickup"
	//   - DropOff
	//   - Pickup
	FirstMileType string `json:"first_mile_type"`
}

// You do not need to specify any parameters in the request. Your company will be identified by the Warehouses ID
func (c Warehouses) GetListOfWarehouses(ctx context.Context) (*GetListOfWarehousesResponse, error) {
	url := "/v1/warehouse/list"

	resp := &GetListOfWarehousesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetListOfDeliveryMethodsParams struct {
	// Search filter for delivery methods
	Filter *GetListOfDeliveryMethodsFilter `json:"filter,omitempty"`

	// Number of items in a response. Maximum is 50, minimum is 1
	Limit int64 `json:"limit"`

	// Number of elements that will be skipped in the response.
	// For example, if offset=10, the response will start with the 11th element found
	Offset int64 `json:"offset"`
}

type GetListOfDeliveryMethodsFilter struct {
	// Delivery service identifier
	ProviderId int64 `json:"provider_id"`

	// Delivery method status:
	//   - NEW—created
	//   - EDITED—being edited
	//   - ACTIVE—active
	//   - DISABLED—inactive
	Status string `json:"status"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

type GetListOfDeliveryMethodsResponse struct {
	core.CommonResponse

	// Indication that only part of delivery methods was returned in the response:
	//   - true — make a request with a new offset parameter value for getting the rest of delivery methods;
	//   - false — all delivery methods were returned
	HasNext bool `json:"has_next"`

	// Method result
	Result []GetListOfDeliveryMethodsResult `json:"result"`
}

type GetListOfDeliveryMethodsResult struct {
	// Company identifier
	CompanyId int64 `json:"company_id"`

	// Date and time of delivery method creation
	CreatedAt time.Time `json:"created_at"`

	// Time before an order must be packaged
	Cutoff string `json:"cutoff"`

	// Delivery method identifier
	Id int64 `json:"id"`

	// Delivery method name
	Name string `json:"name"`

	// Delivery service identifier
	ProviderId int64 `json:"provider_id"`

	// Minimum time to package an order in minutes according to warehouse settings
	SLACutIn int64 `json:"sla_cut_in"`

	// Delivery method status:
	//   - NEW—created,
	//   - EDITED—being edited,
	//   - ACTIVE—active,
	//   - DISABLED—inactive
	Status string `json:"status"`

	// Order delivery service identifier
	TemplateId int64 `json:"template_id"`

	// Date and time when the delivery method was last updated
	UpdatedAt time.Time `json:"updated_at"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

// This methods allows you to get list of all delivery methods that can be applied for this warehouse
func (c Warehouses) GetListOfDeliveryMethods(ctx context.Context, params *GetListOfDeliveryMethodsParams) (*GetListOfDeliveryMethodsResponse, error) {
	url := "/v1/delivery-method/list"

	resp := &GetListOfDeliveryMethodsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListForShippingParams struct {
	// Supply type
	FilterBySupplyType []string `json:"filter_by_supply_type"`

	// Search by warehouse name. To search for pick-up points, specify the full name
	Search string `json:"search"`
}

type ListForShippingResponse struct {
	core.CommonResponse

	// Warehouse search result
	Search []ListForShippingSearch `json:"search"`
}

type ListForShippingSearch struct {
	// Warehouse address
	Address string `json:"address"`

	// Warehouse coordinates
	Coordinates Coordinates `json:"coordinates"`

	// Warehouse name
	Name string `json:"name"`

	// Identifier of the warehouse, pick-up point, or sorting center
	WarehouseId int64 `json:"warehouse_id"`

	// Type of warehouse, pick-up point, or sorting center
	WarehouseType string `json:"warehouse_type"`
}

type Coordinates struct {
	// Latitude
	Latitude float64 `json:"latitude"`

	// Longitude
	Longitude float64 `json:"longitude"`
}

// Use the method to find sorting centres, pick-up points, and drop-off points available for cross-docking and direct supplies.
//
// You can view the addresses of all points on the map and in a table in the Knowledge Base.
func (c Warehouses) ListForShipping(ctx context.Context, params *ListForShippingParams) (*ListForShippingResponse, error) {
	url := "/v1/warehouse/fbo/list"

	resp := &ListForShippingResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
