package ozon

import (
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
	Result []struct {
		// Additional data for shipment list
		AdditionalData []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"additional_data"`

		// Analytical data
		AnalyticsData struct {
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
		} `json:"analytics_data"`

		// Shipment cancellation reason identifier
		CancelReasonId int64 `json:"cancel_reason_id"`

		// Date and time of shipment creation
		CreatedAt time.Time `json:"created_at"`

		// Financial data
		FinancialData struct {
			// Identifier of the cluster, where the shipment is sent from
			ClusterFrom string `json:"cluster_from"`

			// Identifier of the cluster, where the shipment is delivered to
			ClusterTo string `json:"cluster_to"`

			// Services
			PostingServices MarketplaceServices `json:"posting_services"`

			// Products list
			Products []FinancialDataProduct `json:"products"`
		} `json:"financial_data"`

		// Date and time of shipment processing start
		InProccessAt time.Time `json:"in_process_at"`

		// Identifier of the order to which the shipment belongs
		OrderId int64 `json:"order_id"`

		// Number of the order to which the shipment belongs
		OrderNumber string `json:"order_number"`

		// Shipment number
		PostingNumber string `json:"posting_number"`

		// Number of products in the shipment
		Products []struct {
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
		} `json:"products"`

		// Shipment status
		Status string `json:"status"`
	} `json:"result"`
}

// Returns a list of shipments for a specified period of time. You can additionally filter the shipments by their status
func (c FBO) GetFBOShipmentsList(params *GetFBOShipmentsListParams) (*GetFBOShipmentsListResponse, error) {
	url := "/v2/posting/fbo/list"

	resp := &GetFBOShipmentsListResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
