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

type GetRFBSReturnsParams struct {
	// Filter
	Filter *GetRFBSReturnsFilter `json:"filter,omitempty"`

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

	Returns GetRFBSReturn `json:"returns"`
}

type GetRFBSReturn struct {
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

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
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

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
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

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
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

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
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
func (c Returns) ResetGiveoutBarcode(ctx context.Context) (*GetGiveoutResponse, error) {
	url := "/v1/return/giveout/barcode-reset"

	resp := &GetGiveoutResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetGiveoutListParams struct {
	// Identifier of the last value on the page
	LastId int64 `json:"last_id"`

	// Number of values in the response
	Limit int64 `json:"limit"`
}

type GetGiveoutListResponse struct {
	core.CommonResponse

	// Shipment identifier
	Giveouts []GetGiveoutListGiveout `json:"giveouts"`
}

type GetGiveoutListGiveout struct {
	// Number of products in shipment
	ApprovedArticlesCount int32 `json:"approved_articles_count"`

	// Creation date and time
	CreatedAt time.Time `json:"created_at"`

	// Shipment identifier
	GiveoutId int64 `json:"giveout_id"`

	// Return shipment status
	GiveoutStatus GiveoutStatus `json:"giveout_status"`

	// Total number of products to be picked up from the warehouse
	TotalArticlesCount int32 `json:"total_articles_count"`

	// Warehouse address
	WarehouseAddress string `json:"warehouse_address"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`

	// Warehouse name
	WarehouseName string `json:"warehouse_name"`
}

// Return shipments list
func (c Returns) GetGiveoutList(ctx context.Context, params *GetGiveoutListParams) (*GetGiveoutListResponse, error) {
	url := "/v1/return/giveout/list"

	resp := &GetGiveoutListResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetGiveoutInfoParams struct {
	// Shipment identifier
	GiveoutId int64 `json:"giveout_id"`
}

type GetGiveoutInfoResponse struct {
	core.CommonResponse

	// Product IDs
	Articles []GetGiveoutInfoArticle `json:"articles"`

	// Shipment identifier
	GiveoutId int64 `json:"giveout_id"`

	// Return shipment status
	GiveoutStatus GiveoutStatus `json:"giveout_status"`

	// Warehouse address
	WarehouseAddress string `json:"warehouse_address"`

	// Warehouse name
	WarehouseName string `json:"warehouse_name"`
}

type GetGiveoutInfoArticle struct {
	// `true` if the shipment is confirmed
	Approved bool `json:"approved"`

	// Delivery schema
	DeliverySchema GiveoutDeliverySchema `json:"delivery_schema"`

	// Product name
	Name string `json:"name"`

	// Seller identifier
	SellerId int64 `json:"seller_id"`
}

// Information on return shipment
func (c Returns) GetGiveoutInfo(ctx context.Context, params *GetGiveoutInfoParams) (*GetGiveoutInfoResponse, error) {
	url := "/v1/return/giveout/info"

	resp := &GetGiveoutInfoResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetFBSQuantityReturnsParams struct {
	Filter GetFBSQuantityReturnsFilter `json:"filter"`

	// Split the method response
	Pagination GetFBSQuantityReturnsPagination `json:"pagination"`
}

type GetFBSQuantityReturnsFilter struct {
	// Filter by drop-off point identifier
	PlaceId int64 `json:"place_id"`
}

type GetFBSQuantityReturnsPagination struct {
	// Identifier of the last drop-off point on the page. Leave this field blank in the first request.
	//
	// To get the next values, specify id of the last drop-off point from the response of the previous request
	LastId int64 `json:"last_id"`

	// Number of drop-off points per page. Maximum is 500
	Limit int32 `json:"limit"`
}

type GetFBSQuantityReturnsResponse struct {
	core.CommonResponse

	DropoffPoints []GetFBSQuantityDropoffPoint `json:"drop_off_points"`

	// true if there are any other points where sellers have orders waiting
	HasNext bool `json:"has_next"`
}

type GetFBSQuantityDropoffPoint struct {
	// Drop-off point address
	Address string `json:"address"`

	// Drop-off point identifier
	Id int64 `json:"id"`

	// Drop-off point name
	Name string `json:"name"`

	// Pass information
	PassInfo GetFBSQuantityDropoffPointPassInfo `json:"pass_info"`

	// The warehouse identifier to which the shipment will arrive
	PlaceId int64 `json:"place_id"`

	// Quantity of returns at the drop-off point
	ReturnsCount int32 `json:"returns_count"`

	// Seller's warehouses identifiers
	WarehouseIds []string `json:"warehouses_ids"`

	// Number of boxes in drop-off point
	BoxCount int32 `json:"box_count"`

	// Time zone offset of the shipping time from UTC-0
	UTCOffset string `json:"utc_offset"`
}

type GetFBSQuantityDropoffPointPassInfo struct {
	// Quantity of drop-off point passes
	Count int32 `json:"count"`

	// true if you need a pass to the drop-off point
	IsRequired bool `json:"is_required"`
}

func (c Returns) FBSQuantity(ctx context.Context, params *GetFBSQuantityReturnsParams) (*GetFBSQuantityReturnsResponse, error) {
	url := "/v1/returns/company/fbs/info"

	resp := &GetFBSQuantityReturnsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListReturnsParams struct {
	// Filter
	Filter *ListReturnsFilter `json:"filter,omitempty"`

	// Number of loaded returns. The maximum value is 500
	Limit int32 `json:"limit"`

	// Identifier of the last loaded return
	LastId int64 `json:"last_id"`
}

type ListReturnsFilter struct {
	// Filter by return creation date
	LogisticReturnDate *GetFBSReturnsFilterTimeRange `json:"logistic_return_date"`

	// Filter by storage fees start date
	StorageTarifficationDate *GetFBSReturnsFilterTimeRange `json:"storage_tariffication_start_date"`

	// Filter by date the return status changed
	VisualStatusChangeMoment *GetFBSReturnsFilterTimeRange `json:"visual_status_change_moment"`

	// Filter by order identifier
	OrderId int64 `json:"order_id,omitempty"`

	// Filter by shipment number
	PostingNumbers []string `json:"posting_numbers,omitempty"`

	// Filter by product name
	ProductName string `json:"product_name,omitempty"`

	// Filter by product identifier in the seller's system
	OfferId string `json:"offer_id,omitempty"`

	// Filter by return status
	VisualStatusName VisualStatus `json:"visual_status_name,omitempty"`

	// Filter by warehouse identifier
	WarehouseId int64 `json:"warehouse_id,omitempty"`

	// Filter by return label barcode
	Barcode string `json:"barcode,omitempty"`

	// Filter by delivery scheme: FBS or FBO
	ReturnSchema string `json:"return_schema,omitempty"`
}

type ListReturnsResponse struct {
	core.CommonResponse

	// Returns details
	Returns []Return `json:"returns"`

	// true, if the seller has other returns
	HasNext bool `json:"has_next"`
}

type Return struct {
	// Product items data
	Exemplars []ReturnExemplar `json:"exemplars"`

	// Return identifier
	Id int64 `json:"id"`

	// Company identifier
	CompanyId int64 `json:"company_id"`

	// Return reason
	ReturnReasonName string `json:"return_reason_name"`

	// Return type
	Type string `json:"type"`

	// Return scheme
	Schema string `json:"schema"`

	// Order identifier
	OrderId int64 `json:"order_id"`

	// Order number
	OrderNumber string `json:"order_number"`

	// Warehouse where the return is stored
	Place ReturnPlace `json:"place"`

	// Warehouse where returns are sent to
	TargetPlace ReturnPlace `json:"target_place"`

	// Storage details
	Storage ReturnStorage `json:"storage"`

	// Product details
	Product ReturnProduct `json:"product"`

	// Return details
	Logistic ReturnLogistic `json:"logistic"`

	// Return status details
	Visual ReturnVisual `json:"visual"`

	// Additional information
	AdditionalInfo ReturnAdditionalInfo `json:"additional_info"`

	// Previous return identifier
	SourceId int64 `json:"source_id"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Original shipment barcode
	ClearingId int64 `json:"clearing_id"`

	// Package unit identifier in the Ozon logistics system
	ReturnClearingId int64 `json:"return_clearing_id"`
}

type ReturnExemplar struct {
	// Product identifier
	Id int64 `json:"id"`
}

type ReturnPlace struct {
	// Warehouse identifier
	Id int64 `json:"id"`

	// Warehouse name
	Name string `json:"name"`

	// Warehouse address
	Address string `json:"address"`
}

type ReturnStorage struct {
	// Storage cost details
	Sum ReturnSum `json:"sum"`

	// First day of charging for storage
	TarifficationsFirstDate time.Time `json:"tariffication_first_date"`

	// Start date for storage fees
	TarifficationsStartDate time.Time `json:"tariffication_start_date"`

	// Date when the return was ready for handover
	ArrivedMoment time.Time `json:"arrived_moment"`

	// Number of days the return has been waiting for handover
	Days int64 `json:"days"`

	// Disposal cost details
	UtilizationSum ReturnSum `json:"utilization_sum"`

	// Planned disposal date
	UtilizationForecastDate string `json:"utilization_forecast_date"`
}

type ReturnSum struct {
	// Currency
	CurrencyCode string `json:"currency_code"`

	// Disposal cost
	Price float64 `json:"price"`
}

type ReturnProduct struct {
	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// product name
	Name string `json:"name"`

	// Product price details
	Price ReturnSum `json:"price"`

	// Product cost without commission
	PriceWithoutCommission ReturnSum `json:"price_without_commission"`

	// Sales commission by category
	CommissionPercent float64 `json:"commission_percent"`

	// Commission details
	Commission ReturnSum `json:"commission"`

	// Product quantity
	Quantity int32 `json:"quantity"`
}

type ReturnLogistic struct {
	// Date when the order was placed for technical return
	TechnicalReturnMoment time.Time `json:"technical_return_moment"`

	// Date when the return arrived to the warehouse or was handed over to the seller
	FinalMoment time.Time `json:"final_moment"`

	// Date when the seller received compensation for the return
	CancelledWithCompensationMoment time.Time `json:"cancelled_with_compensation_moment"`

	// Date when the customer returned the product
	ReturnDate time.Time `json:"return_date"`

	// Return label barcode
	Barcode string `json:"barcode"`
}

type ReturnVisual struct {
	// Return status
	Status ReturnVisualStatus `json:"status"`

	// Date the return status changed
	ChangeMoment time.Time `json:"change_moment"`
}

type ReturnVisualStatus struct {
	// Return status identifier
	Id int32 `json:"id"`

	// Return status name
	DisplayName string `json:"display_name"`

	// System name of the return status
	SystemName string `json:"sys_name"`
}

type ReturnAdditionalInfo struct {
	// true, if the return package is opened
	IsOpened bool `json:"is_opened"`

	// true, if the return belongs to Super Economy products
	IsSuperEconom bool `json:"is_super_econom"`
}

func (c Returns) List(ctx context.Context, params *ListReturnsParams) (*ListReturnsResponse, error) {
	url := "/v1/returns/list"

	resp := &ListReturnsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
