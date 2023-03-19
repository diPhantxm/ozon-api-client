package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Products struct {
	client *core.Client
}

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

// Returns information about the quantity of products in stock:
//
// * how many items are available,
//
// * how many are reserved by customers.
func (c Products) GetStocksInfo(params *GetStocksInfoParams) (*GetStocksInfoResponse, error) {
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
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

type GetProductDetailsResponse struct {
	core.CommonResponse

	// Request results
	Result ProductDetails `json:"Result"`
}

type ProductDetails struct {
	// Barcode
	Barcode string `json:"barcode"`

	// All product barcodes
	Barcodes []string `json:"barcodes"`

	// Main offer price on Ozon.
	//
	// The field is deprecated. Returns an empty string ""
	BuyboxPrice string `json:"buybox_price"`

	// Category identifier
	CategoryId int64 `json:"category_id"`

	// Marketing color
	ColorImage string `json:"color_image"`

	// Commission fees details
	Commissions []struct {
		// Delivery cost
		DeliveryAmount float64 `json:"deliveryAmount"`

		// Minimum commission fee
		MinValue float64 `json:"minValue"`

		// Commission percentage
		Percent float64 `json:"percent"`

		// Return cost
		ReturnAmount float64 `json:"returnAmount"`

		// Sale scheme
		SaleSchema string `json:"saleSchema"`

		// Commission fee amount
		Value float64 `json:"value"`
	} `json:"commissions"`

	// Date and time when the product was created
	CreatedAt time.Time `json:"created_at"`

	// SKU of the product that is sold from the Ozon warehouse (FBO)
	FBOSKU int64 `json:"fbo_sku"`

	// SKU of the product that is sold from the seller's warehouse (FBS and rFBS)
	FBSSKU int64 `json:"fbs_sku"`

	// Document generation task number
	Id int64 `json:"id"`

	// An array of links to images. The images in the array are arranged in the order of their arrangement on the site. If the `primary_image` parameter is not specified, the first image in the list is the main one for the product
	Images []string `json:"images"`

	// Main product image
	PrimaryImage string `json:"primary_image"`

	// Array of 360 images
	Images360 []string `json:"images360"`

	// true if the product has markdown equivalents at the Ozon warehouse
	HasDiscountedItem bool `json:"has_discounted_item"`

	// Indication of a markdown product:
	//
	// * true if the product was created by the seller as a markdown
	//
	// * false if the product is not markdown or was marked down by Ozon
	IsDiscounted bool `json:"is_discounted"`

	// Markdown products stocks
	DiscountedStocks ProductDiscountedStocks `json:"discounted_stocks"`

	// Indication of a bulky product
	IsKGT bool `json:"is_kgt"`

	// Indication of mandatory prepayment for the product:
	//
	// * true — to buy a product, you need to make a prepayment.
	//
	// * false—prepayment is not required
	IsPrepayment bool `json:"is_prepayment"`

	// If prepayment is possible, the value is true
	IsPrepaymentAllowed bool `json:"is_prepayment_allowed"`

	// Currency of your prices. It matches the currency set in the personal account settings
	CurrencyCode string `json:"currency_code"`

	// The price of the product including all promotion discounts. This value will be shown on the Ozon storefront
	MarketingPrice string `json:"marketing_price"`

	// Minimum price for similar products on Ozon.
	//
	// The field is deprecated. Returns an empty string ""
	MinOzonPrice string `json:"min_ozon_price"`

	// Minimum product price with all promotions applied
	MinPrice string `json:"min_price"`

	// Name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Price before discounts. Displayed strikethrough on the product description page
	OldPrice string `json:"old_price"`

	// Price for customers with an Ozon Premium subscription
	PremiumPrice string `json:"premium_price"`

	// Product price including discounts. This value is shown on the product description page
	Price string `json:"price"`

	// Price index. Learn more in Help Center
	PriceIndex string `json:"price_idnex"`

	// Product price suggested by the system based on similar offers
	RecommendedPrice string `json:"recommended_price"`

	// Product state description
	Status struct {
		// Product state
		State string `json:"state"`

		// Product state on the transition to which an error occurred
		StateFailed string `json:"state_failed"`

		// Moderation status
		ModerateStatus string `json:"moderate_status"`

		// Product decline reasons
		DeclineReasons []string `json:"decline_reasons"`

		// Validation status
		ValidationsState string `json:"validation_state"`

		// Product status name
		StateName string `json:"state_name"`

		// Product state description
		StateDescription string `json:"state_description"`

		// Indiction that there were errors while creating products
		IsFailed bool `json:"is_failed"`

		// Indiction that the product was created
		IsCreated bool `json:"is_created"`

		// Tooltips for the current product state
		StateTooltip string `json:"state_tooltip"`

		// Product loading errors
		ItemErrors []GetProductDetailsResponseItemError `json:"item_errors"`

		// The last time product state changed
		StateUpdatedAt time.Time `json:"state_updated_at"`
	} `json:"status"`

	// Details about the sources of similar offers. Learn more in Help Сenter
	Sources []struct {
		// Indication that the source is taken into account when calculating the market value
		IsEnabled bool `json:"is_enabled"`

		// Product identifier in the Ozon system, SKU
		SKU int64 `json:"sku"`

		// Link to the source
		Source string `json:"source"`
	} `json:"sources"`

	// Details about product stocks
	Stocks struct {
		// Supply expected
		Coming int32 `json:"coming"`

		// Currently at the warehouse
		Present int32 `json:"present"`

		// Reserved
		Reserved int32 `json:"reserved"`
	} `json:"stocks"`

	// Date of the last product update
	UpdatedAt time.Time `json:"updated_at"`

	// Product VAT rate
	VAT string `json:"vat"`

	// Product visibility settings
	VisibilityDetails struct {
		// If the product is active, the value is true
		ActiveProduct bool `json:"active_product"`

		// If the price is set, the value is true
		HasPrice bool `json:"has_price"`

		// If there is stock at the warehouses, the value is true
		HasStock bool `json:"has_stock"`
	} `json:"visibility_details"`

	// If the product is on sale, the value is true
	Visible bool `json:"visible"`

	// Product volume weight
	VolumeWeight float64 `json:"volume_weights"`
}

type ProductDiscountedStocks struct {
	// Quantity of products to be supplied
	Coming int32 `json:"coming"`

	// Quantity of products in warehouse
	Present int32 `json:"present"`

	// Quantity of products reserved
	Reserved int32 `json:"reserved"`
}
type GetProductDetailsResponseItemError struct {
	// Error code
	Code string `json:"code"`

	// Product state in which an error occurred
	State string `json:"state"`

	// Error level
	Level string `json:"level"`

	// Error description
	Description string `json:"description"`

	// Error field
	Field string `json:"field"`

	// Error attribute identifier
	AttributeId int64 `json:"attribute_id"`

	// Attribute name
	AttributeName string `json:"attribute_name"`

	// Additional fields for error description
	OptionalDescriptionElements map[string]string `json:"optional_description_elements"`
}

// Get product details
func (c Products) GetProductDetails(params *GetProductDetailsParams) (*GetProductDetailsResponse, error) {
	url := "/v2/product/info"

	resp := &GetProductDetailsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateStocksParams struct {
	// Stock details
	Stocks []UpdateStocksStock `json:"stocks"`
}

// Stock detail
type UpdateStocksStock struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Quantity of products in stock
	Stock int64 `json:"stocks"`
}

type UpdateStocksResponse struct {
	core.CommonResponse

	// Request results
	Result []struct {
		// An array of errors that occurred while processing the request
		Errors []struct {
			// Error code
			Code string `json:"code"`

			// Error reason
			Message string `json:"message"`
		} `json:"errors"`

		// Product identifier in the seller's system
		OfferId string `json:"offer_id"`

		// Product identifier
		ProductId int64 `json:"product_id"`

		// If the product details have been successfully updated — true
		Updated bool `json:"updated"`
	}
}

// Allows you to change the products in stock quantity. The method is only used for FBS and rFBS warehouses.
//
// With one request you can change the availability for 100 products. You can send up to 80 requests in a minute.
//
// Availability can only be set after the product status has been changed to processed.
func (c Products) UpdateStocks(params *UpdateStocksParams) (*UpdateStocksResponse, error) {
	url := "/v1/product/import/stocks"

	resp := &UpdateStocksResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateQuantityStockProductsParams struct {
	// Information about the products at the warehouses
	Stocks []UpdateQuantityStockProductsStock `json:"stocks"`
}

type UpdateQuantityStockProductsStock struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Quantity
	Stock int64 `json:"stock"`

	// Warehouse identifier derived from the /v1/warehouse/list method
	WarehouseId int64 `json:"warehouse_id"`
}

type UpdateQuantityStockProductsResponse struct {
	core.CommonResponse

	// Method result
	Result []struct {
		// An array of errors that occurred while processing the request
		Errors []struct {
			// Error code
			Code string `json:"code"`

			// Error reason
			Message string `json:"message"`
		} `json:"errors"`

		// Product identifier in the seller's system
		Offerid string `json:"offer_id"`

		// Product identifier
		ProductId int64 `json:"product_id"`

		// If the request was completed successfully and the stocks are updated — true
		Updated bool `json:"updated"`

		// Warehouse identifier derived from the /v1/warehouse/list method
		WarehouseId int64 `json:"warehouse_id"`
	} `json:"result"`
}

// Allows you to change the products in stock quantity.
//
// With one request you can change the availability for 100 products. You can send up to 80 requests in a minute.
//
// You can update the stock of one product in one warehouse only once in 2 minutes, otherwise there will be the TOO_MANY_REQUESTS error in the response.
//
// Availability can only be set after the product status has been changed to processed.
//
// Bulky products stock can only be updated in the warehouses for bulky products.
func (c Products) UpdateQuantityStockProducts(params *UpdateQuantityStockProductsParams) (*UpdateQuantityStockProductsResponse, error) {
	url := "/v2/products/stocks"

	resp := &UpdateQuantityStockProductsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type StocksInSellersWarehouseParams struct {
	// SKU of the product that is sold from the seller's warehouse (FBS and RFBS schemes).
	//
	// Get fbs_sku in the /v2/product/info and /v2/product/info/list methods response.
	//
	// The maximum number of SKUs per request is 500.
	FBSSKU []string `json:"fbs_sku"`
}

type StocksInSellersWarehouseResponse struct {
	core.CommonResponse

	// Method result
	Result []struct {
		// SKU of the product that is sold from the seller's warehouse (FBS and RFBS schemes)
		FBSSKU int64 `json:"fbs_sku"`

		// Total number of items in the warehouse
		Present int64 `json:"present"`

		// The product identifier in the seller's system
		ProductId int64 `json:"product_id"`

		// The number of reserved products in the warehouse
		Reserved int64 `json:"reserved"`

		// Warehouse identifier
		WarehouseId int64 `json:"warehouse_id"`

		// Warehouse name
		WarehouseName string `json:"warehouse_name"`
	}
}

// Get stocks in seller's warehouse
func (c Products) StocksInSellersWarehouse(params *StocksInSellersWarehouseParams) (*StocksInSellersWarehouseResponse, error) {
	url := "/v1/product/info/stocks-by-warehouse/fbs"

	resp := &StocksInSellersWarehouseResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdatePricesParams struct {
	// Product prices details
	Prices []UpdatePricesPrice `json:"prices"`
}

// Product price details
type UpdatePricesPrice struct {
	// Attribute for enabling and disabling promos auto-application
	AutoActionEnabled string `json:"auto_action_enabled"`

	// Currency of your prices. The passed value must be the same as the one set in the personal account settings.
	// By default, the passed value is RUB, Russian ruble
	CurrencyCode string `json:"currency_code"`

	// Minimum product price with all promotions applied
	MinPrice string `json:"min_price"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Price before discounts. Displayed strikethrough on the product description page.
	// Specified in rubles. The fractional part is separated by decimal point, up to two digits after the decimal point.
	OldPrice string `json:"old_price"`

	// Product price including discounts. This value is displayed on the product description page.
	//
	// If the current price of the product is from 400 to 10 000 rubles inclusive, the difference between the values of price and old_price fields should be more than 5%, but not less than 20 rubles.
	Price string `json:"price"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

type UpdatePricesResponse struct {
	core.CommonResponse

	Result []struct {
		// An array of errors that occurred while processing the request
		Errors []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"errors"`

		// Product identifier in the seller's system
		OfferId string `json:"offer_id"`

		// Product ID
		ProductId int64 `json:"product_id"`

		// If the product details have been successfully updated — true
		Updated bool `json:"updated"`
	} `json:"result"`
}

// Allows you to change a price of one or more products.
// You can change prices for 1000 products in one request.
// To reset old_price or premium_price set these parameters to 0.
//
// A new price must differ from the old one by at least 5%.
func (c Products) UpdatePrices(params *UpdatePricesParams) (*UpdatePricesResponse, error) {
	url := "/v1/product/import/prices"

	resp := &UpdatePricesResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CreateOrUpdateProductParams struct {
	// Data array
	Items []CreateOrUpdateProductItem `json:"items"`
}

// Data array
type CreateOrUpdateProductItem struct {
	// Array with the product characteristics. The characteristics depend on category.
	// You can view them in Help Center or via API
	Attributes []CreateOrUpdateAttribute `json:"attributes"`

	// Product barcode
	Barcode string `json:"barcode"`

	// Category identifier
	CategoryId int64 `json:"category_id"`

	// Marketing color.
	//
	// Pass the link to the image in the public cloud storage. The image format is JPG
	ColorImage string `json:"color_image"`

	// Array of characteristics that have nested attributes
	ComplexAttributes []CreateOrUpdateComplexAttribute `json:"complex_attributes"`

	// Package depth
	Depth int32 `json:"depth"`

	// Dimensions measurement units:
	//   - mm — millimeters,
	//   - cm — centimeters,
	//   - in — inches
	DimensionUnit string `json:"dimension_unit"`

	// Geo-restrictions. Pass a list consisting of name values received in the response of the /v1/products/geo-restrictions-catalog-by-filter method
	GeoNames []string `json:"geo_names"`

	// Package height
	Height int32 `json:"height"`

	// Array of images, up to 15 files. The images are displayed on the site in the same order as they are in the array.
	//
	// The first one will be set as the main image for the product if the primary_image parameter is not specified.
	//
	// If you use the primary_image parameter, the maximum number of images is 14. If the primary_image parameter is not specified, you can upload up to 15 images.
	//
	// Pass links to images in the public cloud storage. The image format is JPG or PNG
	Images []string `json:"images"`

	// Link to main product image
	PrimaryImage string `json:"primary_image"`

	// Array of 360 images—up to 70 files.
	//
	// Pass links to images in the public cloud storage. The image format is JPG
	Images360 []string `json:"images_360"`

	// Product name. Up to 500 characters
	Name string `json:"name"`

	// Product identifier in the seller's system.
	//
	// The maximum length of a string is 50 characters
	OfferId string `json:"offer_id"`

	// Currency of your prices. The passed value must be the same as the one set in the personal account settings.
	// By default, the passed value is RUB, Russian ruble.
	//
	// For example, if your currency set in the settings is yuan, pass the value CNY, otherwise an error will be returned
	CurrencyCode string `json:"currency_code"`

	// Price before discounts. Displayed strikethrough on the product description page. Specified in rubles. The fractional part is separated by decimal point, up to two digits after the decimal point.
	//
	// If you specified the old_price before and updated the price parameter you should update the old_price too
	OldPrice string `json:"old_price"`

	// List of PDF files
	PDFList []CreateOrUpdateProductPDF `json:"pdf_list"`

	// Price for customers with an Ozon Premium subscription
	PremiumPrice string `json:"premium_price"`

	// Product price including discounts. This value is shown on the product description card.
	// If there are no discounts on the product, specify the old_price value
	Price string `json:"price"`

	// Default: "IS_CODE_SERVICE"
	// Service type. Pass one of the values in upper case:
	//   - IS_CODE_SERVICE,
	//   - IS_NO_CODE_SERVICE
	ServiceType string `json:"service_type" default:"IS_CODE_SERVICE"`

	// VAT rate for the product:
	//   - 0 — not subject to VAT,
	//   - 0.1 — 10%,
	//   - 0.2 — 20%
	VAT string `json:"vat"`

	// Product weight with the package. The limit value is 1000 kilograms or a corresponding converted value in other measurement units
	Weight int32 `json:"weight"`

	// Weight measurement units:
	//   - g—grams,
	//   - kg—kilograms,
	//   - lb—pounds
	WeightUnit string `json:"weight_unit"`

	// Package width
	Width int32 `json:"width"`
}

// Array with the product characteristics. The characteristics depend on category.
// You can view them in Help Center or via API
type CreateOrUpdateAttribute struct {
	// Identifier of the characteristic that supports nested properties.
	// For example, the "Processor" characteristic has nested characteristics "Manufacturer", "L2 Cache", and others.
	// Each of the nested characteristics can have multiple value variants
	ComplexId int64 `json:"complex_id"`

	// Characteristic identifier
	Id int64 `json:"id"`

	Values []CreateOrUpdateAttributeValue `json:"values"`
}

type CreateOrUpdateAttributeValue struct {
	// Directory identifier
	DictionaryValueId int64 `json:"dictrionary_value_id"`

	// Value from the directory
	Value string `json:"value"`
}

type CreateOrUpdateComplexAttribute struct {
	Attributes []CreateOrUpdateAttribute `json:"attributes"`
}

type CreateOrUpdateProductPDF struct {
	// Storage order index
	Index int64 `json:"index"`

	// File name
	Name string `json:"name"`

	// File address
	URL string `json:"url"`
}

type CreateOrUpdateProductResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Number of task for products upload
		TaskId int64 `json:"task_id"`
	} `json:"result"`
}

// This method allows you to create products and update their details
func (c Products) CreateOrUpdateProduct(params *CreateOrUpdateProductParams) (*CreateOrUpdateProductResponse, error) {
	url := "/v2/product/import"

	resp := &CreateOrUpdateProductResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetListOfProductsParams struct {
	// Filter by product
	Filter GetListOfProductsFilter `json:"filter"`

	// Identifier of the last value on the page. Leave this field blank in the first request.
	//
	// To get the next values, specify last_id from the response of the previous request
	LastId string `json:"last_id"`

	// Number of values per page. Minimum is 1, maximum is 1000
	Limit int64 `json:"limit"`
}

type GetListOfProductsFilter struct {
	// Filter by the offer_id parameter. You can pass a list of values in this parameter
	OfferId []string `json:"offer_id"`

	// Filter by the product_id parameter. You can pass a list of values in this parameter
	ProductId []int64 `json:"product_id"`

	// Filter by product visibility
	Visibility string `json:"visibility"`
}

type GetListOfProductsResponse struct {
	core.CommonResponse

	// Result
	Result struct {
		// Products list
		Items []struct {
			// Product identifier in the seller's system
			OfferId string `json:"offer_id"`

			// Product ID
			ProductId int64 `json:"product_id"`
		} `json:"items"`

		// Identifier of the last value on the page.
		//
		// To get the next values, specify the recieved value in the next request in the last_id parameter
		LastId string `json:"last_id"`

		// Total number of products
		Total int32 `json:"total"`
	} `json:"result"`
}

func (c Products) GetListOfProducts(params *GetListOfProductsParams) (*GetListOfProductsResponse, error) {
	url := "/v2/product/list"

	resp := &GetListOfProductsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductsRatingBySKUParams struct {
	// List of product SKUs for which content rating should be returned
	SKUs []int64 `json:"skus"`
}

type GetProductsRatingBySKUResponse struct {
	core.CommonResponse

	// Products' content rating
	Products []struct {
		// Product identifier in the Ozon system, SKU
		SKU int64 `json:"sku"`

		// Product content rating: 0 to 100
		Rating float64 `json:"rating"`

		// Groups of characteristics that make up the content rating
		Groups []struct {
			// List of conditions that increase the product content rating
			Conditions []struct {
				// Number of content rating points that the condition gives
				Cost float64 `json:"cost"`

				// Condition description
				Description string `json:"description"`

				// Indication that the condition is met
				Fulfilled bool `json:"fulfilled"`

				// Condition identifier
				Key string `json:"key"`
			} `json:"conditions"`

			// Number of attributes you need to fill in to get the maximum score in this characteristics group
			ImproveAtLeast int32 `json:"improve_at_least"`

			// List of attributes that can increase the product content rating
			ImproveAttributes []struct {
				// Attribute identifier
				Id int64 `json:"id"`

				// Attribute name
				Name string `json:"name"`
			} `json:"improve_attributes"`

			// Group identifier
			Key string `json:"key"`

			// Group name
			Name string `json:"name"`

			// Rating in the group
			Rating float64 `json:"rating"`

			// Percentage influence of group characteristics on the content rating
			Weight float64 `json:"weight"`
		} `json:"groups"`
	} `json:"products"`
}

// Method for getting products' content rating and recommendations on how to increase it
func (c Products) GetProductsRatingBySKU(params *GetProductsRatingBySKUParams) (*GetProductsRatingBySKUResponse, error) {
	url := "/v1/product/rating-by-sku"

	resp := &GetProductsRatingBySKUResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductImportStatusParams struct {
	// Importing products task code
	TaskId int64 `json:"task_id"`
}

type GetProductImportStatusResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Product details
		Items []struct {
			// Product identifier in the seller's system.
			//
			// The maximum length of a string is 50 characters
			OfferId string `json:"offer_id"`

			// Product identifier
			ProductId int64 `json:"product_id"`

			// Product creation status. Product information is processed in queues. Possible parameter values:
			//   - pending — product in the processing queue;
			//   - imported — product loaded successfully;
			//   - failed — product loaded with errors
			Status string `json:"status"`

			// Array of errors
			Errors []struct {
				GetProductDetailsResponseItemError

				// Error technical description
				Message string `json:"message"`
			} `json:"errors"`
		} `json:"items"`

		// Product identifier in the seller's system
		Total int32 `json:"total"`
	} `json:"result"`
}

// Allows you to get the status of a product description page creation process
func (c Products) GetProductImportStatus(params *GetProductImportStatusParams) (*GetProductImportStatusResponse, error) {
	url := "/v1/product/import/info"

	resp := &GetProductImportStatusResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CreateProductByOzonIDParams struct {
	// Products details
	Items []CreateProductsByOzonIDItem `json:"items"`
}

type CreateProductsByOzonIDItem struct {
	// Product name. Up to 500 characters
	Name string `json:"name"`

	// Product identifier in the seller's system.
	//
	// The maximum length of a string is 50 characters
	OfferId string `json:"offer_id"`

	// Price before discounts. Displayed strikethrough on the product description page. Specified in rubles.
	// The fractional part is separated by decimal point, up to two digits after the decimal point
	OldPrice string `json:"old_price"`

	// Price for customers with an Ozon Premium subscription
	PremiumPrice string `json:"premium_price"`

	// Product price including discounts. This value is shown on the product description page.
	// If there are no discounts, pass the old_price value in this parameter
	Price string `json:"price"`

	// Currency of your prices. The passed value must be the same as the one set in the personal account settings.
	// By default, the passed value is RUB, Russian ruble.
	//
	// For example, if your currency set in the settings is yuan, pass the value CNY, otherwise an error will be returned
	CurrencyCode string `json:"currency_code"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// VAT rate for the product:
	//   - 0 — not subject to VAT,
	//   - 0.1 — 10%,
	//   - 0.2 — 20%
	VAT string `json:"vat"`
}

type CreateProductByOzonIDResponse struct {
	core.CommonResponse

	// Products import task code
	TaskId int64 `json:"task_id"`

	// Products identifiers list
	UnmatchedSKUList []int64 `json:"unmatched_sku_list"`
}

// Creates a product by the specified Ozon ID. The number of products is unlimited.
//
// It's not possible to update products using Ozon ID
func (c Products) CreateProductByOzonID(params *CreateProductByOzonIDParams) (*CreateProductByOzonIDResponse, error) {
	url := "/v1/product/import-by-sku"

	resp := &CreateProductByOzonIDResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateProductImagesParams struct {
	// Marketing color
	ColorImage string `json:"color_image"`

	// Array of links to images. The images in the array are arranged in the order of their arrangement on the site.
	// The first image in the list is the main one for the product.
	//
	// Pass links to images in the public cloud storage. The image format is JPG
	Images []string `json:"images"`

	// Array of 360 images—up to 70 files
	Images360 []string `json:"images360"`

	// Product identfier
	ProductId int64 `json:"product_id"`
}

type ProductInfoResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Pictures
		Pictures []struct {
			// Attribute of a 360 image
			Is360 bool `json:"is_360"`

			// Attribute of a marketing color
			IsColor bool `json:"is_color"`

			// Attribute of a marketing color
			IsPrimary bool `json:"is_primary"`

			// Product identifier
			ProductId int64 `json:"product_id"`

			// Image uploading status.
			//
			// If the `/v1/product/pictures/import` method was called, the response will always be imported—image not processed.
			// To see the final status, call the `/v1/product/pictures/info` method after about 10 seconds.
			//
			// If you called the `/v1/product/pictures/info` method, one of the statuses will appear:
			//   - uploaded — image uploaded;
			//   - failed — image was not uploaded
			State string `json:"state"`

			// The link to the image in the public cloud storage. The image format is JPG or PNG
			URL string `json:"url"`
		} `json:"pictures"`
	} `json:"result"`
}

// The method for uploading and updating product images.
//
// Each time you call the method, pass all the images that should be on the product description page.
// For example, if you call a method and upload 10 images,
// and then call the method a second time and load one imahe, then all 10 previous ones will be erased.
//
// To upload image, pass a link to it in a public cloud storage. The image format is JPG or PNG.
//
// Arrange the pictures in the images array as you want to see them on the site.
// The first picture in the array will be the main one for the product.
//
// You can upload up to 15 pictures for each product.
//
// To upload 360 images, use the images360 field, and to upload a marketing color use color_image.
//
// If you want to add, remove, or replace some images, or change their order,
// first get the details using `/v2/product/info` or `/v2/product/info/list` methods.
// Using them you can get the current list of images and their order.
// Copy the data from the images, images360, and color_image fields and make the necessary changes to it
func (c Products) UpdateProductImages(params *UpdateProductImagesParams) (*ProductInfoResponse, error) {
	url := "/v1/product/pictures/import"

	resp := &ProductInfoResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CheckImageUploadingStatusParams struct {
	// Product identifiers list
	ProductId []int64 `json:"product_id"`
}

// Check products images uploading status
func (c Products) CheckImageUploadingStatus(params *CheckImageUploadingStatusParams) (*ProductInfoResponse, error) {
	url := "/v1/product/pictures/info"

	resp := &ProductInfoResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListProductsByIDsParams struct {
	// Product identifier in the seller's system
	OfferId []string `json:"offer_id"`

	// Product identifier
	ProductId []int64 `json:"product_id"`

	// Product identifier in the Ozon system, SKU
	SKU []int64 `json:"sku"`
}

type ListProductsByIDsResponse struct {
	core.CommonResponse

	// Request results
	Result struct {
		// Data array
		Items []ProductDetails `json:"items"`
	} `json:"result"`
}

// Method for getting an array of products by their identifiers.
//
// The request body must contain an array of identifiers of the same type. The response will contain an items array.
//
// For each shipment in the items array the fields match the ones recieved in the /v2/product/info method
func (c Products) ListProductsByIDs(params *ListProductsByIDsParams) (*ListProductsByIDsResponse, error) {
	url := "/v2/product/info/list"

	resp := &ListProductsByIDsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
