package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type ListUnprocessedShipmentsParams struct {
	Direction string                         `json:"dir"`
	Filter    ListUnprocessedShipmentsFilter `json:"filter"`
	Limit     int64                          `json:"limit"`
	Offset    int64                          `json:"offset"`
	With      ListUnprocessedShipmentsWith   `json:"with"`
}

type ListUnprocessedShipmentsFilter struct {
	CutoffFrom         time.Time `json:"cutoff_from"`
	CutoffTo           time.Time `json:"cutoff_to"`
	DeliveringDateFrom time.Time `json:"delivering_date_from"`
	DeliveringDateTo   time.Time `json:"delivering_date_to"`
	DeliveryMethodId   []int64   `json:"deliveryMethodId"`
	ProviderId         []int64   `json:"provider_id"`
	Status             string    `json:"status"`
	WarehouseId        []int64   `json:"warehouse_id"`
}

type ListUnprocessedShipmentsWith struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes      bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	Translit      bool `json:"translit"`
}

type ListUnprocessedShipmentsResponse struct {
	core.CommonResponse

	Result ListUnprocessedShipmentsResult `json:"result"`
}

type ListUnprocessedShipmentsResult struct {
	Count    int64        `json:"count"`
	Postings []FBSPosting `json:"postings"`
}

type FBSPosting struct {
	Addressee struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"addressee"`

	AnalyticsData struct {
		City                 string    `json:"city"`
		DeliveryDateBegin    time.Time `json:"delivery_date_begin"`
		DeliveryDateEnd      time.Time `json:"delivery_date_end"`
		DeliveryType         string    `json:"delivery_type"`
		IsLegal              bool      `json:"is_legal"`
		IsPremium            bool      `json:"is_premium"`
		PaymentTypeGroupName string    `json:"payment_type_group_name"`
		Region               string    `json:"region"`
		TPLProvider          string    `json:"tpl_provider"`
		TPLProviderId        int64     `json:"tpl_provider_id"`
		Warehouse            string    `json:"warehouse"`
		WarehouseId          int64     `json:"warehouse_id"`
	} `json:"analytics_data"`

	Barcodes struct {
		LowerBarcode string `json:"lower_barcode"`
		UpperBarcode string `json:"upper_barcode"`
	} `json:"barcodes"`

	Cancellation struct {
		AffectCancellationRating bool   `json:"affect_cancellation_rating"`
		CancelReason             string `json:"cancel_reason"`
		CancelReasonId           int64  `json:"cancel_reason_id"`
		CancellationInitiator    string `json:"cancellation_initiator"`
		CancellationType         string `json:"cancellation_type"`
		CancelledAfterShip       bool   `json:"cancellation_after_ship"`
	} `json:"cancellation"`

	Customer FBSCustomer `json:"customer"`

	DeliveringDate time.Time `json:"delivering_date"`

	DeliveryMethod struct {
		Id            int64  `json:"id"`
		Name          string `json:"name"`
		TPLProvider   string `json:"tpl_provider"`
		TPLProviderId int64  `json:"tpl_provider_id"`
		Warehouse     string `json:"warehouse"`
		WarehouseId   int64  `json:"warehouse_id"`
	} `json:"delivery_method"`

	FinancialData struct {
		ClusterFrom     string              `json:"cluster_from"`
		ClusterTo       string              `json:"cluster_to"`
		PostingServices MarketplaceServices `json:"posting_services"`

		Products []FinancialDataProduct `json:"products"`
	}

	InProccessAt        time.Time `json:"in_process_at"`
	IsExpress           bool      `json:"is_express"`
	IsMultibox          bool      `json:"is_multibox"`
	MultiBoxQuantity    int32     `json:"multi_box_qty"`
	OrderId             int64     `json:"order_id"`
	OrderNumber         string    `json:"order_number"`
	ParentPostingNumber string    `json:"parent_posting_number"`
	PostingNumber       string    `json:"posting_number"`

	Products []struct {
		MandatoryMark []string `json:"mandatory_mark"`
		Name          string   `json:"name"`
		OfferId       string   `json:"offer_id"`
		CurrencyCode  string   `json:"currency_code"`
		Price         string   `json:"price"`
		Quantity      int32    `json:"quantity"`
		SKU           int64    `json:"sku"`
	} `json:"products"`

	Requirements struct {
		ProductsRequiringGTD           []string `json:"products_requiring_gtd"`
		ProductsRequiringCountry       []string `json:"products_requiring_country"`
		ProductsRequiringMandatoryMark []string `json:"products_requiring_mandatory_mark"`
		ProductsRequiringRNPT          []string `json:"products_requiring_rnpt"`
	} `json:"requirements"`

	ShipmentDate       time.Time `json:"shipment_date"`
	Status             string    `json:"status"`
	TPLIntegrationType string    `json:"tpl_integration_type"`
	TrackingNumber     string    `json:"tracking_number"`
}

type FBSCustomer struct {
	Address struct {
		AddressTail     string  `json:"address_tail"`
		City            string  `json:"city"`
		Comment         string  `json:"comment"`
		Country         string  `json:"country"`
		District        string  `json:"district"`
		Latitude        float64 `json:"latitude"`
		Longitude       float64 `json:"longitude"`
		ProviderPVZCode string  `json:"provider_pvz_code"`
		PVZCode         int64   `json:"pvz_code"`
		Region          string  `json:"region"`
		ZIPCode         string  `json:"zip_code"`
	} `json:"customer"`

	CustomerEmail string `json:"customer_email"`
	CustomerId    int64  `json:"customer_id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
}

type MarketplaceServices struct {
	DeliveryToCustomer            float64 `json:"marketplace_service_item_deliv_to_customer"`
	DirectFlowTrans               float64 `json:"marketplace_service_item_direct_flow_trans"`
	DropoffFF                     float64 `json:"marketplace_service_item_item_dropoff_ff"`
	DropoffPVZ                    float64 `json:"marketplace_service_item_dropoff_pvz"`
	DropoffSC                     float64 `json:"marketplace_service_item_dropoff_sc"`
	Fulfillment                   float64 `json:"marketplace_service_item_fulfillment"`
	Pickup                        float64 `json:"marketplace_service_item_pickup"`
	ReturnAfterDeliveryToCustomer float64 `json:"marketplace_service_item_return_after_deliv_to_customer"`
	ReturnFlowTrans               float64 `json:"marketplace_service_item_return_flow_trans"`
	ReturnNotDeliveryToCustomer   float64 `json:"marketplace_service_item_return_not_deliv_to_customer"`
	ReturnPartGoodsCustomer       float64 `json:"marketplace_service_item_return_part_goods_customer"`
}

type FinancialDataProduct struct {
	Actions                 []string            `json:"actions"`
	ClientPrice             string              `json:"client_price"`
	CommissionAmount        float64             `json:"commission_amount"`
	CommissionPercent       int64               `json:"commission_percent"`
	CommissionsCurrencyCode string              `json:"commissions_currency_code"`
	ItemServices            MarketplaceServices `json:"item_services"`
	CurrencyCode            string              `json:"currency_code"`
	OldPrice                float64             `json:"old_price"`
	Payout                  float64             `json:"payout"`
	Picking                 struct {
		Amount float64   `json:"amount"`
		Moment time.Time `json:"moment"`
		Tag    string    `json:"tag"`
	} `json:"picking"`
	Price                float64 `json:"price"`
	ProductId            int64   `json:"product_id"`
	Quantity             int64   `json:"quantity"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	TotalDiscountValue   float64 `json:"total_discount_value"`
}

func (c Client) ListUnprocessedShipments(params *ListUnprocessedShipmentsParams) (*ListUnprocessedShipmentsResponse, error) {
	url := "/v3/posting/fbs/unfulfilled/list"

	resp := &ListUnprocessedShipmentsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetFBSShipmentsListParams struct {
	// Sorting direction
	Direction string `json:"direction"`

	//Filter
	Filter GetFBSShipmentsListFilter `json:"filter"`

	// Number of shipments in the response:
	//   - maximum is 50,
	//   - minimum is 1.
	Limit int64 `json:"limit"`

	// Number of elements that will be skipped in the response. For example, if offset=10, the response will start with the 11th element found
	Offset int64 `json:"offset"`

	// Additional fields that should be added to the response
	With GetFBSShipmentsListWith `json:"with"`
}

type GetFBSShipmentsListFilter struct {
	// Delivery method identifier
	DeliveryMethodId []int64 `json:"delivery_method_id"`

	// Order identifier
	OrderId int64 `json:"order_id"`

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
	Result struct {
		// Indicates that the response returned not the entire array of shipments:
		//
		//   - true — it is necessary to make a new request with a different offset value to get information on the remaining shipments;
		//   - false — the entire array of shipments for the filter specified in the request was returned in the response
		HasNext bool `json:"has_next"`

		// Shipment details
		Postings FBSPosting `json:"postings"`
	} `json:"result"`
}

// Returns a list of shipments for the specified time period: it shouldn't be longer than one year.
//
// You can filter shipments by their status. The list of available statuses is specified in the description of the filter.status parameter.
//
// The true value of the has_next parameter in the response means there is not the entire array of shipments in the response. To get information on the remaining shipments, make a new request with a different offset value.
func (c Client) GetFBSShipmentsList(params *GetFBSShipmentsListParams) (*GetFBSShipmentsListResponse, error) {
	url := "/v3/posting/fbs/list"

	resp := &GetFBSShipmentsListResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
