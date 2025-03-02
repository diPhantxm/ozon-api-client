package ozon

import (
	"context"
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type FBS struct {
	client *core.Client
}

type ListUnprocessedShipmentsParams struct {
	// Sorting direction
	Direction Order `json:"dir,omitempty"`

	// Request filter
	Filter ListUnprocessedShipmentsFilter `json:"filter"`

	// Number of values in the response:
	//
	// maximum — 1000,
	// minimum — 1.
	Limit int64 `json:"limit"`

	// Number of elements that will be skipped in the response.
	// For example, if `offset` = 10, the response will start with the 11th element found
	Offset int64 `json:"offset,omitempty"`

	// Additional fields that should be added to the response
	With *ListUnprocessedShipmentsWith `json:"with,omitempty"`
}

type ListUnprocessedShipmentsFilter struct {
	// Filter by the time by which the seller should pack the order. Period start.
	//
	// Format: YYYY-MM-DDThh: mm:ss. mcsZ. Example: 2020-03-18T07:34:50.359 Z
	CutoffFrom *core.TimeFormat `json:"cutoff_from,omitempty"`

	// Filter by the time by which the seller should pack the order. Period end.
	//
	// Format: YYYY-MM-DDThh: mm:ss. mcsZ. Example: 2020-03-18T07:34:50.359 Z
	CutoffTo *core.TimeFormat `json:"cutoff_to,omitempty"`

	// Minimum date when shipment should be handed over for delivery
	DeliveringDateFrom *core.TimeFormat `json:"delivering_date_from,omitempty"`

	// Maximum date when shipment should be handed over for delivery
	DeliveringDateTo *core.TimeFormat `json:"delivering_date_to,omitempty"`

	// Delivery method identifier
	DeliveryMethodId []int64 `json:"delivery_method_id"`

	// Specify true to get only MOQ shipments.
	//
	// The default value is false, the response contains all shipments
	IsQuantum bool `json:"is_quantum"`

	// Filter for shipments delivered from partner warehouse (FBP). You can pass one of the following values:
	//
	// Default value is all.
	//
	// The FBP scheme is available only for sellers from China
	FBPFilter FBPFilter `json:"fbpFilter" default:"all"`

	// Delivery service identifier
	ProviderId []int64 `json:"provider_id"`

	// Shipment status
	Status string `json:"status"`

	// Warehouse identifier
	WarehouseId []int64 `json:"warehouse_id"`
}

type ListUnprocessedShipmentsWith struct {
	// Add analytics data to the response
	AnalyticsData bool `json:"analytics_data"`

	// Add the shipment barcodes to the response
	Barcodes bool `json:"barcodes"`

	// Add financial data to the response
	FinancialData bool `json:"financial_data"`

	// Transliterate the return values
	Translit bool `json:"translit"`
}

type ListUnprocessedShipmentsResponse struct {
	core.CommonResponse

	// Request result
	Result ListUnprocessedShipmentsResult `json:"result"`
}

type ListUnprocessedShipmentsResult struct {
	// Element counter in the response
	Count int64 `json:"count"`

	// List of shipments and detailed information on each one
	Postings []FBSPosting `json:"postings"`
}

type FBSPosting struct {
	// Recipient details
	Addressee FBSPostingAddressee `json:"addressee"`

	// Analytics data
	AnalyticsData FBSPostingAnalyticsData `json:"analytics_data"`

	// Available actions and shipment information
	AvailableActions []string `json:"available_actions"`

	// Shipment barcodes
	Barcodes FBSBarcode `json:"barcodes"`

	// Cancellation details
	Cancellation FBSCancellation `json:"cancellation"`

	// Customer details
	Customer FBSCustomer `json:"customer"`

	// Date when the shipment was transferred for delivery
	DeliveringDate time.Time `json:"delivering_date"`

	// Delivery method
	DeliveryMethod FBSDeliveryMethod `json:"delivery_method"`

	// Data on the product cost, discount amount, payout and commission
	FinancialData FBSFinancialData `json:"financial_data"`

	// Start date and time of shipment processing
	InProccessAt time.Time `json:"in_process_at"`

	// If Ozon Express fast delivery was used — `true`
	IsExpress bool `json:"is_express"`

	// Indication that there is a multi-box product in the shipment
	// and you need to pass the number of boxes for it
	IsMultibox bool `json:"is_multibox"`

	// Number of boxes in which the product is packed
	MultiBoxQuantity int32 `json:"multi_box_qty"`

	// Identifier of the order to which the shipment belongs
	OrderId int64 `json:"order_id"`

	// Number of the order to which the shipment belongs
	OrderNumber string `json:"order_number"`

	// Number of the parent shipment which split resulted in the current shipment
	ParentPostingNumber string `json:"parent_posting_number"`

	// Date and time of successful verification of the courier code
	PickupCodeVerifiedAt time.Time `json:"pickup_code_verified_at"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// List of products in the shipment
	Products []PostingProduct `json:"products"`

	// The parameter is only relevant for bulky products
	// with a delivery by a third-party or integrated service
	PRROption PRROptionStatus `json:"prr_option"`

	// Array of Ozon Product IDs (SKU) for which you need to pass the
	// customs cargo declaration (CCD) number, the manufacturing country,
	// product batch registration number, or "Chestny ZNAK" labeling to
	// change the shipment status to the next one
	Requirements FBSRequirements `json:"requirements"`

	// Date and time before which the shipment must be packaged.
	// If the shipment is not packaged by this date, it will be
	// canceled automatically
	ShipmentDate time.Time `json:"shipment_date"`

	// Shipment status
	Status string `json:"status"`

	// Shipment substatus
	Substatus string `json:"substatus"`

	// Type of integration with the delivery service
	TPLIntegrationType TPLIntegrationType `json:"tpl_integration_type"`

	// Shipment tracking number
	TrackingNumber string `json:"tracking_number"`

	// Details on shipping rate
	Tariffication []FBSPostingTariffication `json:"tariffication"`

	// Economy product identifier
	QuantumId int64 `json:"quantum_id"`
}

type FBSPostingTariffication struct {
	// Current shipping rate as a percentage
	CurrentTariffRate float64 `json:"current_tariff_rate"`

	// Type of shipping rate adjustment: discount or surcharge
	CurrentTariffType string `json:"current_tariff_type"`

	// Current amount of discount or surcharge
	CurrentTariffCharge string `json:"current_tariff_charge"`

	// Currency of the amount
	CurrencyTariffCurrencyCode string `json:"current_tariff_charge_currency_code"`

	// Percentage by which the shipping rate is adjusted
	// after the time specified in the next_tariff_starts_at parameter
	NextTariffRate float64 `json:"next_tariff_rate"`

	// The adjustment type applied to the shipping rate
	// after the time specified in the next_tariff_starts_at parameter:
	// discount or surcharge
	NextTariffType string `json:"next_tariff_type"`

	// Discount or surcharge amount applied during the next shipping rate adjustment step
	NextTariffCharge string `json:"next_tariff_charge"`

	// Date and time when the new shipping rate is applied
	NextTariffStartsAt time.Time `json:"next_tariff_starts_at"`

	// New shipping rate currency
	NextTariffCurrencyCode string `json:"next_tariff_charge_currency_code"`
}

type FBSPostingAddressee struct {
	// Recipient name
	Name string `json:"name"`

	// Recipient phone number.
	//
	// Returns an empty string
	Phone string `json:"phone"`
}

type FBSPostingAnalyticsData struct {
	// Delivery city. Only for rFBS shipments
	City string `json:"city"`

	// Delivery start date and time
	DeliveryDateBegin time.Time `json:"delivery_date_begin"`

	// Delivery end date and time
	DeliveryDateEnd time.Time `json:"delivery_date_end"`

	// Delivery method
	DeliveryType string `json:"delivery_type"`

	// Indication that the recipient is a legal entity:
	//
	// true — a legal entity,
	// false — a natural person
	IsLegal bool `json:"is_legal"`

	// Premium subscription availability
	IsPremium bool `json:"is_premium"`

	// Payment method
	PaymentTypeGroupName PaymentTypeGroupName `json:"payment_type_group_name"`

	// Delivery region. Only for rFBS shipments
	Region string `json:"region"`

	// Delivery service
	TPLProvider string `json:"tpl_provider"`

	// Delivery service identifier
	TPLProviderId int64 `json:"tpl_provider_id"`

	// Order shipping warehouse name
	Warehouse string `json:"warehouse"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

type FBSBarcode struct {
	// Lower barcode on the shipment label
	LowerBarcode string `json:"lower_barcode"`

	// Upper barcode on the shipment label
	UpperBarcode string `json:"upper_barcode"`
}

type FBSCancellation struct {
	// If the cancellation affects the seller's rating—true
	AffectCancellationRating bool `json:"affect_cancellation_rating"`

	// Cancellation reason
	CancelReason string `json:"cancel_reason"`

	// Cancellation reason identifier
	CancelReasonId int64 `json:"cancel_reason_id"`

	// Shipment cancellation initiator:
	//
	// Клиент—client,
	// Ozon,
	// Продавец—seller
	CancellationInitiator string `json:"cancellation_initiator"`

	// Cancellation type:
	//
	// client—canceled by client.
	// ozon—canceled by Ozon.
	// seller—canceled by seller
	CancellationType string `json:"cancellation_type"`

	// If the cancellation occurred after the shipment had been packaged—true
	CancelledAfterShip bool `json:"cancelled_after_ship"`
}

type FBSDeliveryMethod struct {
	// Delivery method identifier
	Id int64 `json:"id"`

	// Delivery method name
	Name string `json:"name"`

	// Delivery service
	TPLProvider string `json:"tpl_provider"`

	// Delivery service identifier
	TPLProviderId int64 `json:"tpl_provider_id"`

	// Warehouse name
	Warehouse string `json:"warehouse"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

type FBSFinancialData struct {
	// Identifier of the cluster, where the shipment is sent from
	ClusterFrom string `json:"cluster_from"`

	// Identifier of the cluster, where the shipment is delivered
	ClusterTo string `json:"cluster_to"`

	// Services
	PostingServices MarketplaceServices `json:"posting_services"`

	// List of products in the shipment
	Products []FinancialDataProduct `json:"products"`
}

type FBSRequirements struct {
	// Array of Ozon Product IDs (SKU) for which you need to pass the customs cargo declaration (CCD) numbers.
	//
	// To pack the shipment, pass the CCD number for all listed SKUs.
	// If you do not have a CCD number, pass the value `is_gtd_absent` = true
	// via the `/v3/posting/fbs/ship/package`
	ProductsRequiringGTD []string `json:"products_requiring_gtd"`

	// Array of Ozon Product IDs (SKU) for which
	// you need to pass the manufacturing country.
	//
	// To pack the shipment, pass the manufacturing
	// country value for all listed SKUs via the `/v2/posting/fbs/product/country/set` method
	ProductsRequiringCountry []string `json:"products_requiring_country"`

	// Array of Ozon Product IDs (SKU) for which you need to pass the "Chestny ZNAK" labeling
	ProductsRequiringMandatoryMark []string `json:"products_requiring_mandatory_mark"`

	// Array of Ozon Product IDs (SKU) for which you need to pass a product batch registration number
	ProductsRequiringRNPT []string `json:"products_requiring_rnpt"`
}

type PostingProduct struct {
	// Mandatory product labeling
	MandatoryMark []string `json:"mandatory_mark"`

	// Product name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Currency of your prices. It matches the one set in the personal account settings
	CurrencyCode string `json:"currency_code"`

	// Product price
	Price string `json:"price"`

	// Product quantity in the shipment
	Quantity int32 `json:"quantity"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

type FBSCustomer struct {
	// Delivery address details
	Address FBSCustomerAddress `json:"address"`

	// Customer identifier
	CustomerId int64 `json:"customer_id"`

	// Customer name
	Name string `json:"name"`

	// Customer phone number.
	//
	// Returns an empty string ""
	Phone string `json:"phone"`
}

type FBSCustomerAddress struct {
	// Address in text format
	AddressTail string `json:"address_tail"`

	// Delivery city
	City string `json:"city"`

	// Comment on the order
	Comment string `json:"comment"`

	// Delivery country
	Country string `json:"country"`

	// Delivery area
	District string `json:"district"`

	// Latitude
	Latitude float64 `json:"latitude"`

	// Longitude
	Longitude float64 `json:"longitude"`

	// 3PL pick-up point code
	ProviderPVZCode string `json:"provider_pvz_code"`

	// Pick-up point code
	PVZCode int64 `json:"pvz_code"`

	// Delivery region
	Region string `json:"region"`

	// Recipient postal code
	ZIPCode string `json:"zip_code"`
}

type MarketplaceServices struct {
	// Last mile
	DeliveryToCustomer float64 `json:"marketplace_service_item_deliv_to_customer"`

	// Pipeline
	DirectFlowTrans float64 `json:"marketplace_service_item_direct_flow_trans"`

	// Shipment processing in the fulfilment warehouse (FF)
	DropoffFF float64 `json:"marketplace_service_item_dropoff_ff"`

	// Shipment processing at the pick up point
	DropoffPVZ float64 `json:"marketplace_service_item_dropoff_pvz"`

	// Shipment processing at the sorting center
	DropoffSC float64 `json:"marketplace_service_item_dropoff_sc"`

	// Order packaging
	Fulfillment float64 `json:"marketplace_service_item_fulfillment"`

	// Transport arrival to the seller's address for shipments pick-up (Pick-up)
	Pickup float64 `json:"marketplace_service_item_pickup"`

	// Return processing
	ReturnAfterDeliveryToCustomer float64 `json:"marketplace_service_item_return_after_deliv_to_customer"`

	// Reverse pipeline
	ReturnFlowTrans float64 `json:"marketplace_service_item_return_flow_trans"`

	// Cancellations processing
	ReturnNotDeliveryToCustomer float64 `json:"marketplace_service_item_return_not_deliv_to_customer"`

	// Non-purchase processing
	ReturnPartGoodsCustomer float64 `json:"marketplace_service_item_return_part_goods_customer"`
}

type FinancialDataProduct struct {
	// Actions
	Actions []string `json:"actions"`

	// Customer price
	ClientPrice string `json:"client_price"`

	// Commission amount for the product
	CommissionAmount float64 `json:"commission_amount"`

	// Commission percentage
	CommissionPercent int64 `json:"commission_percent"`

	// Code of the currency used to calculate the commissions
	CommissionsCurrencyCode string `json:"commissions_currency_code"`

	// Services
	//
	// Deprecated: The parameter is outdated. To get information on accruals, use the `ListTransactions` method
	ItemServices MarketplaceServices `json:"item_services"`

	// Currency of your prices. It matches the currency set in the personal account settings
	CurrencyCode string `json:"currency_code"`

	// Price before discounts. Displayed strikethrough on the product description page
	OldPrice float64 `json:"old_price"`

	// Payment to the seller
	Payout float64 `json:"payout"`

	// Delivery details.
	//
	// Returns `null`
	Picking FinancialDataProductPicking `json:"picking"`

	// Product price including discounts. This value is shown on the product description page
	Price float64 `json:"price"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product quantity in the shipment
	Quantity int64 `json:"quantity"`

	// Discount percentage
	TotalDiscountPercent float64 `json:"total_discount_percent"`

	// Discount amount
	TotalDiscountValue float64 `json:"total_discount_value"`
}

type FinancialDataProductPicking struct {
	// Delivery cost
	Amount float64 `json:"amount"`

	// Delivery date and time
	Moment time.Time `json:"moment"`

	// Bulky products or not
	Tag string `json:"tag"`
}

func (c FBS) ListUnprocessedShipments(ctx context.Context, params *ListUnprocessedShipmentsParams) (*ListUnprocessedShipmentsResponse, error) {
	url := "/v3/posting/fbs/unfulfilled/list"

	resp := &ListUnprocessedShipmentsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetFBSShipmentsListParams struct {
	// Sorting direction
	Direction Order `json:"dir,omitempty"`

	// Filter
	Filter GetFBSShipmentsListFilter `json:"filter"`

	// Number of shipments in the response:
	//   - maximum is 50,
	//   - minimum is 1.
	Limit int64 `json:"limit"`

	// Number of elements that will be skipped in the response. For example, if offset=10, the response will start with the 11th element found
	Offset int64 `json:"offset,omitempty"`

	// Additional fields that should be added to the response
	With *GetFBSShipmentsListWith `json:"with,omitempty"`
}

type GetFBSShipmentsListFilter struct {
	// Delivery method identifier
	DeliveryMethodId []int64 `json:"delivery_method_id"`

	// Filter for shipments delivered from partner warehouse (FBP)
	//
	// Default value is all.
	//
	// The FBP scheme is available only for sellers from China
	FBPFilter FBPFilter `json:"fbpFilter" default:"all"`

	// Order identifier
	OrderId int64 `json:"order_id"`

	// Specify true to get only MOQ shipments.
	//
	// The default value is false, the response contains all shipments
	IsQuantum bool `json:"is_quantum"`

	// Delivery service identifier
	ProviderId []int64 `json:"provider_id"`

	// Start date of the period for which a list of shipments should be generated.
	//
	// Format: YYYYY-MM-DDTHH:MM:SSZ.
	//
	// Example: 2019-08-24T14:15:22Z
	Since time.Time `json:"since"`

	// End date of the period for which a list of shipments should be generated.
	//
	// Format: YYYYY-MM-DDTHH:MM:SSZ.
	//
	// Example: 2019-08-24T14:15:22Z.
	To time.Time `json:"to"`

	// Shipment status
	Status string `json:"status"`

	// Warehouse identifier
	WarehouseId []int64 `json:"warehouse_id"`

	LastChangedStatusDate GetFBSShipmentsListFilterLastChangeDate `json:"last_changed_status_date"`
}

type GetFBSShipmentsListFilterLastChangeDate struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type GetFBSShipmentsListWith struct {
	// Add analytics data to the response
	AnalyticsData bool `json:"analytics_data"`

	// Add the shipment barcodes to the response
	Barcodes bool `json:"barcodes"`

	// Add financial data to the response
	FinancialData bool `json:"financial_data"`

	// Transliterate the return values
	Translit bool `json:"translit"`
}

type GetFBSShipmentsListResponse struct {
	core.CommonResponse

	// Array of shipments
	Result GetFBSShipmentsListResult `json:"result"`
}

type GetFBSShipmentsListResult struct {
	// Indicates that the response returned not the entire array of shipments:
	//
	//   - true — it is necessary to make a new request with a different offset value to get information on the remaining shipments;
	//   - false — the entire array of shipments for the filter specified in the request was returned in the response
	HasNext bool `json:"has_next"`

	// Shipment details
	Postings []FBSPosting `json:"postings"`
}

// Returns a list of shipments for the specified time period: it shouldn't be longer than one year.
//
// You can filter shipments by their status. The list of available statuses is specified in the description of the filter.status parameter.
//
// The true value of the has_next parameter in the response means there is not the entire array of shipments in the response. To get information on the remaining shipments, make a new request with a different offset value.
func (c FBS) GetFBSShipmentsList(ctx context.Context, params *GetFBSShipmentsListParams) (*GetFBSShipmentsListResponse, error) {
	url := "/v3/posting/fbs/list"

	resp := &GetFBSShipmentsListResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type PackOrderParams struct {
	// List of packages. Each package contains a list of shipments that the order was divided into
	Packages []PackOrderPackage `json:"packages"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Additional information
	With *PackOrderWith `json:"with,omitempty"`
}

type PackOrderPackage struct {
	Products []PackOrderPackageProduct `json:"products"`
}

type PackOrderPackageProduct struct {
	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product items quantity
	Quantity int32 `json:"quantity"`
}

type PackOrderWith struct {
	// Pass true to get additional information
	AdditionalData bool `json:"additional_data"`
}

type PackOrderResponse struct {
	core.CommonResponse

	// Additional information about shipments
	AdditionalData []PackOrderAdditionalData `json:"additional_data"`

	// Order packaging result
	Result []string `json:"result"`
}

type PackOrderAdditionalData struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// List of products in the shipment
	Products []PostingProduct `json:"products"`
}

// Divides the order into shipments and changes its status to awaiting_deliver.
//
// Each element of the packages may contain several instances of the products. One instance of the products is one shipment. Each element of the products is a product included into the shipment.
//
// It is necessary to split the order if:
//
// the products do not fit in one package,
// the products cannot be put in one package.
// Differs from /v2/posting/fbs/ship by the presence of the field `exemplar_info` in the request.
//
// If necessary, specify the number of the cargo customs declaration in the gtd parameter. If it is missing, pass the value is_gtd_absent = true
func (c FBS) PackOrder(ctx context.Context, params *PackOrderParams) (*PackOrderResponse, error) {
	url := "/v4/posting/fbs/ship"

	resp := &PackOrderResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ValidateLabelingCodesParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Products list
	Products []ValidateLabelingCodesProduct `json:"products"`
}

type ValidateLabelingCodesProduct struct {
	// Product items data
	Exemplars []ValidateLabelingCodesExemplar `json:"exemplars"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

type ValidateLabelingCodesExemplar struct {
	// Сustoms cargo declaration (CCD) number
	GTD string `json:"gtd"`

	// Mandatory “Chestny ZNAK” labeling
	Marks []SetProductItemsDataProductMark `json:"marks"`

	// Product batch registration number
	RNPT string `json:"rnpt"`
}

type ValidateLabelingCodesResponse struct {
	core.CommonResponse

	// Products list
	Products []ValidateLabelingCodesResultProduct `json:"products"`
}

type ValidateLabelingCodesResultProduct struct {
	// Error code
	Error string `json:"error"`

	// Product items data
	Exemplars []ValidateLabelingCodesResultExemplar `json:"exemplars"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Check result. true if the labeling codes of all product items meet the requirements
	Valid bool `json:"valid"`
}

type ValidateLabelingCodesResultExemplar struct {
	// Product item validation errors
	Errors []string `json:"errors"`

	// Сustoms cargo declaration (CCD) number
	GTD string `json:"gtd"`

	// List of Control Identification Marks in one copy
	Marks []ValidateLabelingCodesMark `json:"marks"`

	// Product batch registration number
	RNPT string `json:"rnpt"`

	// Check result. true if the labeling codes of all product items meet the requirements
	Valid bool `json:"valid"`
}

type ValidateLabelingCodesMark struct {
	// Errors that appeared during verification of Control Identification Marks
	Errors []string `json:"errors"`

	// Labeling code meaning
	Mark string `json:"mark"`

	// Labeling code type
	MarkType string `json:"mark_type"`

	// Check result. true if the labeling
	// codes of all product items meet the requirements
	Valid bool `json:"valid"`
}

// Method for checking whether labeling codes meet the "Chestny ZNAK" system requirements on length and symbols.
//
// If you don't have the customs cargo declaration (CCD) number, you don't have to specify it
func (c FBS) ValidateLabelingCodes(ctx context.Context, params *ValidateLabelingCodesParams) (*ValidateLabelingCodesResponse, error) {
	url := "/v5/fbs/posting/product/exemplar/validate"

	resp := &ValidateLabelingCodesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetShipmentDataByBarcodeParams struct {
	// Shipment barcode
	Barcode string `json:"barcode"`
}

type GetShipmentDataByBarcodeResponse struct {
	core.CommonResponse

	// Method result
	Result GetShipmentDataByBarcodeResult `json:"result"`
}

type GetShipmentDataByBarcodeResult struct {
	// Analytical data
	AnalyticsData GetShipmentDataByBarcodeResultAnalyticsData `json:"analytics_data"`

	// Shipment barcodes
	Barcodes FBSBarcode `json:"barcodes"`

	// Cancellation reason identifier
	CancelReasonId int64 `json:"cancel_reason_id"`

	// Date and time when the shipment was created
	CreatedAt time.Time `json:"created_at"`

	// Financial data
	FinancialData GetShipmentDataByBarcodeResultFinancialData `json:"financial_data"`

	// Start date and time of shipment processing
	InProcessAt time.Time `json:"in_process_at"`

	// Order identifier to which the shipment belongs
	OrderId int64 `json:"order_id"`

	// Order number to which the shipment belongs
	OrderNumber string `json:"order_number"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// List of products in the shipment
	Products []PostingProduct `json:"products"`

	// Date and time before which the shipment must be packaged.
	// If the shipment is not packaged by this date, it will be canceled automatically
	ShipmentDate time.Time `json:"shipment_date"`

	// Shipment status
	Status string `json:"status"`
}

type GetShipmentDataByBarcodeResultAnalyticsData struct {
	// Delivery city
	City string `json:"city"`

	// Delivery method
	DeliveryType string `json:"delivery_type"`

	// Indication that the recipient is a legal entity:
	//   - true — a legal entity
	//   - false — a natural person
	IsLegal bool `json:"is_legal"`

	// Premium subscription availability
	IsPremium bool `json:"is_premium"`

	// Payment method
	PaymentTypeGroupName string `json:"payment_type_group_name"`

	// Delivery region
	Region string `json:"region"`
}

type GetShipmentDataByBarcodeResultFinancialData struct {
	// Identifier of the cluster, where the shipment is sent from
	ClusterFrom string `json:"cluster_from"`

	// Identifier of the cluster, where the shipment is delivered to
	ClusterTo string `json:"cluster_to"`

	// Services
	PostingServices []MarketplaceServices `json:"posting_services"`

	// Products list
	Products []FinancialDataProduct `json:"products"`
}

// Method for getting shipments data by barcode
func (c FBS) GetShipmentDataByBarcode(ctx context.Context, params *GetShipmentDataByBarcodeParams) (*GetShipmentDataByBarcodeResponse, error) {
	url := "/v2/posting/fbs/get-by-barcode"

	resp := &GetShipmentDataByBarcodeResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetShipmentDataByIdentifierParams struct {
	// Shipment identifier
	PostingNumber string `json:"posting_number"`

	// Additional fields that should be added to the response
	With *GetShipmentDataByIdentifierWith `json:"with,omitempty"`
}

type GetShipmentDataByIdentifierWith struct {
	// Add analytics data to the response
	AnalyticsData bool `json:"analytics_data"`

	// Add the shipment barcodes to the response
	Barcodes bool `json:"barcodes"`

	// Add financial data to the response
	FinancialData bool `json:"financial_data"`

	// Add data on products and their instances to the response
	ProductExemplars bool `json:"product_exemplars"`

	// Add related shipment numbers to the response.
	// Related shipments are ones into which the parent shipment was split during packaging
	RelatedPostings bool `json:"related_postings"`

	// Transliterate the return values
	Translit bool `json:"translit"`
}

type GetShipmentDataByIdentifierResponse struct {
	core.CommonResponse

	// Method result
	Result GetShipmentDataByIdentifierResult `json:"result"`
}

type GetShipmentDataByIdentifierResult struct {
	// Additional Data Key-Value
	AdditionalData []GetShipmentDataByIdentifierResultAdditionalData `json:"additional_data"`

	// Recipient details
	Addressee GetShipmentDataByIdentifierResultAddressee `json:"addressee"`

	// Analytics data
	AnalyticsData GetShipmentDataByIdentifierResultAnalyticsData `json:"analytics_data"`

	// Available actions and shipment information
	AvailableActions []string `json:"available_actions"`

	// Shipment barcodes
	Barcodes FBSBarcode `json:"barcodes"`

	// Cancellation details
	Cancellation FBSCancellation `json:"cancellation"`

	// Courier information
	Courier GetShipmentDataByIdentifierResultCourier `json:"courier"`

	// Customer details
	Customer FBSCustomer `json:"customer"`

	// Date when the shipment was transferred for delivery
	DeliveringDate time.Time `json:"delivering_date"`

	// Delivery method
	DeliveryMethod FBSDeliveryMethod `json:"delivery_method"`

	// Delivery cost
	DeliveryPrice string `json:"delivery_price"`

	// Data on the product cost, discount amount, payout and commission
	FinancialData FBSFinancialData `json:"financial_data"`

	// Start date and time of shipment processing
	InProcessAt time.Time `json:"in_process_at"`

	// If Ozon Express fast delivery was used—true
	IsExpress bool `json:"is_express"`

	// Indication that there is a multi-box product in the shipment and you need to pass the number of boxes for it:
	//   - true — before packaging pass the number of boxes using the /v3/posting/multiboxqty/set method.
	//   - false — you packed the shipment specifying the number of boxes in the multi_box_qty parameter, or there is no multi-box product in the shipment
	IsMultibox bool `json:"is_multibox"`

	// Number of boxes in which the product is packed
	MultiBoxQuantity int32 `json:"multi_box_qty"`

	// Order identifier to which the shipment belongs
	OrderId int64 `json:"order_id"`

	// Order number to which the shipment belongs
	OrderNumber string `json:"order_number"`

	// Number of the parent shipment which split resulted in the current shipment
	ParentPostingNumber string `json:"parent_posting_number"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Date and time of successful verification of the courier code
	PickupCodeVerifiedAt time.Time `json:"pickup_code_verified_at"`

	// Information on products and their instances.
	//
	// The response contains the field product_exemplars, if the attribute with.product_exemplars = true is passed in the request
	ProductExemplars GetShipmentDataByIdentifierResultExemplars `json:"product_exemplars"`

	// Array of products in the shipment
	Products []ShipmentProduct `json:"products"`

	// Delivery service status
	ProviderStatus string `json:"provider_status"`

	// Previous sub-status of the shipment
	PreviousSubstatus string `json:"previous_substatus"`

	// Information on lifting service. Only relevant for bulky products
	// with a delivery by a third-party or integrated service
	PRROption GetShipmentDataByIdentifierResultPRROption `json:"prr_option"`

	// Related shipments
	RelatedPostings GetShipmentDataByIdentifierResultRelatedPostings `json:"related_postings"`

	// Array of Ozon Product IDs (SKU) for which you need to pass the customs cargo declaration (CCD) number, the manufacturing country,
	// product batch registration number, or "Chestny ZNAK" labeling to change the shipment status to the next one
	Requirements FBSRequirements `json:"requirements"`

	// Date and time before which the shipment must be packaged.
	// If the shipment is not packaged by this date, it will be canceled automatically
	ShipmentDate time.Time `json:"shipment_date"`

	// Shipment status
	Status ShipmentStatus `json:"status"`

	// Shipment substatus
	Substatus ShipmentSubstatus `json:"substatus"`

	// Type of integration with the delivery service
	TPLIntegrationType TPLIntegrationType `json:"tpl_integration_type"`

	// Shipment tracking number
	TrackingNumber string `json:"tracking_number"`

	// Details on shipping rate
	Tariffication []FBSPostingTariffication `json:"tariffication"`
}

type GetShipmentDataByIdentifierResultAdditionalData struct {
	// Key
	Key string `json:"key"`

	// value
	Value string `json:"value"`
}

type GetShipmentDataByIdentifierResultAddressee struct {
	// Recipient name
	Name string `json:"name"`

	// Recipient phone number.
	//
	// Returns an empty string ""
	Phone string `json:"phone"`
}

type GetShipmentDataByIdentifierResultAnalyticsData struct {
	// Delivery city. Only for rFBS shipments
	City string `json:"city"`

	// Delivery start date and time
	DeliveryDateBegin time.Time `json:"delivery_date_begin"`

	// Delivery end date and time
	DeliveryDateEnd time.Time `json:"delivery_date_end"`

	// Delivery method
	DeliveryType string `json:"delivery_type"`

	// Indication that the recipient is a legal entity:
	//   - true — a legal entity,
	//   - false — a natural person
	IsLegal bool `json:"is_legal"`

	// Premium subscription availability
	IsPremium bool `json:"is_premium"`

	// Payment method
	PaymentTypeGroupName string `json:"payment_type_group_name"`

	// Delivery region. Only for rFBS shipments
	Region string `json:"region"`

	// Delivery service
	TPLProvider string `json:"tpl_provider"`

	// Delivery service identifier
	TPLProviderId int64 `json:"tpl_provider_id"`

	// Order shipping warehouse name
	Warehouse string `json:"warehouse"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

type GetShipmentDataByIdentifierResultCourier struct {
	// Car model
	CarModel string `json:"car_model"`

	// Car number
	CarNumber string `json:"car_number"`

	// Courier's full name
	Name string `json:"name"`

	// Courier's phone number
	//
	// Returns an empty string ""
	Phone string `json:"phone"`
}

type GetShipmentDataByIdentifierResultExemplars struct {
	// Products
	Products []GetShipmentDataByIdentifierResultExemplarsProduct `json:"products"`
}

type GetShipmentDataByIdentifierResultExemplarsProduct struct {
	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Array of exemplars
	Exemplars []FBSProductExemplar `json:"exemplars"`
}

type GetShipmentDataByIdentifierResultPRROption struct {
	// Lifting service code
	Code PRROptionStatus `json:"code"`

	// Service cost, which Ozon reimburses to the seller
	Price string `json:"price"`

	// Currency
	CurrencyCode string `json:"currency_code"`

	// Floor to which you need to lift the product
	Floor string `json:"floor"`
}

type GetShipmentDataByIdentifierResultRelatedPostings struct {
	RelatedPostingNumbers []string `json:"related_posting_numbers"`
}

type ShipmentProduct struct {
	PostingProduct

	// Product dimensions
	Dimensions ProductDimension `json:"dimensions"`
}

type ProductDimension struct {
	// Package height
	Height string `json:"height"`

	// Product length
	Length string `json:"length"`

	// Weight of product in the package
	Weight string `json:"weight"`

	// Package width
	Width string `json:"width"`
}

type FBSProductExemplar struct {
	// Item identifier
	ExemplarId int64 `json:"exemplar_id"`

	// Mandatory “Chestny ZNAK” labeling
	MandatoryMark string `json:"mandatory_mark"`

	// Сustoms cargo declaration (CCD) number
	GTD string `json:"gtd"`

	// Indication that a сustoms cargo declaration (CCD) number hasn't been specified
	IsGTDAbsest bool `json:"is_gtd_absent"`

	// Product batch registration number
	RNPT string `json:"rnpt"`

	// Indication that a product batch registration number hasn't been specified
	IsRNPTAbsent bool `json:"is_rnpt_absent"`
}

// Method for getting shipment details by identifier
func (c FBS) GetShipmentDataByIdentifier(ctx context.Context, params *GetShipmentDataByIdentifierParams) (*GetShipmentDataByIdentifierResponse, error) {
	url := "/v3/posting/fbs/get"

	resp := &GetShipmentDataByIdentifierResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type AddTrackingNumbersParams struct {
	// An array with shipment identifier—tracking number pairs
	TrackingNumbers []FBSTrackingNumbersParams `json:"tracking_numbers"`
}

type FBSTrackingNumbersParams struct {
	// Shipment identifier
	PostingNumber string `json:"posting_number"`

	// Shipment tracking number
	TrackingNumber string `json:"tracking_number"`
}

type AddTrackingNumbersResponse struct {
	core.CommonResponse

	// Method result
	Result []AddTrackingNumbersResponseResult `json:"result"`
}

type AddTrackingNumbersResponseResult struct {
	// Error when processing the request
	Error []string `json:"error"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// If the request is executed without errors — true
	Result bool `json:"result"`
}

// Add tracking numbers to shipments
func (c FBS) AddTrackingNumbers(ctx context.Context, params *AddTrackingNumbersParams) (*AddTrackingNumbersResponse, error) {
	url := "/v2/fbs/posting/tracking-number/set"

	resp := &AddTrackingNumbersResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListOfShipmentCertificatesParams struct {
	// Filter parameters
	Filter ListOfShipmentCertificatesFilter `json:"filter"`

	// Maximum number of certificates in the response
	Limit int64 `json:"limit"`
}

type ListOfShipmentCertificatesFilter struct {
	// Initial date of shipment creation
	DateFrom string `json:"date_from"`

	// Final date of shipment creation
	DateTo string `json:"date_to"`

	// Type of integration with the delivery service:
	//   - ozon — delivery by the Ozon service,
	//   - 3pl_tracking — delivery by the integrated service,
	//   - non_integrated — delivery by a third-party service,
	//   - aggregator — delivery by Ozon partner delivery
	IntegrationType string `json:"integration_type"`

	// Freight statuses
	Status []ShipmentCertificateFilterStatus `json:"status"`
}

type ListOfShipmentCertificatesResponse struct {
	core.CommonResponse

	// Request result
	Result []ListOfShipmentCertificatesResult `json:"result"`
}

type ListOfShipmentCertificatesResult struct {
	// Shipment identifier
	Id int64 `json:"id"`

	// Delivery method identifier
	DeliveryMethodId int64 `json:"delivery_method_id"`

	// Delivery method name
	DeliveryMethodName string `json:"delivery_method_name"`

	// Type of integration with the delivery service:
	//   - ozon — delivery by the Ozon service,
	//   - 3pl — delivery by the integrated service
	IntegrationType string `json:"integration_type"`

	// Number of package units
	ContainersCount int32 `json:"containers_count"`

	// Shipping status
	Status string `json:"status"`

	// Shipping date
	DepartureDate string `json:"departure_date"`

	// Shipping record creation date
	CreatedAt time.Time `json:"created_at"`

	// Shipping record update date
	UpdatedAt time.Time `json:"updated_at"`

	// Acceptance certificate type for FBS sellers
	ActType string `json:"act_type"`

	// Indication of a partial freight. true if the freigth is partial.
	//
	// Partial freigt means that the shipment was splitted into several parts
	// and for each of them you need to generate separate acts
	IsPartial bool `json:"is_partial"`

	// Indication that there are shipments subject to shipping that are not in the current freight.
	// true if there are such shipments
	HasPostingsForNextCarriage bool `json:"has_postings_for_next_carriage"`

	// Serial number of the partial freight
	PartialNum int64 `json:"partial_num"`

	// Information about shipment certificates
	RelatedDocs ListOfShipmentCertificatesResultDocs `json:"related_docs"`
}

type ListOfShipmentCertificatesResultDocs struct {
	// Information about acceptance certificate
	ActOfAcceptance FBSAct `json:"act_of_acceptance"`

	// Information about discrepancy certificate
	ActOfMismatch FBSAct `json:"act_of_mismatch"`

	// Information about surplus certificate
	ActOfExcess FBSAct `json:"act_of_excess"`
}

type FBSAct struct {
	// Certificate creation date
	CreatedAt time.Time `json:"created_at"`

	// Certificate status:
	//   - FORMING — not ready yet
	//   - FORMED — formed
	//   - CONFIRMED — signed by Ozon
	//   - CONFIRMED_WITH_MISMATCH — signed with discrepancies by Ozon
	//   - ACCEPTED_BY_CARGO_PLACES — accepted by package units
	//   - PRINTED_CARRIAGE — digital signature not required
	//   - ERROR, UNKNOWN_ERROR — error
	DocumentStatus string `json:"document_status"`
}

// Returns a list of shipment certificates allowing to filter them by time period, status, and integration type
func (c FBS) ListOfShipmentCertificates(ctx context.Context, params *ListOfShipmentCertificatesParams) (*ListOfShipmentCertificatesResponse, error) {
	url := "/v2/posting/fbs/act/list"

	resp := &ListOfShipmentCertificatesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ChangeStatusToParams struct {
	// Shipment identifier
	PostingNumber []string `json:"posting_number"`
}

type ChangeStatusToResponse struct {
	core.CommonResponse

	// Method result
	Result []ChangeStatusToResponseResult `json:"result"`
}

type ChangeStatusToResponseResult struct {
	// Error when processing the request
	Error []string `json:"error"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// If the request is executed without errors — true
	Result bool `json:"result"`
}

// Changes the shipment status to "Delivering" if a third-party delivery service is being used
func (c FBS) ChangeStatusToDelivering(ctx context.Context, params *ChangeStatusToParams) (*ChangeStatusToResponse, error) {
	url := "/v2/fbs/posting/delivering"

	resp := &ChangeStatusToResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

// Changes the shipment status to "Last mile" if a third-party delivery service is being used
func (c FBS) ChangeStatusToLastMile(ctx context.Context, params *ChangeStatusToParams) (*ChangeStatusToResponse, error) {
	url := "/v2/fbs/posting/last-mile"

	resp := &ChangeStatusToResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

// Changes the shipment status to "Delivered" if a third-party delivery service is being used
func (c FBS) ChangeStatusToDelivered(ctx context.Context, params *ChangeStatusToParams) (*ChangeStatusToResponse, error) {
	url := "/v2/fbs/posting/delivered"

	resp := &ChangeStatusToResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

// Change shipment status to "Sent by seller". Status is only available to sellers with a first mile selling from abroad
func (c FBS) ChangeStatusToSendBySeller(ctx context.Context, params *ChangeStatusToParams) (*ChangeStatusToResponse, error) {
	url := "/v2/fbs/posting/sent-by-seller"

	resp := &ChangeStatusToResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type PassShipmentToShippingParams struct {
	// Shipment identifier. The maximum number of values in one request is 100
	PostingNumber []string `json:"posting_number"`
}

type PassShipmentToShippingResponse struct {
	core.CommonResponse

	// Request processing result. true, if the request was executed without errors
	Result bool `json:"result"`
}

// Transfers disputed orders to shipping. The shipment status will change to `awaiting_deliver`
func (c FBS) PassShipmentToShipping(ctx context.Context, params *PassShipmentToShippingParams) (*PassShipmentToShippingResponse, error) {
	url := "/v2/posting/fbs/awaiting-delivery"

	resp := &PassShipmentToShippingResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CancelShipmentParams struct {
	// Shipment cancellation reason identifier
	CancelReasonId int64 `json:"cancel_reason_id"`

	// Additional information on cancellation. If `cancel_reason_id` = 402, the parameter is required
	CancelReasonMessage string `json:"cancel_reason_message"`

	// Shipment identifier
	PostingNumber string `json:"posting_number"`
}

type CancelShipmentResponse struct {
	core.CommonResponse

	// Request processing result. true, if the request was executed without errors
	Result bool `json:"result"`
}

// Change shipment status to `cancelled`.
//
// If you are using the rFBS scheme, you have the following cancellation reason identifiers (`cancel_reason_id`) available:
//
// 352—product is out of stock ;
// 400—only defective products left;
// 401—cancellation from arbitration;
// 402—other reason;
// 665—the customer didn't pick the order;
// 666—delivery isn't available in the region;
// 667—order was lost by the delivery service.
// The last 4 reasons are available for shipments in the "Delivering" and "Courier on the way" statuses.
//
// You can't cancel presumably delivered orders.
//
// If `cancel_reason_id` parameter value is 402, fill the `cancel_reason_message` field.
func (c FBS) CancelShipment(ctx context.Context, params *CancelShipmentParams) (*CancelShipmentResponse, error) {
	url := "/v2/posting/fbs/cancel"

	resp := &CancelShipmentResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CreateActParams struct {
	// Number of package units.
	//
	// Use this parameter if you have trusted acceptance enabled and ship orders by package units.
	// If you do not have trusted acceptance enabled, skip it
	ContainersCount int32 `json:"containers_count"`

	// Delivery method identifier
	DeliveryMethodId int64 `json:"delivery_method_id"`

	// Shipping date.
	//
	// To make documents printing available before the shipping day,
	// enable Printing the acceptance certificate in advance in your personal account under the method settings.
	// The time for packaging (packaging SLA) should be more than 13 hours
	DepartureDate time.Time `json:"departure_date"`
}

type CreateActResponse struct {
	core.CommonResponse

	// Method result
	Result CreateActResult `json:"result"`
}

type CreateActResult struct {
	// Document generation task number
	Id int64 `json:"id"`
}

// Launches the procedure for generating the transfer documents: acceptance and transfer certificate and the waybill.
//
// To generate and receive transfer documents, transfer the shipment to the `awaiting_deliver` status
func (c FBS) CreateAct(ctx context.Context, params *CreateActParams) (*CreateActResponse, error) {
	url := "/v2/posting/fbs/act/create"

	resp := &CreateActResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetLabelingParams struct {
	// Task identifier for labeling generation from the /v1/posting/fbs/package-label/create method response
	TaskId int64 `json:"task_id"`
}

type GetLabelingResponse struct {
	core.CommonResponse

	// Method result
	Result GetLabelingResult `json:"result"`
}

type GetLabelingResult struct {
	// Error code
	Error string `json:"error"`

	// Link to a labeling file
	FileUrl string `json:"file_url"`

	// Status of labeling generation:
	//   - pending — task is in the queue.
	//   - in_progress — being generated.
	//   - completed — labeling file is ready.
	//   - error — error occurred during file generation
	Status string `json:"status"`
}

// Method for getting labeling after using the /v1/posting/fbs/package-label/create method
func (c FBS) GetLabeling(ctx context.Context, params *GetLabelingParams) (*GetLabelingResponse, error) {
	url := "/v1/posting/fbs/package-label/get"

	resp := &GetLabelingResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, map[string]string{
		"Content-Type": "application/pdf",
	})
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type PrintLabelingParams struct {
	// Shipment identifier
	PostingNumber []string `json:"posting_number"`
}

type PrintLabelingResponse struct {
	core.CommonResponse

	// Order content
	Content string `json:"file_content"`

	// File name
	FileName string `json:"file_name"`

	// File type
	ContentType string `json:"content_type"`
}

// Generates a PDF file with a labeling for the specified shipments. You can pass a maximum of 20 identifiers in one request.
// If an error occurs for at least one shipment, the labeling will not be generated for all shipments in the request.
//
// We recommend requesting the label 45-60 seconds after the shipments were packed.
//
// The next postings are not ready error means that the label is not ready. Try again later
func (c FBS) PrintLabeling(ctx context.Context, params *PrintLabelingParams) (*PrintLabelingResponse, error) {
	url := "/v2/posting/fbs/package-label"

	resp := &PrintLabelingResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, map[string]string{
		"Content-Type": "application/pdf",
	})
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CreateTaskForGeneratingLabelParams struct {
	// Numbers of shipments that need labeling
	PostingNumber []string `json:"posting_number"`
}

type CreateTaskForGeneratingLabelResponse struct {
	core.CommonResponse

	// Method result
	Result CreateTaskForGeneratingLabelResult `json:"result"`
}

type CreateTaskForGeneratingLabelResult struct {
	Tasks []CreateTaskForGeneratingLabelTask `json:"tasks"`
}

type CreateTaskForGeneratingLabelTask struct {
	// Task identifier for labeling generation
	TaskId int64 `json:"task_id"`

	// Type of label generation task:
	//
	// 	big_label — for a regular label,
	// 	small_label — for a small label
	TaskType string `json:"task_type"`
}

// Method for creating a task for asynchronous labeling generation.
//
// To get labels created as a result of the method, use the /v1/posting/fbs/package-label/get method
func (c FBS) CreateTaskForGeneratingLabel(ctx context.Context, params *CreateTaskForGeneratingLabelParams) (*CreateTaskForGeneratingLabelResponse, error) {
	url := "/v2/posting/fbs/package-label/create"

	resp := &CreateTaskForGeneratingLabelResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, map[string]string{
		"Content-Type": "application/pdf",
	})
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetDropOffPointRestrictionsParams struct {
	// The number of shipment for which you want to determine the restrictions
	PostingNumber string `json:"posting_number"`
}

type GetDropOffPointRestrictionsResponse struct {
	core.CommonResponse

	// Method result
	Result GetDropOffPointRestrictionsResult `json:"result"`
}

type GetDropOffPointRestrictionsResult struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Maximum weight limit in grams
	MaxPostingWeight float64 `json:"max_posting_weight"`

	// Minimum weight limit in grams
	MinPostingWeight float64 `json:"min_posting_weight"`

	// Width limit in centimeters
	Width float64 `json:"width"`

	// Length limit in centimeters
	Length float64 `json:"length"`

	// Height limit in centimeters
	Height float64 `json:"height"`

	// Maximum shipment cost limit in rubles
	MaxPostingPrice float64 `json:"max_posting_price"`

	// Minimum shipment cost limit in rubles
	MinPostingPrice float64 `json:"min_posting_price"`
}

// Method for getting dimensions, weight, and other restrictions of the drop-off point by the shipment number.
// The method is applicable only for the FBS scheme
func (c FBS) GetDropOffPointRestrictions(ctx context.Context, params *GetDropOffPointRestrictionsParams) (*GetDropOffPointRestrictionsResponse, error) {
	url := "/v1/posting/fbs/restrictions"

	resp := &GetDropOffPointRestrictionsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type SetProductItemsDataParams struct {
	// Quantity of boxes the product is packed in
	MultiBoxQuantity int32 `json:"multi_box_qty"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Product list
	Products []SetProductItemsDataProduct `json:"products"`
}

type SetProductItemsDataProduct struct {
	// Product items data
	Exemplars []SetProductItemsDataProductExemplar `json:"exemplars"`

	// Indication that you need to pass the сustoms cargo declaration
	// (CCD) number for the product and shipment
	IsGTDNeeded bool `json:"is_gtd_needed"`

	// Indication that you need to pass the unique identifier of charges of the jewelry
	IsJwUINNeeded bool `json:"is_jw_uin_needed"`

	// Indication that you need to pass the "Chestny ZNAK" labeling
	IsMandatoryMarkNeeded bool `json:"is_mandatory_mark_needed"`

	// Indication that you can pass the "Chestny ZNAK" labeling, but it's not mandatory
	IsMandatoryMarkPossible bool `json:"is_mandatory_mark_possible"`

	// Indication that you need to pass the product batch registration number
	IsRNPTNeeded bool `json:"is_rnpt_needed"`

	// Product ID
	ProductId int64 `json:"product_id"`

	// Items quantity
	Quantity int32 `json:"quantity"`
}

type SetProductItemsDataProductExemplar struct {
	// Item identifier
	ExemplarId int64 `json:"exemplar_id"`

	// Customs cargo declaration (CCD) number
	GTD string `json:"gtd"`

	// Indication that the customs cargo declaration (CCD) number isn't specified
	IsGTDAbsent bool `json:"is_gtd_absent"`

	// Indication that the product batch registration number isn't specified
	IsRNPTAbsent bool `json:"is_rnpt_absent"`

	// Product batch registration number
	RNPT string `json:"rnpt"`

	// Errors that appeared during verification of Control Identification Marks
	Marks []SetProductItemsDataProductMark `json:"marks"`
}

type SetProductItemsDataProductMark struct {
	// Labeling code meaning
	Mark string `json:"mark"`

	// Labeling code type
	MarkType string `json:"mark_type"`
}

type SetProductItemsDataResponse struct {
	core.CommonResponse

	// Method result. true if the request was processed successfully
	Result bool `json:"result"`
}

// Asynchronous method:
//
// for checking the availability of product items in the “Chestny ZNAK” labeling system;
// for saving product items data.
// To get the checks results, use the /v5/fbs/posting/product/exemplar/status method. To get data about created items, use the /v6/fbs/posting/product/exemplar/create-or-get method.
//
// If you have multiple identical products in a shipment, specify one product_id and exemplars array for each product in the shipment.
//
// Always pass a complete set of product items data.
//
// For example, you have 10 product items in your system. You pass them for checking and saving. Then you add another 60 product items to your system. When you pass product items for checking and saving again, pass all of them: both old and newly added.
//
// The 200 response code doesn't guarantee that instance data has been received. It indicates that a task for adding the information has been created. To check the task status, use the /v5/fbs/posting/product/exemplar/status method.
func (c FBS) SetProductItemsData(ctx context.Context, params *SetProductItemsDataParams) (*SetProductItemsDataResponse, error) {
	url := "/v6/fbs/posting/product/exemplar/set"

	resp := &SetProductItemsDataResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductItemsCheckStatusesParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`
}

type GetProductItemsCheckStatusesResponse struct {
	core.CommonResponse

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Products list
	Products []GetProductItemsCheckStatusProduct `json:"products"`

	// Product items check statuses and order collection availability
	Status string `json:"status"`
}

type GetProductItemsCheckStatusProduct struct {
	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product items data
	Exemplars []GetProductItemsCheckStatusExemplar `json:"exemplars"`
}

type GetProductItemsCheckStatusExemplar struct {
	// Item identifier
	ExemplarId int64 `json:"exemplar_id"`

	// Customs cargo declaration (CCD) number
	GTD string `json:"gtd"`

	// Сustoms cargo declaration (CCD) check status
	GTDCheckStatus string `json:"gtd_check_status"`

	// Сustoms cargo declaration (CCD) check error codes
	GTDErrorCodes []string `json:"gtd_error_codes"`

	// Indication that the customs cargo declaration (CCD) number isn't specified
	IsGTDAbsent bool `json:"is_gtd_absent"`

	// Indication that the product batch registration number isn't specified
	IsRNPTAbsent bool `json:"is_rnpt_absent"`

	// List of Control Identification Marks in one copy
	Marks []GetProductItemsCheckStatusMark `json:"marks"`

	// Product batch registration number
	RNPT string `json:"rnpt"`

	// Product batch registration number check status
	RNPTCheckStatus string `json:"rnpt_check_status"`

	// Product batch registration number check error codes
	RNPTErrorCodes []string `json:"rnpt_error_codes"`
}

type GetProductItemsCheckStatusMark struct {
	// Check status
	CheckStatus string `json:"check_status"`

	// Errors that appeared during verification of Control Identification Marks
	ErrorCodes []string `json:"error_codes"`

	// Labeling code meaning
	Mark string `json:"mark"`

	// Labeling code type
	MarkType string `json:"mark_type"`
}

// Method for getting product items addition statuses
// that were passed in the /v6/fbs/posting/product/exemplar/set method.
// Also returns data on these product items.
func (c FBS) GetProductItemsCheckStatuses(ctx context.Context, params *GetProductItemsCheckStatusesParams) (*GetProductItemsCheckStatusesResponse, error) {
	url := "/v5/fbs/posting/product/exemplar/status"

	resp := &GetProductItemsCheckStatusesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type RescheduleShipmentDeliveryDateParams struct {
	// New delivery date period
	NewTimeslot RescheduleShipmentDeliveryDateTimeslot `json:"new_timeslot"`

	// Shipment number
	PostingNumber string `json:"posting_number"`
}

type RescheduleShipmentDeliveryDateTimeslot struct {
	// Period start date
	From time.Time `json:"from"`

	// Period end date
	To time.Time `json:"to"`
}

type RescheduleShipmentDeliveryDateResponse struct {
	core.CommonResponse

	// true, if the date was changed
	Result bool `json:"result"`
}

// You can change the delivery date of a shipment up to two times
func (c FBS) RescheduleShipmentDeliveryDate(ctx context.Context, params *RescheduleShipmentDeliveryDateParams) (*RescheduleShipmentDeliveryDateResponse, error) {
	url := "/v1/posting/fbs/timeslot/set"

	resp := &RescheduleShipmentDeliveryDateResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type DateAvailableForDeliveryScheduleParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`
}

type DateAvailableForDeliveryScheduleResponse struct {
	core.CommonResponse

	// Number of delivery date reschedules made
	AvailableChangecount int64 `json:"available_change_count"`

	// Period of dates available for reschedule
	DeliveryInterval DateAvailableForDeliveryScheduleDeliveryInterval `json:"delivery_interval"`

	// Number of delivery date reschedules left
	RemainingChangeCount int64 `json:"remaining_change_count"`
}

type DateAvailableForDeliveryScheduleDeliveryInterval struct {
	// Period start date
	Begin time.Time `json:"begin"`

	// Period end date
	End time.Time `json:"end"`
}

// Method for getting the dates and number of times available for delivery reschedule
func (c FBS) DateAvailableForDeliverySchedule(ctx context.Context, params *DateAvailableForDeliveryScheduleParams) (*DateAvailableForDeliveryScheduleResponse, error) {
	url := "/v1/posting/fbs/timeslot/change-restrictions"

	resp := &DateAvailableForDeliveryScheduleResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListManufacturingCountriesParams struct {
	// Filtering by line
	NameSearch string `json:"name_search"`
}

type ListManufacturingCountriesResponse struct {
	core.CommonResponse

	// List of manufacturing countries and their ISO codes
	Result []ListManufacturingCountriesResult `json:"result"`
}

type ListManufacturingCountriesResult struct {
	// Country name in Russian
	Name string `json:"name"`

	// Country ISO code
	CountriISOCode string `json:"country_iso_code"`
}

// Method for getting a list of available manufacturing countries and their ISO codes
func (c FBS) ListManufacturingCountries(ctx context.Context, params *ListManufacturingCountriesParams) (*ListManufacturingCountriesResponse, error) {
	url := "/v2/posting/fbs/product/country/list"

	resp := &ListManufacturingCountriesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type SetManufacturingCountryParams struct {
	// Shipment identifier
	PostingNumber string `json:"posting_number"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Country ISO code from the `/v2/posting/fbs/product/country/list` method response
	CountryISOCode string `json:"country_iso_code"`
}

type SetManufacturingCountryResponse struct {
	core.CommonResponse

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Indication that you need to pass the сustoms cargo declaration (CCD) number for the product and shipment
	IsGTDNeeded bool `json:"is_gtd_needed"`
}

// The method to set the manufacturing country to the product if it hasn't been specified
func (c FBS) SetManufacturingCountry(ctx context.Context, params *SetManufacturingCountryParams) (*SetManufacturingCountryResponse, error) {
	url := "/v2/posting/fbs/product/country/set"

	resp := &SetManufacturingCountryResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type PartialPackOrderParams struct {
	// Shipment ID
	PostingNumber string `json:"posting_number"`

	// Array of products
	Products []PartialPackOrderProduct `json:"products"`
}

type PartialPackOrderProduct struct {
	// Product item identifiers
	ExemplarIds []string `json:"exemplarIds"`

	// FBS product identifier in the Ozon system, SKU
	ProductId int64 `json:"product_id"`

	// Product quantity
	Quantity int32 `json:"quantity"`
}

type PartialPackOrderResponse struct {
	core.CommonResponse

	// Shipments numbers formed after packaging
	Result string `json:"result"`
}

type PartialPackOrderAdditionalData struct {
	// Shipment identifier
	PostingNumber string `json:"posting_number"`

	// List of products in the shipment
	Products []PostingProduct `json:"products"`
}

// If you pass some of the shipped products through the request,
// the primary shipment will split into two parts.
// The primary unassembled shipment will contain some of the products
// that weren't passed to the request.
//
// Default status of created shipments is awaiting_packaging,
// which indicates that the shipment is awaiting assembly.
//
// The status of the original shipment will only change
// when the split shipments status changes
func (c FBS) PartialPackOrder(ctx context.Context, params *PartialPackOrderParams) (*PartialPackOrderResponse, error) {
	url := "/v4/posting/fbs/ship/package"

	resp := &PartialPackOrderResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type AvailableFreightsListParams struct {
	// Filter by delivery method identifier
	DeliveryMethodId int64 `json:"delivery_method_id"`

	// Shipping date. The default value is current date
	DepartureDate time.Time `json:"departure_date"`
}

type AvailableFreightsListResponse struct {
	core.CommonResponse

	// Method result
	Result []AvailableFreightsListResult `json:"result"`
}

type AvailableFreightsListResult struct {
	// Freight identifier (document generation task number)
	CarriageId int64 `json:"carriage_id"`

	// Number of shipments in the freight
	CarriagePostingsCount int32 `json:"carriage_postings_count"`

	// Freight status for requested delivery method and shipping date
	CarriageStatus string `json:"carriage_status"`

	// Date and time before a shipment must be packaged
	CutoffAt time.Time `json:"cutoff_at"`

	// Delivery method identifier
	DeliveryMethodId int64 `json:"delivery_method_id"`

	// Delivery method name
	DeliveryMethodName string `json:"delivery_method_name"`

	// Errors list
	Errors []AvailableFreightsListResultError `json:"errors"`

	// First mile type
	FirstMileType string `json:"first_mile_type"`

	// Trusted acceptance attribute. true if trusted acceptance is enabled in the warehouse
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`

	// Number of shipments to be packaged
	MandatoryPostingsCount int32 `json:"mandatory_postings_count"`

	// Number of already packaged shipments
	MandatoryPackagedCount int32 `json:"mandatory_packaged_count"`

	// Recommended local time of shipping to the pick-up point
	RecommendedTimeLocal string `json:"recommended_time_local"`

	// Time zone offset of the recommended shipping time from UTC-0 in minutes
	RecommendedTimeUTCOffset int32 `json:"recommended_time_utc_offset_in_minutes"`

	// Delivery service icon link
	TPLProviderIconURL string `json:"tpl_provider_icon_url"`

	// Delivery service name
	TPLProviderName string `json:"tpl_provider_name"`

	// Warehouse city
	WarehouseCity string `json:"warehouse_city"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`

	// Warehouse name
	WarehouseName string `json:"warehouse_name"`

	// Warehouse timezone
	WarehouseTimezone string `json:"warehouse_timezone"`
}

type AvailableFreightsListResultError struct {
	// Error code
	Code string `json:"code"`

	// Error type:
	//   - warning
	//   - critical
	Status string `json:"status"`
}

// Method for getting freights that require printing acceptance and transfer certificates and a waybill
func (c FBS) AvailableFreightsList(ctx context.Context, params *AvailableFreightsListParams) (*AvailableFreightsListResponse, error) {
	url := "/v1/posting/carriage-available/list"

	resp := &AvailableFreightsListResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GenerateActParams struct {
	// Document generation task number (freight identifier) received from the POST `/v2/posting/fbs/act/create` method
	Id int64 `json:"id"`
}

type GenerateActResponse struct {
	core.CommonResponse

	// Document generation task number
	Id int64 `json:"id"`

	// Documents generation status:
	//   - FORMING—in process,
	//   - FORMED—generated successfully,
	//   - CONFIRMED—signed by Ozon,
	//   - CONFIRMED_WITH_MISMATCH—signed by Ozon with mismatches,
	//   - NOT_FOUND—documents are not found,
	//   - UNKNOWN_ERROR—unknown error
	Status string `json:"status"`
}

// Get current status of generating digital acceptance and transfer certificate and waybill
func (c FBS) GenerateAct(ctx context.Context, params *GenerateActParams) (*GenerateActResponse, error) {
	url := "/v2/posting/fbs/digital/act/check-status"

	resp := &GenerateActResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetDigitalActParams struct {
	// Document generation task number (freight identifier) received from the POST `/v2/posting/fbs/act/create` method
	Id int64 `json:"id"`

	// Type of shipment certificate:
	//   - act_of_acceptance — acceptance certificate,
	//   - act_of_mismatch — discrepancy certificate,
	//   - act_of_excess — surplus certificate
	DocType DigitalActType `json:"doc_type"`
}

type GetDigitalActResponse struct {
	core.CommonResponse

	// File content in binary format
	Content string `json:"file_content"`

	// File name
	Name string `json:"file_name"`

	// File type
	Type string `json:"content_type"`
}

// Specify the type of a certificate in the doc_type parameter: `act_of_acceptance`, `act_of_mismatch`, `act_of_excess`
func (c FBS) GetDigitalAct(ctx context.Context, params *GetDigitalActParams) (*GetDigitalActResponse, error) {
	url := "/v2/posting/fbs/digital/act/get-pdf"

	resp := &GetDigitalActResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, map[string]string{
		"Content-Type": "application/json",
	})
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type PackageUnitLabelsParams struct {
	// Document generation task number (freight identifier) received from the POST `/v2/posting/fbs/act/create` method.
	Id int64 `json:"id"`
}

type PackageUnitLabelsResponse struct {
	core.CommonResponse

	// Content
	Content string `json:"content"`

	// Name
	Name string `json:"name"`

	// Type
	Type string `json:"type"`
}

// Method creates package unit labels
func (c FBS) PackageUnitLabel(ctx context.Context, params *PackageUnitLabelsParams) (*PackageUnitLabelsResponse, error) {
	url := "/v2/posting/fbs/act/get-container-labels"

	resp := &PackageUnitLabelsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, map[string]string{
		"Content-Type": "application/json",
	})
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type OpenDisputeOverShipmentParams struct {
	// Shipment identifier
	PostingNumber []string `json:"posting_number"`
}

type OpenDisputeOverShipmentResponse struct {
	core.CommonResponse

	// Request processing result. true, if the request was executed without errors
	Result bool `json:"result"`
}

// If the shipment has been handed over for delivery, but has not been scanned at the sorting center, you can open a dispute.
// Opened dispute will put the shipment into the `arbitration` status
func (c FBS) OpenDisputeOverShipment(ctx context.Context, params *OpenDisputeOverShipmentParams) (*OpenDisputeOverShipmentResponse, error) {
	url := "/v2/posting/fbs/arbitration"

	resp := &OpenDisputeOverShipmentResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ShipmentCancellationReasonsParams struct {
	// Shipment numbers
	RelatedPostingNumbers []string `json:"related_posting_numbers"`
}

type ShipmentCancellationReasonsResponse struct {
	core.CommonResponse

	// Request result
	Result []ShipmentCancellationReasonsResult `json:"result"`
}

type ShipmentCancellationReasonsResult struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Information about cancellation reasons
	Reasons []ShipmentCancellationReasonsResultReason `json:"reasons"`
}

type ShipmentCancellationReasonsResultReason struct {
	// Cancellation reasons
	Id int64 `json:"id"`

	// Reason description
	Title string `json:"title"`

	// Shipment cancellation initiator
	TypeId string `json:"type_id"`
}

// Returns a list of cancellation reasons for particular shipments
func (c FBS) ShipmentCancellationReasons(ctx context.Context, params *ShipmentCancellationReasonsParams) (*ShipmentCancellationReasonsResponse, error) {
	url := "/v1/posting/fbs/cancel-reason"

	resp := &ShipmentCancellationReasonsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ShipmentsCancellationReasonsResponse struct {
	core.CommonResponse

	// Method result
	Result []ShipmentsCancellatinoReasonsResult `json:"result"`
}

type ShipmentsCancellatinoReasonsResult struct {
	// Cancellation reason
	Id int64 `json:"id"`

	// Shipment cancellation result. true if the request is available for cancellation
	IsAvailableForCancellation bool `json:"is_available_for_cancellation"`

	// Category name
	Title string `json:"title"`

	// Shipment cancellation initiator:
	//   - buyer
	//   - seller
	TypeId string `json:"type_id"`
}

// Returns a list of cancellation reasons for particular shipments
func (c FBS) ShipmentsCancellationReasons(ctx context.Context) (*ShipmentsCancellationReasonsResponse, error) {
	url := "/v2/posting/fbs/cancel-reason/list"

	resp := &ShipmentsCancellationReasonsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type AddWeightForBulkProductParams struct {
	// Products information
	Items AddWeightForBulkProductItem `json:"items"`

	// Shipment identifier
	PostingNumber string `json:"posting_number"`
}

type AddWeightForBulkProductItem struct {
	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// List with weights of the products in the posting
	WeightReal []float64 `json:"weightReal"`
}

type AddWeightForBulkProductResponse struct {
	core.CommonResponse

	// Shipment identifier
	Result string `json:"result"`
}

// Add weight for bulk products in a shipment
func (c FBS) AddWeightForBulkProduct(ctx context.Context, params *AddWeightForBulkProductParams) (*AddWeightForBulkProductResponse, error) {
	url := "/v2/posting/fbs/product/change"

	resp := &AddWeightForBulkProductResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CancelSendingParams struct {
	// Product shipping cancellation reason identifier
	CancelReasonId int64 `json:"cancel_reason_id"`

	// Additional information on cancellation. Required parameter
	CancelReasonMessage string `json:"cancel_reason_message"`

	// Products information
	Items []CancelSendingItem `json:"items"`

	// Shipment identifier
	PostingNumber string `json:"posting_number"`
}

type CancelSendingItem struct {
	// Number of products in the shipment
	Quantity int32 `json:"quantity"`

	// Product identifier in the seller's system
	SKU int64 `json:"sku"`
}

type CancelSendingResponse struct {
	core.CommonResponse

	// Shipment number
	Result string `json:"result"`
}

// Use this method if you cannot send some of the products from the shipment.
//
// To get the cancel_reason_id cancellation reason identifiers
// when working with the FBS or rFBS schemes,
// use the /v2/posting/fbs/cancel-reason/list method.
//
// You can't cancel presumably delivered orders.
func (c FBS) CancelSending(ctx context.Context, params *CancelSendingParams) (*CancelSendingResponse, error) {
	url := "/v2/posting/fbs/product/cancel"

	resp := &CancelSendingResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListShipmentInCertificateParams struct {
	// Certificate identifier
	Id int64 `json:"id"`
}

type ListShipmentInCertificateResponse struct {
	core.CommonResponse

	// Information about shipments
	Result []ListShipmentInCertificateResult `json:"result"`
}

type ListShipmentInCertificateResult struct {
	// Certificate identifier
	Id int64 `json:"id"`

	// Number of boxes in which the product is packed
	MultiBoxQuantity int32 `json:"multi_box_qty"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Shipment status
	Status string `json:"status"`

	// Error code explanation
	SellerError string `json:"seller_error"`

	// Shipment record update date and time
	UpdatedAt time.Time `json:"updated_at"`

	// Shipment record creation date and time
	CreatedAt time.Time `json:"created_at"`

	// List of products in the shipment
	Products []ListShipmentInCertificateResultProduct `json:"products"`
}

type ListShipmentInCertificateResultProduct struct {
	// Product name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product price
	Price string `json:"price"`

	// Product number in the shipment
	Quantity int32 `json:"quantity"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

// Returns a list of shipments in the certificate by certificate identifier
func (c FBS) ListShipmentInCertificate(ctx context.Context, params *ListShipmentInCertificateParams) (*ListShipmentInCertificateResponse, error) {
	url := "/v2/posting/fbs/act/get-postings"

	resp := &ListShipmentInCertificateResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type SpecifyNumberOfBoxesParams struct {
	// Multi-box shipment identifier
	PostingNumber string `json:"posting_number"`

	// Number of boxes in which the product is packed
	MultiBoxQuantity int64 `json:"multi_box_qty"`
}

type SpecifyNumberOfBoxesResponse struct {
	core.CommonResponse

	// Result of transferring the boxes number
	Result SpecifyNumberOfBoxesResult `json:"result"`
}

type SpecifyNumberOfBoxesResult struct {
	// Possible values:
	//   - true — the number is successfully passed.
	//   - false — an error occurred while passing the number. Please try again
	Result bool `json:"result"`
}

// Method for passing the number of boxes for multi-box shipments when working under the rFBS Aggregator scheme (using the Ozon partner delivery)
func (c FBS) SpecifyNumberOfBoxes(ctx context.Context, params *SpecifyNumberOfBoxesParams) (*SpecifyNumberOfBoxesResponse, error) {
	url := "/v3/posting/multiboxqty/set"

	resp := &SpecifyNumberOfBoxesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type StatusOfActParams struct {
	// Document generation task number (freight identifier) received from the POST `/v2/posting/fbs/act/create` method
	Id int64 `json:"id"`
}

type StatusOfActResponse struct {
	core.CommonResponse

	// Method result
	Result StatusOfActResponseResult `json:"result"`
}

type StatusOfActResponseResult struct {
	// Acceptance and transfer certificate and a waybill type.
	//
	// If the value is ozon_digital,
	// use the `/v2/posting/fbs/digital/act/check-status` and `/v2/posting/fbs/digital/act/get-pdf` methods for getting digital acceptance
	// and transfer certificate and waybill
	ActType string `json:"act_type"`

	// List with numbers of shipments that are included in the acceptance and transfer certificate.
	// You should hand these shipments over today
	AddedToAct []string `json:"added_to_act"`

	// List with numbers of shipments that are not included in the acceptance and transfer certificate.
	// You should hand these shipments over in the next shipping
	RemovedFromAct []string `json:"removed_from_act"`

	// Request status:
	//
	//   - in_process — documents generation in process, please wait.
	//   - ready — documents are ready for downloading.
	//   - error — error occured during document geneartion process. Send a request again.
	//   - cancelled — documents generation was canceled. Send a request again.
	//   - The next postings are not ready — error occured, shipmants are not included in the shipping. Wait and check request results again. If you see the error again, contact our support team
	Status string `json:"status"`

	// Indication of a partial freight. true if the freigth is partial.
	//
	// Partial freigt means that the shipment was splitted into several parts and
	// for each of them you need to generate separate acceptance and transfer certificates
	IsPartial bool `json:"is_partial"`

	// Indication that there are shipments subject to shipping that are not in the current freight. true if there are such shipments.
	//
	// If there are such shipments, create a new acceptance and transfer certificate
	// using the `/v2/posting/fbs/act/create` method and check the creation status. Create acts until this field returns false
	HasPostingsForNextCarriage bool `json:"has_postings_for_next_carriage"`

	// Serial number of the partial freight
	PartialSum int64 `json:"partial_sum"`
}

// If you are not connected to electronic document circulation (EDC),
// the method returns status of generating an acceptance and transfer certificate and a waybill.
//
// If you are connected to EDC, the method returns status of generating a waybill only
func (c FBS) StatusOfAct(ctx context.Context, params *StatusOfActParams) (*StatusOfActResponse, error) {
	url := "/v2/posting/fbs/act/check-status"

	resp := &StatusOfActResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ETGBCustomsDeclarationsParams struct {
	// Filter by period of declaration creation
	Date ETGBCustomsDeclarationsDate `json:"date"`
}

type ETGBCustomsDeclarationsDate struct {
	// Start date
	From time.Time `json:"from"`

	// End date
	To time.Time `json:"to"`
}

type ETGBCustomsDeclarationsResponse struct {
	core.CommonResponse

	// Request result
	Result []ETGBCustomDeclarationsResult `json:"result"`
}

type ETGBCustomDeclarationsResult struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Declaration information
	ETGB ETGBCustomDeclarationsResultETGB `json:"etgb"`
}

type ETGBCustomDeclarationsResultETGB struct {
	// Number
	Number string `json:"number"`

	// Creation date
	Date string `json:"date"`

	// Link to file.
	//
	// If the field is empty and you need the file, contact Ozon support
	URL string `json:"url"`
}

// Method for getting Elektronik Ticaret Gümrük Beyannamesi (ETGB) customs declarations for sellers from Turkey
func (c FBS) ETGBCustomsDeclarations(ctx context.Context, params *ETGBCustomsDeclarationsParams) (*ETGBCustomsDeclarationsResponse, error) {
	url := "/v1/posting/global/etgb"

	resp := &ETGBCustomsDeclarationsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type BarcodeFromProductShipmentParams struct {
	// Freight identifier
	Id int64 `json:"id"`
}

type BarcodeFromProductShipmentResponse struct {
	core.CommonResponse

	// Link to barcode image
	Content string `json:"file_content"`

	// File name
	Name string `json:"file_name"`

	// File type
	Type string `json:"content_type"`
}

// Method for getting a barcode to show at a pick-up point or sorting center during the shipment
func (c FBS) BarcodeFromProductShipment(ctx context.Context, params *BarcodeFromProductShipmentParams) (*BarcodeFromProductShipmentResponse, error) {
	url := "/v2/posting/fbs/act/get-barcode"

	resp := &BarcodeFromProductShipmentResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, map[string]string{
		"Content-Type": "image/png",
	})
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type BarcodeValueFromProductShipmentParams struct {
	// Freight identifier
	Id int64 `json:"id"`
}

type BarcodeValueFromProductShipmentResponse struct {
	core.CommonResponse

	// Barcode in text format
	Result string `json:"result"`
}

// Use this method to get the barcode from the /v2/posting/fbs/act/get-barcode response in text format.
func (c FBS) BarcodeValueFromProductShipment(ctx context.Context, params *BarcodeValueFromProductShipmentParams) (*BarcodeValueFromProductShipmentResponse, error) {
	url := "/v2/posting/fbs/act/get-barcode/text"

	resp := &BarcodeValueFromProductShipmentResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetActPDFParams struct {
	// Document generation task number (freight identifier) received from
	// the POST `/v2/posting/fbs/act/create` method
	Id int64 `json:"id"`
}

type GetActPDFResponse struct {
	core.CommonResponse

	// File content in binary format
	FileContent string `json:"file_content"`

	// File name
	FileName string `json:"file_name"`

	// File type
	ContentType string `json:"content_type"`
}

// Get the generated transfer documents in PDF format: an acceptance and transfer certificate and a waybill.
//
// If you are not connected to electronic document circulation (EDC), the method returns an acceptance and transfer certificate and a waybill.
//
// If you are connected to EDC, the method returns a waybill only.
func (c FBS) GetActPDF(ctx context.Context, params *GetActPDFParams) (*GetActPDFResponse, error) {
	url := "/v2/posting/fbs/act/get-pdf"

	resp := &GetActPDFResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, map[string]string{
		"Content-Type": "application/pdf",
	})
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CreateOrGetProductExemplarParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`
}

type CreateOrGetProductExemplarResponse struct {
	core.CommonResponse

	// Quantity of boxes the product is packed in
	MultiBoxQuantity int32 `json:"multi_box_qty"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Product list
	Products []CheckProductItemsDataProduct `json:"products"`
}

type CheckProductItemsDataProduct struct {
	// Data about items
	Exemplars []SetProductItemsDataProductExemplar `json:"exemplars"`

	// Indication that you need to pass the сustoms cargo declaration (CCD) number for the product and shipment
	IsGTDNeeded bool `json:"is_gtd_needed"`

	// Indication that you need to pass the unique identifier of charges of the jewelry
	IsJwUINNeeded bool `json:"is_jw_uin_needed"`

	// Indication that you need to pass the "Chestny ZNAK" labeling
	IsMandatoryMarkNeeded bool `json:"is_mandatory_mark_needed"`

	// Indication that you can pass the "Chestny ZNAK" labeling, but it's not mandatory
	IsMandatoryMarkPossible bool `json:"is_mandatory_mark_possible"`

	// Indication that you need to pass the product batch registration number
	IsRNPTNeeded bool `json:"is_rnpt_needed"`

	// Product ID
	ProductId int64 `json:"product_id"`

	// Items quantity
	Quantity int32 `json:"quantity"`
}

// Method returns the created items data passed in the `/v5/fbs/posting/product/exemplar/set` method.
//
// Use this method to get the `exemplar_id`
func (c FBS) CreateOrGetProductExemplar(ctx context.Context, params *CreateOrGetProductExemplarParams) (*CreateOrGetProductExemplarResponse, error) {
	url := "/v6/fbs/posting/product/exemplar/create-or-get"

	resp := &CreateOrGetProductExemplarResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetCarriageParams struct {
	CarriageId int64 `json:"carriage_id"`
}

type GetCarriageResponse struct {
	core.CommonResponse

	// Acceptance certificate type for FBS sellers
	ActType string `json:"act_type"`

	// Pass identifiers for the freight
	ArrivalPassIds []string `json:"arrival_pass_ids"`

	// List of available actions on the freight
	AvailableActions []string `json:"available_actions"`

	// Cancel availability
	CancelAvailability GetCarriageCancelAvailability `json:"cancel_availability"`

	// Freight identifier
	CarriageId int64 `json:"carriage_id"`

	// Company identifier
	CompanyId int64 `json:"company_id"`

	// Number of package units
	ContainersCount int32 `json:"containers_count"`

	// Date and time of freight creation
	CreatedAt time.Time `json:"created_at"`

	// Delivery method identifier
	DeliveryMethodId int64 `json:"delivery_method_id"`

	// Shipping date
	DepartureDate string `json:"departure_date"`

	// First mile type
	FirstMileType string `json:"first_mile_type"`

	// true if there are shipments subject to shipping that are not in the current freight
	HasPostingsForNextCarriage bool `json:"has_postings_for_next_carriage"`

	// Delivery service integration type
	IntegrationType string `json:"integration_type"`

	// true if you already printed shipping labels
	IsContainerLabelPrinted bool `json:"is_container_label_printed"`

	// true if the freight is partial
	IsPartial bool `json:"is_partial"`

	// Serial number of the partial freight
	PartialNum int64 `json:"partial_num"`

	// The number of retries to create a freight
	RetryCount int32 `json:"retry_count"`

	// Freight status
	Status GetCarriageStatus `json:"status"`

	// Delivery method identifier
	TPLProviderId int64 `json:"tpl_provider_id"`

	// Date and time when the freight was last updated
	UpdatedAt time.Time `json:"updated_at"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

type GetCarriageCancelAvailability struct {
	// true if the freight can be cancelled
	IsCancelAvailable bool `json:"is_cancel_available"`

	// Reason why freight can't be cancelled
	Reason string `json:"reason"`
}

func (c FBS) GetCarriage(ctx context.Context, params *GetCarriageParams) (*GetCarriageResponse, error) {
	url := "/v1/carriage/get"

	resp := &GetCarriageResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetCancellationReasonsResponse struct {
	core.CommonResponse

	// Method result
	Result []GetCancellationReasonsResult `json:"result"`
}

type GetCancellationReasonsResult struct {
	// Cancellation reasons identifier
	Id int64 `json:"id"`

	// Shipment cancellation result. true if the request is available for cancellation
	IsAvailableForCancellation bool `json:"is_available_for_cancellation"`

	// Category name
	Title string `json:"title"`

	// Shipment cancellation initiator
	TypeId string `json:"type_id"`
}

func (c FBS) GetCancellationReasons(ctx context.Context) (*GetCancellationReasonsResponse, error) {
	url := "/v1/posting/fbo/cancel-reason/list"

	resp := &GetCancellationReasonsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type SetShippingDateParams struct {
	// New shipping date
	NewCutoffDate time.Time `json:"new_cutoff_date"`

	// Shipment number
	PostingNumber string `json:"posting_number"`
}

type SetShippingDateResponse struct {
	core.CommonResponse

	// true if the new date is set
	Result bool `json:"result"`
}

func (c FBS) SetShippingDate(ctx context.Context, params *SetShippingDateParams) (*SetShippingDateResponse, error) {
	url := "/v1/posting/cutoff/set"

	resp := &SetShippingDateResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type SplitOrderParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Shipments list the order will be split into. You can split one order per one request
	Postings []SplitOrderParamPosting `json:"postings"`
}

type SplitOrderParamPosting struct {
	Products []SplitOrderPostingProduct `json:"products"`
}

type SplitOrderResponse struct {
	core.CommonResponse

	// Original shipment details
	ParentPosting SplitOrderPosting `json:"parent_posting"`

	// List of shipments the order was split into
	Postings []SplitOrderPosting `json:"postings"`
}

type SplitOrderPosting struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// List of products in the shipment
	Products []SplitOrderPostingProduct `json:"products"`
}

type SplitOrderPostingProduct struct {
	// FBS product identifier in the Ozon system, SKU
	ProductId int64 `json:"product_id"`

	// Product quantity
	Quantity int64 `json:"quantity"`
}

func (c FBS) SplitOrder(ctx context.Context, params *SplitOrderParams) (*SplitOrderResponse, error) {
	url := "/v1/posting/fbs/split"

	resp := &SplitOrderResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListUnpaidProductsParams struct {
	// Cursor for the next data sample
	Cursor string `json:"cursor"`

	// Number of values in the response
	Limit int32 `json:"limit,omitempty"`
}

type ListUnpaidProductsResponse struct {
	core.CommonResponse

	Products []UnpaidProduct `json:"products"`

	// Cursor for the next data sample
	Cursor string `json:"cursor"`
}

type UnpaidProduct struct {
	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product quantity, pcs
	Quantity int32 `json:"quantity"`

	// Product name
	Name string `json:"name"`

	// Link to product image
	ImageURL string `json:"image_url"`
}

func (c FBS) ListUnpaidProducts(ctx context.Context, params *ListUnpaidProductsParams) (*ListUnpaidProductsResponse, error) {
	url := "/v1/posting/unpaid-legal/product/list"

	resp := &ListUnpaidProductsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ChangeShipmentCompositionParams struct {
	// Freight identifier
	CarriageId int64 `json:"carriage_id"`

	// Current list of shipments
	PostingNumbers []string `json:"posting_numbers"`
}

type ChangeShipmentCompositionResponse struct {
	core.CommonResponse

	Result []ChangeShipmentCompositionResult `json:"result"`
}

type ChangeShipmentCompositionResult struct {
	// Error message
	Error string `json:"error"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Request processing result. true if the request was executed successfully
	Result bool `json:"result"`
}

// Method overwrites the list of orders in the shipment. Pass only orders in the awaiting_deliver status and ones which are ready for shipping.
func (c FBS) ChangeShipmentComposition(ctx context.Context, params *ChangeShipmentCompositionParams) (*ChangeShipmentCompositionResponse, error) {
	url := "/v1/carriage/set-postings"

	resp := &ChangeShipmentCompositionResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type DeleteShipmentParams struct {
	// Freight identifier
	CarriageId int64 `json:"carriage_id"`
}

type DeleteShipmentResponse struct {
	core.CommonResponse

	// Error message
	Error string `json:"error"`

	// Carriage status
	Status string `json:"carriage_status"`
}

func (c FBS) DeleteShipment(ctx context.Context, params *DeleteShipmentParams) (*DeleteShipmentResponse, error) {
	url := "/v1/carriage/cancel"

	resp := &DeleteShipmentResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type VerifyCourierCodeParams struct {
	// Courier code
	PickupCode string `json:"pickup_code"`

	// Shipment number
	PostingNumber string `json:"posting_number"`
}

type VerifyCourierCodeResponse struct {
	core.CommonResponse

	// true, if the code is correct
	Valid bool `json:"valid"`
}

// Use this method to verify the courier code when handing over realFBS Express shipments
func (c FBS) VerifyCourierCode(ctx context.Context, params *VerifyCourierCodeParams) (*VerifyCourierCodeResponse, error) {
	url := "/v1/posting/fbs/pick-up-code/verify"

	resp := &VerifyCourierCodeResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateProductsDataParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`
}

type UpdateProductsDataResponse struct {
	core.CommonResponse
}

// Use the method after passing item data using the
// /v6/fbs/posting/product/exemplar/set method to
// save updated item data for shipments in the “Awaiting shipment” status
func (c FBS) UpdateProductsData(ctx context.Context, params *UpdateProductsDataParams) (*UpdateProductsDataResponse, error) {
	url := "/v1/fbs/posting/product/exemplar/update"

	resp := &UpdateProductsDataResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
