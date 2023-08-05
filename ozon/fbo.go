package ozon

import (
	"context"
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type FBO struct {
	client *core.Client
}

type GetFBOShipmentsListParams struct {
	// Sorting direction
	Direction string `json:"dir"`

	// Shipment search filter
	Filter GetFBOShipmentsListFilter `json:"filter"`

	// Number of values in the response. Maximum is 1000, minimum is 1
	Limit int64 `json:"limit"`

	// Number of elements that will be skipped in the response. For example, if offset=10, the response will start with the 11th element found
	Offset int64 `json:"offset"`

	// true if the address transliteration from Cyrillic to Latin is enabled
	Translit bool `json:"translit"`

	// Additional fields to add to the response
	With GetFBOShipmentsListWith `json:"with"`
}

// Shipment search filter
type GetFBOShipmentsListFilter struct {
	// Period start in YYYY-MM-DD format
	Since time.Time `json:"since"`

	// Shipment status
	Status string `json:"status"`

	// Period end in YYYY-MM-DD format
	To time.Time `json:"to"`
}

// Additional fields to add to the response
type GetFBOShipmentsListWith struct {
	// Specify true to add analytics data to the response
	AnalyticsData bool `json:"analytics_data"`

	// Specify true to add financial data to the response
	FinancialData bool `json:"financial_data"`
}

type GetFBOShipmentsListResponse struct {
	core.CommonResponse

	// Shipments list
	Result []GetFBOShipmentsListResult `json:"result"`
}

type GetFBOShipmentsListResult struct {
	// Additional data for shipment list
	AdditionalData []GetFBOShipmentsListResultAdditionalData `json:"additional_data"`

	// Analytical data
	AnalyticsData GetFBOShipmentsListResultAnalyticsData `json:"analytics_data"`

	// Shipment cancellation reason identifier
	CancelReasonId int64 `json:"cancel_reason_id"`

	// Date and time of shipment creation
	CreatedAt time.Time `json:"created_at"`

	// Financial data
	FinancialData FBOFinancialData `json:"financial_data"`

	// Date and time of shipment processing start
	InProccessAt time.Time `json:"in_process_at"`

	// Identifier of the order to which the shipment belongs
	OrderId int64 `json:"order_id"`

	// Number of the order to which the shipment belongs
	OrderNumber string `json:"order_number"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Number of products in the shipment
	Products []FBOPostingProduct `json:"products"`

	// Shipment status
	Status string `json:"status"`
}

type GetFBOShipmentsListResultAdditionalData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetFBOShipmentsListResultAnalyticsData struct {
	// Delivery city
	City string `json:"city"`

	// Delivery method
	DeliveryType string `json:"delivery_type"`

	// Indication that the recipient is a legal person
	//   * true — a legal person,
	//   * false — a natural person.
	IsLegal bool `json:"is_legal"`

	// Premium subscription
	IsPremium bool `json:"is_premium"`

	// Payment method
	PaymentTypeGroupName string `json:"payment_type_group_name"`

	// Delivery region
	Region string `json:"region"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`

	// Name of the warehouse from which the order is shipped
	WarehouseName string `json:"warehouse_name"`
}

type FBOPostingProduct struct {
	// Activation codes for services and digital products
	DigitalCodes []string `json:"digital_codes"`

	// Currency of your prices. It matches the currency set in the personal account settings
	CurrencyCode string `json:"currency_code"`

	// Product name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product price
	Price string `json:"price"`

	// Quantity of products in the shipment
	Quantity int64 `json:"quantity"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

type FBOFinancialData struct {
	// Identifier of the cluster, where the shipment is sent from
	ClusterFrom string `json:"cluster_from"`

	// Identifier of the cluster, where the shipment is delivered to
	ClusterTo string `json:"cluster_to"`

	// Services
	PostingServices MarketplaceServices `json:"posting_services"`

	// Products list
	Products []FinancialDataProduct `json:"products"`
}

// Returns a list of shipments for a specified period of time. You can additionally filter the shipments by their status
func (c FBO) GetShipmentsList(ctx context.Context, params *GetFBOShipmentsListParams) (*GetFBOShipmentsListResponse, error) {
	url := "/v2/posting/fbo/list"

	resp := &GetFBOShipmentsListResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetShipmentDetailsParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// true if the address transliteration from Cyrillic to Latin is enabled
	Translit bool `json:"translit"`

	// Additional fields to add to the response
	With GetShipmentDetailsWith `json:"with"`
}

type GetShipmentDetailsWith struct {
	// Specify true to add analytics data to the response
	AnalyticsData bool `json:"analytics_data"`

	// Specify true to add financial data to the response
	FinancialData bool `json:"financial_data"`
}

type GetShipmentDetailsResponse struct {
	core.CommonResponse

	// Method result
	Result GetShipmentDetailsResult `json:"result"`
}

type GetShipmentDetailsResult struct {
	// Additional data
	AdditionalData []GetShipmentDetailsResultAdditionalData `json:"additional_data"`

	// Analytical data
	AnalyticsData GetShipmentDetailsResultAnalyticsData `json:"analytics_data"`

	// Shipment cancellation reason identifier
	CancelReasonId int64 `json:"cancel_reason_id"`

	// Date and time of shipment creation
	CreatedAt time.Time `json:"created_at"`

	// Financial data
	FinancialData FBOFinancialData `json:"financial_data"`

	// Date and time of shipment processing start
	InProcessAt time.Time `json:"in_process_at"`

	// Identifier of the order to which the shipment belongs
	OrderId int64 `json:"order_id"`

	// Number of the order to which the shipment belongs
	OrderNumber string `json:"order_number"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Number of products in the shipment
	Products []FBOPostingProduct `json:"products"`

	// Shipment status
	Status string `json:"status"`
}

type GetShipmentDetailsResultAdditionalData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetShipmentDetailsResultAnalyticsData struct {
	// Delivery city
	City string `json:"Delivery city"`

	// Delivery method
	DeliveryType string `json:"delivery_type"`

	// Indication that the recipient is a legal person:
	//   - true — a legal person
	//   - false — a natural person
	IsLegal bool `json:"is_legal"`

	// Premium subscription
	IsPremium bool `json:"is_premium"`

	// Payment method
	PaymentTypeGroupName string `json:"payment_type_group_name"`

	// Delivery region
	Region string `json:"region"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`

	// Name of the warehouse from which the order is shipped
	WarehouseName string `json:"warehouse_name"`
}

// Returns information about the shipment by its identifier
func (c FBO) GetShipmentDetails(ctx context.Context, params *GetShipmentDetailsParams) (*GetShipmentDetailsResponse, error) {
	url := "/v2/posting/fbo/get"

	resp := &GetShipmentDetailsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListSupplyRequestsParams struct {
	// Number of the page returned in the request
	Page int32 `json:"page"`

	// Number of elements on the page
	PageSize int32 `json:"page_size"`

	// Filter on status of a supply by request
	States []SupplyRequestState `json:"states"`
}

type ListSupplyRequestsResponse struct {
	core.CommonResponse

	// Indicates that the response contains not the entire array of supply requests:
	//   - true — make a new request with a different page and page_size values to get information on the remaining requests;
	//   - false — the entire array of requests for the filter specified in the request was returned in the response
	HasNext bool `json:"has_next"`

	// Supply requests list
	SupplyOrders []SupplyRequestCommonResponse `json:"supply_orders"`

	// Total requests number
	TotalSupplyOrdersCount int32 `json:"total_supply_orders_count"`
}

type SupplyRequestCommonResponse struct {
	// Supply request creation date
	CreatedAt string `json:"created_at"`

	// Local time supply interval
	LocalTimeslot SupplyRequestCommonResponseLocalTimeslot `json:"local_timeslot"`

	// Date from which you want to bring the supply to the warehouse. Only for supplies via vDC
	PreferredSupplyDataFrom string `json:"preferred_supply_data_from"`

	// Date by which you want to bring the supply to the warehouse. Only for supplies via vDC
	PreferredSupplyDataTo string `json:"preferred_supply_data_to"`

	// Status of a supply by request
	State string `json:"state"`

	// Supply request identifier
	SupplyOrderId int64 `json:"supply_order_id"`

	// Supply request number
	SupplyOrderNumber string `json:"supply_order_number"`

	// Supply warehouse
	SupplyWarehouse SupplyRequestCommonResponseSupplyWarehouse `json:"supply_warehouse"`

	// time_left_to_prepare_supply
	TimeLeftToPrepareSupply int64 `json:"time_left_to_prepare_supply"`

	// Time in seconds left to select the supply option. Only for supplies via vDC
	TimeLeftToSelectSupplyVariant int64 `json:"time_left_to_select_supply_variant"`

	// total_items_count
	TotalItemsCount int32 `json:"total_items_count"`

	// Total number of items in the request
	TotalQuantity int32 `json:"total_quantity"`
}

type SupplyRequestCommonResponseLocalTimeslot struct {
	// Interval start
	From string `json:"from"`

	// Interval end
	To string `json:"to"`
}

type SupplyRequestCommonResponseSupplyWarehouse struct {
	// Warehouse address
	Address string `json:"address"`

	// Warehouse name
	Name string `json:"name"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

// Method for getting a list of supply requests to the Ozon warehouse.
// Requests with supply both to a specific warehouse and via a virtual
// distribution center (vDC) are taken into account
func (c FBO) ListSupplyRequests(ctx context.Context, params *ListSupplyRequestsParams) (*ListSupplyRequestsResponse, error) {
	url := "/v1/supply-order/list"

	resp := &ListSupplyRequestsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetSupplyRequestInfoParams struct {
	// Supply request identifier
	SupplyOrderId int64 `json:"supply_order_id"`
}

type GetSupplyRequestInfoResponse struct {
	core.CommonResponse

	SupplyRequestCommonResponse

	// Driver and car information
	VehicleInfo GetSupplyRequestInfoVehicle `json:"vehicle_info"`
}

type GetSupplyRequestInfoVehicle struct {
	// Driver name
	DriverName string `json:"driver_name"`

	// Driver phone number
	DriverPhone string `json:"driver_phone"`

	// Car model
	VehicleModel string `json:"vehicle_model"`

	// Car number
	VehicleNumber string `json:"vehicle_number"`
}

// Method for getting detailed information on a supply request.
// Requests with supply both to a specific warehouse and via a
// virtual distribution center (vDC) are taken into account
func (c FBO) GetSupplyRequestInfo(ctx context.Context, params *GetSupplyRequestInfoParams) (*GetSupplyRequestInfoResponse, error) {
	url := "/v1/supply-order/get"

	resp := &GetSupplyRequestInfoResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListProductsInSupplyRequestParams struct {
	// Number of the page returned in the query
	Page int32 `json:"page"`

	// Number of elements on the page
	PageSize int32 `json:"page_size"`

	// Supply request identifier
	SupplyOrderId int64 `json:"supply_order_id"`
}

type ListProductsInSupplyRequestResponse struct {
	core.CommonResponse

	// Indicates that the response contains not the entire array of supply requests:
	//   - true — make a new request with a different page and page_size values to get the remaining products;
	//   - false — the entire array of product was returned in the response
	HasNext bool `json:"has_next"`

	// Products list
	Items []ListProductsInSupplyRequestItem `json:"items"`

	// Total number of products in the request
	TotalItemsCount int32 `json:"total_items_count"`
}

type ListProductsInSupplyRequestItem struct {
	// Link to product image
	IconPath string `json:"icon_path"`

	// Product name
	Name string `json:"name"`

	// Product ID
	OfferId string `json:"offer_id"`

	// Product quantity
	Quantity int64 `json:"quantity"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

// List of products in the sullpy request
func (c FBO) ListProductsInSupplyRequest(ctx context.Context, params *ListProductsInSupplyRequestParams) (*ListProductsInSupplyRequestResponse, error) {
	url := "/v1/supply-order/items"

	resp := &ListProductsInSupplyRequestResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetWarehouseWorkloadResponse struct {
	core.CommonResponse

	// Method result
	Result []GetWarehouseWorkloadResult `json:"result"`
}

type GetWarehouseWorkloadResult struct {
	// Workload
	Schedule GetWarehouseWorkloadResultSchedule `json:"schedule"`

	// Warehouse
	Warehouse GetWarehouseWorkloadResultWarehouse `json:"warehouse"`
}

type GetWarehouseWorkloadResultSchedule struct {
	// Data on the products quantity supplied to the warehouse
	Capacity []GetWarehouseWorkloadResultScheduleCapacity `json:"capacity"`

	// The closest available date for supply, local time
	Date time.Time `json:"date"`
}

type GetWarehouseWorkloadResultScheduleCapacity struct {
	// Period start, local time
	Start time.Time `json:"start"`

	// Period end, local time
	End time.Time `json:"end"`

	// Average number of products that the warehouse can accept per day for the period
	Value int32 `json:"value"`
}

type GetWarehouseWorkloadResultWarehouse struct {
	// Warehouse identifier
	Id string `json:"id"`

	// Warehouse name
	Name string `json:"name"`
}

// Method returns a list of active Ozon warehouses with information about their average workload in the nearest future
func (c FBO) GetWarehouseWorkload(ctx context.Context) (*GetWarehouseWorkloadResponse, error) {
	url := "/v1/supplier/available_warehouses"

	resp := &GetWarehouseWorkloadResponse{}

	response, err := c.client.Request(ctx, http.MethodGet, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
