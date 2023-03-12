package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type GetStocksInfoParams struct {
	// Identifier of the last value on the page. Leave this field blank in the first request.
	//
	// To get the next values, specify last_id from the response of the previous request.
	LastId string `json:"last_id,omitempty"`

	// Number of values per page. Minimum is 1, maximum is 1000
	Limit int64 `json:"limit,omitempty"`

	// Filter by product
	Filter GetStocksInfoFilter `json:"filter,omitempty"`
}

type GetStocksInfoFilter struct {
	// Filter by the offer_id parameter. It is possible to pass a list of values
	OfferId string `json:"offer_id,omitempty"`

	// Filter by the product_id parameter. It is possible to pass a list of values
	ProductId int64 `json:"product_id,omitempty"`

	// Filter by product visibility
	Visibility string `json:"visibility,omitempty"`
}

type GetStocksInfoResponse struct {
	core.CommonResponse

	// Method Result
	Result struct {
		// Identifier of the last value on the page
		//
		// To get the next values, specify the recieved value in the next request in the last_id parameter
		LastId string `json:"last_id,omitempty"`

		// The number of unique products for which information about stocks is displayed
		Total int32 `json:"total,omitempty"`

		// Product details
		Items []struct {
			// Product identifier in the seller's system
			OfferId string `json:"offer_id,omitempty"`

			// Product identifier
			ProductId int64 `json:"product_id,omitempty"`

			// Stock details
			Stocks []struct {
				// In a warehouse
				Present int32 `json:"present,omitempty"`

				// Reserved
				Reserved int32 `json:"reserved,omitempty"`

				// Warehouse type
				Type string `json:"type,omitempty" default:"ALL"`
			} `json:"stocks,omitempty"`
		} `json:"items,omitempty"`
	} `json:"result,omitempty"`
}

func (c Client) GetStocksInfo(params *GetStocksInfoParams) (*GetStocksInfoResponse, error) {
	url := "/v3/product/info/stocks"

	resp := &GetStocksInfoResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductDetailsParams struct {
	OfferId   string `json:"offer_id"`
	ProductId int64  `json:"product_id"`
	SKU       int64  `json:"sku"`
}

type GetProductDetailsResponse struct {
	core.CommonResponse

	Result struct {
		Barcode     string   `json:"barcode"`
		Barcodes    []string `json:"barcodes"`
		BuyboxPrice string   `json:"buybox_price"`
		CategoryId  int64    `json:"category_id"`
		ColorImage  string   `json:"color_image"`
		Commissions []struct {
			DeliveryAmount float64 `json:"deliveryAmount"`
			MinValue       float64 `json:"minValue"`
			Percent        float64 `json:"percent"`
			ReturnAmount   float64 `json:"returnAmount"`
			SaleSchema     string  `json:"saleSchema"`
			Value          float64 `json:"value"`
		} `json:"commissions"`
		CreatedAt         time.Time `json:"created_at"`
		FBOSKU            int64     `json:"fbo_sku"`
		FBSSKU            int64     `json:"fbs_sku"`
		Id                int64     `json:"id"`
		Images            []string  `json:"images"`
		PrimaryImage      string    `json:"primary_image"`
		Images360         []string  `json:"images360"`
		HasDiscountedItem bool      `json:"has_discounted_item"`
		IsDiscounted      bool      `json:"is_discounted"`
		DiscountedStocks  struct {
			Coming   int32 `json:"coming"`
			Present  int32 `json:"present"`
			Reserved int32 `json:"reserved"`
		} `json:"discounted_stocks"`
		IsKGT               bool   `json:"is_kgt"`
		IsPrepayment        bool   `json:"is_prepayment"`
		IsPrepaymentAllowed bool   `json:"is_prepayment_allowed"`
		CurrencyCode        string `json:"currency_code"`
		MarketingPrice      string `json:"marketing_price"`
		MinOzonPrice        string `json:"min_ozon_price"`
		MinPrice            string `json:"min_price"`
		Name                string `json:"name"`
		OfferId             string `json:"offer_id"`
		OldPrice            string `json:"old_price"`
		PremiumPrice        string `json:"premium_price"`
		Price               string `json:"price"`
		PriceIndex          string `json:"price_idnex"`
		RecommendedPrice    string `json:"recommended_price"`
		Status              struct {
			State            string                               `json:"state"`
			StateFailed      string                               `json:"state_failed"`
			ModerateStatus   string                               `json:"moderate_status"`
			DeclineReasons   []string                             `json:"decline_reasons"`
			ValidationsState string                               `json:"validation_state"`
			StateName        string                               `json:"state_name"`
			StateDescription string                               `json:"state_description"`
			IsFailed         bool                                 `json:"is_failed"`
			IsCreated        bool                                 `json:"is_created"`
			StateTooltip     string                               `json:"state_tooltip"`
			ItemErrors       []GetProductDetailsResponseItemError `json:"item_errors"`
			StateUpdatedAt   time.Time                            `json:"state_updated_at"`
		} `json:"status"`
		Sources []struct {
			IsEnabled bool   `json:"is_enabled"`
			SKU       int64  `json:"sku"`
			Source    string `json:"source"`
		} `json:"sources"`
		Stocks struct {
			Coming   int32 `json:"coming"`
			Present  int32 `json:"present"`
			Reserved int32 `json:"reserved"`
		} `json:"stocks"`
		UpdatedAt         time.Time `json:"updated_at"`
		VAT               string    `json:"vat"`
		VisibilityDetails struct {
			ActiveProduct bool `json:"active_product"`
			HasPrice      bool `json:"has_price"`
			HasStock      bool `json:"has_stock"`
		} `json:"visibility_details"`
		Visible      bool    `json:"visible"`
		VolumeWeight float64 `json:"volume_weights"`
	} `json:"Result"`
}

type GetProductDetailsResponseItemError struct {
	Code                        string `json:"code"`
	State                       string `json:"state"`
	Level                       string `json:"level"`
	Description                 string `json:"description"`
	Field                       string `json:"field"`
	AttributeId                 int64  `json:"attribute_id"`
	AttributeName               string `json:"attribute_name"`
	OptionalDescriptionElements struct {
		PropertyName string `json:"property_name"`
	} `json:"optional_description_elements"`
}

func (c Client) GetProductDetails(params *GetProductDetailsParams) (*GetProductDetailsResponse, error) {
	url := "/v2/product/info"

	resp := &GetProductDetailsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
