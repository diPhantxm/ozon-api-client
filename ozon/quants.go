package ozon

import (
	"context"
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type Quants struct {
	client *core.Client
}

type ListQuantsParams struct {
	// Cursor for the next data sample
	Cursor string `json:"cursor"`

	// Filter
	Filter ListQuantsFilter `json:"filter"`

	// Maximum number of values in the response
	Limit int32 `json:"limit"`

	// Parameter by which products will be sorted
	Sort string `json:"sort"`

	// Sorting direction
	SortDir string `json:"sort_dir"`
}

type ListQuantsFilter struct {
	// MOQ creation period
	CreatedAt *ListQuantsFilterTime `json:"created_at"`

	// Time for MOQ assembly
	Cutoff *ListQuantsFilterTime `json:"cutoff"`

	// Destination point identifier
	DestinationPlaceId int64 `json:"destination_place_id"`

	// MOQ inventory identifiers
	InvQuantIds []string `json:"inv_quants_ids"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product name
	SKUName string `json:"sku_name"`

	// MOQ statuses
	Statuses []string `json:"statuses"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

type ListQuantsFilterTime struct {
	// Start date
	From string `json:"from"`

	// End date
	To string `json:"to"`
}

type ListQuantsResponse struct {
	core.CommonResponse

	Result ListQuantsResult `json:"result"`
}

type ListQuantsResult struct {
	// Cursor for the next data sample
	Cursor string `json:"cursor"`

	// Indication that the response returned only a part of characteristic values
	HasNext bool `json:"has_next"`

	// MOQs list
	Quants []Quant `json:"quants"`
}

type Quant struct {
	// List of available actions with MOQ
	AvailableActions []string `json:"available_actions"`

	// Date until which the leftover stock amount must be specified
	AwaitingStockDueDate string `json:"awaiting_stock_due_date"`

	// MOQ cancellation reason
	CancelReason `json:"cancel_reason"`

	// Seller company identifier
	CompanyId int64 `json:"company_id"`

	// MOQ creation date
	CreatedAt string `json:"created_at"`

	// Current number of shipments in the MOQ
	CurrentPostingsCount int64 `json:"current_postings_count"`

	// Time until which the MOQ must be assembled
	Cutoff string `json:"cutoff"`

	// Delivery method name
	DeliveryMethod string `json:"delivery_method_name"`

	// Destination point identifier
	DestinationPlaceId int64 `json:"destination_place_id"`

	// Destination point name
	DestinationPlaceName string `json:"destination_place_name"`

	// MOQ filling percentage
	FillingPercent float32 `json:"filling_percent"`

	// Date when the shipments start to get canceled if the MOQ is not reserved
	FirstPostingCancellationDate string `json:"first_posting_cancellation_date"`

	// MOQ identifier in Ozon system
	Id int64 `json:"id"`

	// MOQ inventory identifier
	InvQuantId int64 `json:"inv_quant_id"`

	// Date of the last MOQ status change
	LastStatusChangeAt string `json:"last_status_change_at"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Total cost of products in the MOQ
	ProductsPrice float32 `json:"products_price"`

	// Start date of MOQ filling
	QuantumStartDate string `json:"quantum_start_date"`

	// Product SKU
	SKU int64 `json:"sku"`

	// Product name
	SKUName string `json:"sku_name"`

	// MOQ statuses
	Status string `json:"status"`

	// Required number of products in the MOQ
	TargetPostingsCount int64 `json:"target_postings_count"`

	// Delivery service name
	TPLProviderName string `json:"tpl_provider_name"`

	// MOQ type: box or pallet
	Type string `json:"type"`

	// Seller warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`

	// Seller warehouse name
	WarehouseName string `json:"warehouse_name"`
}

type CancelReason struct {
	// Identifier of MOQ cancellation reason
	Id int64 `json:"cancel_reason_id"`

	// Cancellation reason name
	Name string `json:"cancel_reason_name"`

	// Cancellation initiator
	Responsible string `json:"responsible"`
}

// You can leave feedback on this method in the comments section to the discussion in the Ozon for dev community.
func (q Quants) List(ctx context.Context, params *ListQuantsParams) (*ListQuantsResponse, error) {
	url := "/v1/quant/list"

	resp := &ListQuantsResponse{}

	response, err := q.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
