package ozon

import (
	"context"
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Returns struct {
	client *core.Client
}

type GetFBOReturnsParams struct {
	// Filter
	Filter GetFBOReturnsFilter `json:"filter"`

	// Identifier of the last value on the page. Leave this field blank in the first request.
	//
	// To get the next values, specify the recieved value in the next request in the `last_id` parameter
	LastId int64 `json:"last_id"`

	// Number of values in the response
	Limit int64 `json:"limit"`
}

type GetFBOReturnsFilter struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Return status
	Status []GetFBOReturnsFilterStatus `json:"status"`
}

type GetFBOReturnsResponse struct {
	core.CommonResponse

	// Identifier of the last value on the page
	LastId int64 `json:"last_id"`

	// Returns information
	Returns []GetFBOReturnsReturn `json:"returns"`
}

type GetFBOReturnsReturn struct {
	// Time when a return was received from the customer
	AcceptedFromCustomerMoment time.Time `json:"accepted_from_customer_moment"`

	// Seller identifier
	CompanyId int64 `json:"company_id"`

	// Current return location
	CurrentPlaceName string `json:"current_place_name"`

	// Return destination
	DestinationPlaceName string `json:"dst_place_name"`

	// Return identifier
	Id int64 `json:"id"`

	// Indication that the package has been opened. true, if it has been
	IsOpened bool `json:"is_opened"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Return reason
	ReturnReasonName string `json:"return_reason_name"`

	// Return delivery time to the Ozon warehouse
	ReturnedToOzonMoment time.Time `json:"returned_to_ozon_moment"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Return status
	Status GetFBOReturnsReturnStatus `json:"status_name"`
}

// Method for getting information on returned products that are sold from the Ozon warehouse
func (c Returns) GetFBOReturns(ctx context.Context, params *GetFBOReturnsParams) (*GetFBOReturnsResponse, error) {
	url := "/v3/returns/company/fbo"

	resp := &GetFBOReturnsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetFBSReturnsParams struct {
	// Filter
	Filter GetFBSReturnsFilter `json:"filter"`

	// Number of values in the response:
	//   - maximum — 1000,
	//   - minimum — 1
	Limit int64 `json:"limit"`

	// Return identifier that was loaded the last time.
	// Return identifiers with the higher value than `last_id`
	// will be returned in the response.
	LastId int64 `json:"offset"`
}

type GetFBSReturnsFilter struct {
	// Time of receiving the return from the customer
	AcceptedFromCustomerMoment GetFBSReturnsFilterTimeRange `json:"accepted_from_customer_moment"`

	// Last day of free storage
	LastFreeWaitingDay GetFBSReturnsFilterTimeRange `json:"last_free_waiting_dat"`

	// Order ID
	OrderId int64 `json:"order_id"`

	// Shipment ID
	PostingNumber []string `json:"posting_number"`

	// Product name
	ProductName string `json:"product_name"`

	// Product ID
	ProductOfferId string `json:"product_offer_id"`

	// Return status
	Status GetFBSReturnsFilterStatus `json:"status"`
}

type GetFBSReturnsFilterTimeRange struct {
	// The beginning of the period.
	//
	// Format: YYYY-MM-DDTHH:mm:ss.sssZ.
	//
	// Example: 2019-11-25T10:43:06.51
	TimeFrom time.Time `json:"time_from"`

	// The end of the period
	//
	// Format: YYYY-MM-DDTHH:mm:ss.sssZ.
	//
	// Example: 2019-11-25T10:43:06.51
	TimeTo time.Time `json:"time_to"`
}

type GetFBSReturnsResponse struct {
	core.CommonResponse

	// Return identifier that was loaded the last time.
	// Return identifiers with the higher value than `last_id`
	// will be returned in the response
	LastId int64 `json:"last_id"`

	// Returns information
	Returns []GetFBSReturnResultReturn `json:"returns"`
}

type GetFBSReturnResultReturn struct {
	// Time of receiving the return from the customer
	AcceptedFromCustomerAmount string `json:"accepted_from_customer_amount"`

	// Bottom barcode on the product label
	ClearingId int64 `json:"clearing_id"`

	// Commission fee
	Commission float64 `json:"commission"`

	// Commission percentage
	CommissionPercent float64 `json:"commission_percent"`

	// Product item identifier in the Ozon logistics system
	ExemplarId int64 `json:"exemplar_id"`

	// Return identifier in the Ozon accounting system
	Id int64 `json:"id"`

	// If the product is in transit — true
	IsMoving bool `json:"is_moving"`

	// Indication that the package has been opened. true, if it has been
	IsOpened bool `json:"is_opened"`

	// Last day of free storage
	LastFreeWaitingDay string `json:"last_free_waiting_day"`

	// ID of the warehouse the product is being transported to
	PlaceId int64 `json:"place_id"`

	// Name of the warehouse the product is being transported to
	MovingToPlaceName string `json:"moving_to_place_name"`

	// Delivery cost
	PickingAmount float64 `json:"picking_amount"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	PickingTag string `json:"picking_tag"`

	// Current product price without a discount
	Price float64 `json:"price"`

	// Product price without commission
	PriceWithoutCommission float64 `json:"price_without_commission"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product name
	ProductName string `json:"product_name"`

	// Product quantity
	Quantity int64 `json:"quantity"`

	// Barcode on the return label. Use this parameter value to work with the return label
	ReturnBarcode string `json:"return_barcode"`

	// Package unit identifier in the Ozon logistics system
	ReturnClearingId int64 `json:"return_clearing_id"`

	// Product return date
	ReturnDate string `json:"return_date"`

	// Return reason
	ReturnReasonName string `json:"return_reason_name"`

	// Date when the product is ready to be handed over to the seller
	WaitingForSellerDate string `json:"waiting_for_seller_date_time"`

	// Date of handing over the product to the seller
	ReturnedToSellerDate string `json:"returned_to_seller_date_time"`

	// Return storage period in days
	WaitingForSellerDays int64 `json:"waiting_for_seller_days"`

	// Return storage cost
	ReturnsKeepingCost float64 `json:"returns_keeping_cost"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Return status
	Status string `json:"status"`
}

// Method for getting information on returned products that are sold from the seller's warehouse
func (c Returns) GetFBSReturns(ctx context.Context, params *GetFBSReturnsParams) (*GetFBSReturnsResponse, error) {
	url := "/v3/returns/company/fbs"

	resp := &GetFBSReturnsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetRFBSReturnsParams struct {
	// Filter
	Filter GetRFBSReturnsFilter `json:"filter"`

	// Identifier of the last value on the page.
	// Leave this field blank in the first request
	LastId int32 `json:"last_id"`

	// Number of values per page
	Limit int32 `json:"limit"`
}

type GetRFBSReturnsFilter struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Filter by request statuses
	GroupState []RFBSReturnsGroupState `json:"group_state"`

	// Period of request creation
	CreatedAt GetRFBSReturnsFilterCreatedAt `json:"created_at"`
}

type GetRFBSReturnsFilterCreatedAt struct {
	// Period start date
	From time.Time `json:"from"`

	// Period end date
	To time.Time `json:"to"`
}

type GetRFBSReturnsResponse struct {
	core.CommonResponse

	// Information on return requests
	Returns GetRFBSReturnsReturn `json:"returns"`
}

type GetRFBSReturnsReturn struct {
	// Customer name
	ClientName string `json:"client_name"`

	// Request creation date
	CreatedAt time.Time `json:"created_at"`

	// Order number
	OrderNumber string `json:"order_number"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Product details
	Product GetRFBSReturnsProduct `json:"product"`

	// Return request identifier
	ReturnId int64 `json:"return_id"`

	// Return request number
	ReturnNumber string `json:"return_number"`

	// Request and refund statuses
	State GetRFBSReturnsState `json:"state"`
}

type GetRFBSReturnsProduct struct {
	// Product name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Currency of your prices. It matches the currency set in your personal account
	CurrencyCode GetRFBSReturnsCurrency `json:"currency_code"`

	// Product price
	Price string `json:"price"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

type GetRFBSReturnsState struct {
	// Request status by the applied filter
	GroupState RFBSReturnsGroupState `json:"group_state"`

	// Refund status
	MoneyReturnStateName string `json:"money_return_state_name"`

	// Request status
	State string `json:"state"`

	// Request status name in Russian
	StateName string `json:"state_name"`
}

// Get a list of return requests
func (c Returns) GetRFBSReturns(ctx context.Context, params *GetRFBSReturnsParams) (*GetRFBSReturnsResponse, error) {
	url := "/v2/returns/rfbs/list"

	resp := &GetRFBSReturnsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetRFBSReturnParams struct {
	// Request identifier
	ReturnId int64 `json:"return_id"`
}

type GetRFBSReturnResponse struct {
	core.CommonResponse

	// List of available actions on the request
	AvailableActions []GetRFBSReturnAction `json:"available_actions"`

	// Customer name
	ClientName string `json:"client_name"`

	// Links to product images
	ClientPhoto []string `json:"client_photo"`

	// Information on return method
	ClientReturnMethodType GetRFBSReturnMethodType `json:"client_return_method_type"`

	// Customer comment
	Comment string `json:"comment"`

	// Request creation date
	CreatedAt time.Time `json:"created_at"`

	// Order number
	OrderNumber string `json:"order_number"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Product details
	Product GetRFBSReturnsProduct `json:"product"`

	// Comment on request rejection
	RejectionComment string `json:"rejection_comment"`

	// Information on rejection reason
	RejectionReason []GetRFBSReturnRejectionReason `json:"rejection_reason"`

	// Method of product return
	ReturnMethodDescription string `json:"return_method_description"`

	// Return request number
	ReturnNumber string `json:"return_number"`

	// Information on return reason
	ReturnReason GetRFBSReturnReason `json:"return_reason"`

	// Postal tracking number
	RUPostTrackingNumber string `json:"ru_post_tracking_number"`

	// Information on return status
	State GetRFBSReturnState `json:"state"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

type GetRFBSReturnAction struct {
	// Action identifier
	Id int32 `json:"id"`

	// Action name
	Name string `json:"name"`
}

type GetRFBSReturnMethodType struct {
	// Identifier
	Id int32 `json:"id"`

	// Name
	Name string `json:"name"`
}

type GetRFBSReturnRejectionReason struct {
	// Hint on further actions with the return
	Hint string `json:"hint"`

	// Reason identifier
	Id int32 `json:"id"`

	// `true` if you need to attach a comment
	IsCommentRequired bool `json:"is_comment_required"`

	// Reason description
	Name string `json:"name"`
}

type GetRFBSReturnReason struct {
	// Reason identifier
	Id int32 `json:"id"`

	// `true` if the product is defective
	IsDefect bool `json:"is_defect"`

	// Reason description
	Name string `json:"name"`
}

type GetRFBSReturnState struct {
	// Status
	State string `json:"state"`

	// Status name in Russian
	StateName string `json:"state_name"`
}

// Get information about a return request
func (c Returns) GetRFBSReturn(ctx context.Context, params *GetRFBSReturnParams) (*GetRFBSReturnResponse, error) {
	url := "/v2/returns/rfbs/get"

	resp := &GetRFBSReturnResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type RejectRFBSReturnParams struct {
	// Return request identifier
	ReturnId int64 `json:"return_id"`

	// Comment
	//
	// The comment is required if the
	// `rejection_reason.is_comment_required` parameter is `true`
	// in the response of the `/v2/returns/rfbs/get` method
	Comment string `json:"comment"`

	// Rejection reason identifier.
	//
	// Pass the value from the list of reasons received in the response
	// of the `/v2/returns/rfbs/get` method in the `rejection_reason` parameter
	RejectionReasonId int64 `json:"rejection_reason_id"`
}

type RejectRFBSReturnResponse struct {
	core.CommonResponse
}

// The method rejects an rFBS return request. Explain your decision in the `comment` parameter
func (c Returns) RejectRFBSReturn(ctx context.Context, params *RejectRFBSReturnParams) (*RejectRFBSReturnResponse, error) {
	url := "/v2/returns/rfbs/reject"

	resp := &RejectRFBSReturnResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CompensateRFBSReturnParams struct {
	// Compensation amount
	CompensationAmount string `json:"compensation_amount"`

	// Return request identifier
	ReturnId int64 `json:"return_id"`
}

type CompensateRFBSReturnResponse struct {
	core.CommonResponse
}

// Using this method you can confirm the partial compensation and agree to keep the product with the customer
func (c Returns) CompensateRFBSReturn(ctx context.Context, params *CompensateRFBSReturnParams) (*CompensateRFBSReturnResponse, error) {
	url := "/v2/returns/rfbs/compensate"

	resp := &CompensateRFBSReturnResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ApproveRFBSReturnParams struct {
	// Return request identifier
	ReturnId int64 `json:"return_id"`

	// Method of product return
	ReturnMethodDescription string `json:"return_method_description"`
}

type ApproveRFBSReturnResponse struct {
	core.CommonResponse
}

// The method allows to approve an rFBS return request and agree to receive products for verification.
//
// Confirm that you've received the product using the `/v2/returns/rfbs/receive-return` method.
func (c Returns) ApproveRFBSReturn(ctx context.Context, params *ApproveRFBSReturnParams) (*ApproveRFBSReturnResponse, error) {
	url := "/v2/returns/rfbs/verify"

	resp := &ApproveRFBSReturnResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ReceiveRFBSReturnParams struct {
	// Return request identifier
	ReturnId int64 `json:"return_id"`
}

type ReceiveRFBSReturnResponse struct {
	core.CommonResponse
}

// Confirm receipt of a product for check
func (c Returns) ReceiveRFBSReturn(ctx context.Context, params *ReceiveRFBSReturnParams) (*ReceiveRFBSReturnResponse, error) {
	url := "/v2/returns/rfbs/receive-return"

	resp := &ReceiveRFBSReturnResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type RefundRFBSParams struct {
	// Return request identifier
	ReturnId int64 `json:"return_id"`

	// Refund amount for shipping the product
	ReturnForBackWay int64 `json:"return_for_back_way"`
}

type RefundRFBSResponse struct {
	core.CommonResponse
}

// The method confirms the refund of the full product cost.
// Use the method if you agree to refund the customer:
//
// Immediately without receiving the product.
// After you received and checked the product.
// If the product is defective or damaged, you also refund its return shipment cost.
func (c Returns) RefundRFBS(ctx context.Context, params *RefundRFBSParams) (*RefundRFBSResponse, error) {
	url := "/v2/returns/rfbs/return-money"

	resp := &RefundRFBSResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type IsGiveoutEnabledResponse struct {
	core.CommonResponse

	// `true` if you can pick up a return shipment by barcode.
	Enabled bool `json:"enabled"`
}

// Check the ability to receive return shipments by barcode
//
// The `enabled` parameter is true if you can pick up return shipments by barcode.
func (c Returns) IsGiveoutEnabled(ctx context.Context) (*IsGiveoutEnabledResponse, error) {
	url := "/v1/return/giveout/is-enabled"

	resp := &IsGiveoutEnabledResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, struct{}{}, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetGiveoutResponse struct {
	core.CommonResponse

	// PDF file with barcode in binary format
	FileContent string `json:"file_content"`

	// File name
	FileName string `json:"file_name"`

	// File type
	ContentType string `json:"content_type"`
}

// Barcode for return shipment in PDF format
//
// Returns a PDF file with a barcode
func (c Returns) GetGiveoutPDF(ctx context.Context) (*GetGiveoutResponse, error) {
	url := "/v1/return/giveout/get-pdf"

	resp := &GetGiveoutResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, struct{}{}, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

// Barcode for return shipment in PNG format
//
// Returns a PNG file with a barcode
func (c Returns) GetGiveoutPNG(ctx context.Context) (*GetGiveoutResponse, error) {
	url := "/v1/return/giveout/get-png"

	resp := &GetGiveoutResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, struct{}{}, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetGiveoutBarcodeResponse struct {
	core.CommonResponse

	// Barcode value in text format
	Barcode string `json:"barcode"`
}

// Value of barcode for return shipments
//
// Use this method to get the barcode from the response of the
// `/v1/return/giveout/get-png` and `/v1/return/giveout/get-pdf` methods in text format
func (c Returns) GetGiveoutBarcode(ctx context.Context) (*GetGiveoutBarcodeResponse, error) {
	url := "/v1/return/giveout/barcode"

	resp := &GetGiveoutBarcodeResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, struct{}{}, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

// Use this method if an unauthorized person has gained access to your barcode.
//
// The method returns a PNG file with the new barcode. Once the method is used,
// you won't be able to get a return shipment using the old barcodes.
// To get a new barcode in PDF format, use the /v1/return/giveout/get-pdf method
func (c Returns) ResetGiveoutBarcode(ctx context.Context) (*GetGiveoutBarcodeResponse, error) {
	url := "/v1/return/giveout/barcode-reset"

	resp := &GetGiveoutBarcodeResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, struct{}{}, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
