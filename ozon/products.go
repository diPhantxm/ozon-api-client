package ozon

import (
	"context"
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
	LastId string `json:"last_id"`

	// Number of values per page. Minimum is 1, maximum is 1000
	Limit int64 `json:"limit"`

	// Filter by product
	Filter GetStocksInfoFilter `json:"filter"`
}

type GetStocksInfoFilter struct {
	// Filter by the offer_id parameter. It is possible to pass a list of values
	OfferId []string `json:"offer_id,omitempty"`

	// Filter by the product_id parameter. It is possible to pass a list of values
	ProductId []int64 `json:"product_id,omitempty"`

	// Filter by product visibility
	Visibility string `json:"visibility,omitempty"`
}

type GetStocksInfoResponse struct {
	core.CommonResponse

	// Method Result
	Result GetStocksInfoResult `json:"result"`
}

type GetStocksInfoResult struct {
	// Identifier of the last value on the page
	//
	// To get the next values, specify the recieved value in the next request in the last_id parameter
	LastId string `json:"last_id"`

	// The number of unique products for which information about stocks is displayed
	Total int32 `json:"total"`

	// Product details
	Items []GetStocksInfoResultItem `json:"items"`
}

type GetStocksInfoResultItem struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Stock details
	Stocks []GetStocksInfoResultItemStock `json:"stocks"`
}

type GetStocksInfoResultItemStock struct {
	// In a warehouse
	Present int32 `json:"present"`

	// Reserved
	Reserved int32 `json:"reserved"`

	// Warehouse type
	Type string `json:"type" default:"ALL"`
}

// Returns information about the quantity of products in stock:
//
// * how many items are available,
//
// * how many are reserved by customers.
func (c Products) GetStocksInfo(ctx context.Context, params *GetStocksInfoParams) (*GetStocksInfoResponse, error) {
	url := "/v3/product/info/stocks"

	resp := &GetStocksInfoResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductDetailsParams struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id,omitempty"`

	// Product identifier
	ProductId int64 `json:"product_id,omitempty"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku,omitempty"`
}

type GetProductDetailsResponse struct {
	core.CommonResponse

	// Request results
	Result ProductDetails `json:"result"`
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
	DescriptionCategoryId int64 `json:"description_category_id"`

	// Product type identifier
	TypeId int64 `json:"type_id"`

	// Marketing color
	ColorImage string `json:"color_image"`

	// Commission fees details
	Commissions []ProductDetailCommission `json:"commissions"`

	// Date and time when the product was created
	CreatedAt time.Time `json:"created_at"`

	// Product SKU
	SKU int64 `json:"sku"`

	// SKU of the product that is sold from the Ozon warehouse (FBO)
	FBOSKU int64 `json:"fbo_sku,omitempty"`

	// SKU of the product that is sold from the seller's warehouse (FBS and rFBS)
	FBSSKU int64 `json:"fbs_sku,omitempty"`

	// Product identifier
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

	// Product price including discounts. This value is shown on the product description page
	Price string `json:"price"`

	// Product price indexes
	PriceIndexes ProductDetailPriceIndex `json:"price_indexes"`

	// Deprecated: Price index. Learn more in Help Center
	//
	// Use PriceIndexes instead
	PriceIndex string `json:"price_index"`

	// Product state description
	Status ProductDetailStatus `json:"status"`

	// Details about the sources of similar offers. Learn more in Help Сenter
	Sources []ProductDetailSource `json:"sources"`

	// Details about product stocks
	Stocks ProductDetailStock `json:"stocks"`

	// Date of the last product update
	UpdatedAt time.Time `json:"updated_at"`

	// Product VAT rate
	VAT string `json:"vat"`

	// Product visibility settings
	VisibilityDetails ProductDetailVisibilityDetails `json:"visibility_details"`

	// If the product is on sale, the value is true
	Visible bool `json:"visible"`

	// Product volume weight
	VolumeWeight float64 `json:"volume_weight"`

	// 'true' if the item is archived manually.
	IsArchived bool `json:"is_archived"`

	// 'true' if the item is archived automatically.
	IsArchivedAuto bool `json:"is_autoarchived"`
}

type ProductDetailCommission struct {
	// Delivery cost
	DeliveryAmount float64 `json:"deliveryAmount"`

	// Commission percentage
	Percent float64 `json:"percent"`

	// Return cost
	ReturnAmount float64 `json:"returnAmount"`

	// Sale scheme
	SaleSchema string `json:"saleSchema"`

	// Commission fee amount
	Value float64 `json:"value"`
}

type ProductDetailPriceIndex struct {
	// Competitors' product price on other marketplaces
	ExternalIndexData ProductDetailPriceIndexExternal `json:"external_index_data"`

	// Competitors' product price on Ozon
	OzonIndexData ProductDetailPriceIndexOzon `json:"ozon_index_data"`

	// Resulting price index of the product
	PriceIndex string `json:"price_index"`

	// Price of your product on other marketplaces
	SelfMarketplaceIndexData ProductDetailPriceIndexSelfMarketplace `json:"self_marketplaces_index_data"`
}

type ProductDetailPriceIndexExternal struct {
	// Minimum competitors' product price on other marketplaces
	MinimalPrice string `json:"minimal_price"`

	// Price currency
	MinimalPriceCurrency string `json:"minimal_price_currency"`

	// Price index value
	PriceIndexValue float64 `json:"price_index_value"`
}

type ProductDetailPriceIndexOzon struct {
	// Minimum competitors' product price on Ozon
	MinimalPrice string `json:"minimal_price"`

	// Price currency
	MinimalPriceCurrency string `json:"minimal_price_currency"`

	// Price index value
	PriceIndexValue float64 `json:"price_index_value"`
}

type ProductDetailPriceIndexSelfMarketplace struct {
	// Minimum price of your product on other marketplaces
	MinimalPrice string `json:"minimal_price"`

	// Price currency
	MinimalPriceCurrency string `json:"minimal_price_currency"`

	// Price index value
	PriceIndexValue float64 `json:"price_index_value"`
}

type ProductDetailStatus struct {
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
}

type ProductDetailSource struct {
	// Indication that the source is taken into account when calculating the market value
	IsEnabled bool `json:"is_enabled"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Link to the source
	Source string `json:"source"`
}

type ProductDetailStock struct {
	// Supply expected
	Coming int32 `json:"coming"`

	// Currently at the warehouse
	Present int32 `json:"present"`

	// Reserved
	Reserved int32 `json:"reserved"`
}

type ProductDetailVisibilityDetails struct {
	// If the product is active, the value is true
	//
	// Deprecated: Use `visible` parameter of `ProductDetails`
	ActiveProduct bool `json:"active_product"`

	// If the price is set, the value is true
	HasPrice bool `json:"has_price"`

	// If there is stock at the warehouses, the value is true
	HasStock bool `json:"has_stock"`

	// Reason why the product is hidden
	Reasons map[string]interface{} `json:"reasons"`
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
//
// Check a minimum product price with all promotions applied in your personal account.
// The min_price parameter from the method response is in development and returns 0
func (c Products) GetProductDetails(ctx context.Context, params *GetProductDetailsParams) (*GetProductDetailsResponse, error) {
	url := "/v2/product/info"

	resp := &GetProductDetailsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Result []UpdateStocksResult `json:"result"`
}

type UpdateStocksResult struct {
	// An array of errors that occurred while processing the request
	Errors []UpdateStocksResultError `json:"errors"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// If the product details have been successfully updated — true
	Updated bool `json:"updated"`
}

type UpdateStocksResultError struct {
	// Error code
	Code string `json:"code"`

	// Error reason
	Message string `json:"message"`
}

// Allows you to change the products in stock quantity. The method is only used for FBS and rFBS warehouses.
//
// With one request you can change the availability for 100 products. You can send up to 80 requests in a minute.
//
// Availability can only be set after the product status has been changed to processed.
func (c Products) UpdateStocks(ctx context.Context, params *UpdateStocksParams) (*UpdateStocksResponse, error) {
	url := "/v1/product/import/stocks"

	resp := &UpdateStocksResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Result []UpdateQuantityStockProductsResult `json:"result"`
}

type UpdateQuantityStockProductsResult struct {
	// An array of errors that occurred while processing the request
	Errors []UpdateQuantityStockProductsResultError `json:"errors"`

	// Product identifier in the seller's system
	Offerid string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// If the request was completed successfully and the stocks are updated — true
	Updated bool `json:"updated"`

	// Warehouse identifier derived from the /v1/warehouse/list method
	WarehouseId int64 `json:"warehouse_id"`
}

type UpdateQuantityStockProductsResultError struct {
	// Error code
	Code string `json:"code"`

	// Error reason
	Message string `json:"message"`
}

// Allows you to change the products in stock quantity.
//
// With one request you can change the availability for 100 products. You can send up to 80 requests in a minute.
//
// You can update the stock of one product in one warehouse only once in 2 minutes, otherwise there will be the TOO_MANY_REQUESTS error in the response.
//
// You can set the availability of an item only after the product status is changed to price_sent
//
// Bulky products stock can only be updated in the warehouses for bulky products.
func (c Products) UpdateQuantityStockProducts(ctx context.Context, params *UpdateQuantityStockProductsParams) (*UpdateQuantityStockProductsResponse, error) {
	url := "/v2/products/stocks"

	resp := &UpdateQuantityStockProductsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type StocksInSellersWarehouseParams struct {
	// Product SKU
	SKU []string `json:"sku"`
}

type StocksInSellersWarehouseResponse struct {
	core.CommonResponse

	// Method result
	Result []StocksInSellersWarehouseResult `json:"result"`
}

type StocksInSellersWarehouseResult struct {
	// SKU of the product that is sold from the seller's warehouse (FBS and RFBS schemes)
	SKU int64 `json:"sku"`

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

// Get stocks in seller's warehouse
func (c Products) StocksInSellersWarehouse(ctx context.Context, params *StocksInSellersWarehouseParams) (*StocksInSellersWarehouseResponse, error) {
	url := "/v1/product/info/stocks-by-warehouse/fbs"

	resp := &StocksInSellersWarehouseResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	// Specified in rubles.
	// The fractional part is separated by decimal point,
	// up to two digits after the decimal point.
	//
	// If there are no discounts on the product, pass 0 to this field and specify the correct price in the price field
	OldPrice string `json:"old_price"`

	// Product price including discounts. This value is displayed on the product description page.
	//
	// If the old_price parameter value is greater than 0,
	// there should be a certain difference between price and old_price.
	// It depends on the price value
	//
	// < 400 - min diff. 20 rubles
	//
	// 400-10,000 - min diff. 5%
	//
	// > 10,000 - min diff. 500 rubles
	Price string `json:"price"`

	// Attribute for enabling and disabling pricing strategies auto-application
	//
	// If you've previously enabled automatic application of pricing strategies and don't want to disable it, pass UNKNOWN in the next requests.
	//
	// If you pass `ENABLED` in this parameter, pass `strategy_id` in the `/v1/pricing-strategy/products/add` method request.
	//
	// If you pass `DISABLED` in this parameter, the product is removed from the strategy
	PriceStrategyEnabled PriceStrategy `json:"price_strategy_enabled"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

type UpdatePricesResponse struct {
	core.CommonResponse

	Result []UpdatePricesResult `json:"result"`
}

type UpdatePricesResult struct {
	// An array of errors that occurred while processing the request
	Errors []UpdatePricesResultError `json:"errors"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product ID
	ProductId int64 `json:"product_id"`

	// If the product details have been successfully updated — true
	Updated bool `json:"updated"`
}

type UpdatePricesResultError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Allows you to change a price of one or more products.
// You can change prices for 1000 products in one request.
// To reset old_price, set 0 for this parameter
func (c Products) UpdatePrices(ctx context.Context, params *UpdatePricesParams) (*UpdatePricesResponse, error) {
	url := "/v1/product/import/prices"

	resp := &UpdatePricesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	DescriptionCategoryId int64 `json:"description_category_id"`

	// New category identifier. Specify it if you want to change the current product category
	NewDescriptinoCategoryId int64 `json:"new_description_category_id"`

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

	// Product price including discounts. This value is shown on the product description card.
	// If there are no discounts on the product, specify the old_price value
	Price string `json:"price"`

	// Default: "IS_CODE_SERVICE"
	// Service type. Pass one of the values in upper case:
	//   - IS_CODE_SERVICE,
	//   - IS_NO_CODE_SERVICE
	ServiceType string `json:"service_type" default:"IS_CODE_SERVICE"`

	// Product type identifier.
	// You can get values from the type_id parameter in the `/v1/description-category/tree` method response.
	// When filling this parameter in,
	// you can leave out the attributes attribute if it has the `id:8229` parameter
	TypeId int64 `json:"type_id"`

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
	DictionaryValueId int64 `json:"dictionary_value_id"`

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
	Result CreateOrUpdateProductResult `json:"result"`
}

type CreateOrUpdateProductResult struct {
	// Number of task for products upload
	TaskId int64 `json:"task_id"`
}

// This method allows you to create products and update their details
// More info: https://docs.ozon.ru/api/seller/en/#operation/ProductAPI_ImportProductsV3
func (c Products) CreateOrUpdateProduct(ctx context.Context, params *CreateOrUpdateProductParams) (*CreateOrUpdateProductResponse, error) {
	url := "/v3/product/import"

	resp := &CreateOrUpdateProductResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Result GetListOfProductsResult `json:"result"`
}

type GetListOfProductsResult struct {
	// Products list
	Items []GetListOfProductsResultItem `json:"items"`

	// Identifier of the last value on the page.
	//
	// To get the next values, specify the recieved value in the next request in the last_id parameter
	LastId string `json:"last_id"`

	// Total number of products
	Total int32 `json:"total"`
}

type GetListOfProductsResultItem struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product ID
	ProductId int64 `json:"product_id"`
}

func (c Products) GetListOfProducts(ctx context.Context, params *GetListOfProductsParams) (*GetListOfProductsResponse, error) {
	url := "/v2/product/list"

	resp := &GetListOfProductsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Products []GetProductsRatingbySKUProduct `json:"products"`
}

type GetProductsRatingbySKUProduct struct {
	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Product content rating: 0 to 100
	Rating float64 `json:"rating"`

	// Groups of characteristics that make up the content rating
	Groups []GetProductsRatingBySKUProductGroup `json:"groups"`
}

type GetProductsRatingBySKUProductGroup struct {
	// List of conditions that increase the product content rating
	Conditions []GetProductsRatingBySKUProductGroupCondition `json:"conditions"`

	// Number of attributes you need to fill in to get the maximum score in this characteristics group
	ImproveAtLeast int32 `json:"improve_at_least"`

	// List of attributes that can increase the product content rating
	ImproveAttributes []GetProductsRatingBySKUProductGroupImproveAttr `json:"improve_attributes"`

	// Group identifier
	Key string `json:"key"`

	// Group name
	Name string `json:"name"`

	// Rating in the group
	Rating float64 `json:"rating"`

	// Percentage influence of group characteristics on the content rating
	Weight float64 `json:"weight"`
}

type GetProductsRatingBySKUProductGroupCondition struct {
	// Number of content rating points that the condition gives
	Cost float64 `json:"cost"`

	// Condition description
	Description string `json:"description"`

	// Indication that the condition is met
	Fulfilled bool `json:"fulfilled"`

	// Condition identifier
	Key string `json:"key"`
}

type GetProductsRatingBySKUProductGroupImproveAttr struct {
	// Attribute identifier
	Id int64 `json:"id"`

	// Attribute name
	Name string `json:"name"`
}

// Method for getting products' content rating and recommendations on how to increase it
func (c Products) GetProductsRatingBySKU(ctx context.Context, params *GetProductsRatingBySKUParams) (*GetProductsRatingBySKUResponse, error) {
	url := "/v1/product/rating-by-sku"

	resp := &GetProductsRatingBySKUResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Result GetProductImportStatusResult `json:"result"`
}

type GetProductImportStatusResult struct {
	// Product details
	Items []GetProductImportStatusResultItem `json:"items"`

	// Product identifier in the seller's system
	Total int32 `json:"total"`
}

type GetProductImportStatusResultItem struct {
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
	Errors []GetProductImportStatusResultItemError `json:"errors"`
}

type GetProductImportStatusResultItemError struct {
	GetProductDetailsResponseItemError

	// Error technical description
	Message string `json:"message"`
}

// Allows you to get the status of a product description page creation process
func (c Products) GetProductImportStatus(ctx context.Context, params *GetProductImportStatusParams) (*GetProductImportStatusResponse, error) {
	url := "/v1/product/import/info"

	resp := &GetProductImportStatusResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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

	Result CreateProductByOzonIDResult `json:"result"`
}

type CreateProductByOzonIDResult struct {
	// Products import task code
	TaskId int64 `json:"task_id"`

	// Products identifiers list
	UnmatchedSKUList []int64 `json:"unmatched_sku_list"`
}

// Creates a product by the specified Ozon ID. The number of products is unlimited.
//
// It's not possible to update products using Ozon ID
func (c Products) CreateProductByOzonID(ctx context.Context, params *CreateProductByOzonIDParams) (*CreateProductByOzonIDResponse, error) {
	url := "/v1/product/import-by-sku"

	resp := &CreateProductByOzonIDResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Result ProductInfoResult `json:"result"`
}

type ProductInfoResult struct {
	// Pictures
	Pictures []ProductInfoResultPicture `json:"pictures"`
}

type ProductInfoResultPicture struct {
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
	//   - pending — image was not uploaded
	State string `json:"state"`

	// The link to the image in the public cloud storage. The image format is JPG or PNG
	URL string `json:"url"`
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
func (c Products) UpdateProductImages(ctx context.Context, params *UpdateProductImagesParams) (*ProductInfoResponse, error) {
	url := "/v1/product/pictures/import"

	resp := &ProductInfoResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
func (c Products) CheckImageUploadingStatus(ctx context.Context, params *CheckImageUploadingStatusParams) (*ProductInfoResponse, error) {
	url := "/v1/product/pictures/info"

	resp := &ProductInfoResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Result ListProductsByIDsResult `json:"result"`
}

type ListProductsByIDsResult struct {
	// Data array
	Items []ProductDetails `json:"items"`
}

// Method for getting an array of products by their identifiers.
//
// The request body must contain an array of identifiers of the same type. The response will contain an items array.
//
// For each shipment in the items array the fields match the ones recieved in the /v2/product/info method
func (c Products) ListProductsByIDs(ctx context.Context, params *ListProductsByIDsParams) (*ListProductsByIDsResponse, error) {
	url := "/v2/product/info/list"

	resp := &ListProductsByIDsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetDescriptionOfProductParams struct {
	// Filter by product
	Filter GetDescriptionOfProductFilter `json:"filter"`

	// Identifier of the last value on the page. Leave this field blank in the first request.
	//
	// To get the next values, specify `last_id` from the response of the previous request
	LastId string `json:"last_id"`

	// Number of values per page. Minimum is 1, maximum is 1000
	Limit int64 `json:"limit"`

	// The parameter by which the products will be sorted
	SortBy string `json:"sort_by,omitempty"`

	// Sorting direction
	SortDirection string `json:"sort_dir,omitempty"`
}

type GetDescriptionOfProductFilter struct {
	// Filter by the `offer_id` parameter. It is possible to pass a list of values
	OfferId []string `json:"offer_id"`

	// Filter by the product_id parameter. It is possible to pass a list of values
	ProductId []int64 `json:"product_id"`

	// Filter by product visibility
	Visibility string `json:"visibility"`
}

type GetDescriptionOfProductResponse struct {
	core.CommonResponse

	// Request results
	Result []GetDescriptionOfProductResult `json:"result"`

	// Identifier of the last value on the page.
	//
	// To get the next values, specify the recieved value in the next request in the last_id parameter
	LastId string `json:"last_id"`

	// Number of products in the list
	Total int32 `json:"total"`
}

type GetDescriptionOfProductResult struct {
	// Array of product characteristics
	Attributes []GetDescriptionOfProductResultAttr `json:"attributes"`

	// Barcode
	Barcode string `json:"barcode"`

	// Category identifier
	DescriptionCategoryId int64 `json:"description_category_id"`

	// Marketing color
	ColorImage string `json:"color_image"`

	// Array of nested characteristics
	ComplexAttributes []GetDescriptionOfProductResultComplexAttrs `json:"complex_attributes"`

	// Depth
	Depth int32 `json:"depth"`

	// Dimension measurement units:
	//   - mm — millimeters,
	//   - cm — centimeters,
	//   - in — inches
	DimensionUnit string `json:"dimension_unit"`

	// Package height
	Height int32 `json:"height"`

	// Product characteristic identifier
	Id int64 `json:"id"`

	// Array of links to product images
	Images []GetDescriptionOfProductResultImage `json:"images"`

	// Array of 360 images
	Images360 []GetDescriptionOfProductResultImage360 `json:"images360"`

	// Product name. Up to 500 characters
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Array of PDF files
	PDFList []GetDescriptionOfProductResultPDF `json:"pdf_list"`

	// Product type identifier
	TypeId int64 `json:"type_id"`

	// Weight of product in the package
	Weight int32 `json:"weight"`

	// Weight measurement unit
	WeightUnit string `json:"weight_unit"`

	// Package width
	Width int32 `json:"width"`
}

type GetDescriptionOfProductResultAttr struct {
	// Characteristic identifier
	AttributeId int64 `json:"attribute_id"`

	// Identifier of the characteristic that supports nested properties.
	// For example, the "Processor" characteristic has nested characteristics "Manufacturer" and "L2 Cache".
	// Each of the nested characteristics can have multiple value variants
	ComplexId int64 `json:"complex_id"`

	// Array of characteristic values
	Values []GetDescriptionOfProductResultAttrValue `json:"values"`
}

type GetDescriptionOfProductResultAttrValue struct {
	// Characteristic identifier in the dictionary
	DictionaryValueId int64 `json:"dictionary_value_id"`

	// Product characteristic value
	Value string `json:"value"`
}

type GetDescriptionOfProductResultComplexAttrs struct {
	// Array of product characteristics
	Attributes []GetDescriptionOfProductResultComplexAttr `json:"attributes"`
}

type GetDescriptionOfProductResultComplexAttr struct {
	// Characteristic identifier
	AttributeId int64 `json:"attribute_id"`

	// Identifier of the characteristic that supports nested properties.
	// For example, the "Processor" characteristic has nested characteristics "Manufacturer" and "L2 Cache".
	// Each of the nested characteristics can have multiple value variants
	ComplexId int64 `json:"complex_id"`

	// Array of characteristic values
	Values []GetDescriptionOfProductResultComplexAttrValue `json:"values"`
}

type GetDescriptionOfProductResultComplexAttrValue struct {
	// Characteristic identifier in the dictionary
	DictionaryValueId int64 `json:"dictionary_value_id"`

	// Product characteristic value
	Value string `json:"value"`
}

type GetDescriptionOfProductResultImage struct {
	Default  bool   `json:"default"`
	FileName string `json:"file_name"`
	Index    int64  `json:"index"`
}

type GetDescriptionOfProductResultImage360 struct {
	FileName string `json:"file_name"`
	Index    int64  `json:"index"`
}

type GetDescriptionOfProductResultPDF struct {
	// Path to PDF file
	FileName string `json:"file_name"`

	// Storage order index
	Index int64 `json:"index"`

	// File name
	Name string `json:"name"`
}

// Returns a product characteristics description by product identifier. You can search for the product by `offer_id` or `product_id`
func (c Products) GetDescriptionOfProduct(ctx context.Context, params *GetDescriptionOfProductParams) (*GetDescriptionOfProductResponse, error) {
	url := "/v3/products/info/attributes"

	resp := &GetDescriptionOfProductResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductDescriptionParams struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

type GetProductDescriptionResponse struct {
	core.CommonResponse

	// Method result
	Result GetProductDescriptionResult `json:"result"`
}

type GetProductDescriptionResult struct {
	// Description
	Description string `json:"description"`

	// Identifier
	Id int64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`
}

// Get product description
func (c Products) GetProductDescription(ctx context.Context, params *GetProductDescriptionParams) (*GetProductDescriptionResponse, error) {
	url := "/v1/product/info/description"

	resp := &GetProductDescriptionResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductRangeLimitResponse struct {
	core.CommonResponse

	// Daily product creation limit
	DailyCreate GetProductRangeLimitUploadQuota `json:"daily_create"`

	// Daily product update limit
	DailyUpdate GetProductRangeLimitUploadQuota `json:"daily_update"`

	// Product range limit
	Total GetProductRangeLimitTotal `json:"total"`
}

type GetProductRangeLimitTotal struct {
	// How many products you can create in your personal account
	Limit int64 `json:"limit"`

	// How many products you've already created
	Usage int64 `json:"usage"`
}

type GetProductRangeLimitUploadQuota struct {
	// How many products you can create in one day
	Limit int64 `json:"limit"`

	// Counter reset time for the current day in UTC format
	ResetAt time.Time `json:"reset_at"`

	// How many products you've created in the current day
	Usage int64 `json:"usage"`
}

// Method for getting information about the following limits:
//   - Product range limit: how many products you can create in your personal account.
//   - Products creation limit: how many products you can create per day.
//   - Products update limit: how many products you can update per day.
//
// If you have a product range limit and you exceed it, you won't be able to create new products
func (c Products) GetProductRangeLimit(ctx context.Context) (*GetProductRangeLimitResponse, error) {
	url := "/v4/product/info/limit"

	resp := &GetProductRangeLimitResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ChangeProductIDsParams struct {
	// List of pairs with new and old values of product identifiers
	UpdateOfferId []ChangeProductIDsUpdateOffer `json:"update_offer_id"`
}

type ChangeProductIDsUpdateOffer struct {
	// New product identifier
	//
	// The maximum length of a string is 50 characters
	NewOfferId string `json:"new_offer_id"`

	// Old product identifier
	OfferId string `json:"offer_id"`
}

type ChangeProductIDsResponse struct {
	core.CommonResponse

	// Errors list
	Errors []ChangeProductIDsError `json:"errors"`
}

type ChangeProductIDsError struct {
	// Error message
	Message string `json:"message"`

	// Product identifier that wasn't changed
	OfferId string `json:"offer_id"`
}

// Method for changing the offer_id linked to products. You can change multiple offer_id in this method.
//
// We recommend transmitting up to 250 values in an array
func (c Products) ChangeProductIDs(ctx context.Context, params *ChangeProductIDsParams) (*ChangeProductIDsResponse, error) {
	url := "/v1/product/update/offer-id"

	resp := &ChangeProductIDsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ArchiveProductParams struct {
	// Product identifier
	ProductId []int64 `json:"product_id"`
}

type ArchiveProductResponse struct {
	core.CommonResponse

	// The result of processing the request. true if the request was executed without errors
	Result bool `json:"result"`
}

// Archive product
func (c Products) ArchiveProduct(ctx context.Context, params *ArchiveProductParams) (*ArchiveProductResponse, error) {
	url := "/v1/product/archive"

	resp := &ArchiveProductResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

// Warning: Since June 14, 2023 the method is disabled.
//
// Unarchive product
func (c Products) UnarchiveProduct(ctx context.Context, params *ArchiveProductParams) (*ArchiveProductResponse, error) {
	url := "/v1/product/unarchive"

	resp := &ArchiveProductResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type RemoveProductWithoutSKUParams struct {
	// Product identifier
	Products []RemoveProductWithoutSKUProduct `json:"products"`
}

type RemoveProductWithoutSKUProduct struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`
}

type RemoveProductWithoutSKUResponse struct {
	core.CommonResponse

	// Product processing status
	Status []RemoveProductWithoutSKUStatus `json:"status"`
}

type RemoveProductWithoutSKUStatus struct {
	// Reason of the error that occurred while processing the request
	Error string `json:"error"`

	// If the request was executed without errors and the products were deleted, the value is true
	IsDeleted bool `json:"is_deleted"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`
}

// Remove a product without an SKU from the archive
//
// You can pass up to 500 identifiers in one request
func (c Products) RemoveProductWithoutSKU(ctx context.Context, params *RemoveProductWithoutSKUParams) (*RemoveProductWithoutSKUResponse, error) {
	url := "/v2/products/delete"

	resp := &RemoveProductWithoutSKUResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListGeoRestrictionsParams struct {
	// Filter. To get all geo-restrictions, leave names blank and specify true in the only_visible parameter
	Filter ListGeoRestrictionsFilter `json:"filter"`

	// Order number of geo-restriction from which to output data in the response.
	//
	// If you specify 23 in this parameter, the first item in the restrictions list will output order_number = 24.
	// If you want to get all geo-restrictions, pass 0 in this parameter
	LastOrderNumber int64 `json:"last_order_number"`

	// Number of items in the response
	Limit int64 `json:"limit"`
}

type ListGeoRestrictionsFilter struct {
	// List with city names
	Names []string `json:"names"`

	// Value visibility. We recommend always passing true in this parameter
	OnlyVisible bool `json:"only_visible"`
}

type ListGeoRestrictionsResponse struct {
	core.CommonResponse

	// Restrictions
	Restrictions []ListGeoRestrictionsRestriction `json:"restrictions"`
}

type ListGeoRestrictionsRestriction struct {
	// Geo-restriction identifier
	Id string `json:"id"`

	// Item visibility
	IsVisible bool `json:"is_visible"`

	// City name
	Name string `json:"name"`

	// Geo-restriction order number.
	//
	// If you specify 23 in the last_order_number parameter in the request,
	// the first item in the restrictions list will have order_number = 24
	OrderNumber int64 `json:"order_number"`
}

// Deprecated: Get a list of geo-restrictions for services
func (c Products) ListGeoRestrictions(ctx context.Context, params *ListGeoRestrictionsParams) (*ListGeoRestrictionsResponse, error) {
	url := "/v1/products/geo-restrictions-catalog-by-filter"

	resp := &ListGeoRestrictionsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UploadActivationCodesParams struct {
	// Digital activation codes
	DigitalCodes []string `json:"digital_codes"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

type UploadActivationCodesResponse struct {
	core.CommonResponse

	// Method result
	Result UploadActivationCodesResult `json:"result"`
}

type UploadActivationCodesResult struct {
	// Uploading digital code task identifier
	TaskId int64 `json:"task_id"`
}

// Upload activation codes when you upload service or digital products. Activation code is associated with the digital product card
func (c Products) UploadActivationCodes(ctx context.Context, params *UploadActivationCodesParams) (*UploadActivationCodesResponse, error) {
	url := "/v1/product/upload_digital_codes"

	resp := &UploadActivationCodesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type StatusOfUploadingActivationCodesParams struct {
	// Uploading activation codes task identifier that was received from the `/v1/product/upload_digital_codes` method
	TaskId int64 `json:"task_id"`
}

type StatusOfUploadingActivationCodesResponse struct {
	core.CommonResponse

	// Method result
	Result StatusOfUploadingActivationCodesResult `json:"result"`
}

type StatusOfUploadingActivationCodesResult struct {
	// Upload status:
	//   - pending — products in queue for processing.
	//   - imported — the product has been successfully uploaded.
	//   - failed — the product was uploaded with errors
	Status string `json:"status"`
}

// Get status of uploading activation codes task for services and digital products
func (c Products) StatusOfUploadingActivationCodes(ctx context.Context, params *StatusOfUploadingActivationCodesParams) (*StatusOfUploadingActivationCodesResponse, error) {
	url := "/v1/product/upload_digital_codes/info"

	resp := &StatusOfUploadingActivationCodesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductPriceInfoParams struct {
	// Filter by product
	Filter GetProductPriceInfoFilter `json:"filter"`

	// Identifier of the last value on page.
	//
	// To get the next values, specify the recieved value in the next request in the `last_id` parameter
	LastId string `json:"last_id"`

	// Number of values per page. Minimum is 1, maximum is 1000
	Limit int32 `json:"limit"`
}

type GetProductPriceInfoFilter struct {
	// Filter by the `offer_id` parameter. It is possible to pass a list of values
	OfferId []string `json:"offer_id"`

	// Filter by the `product_id` parameter. It is possible to pass a list of up to 1000 values
	ProductId []int64 `json:"product_id"`

	// Filter by product visibility
	Visibility string `json:"visibility" default:"ALL"`
}

type GetProductPriceInfoResponse struct {
	core.CommonResponse

	// Result
	Result GetProductPriceInfoResult `json:"result"`
}

type GetProductPriceInfoResult struct {
	// Products list
	Items []GetProductPriceInfoResultItem `json:"items"`

	// Identifier of the last value on page. Leave this field blank in the first request.
	//
	// To get the next values, specify last_id from the response of the previous request
	LastId string `json:"last_id"`

	// Products number in the list
	Total int32 `json:"total"`
}

type GetProductPriceInfoResultItem struct {
	// Maximum acquiring fee
	Acquiring int32 `json:"acquiring"`

	// Commissions information
	Commissions GetProductPriceInfoResultItemCommission `json:"commissions"`

	// Promotions information
	MarketingActions *GetProductPriceInfoResultItemMarketingActions `json:"marketing_actions"`

	// Seller product identifier
	OfferId string `json:"offer_id"`

	// Product price
	Price GetProductPriceInfoResultItemPrice `json:"price"`

	// Deprected: price index
	//
	// Use PriceIndexes instead
	PriceIndex string `json:"price_index"`

	// Product price indexes
	PriceIndexes GetProductPriceInfoResultItemPriceIndexes `json:"price_indexes"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product volume weight
	VolumeWeight float64 `json:"volume_weight"`
}

type GetProductPriceInfoResultItemCommission struct {
	// Last mile (FBO)
	FBOLastMile float64 `json:"fbo_deliv_to_customer_amount"`

	// Pipeline to (FBO)
	FBOPipelineTo float64 `json:"fbo_direct_flow_trans_max_amount"`

	// Pipeline from (FBO)
	FBOPipelineFrom float64 `json:"fbo_direct_flow_trans_min_amount"`

	// Order packaging fee (FBO)
	FBOOrderPackagingFee float64 `json:"fbo_fulfillment_amount"`

	// Return and cancellation fees (FBO)
	FBOReturnCancellationFee float64 `json:"fbo_return_flow_amount"`

	// Reverse logistics fee from (FBO)
	FBOReverseLogisticsFeeFrom float64 `json:"fbo_return_flow_trans_min_amount"`

	// Reverse logistics fee to (FBO)
	FBOReverseLogisticsFeeTo float64 `json:"fbo_return_flow_trans_max_amount"`

	// Last mile (FBS)
	FBSLastMile float64 `json:"fbs_deliv_to_customer_amount"`

	// Pipeline to (FBS)
	FBSPipelineTo float64 `json:"fbs_direct_flow_trans_max_amount"`

	// Pipeline from (FBS)
	FBSPipelineFrom float64 `json:"fbs_direct_flow_trans_min_amount"`

	// Minimal shipment processing fee (FBS) — 0 rubles
	FBSShipmentProcessingToFee float64 `json:"fbs_first_mile_min_amount"`

	// Maximal shipment processing fee (FBS) — 25 rubles
	FBSShipmentProcessingFromFee float64 `json:"fbs_first_mile_max_amount"`

	// Return and cancellation fees, shipment processing (FBS)
	FBSReturnCancellationProcessingFee float64 `json:"fbs_return_flow_amount"`

	// Return and cancellation fees, pipeline to (FBS)
	FBSReturnCancellationToFees float64 `json:"fbs_return_flow_trans_max_amount"`

	// Return and cancellation fees, pipeline from (FBS)
	FBSReturnCancellationFromFees float64 `json:"fbs_return_flow_trans_min_amount"`

	// Sales commission percentage (FBO)
	SalesCommissionFBORate float64 `json:"sales_percent_fbo"`

	// Sales commission percentage (FBS)
	SalesCommissionFBSRate float64 `json:"sales_percent_fbs"`

	// Larger sales commission percentage among FBO and FBS
	SalesCommissionRate float64 `json:"sales_percent"`
}

type GetProductPriceInfoResultItemMarketingActions struct {
	// Seller's promotions. The parameters date_from, date_to, discount_value and title are specified for each seller's promotion
	Actions []GetProductPriceInfoResultItemMarketingActionsAction `json:"actions"`

	// Current period start date and time for all current promotions
	CurrentPeriodFrom time.Time `json:"current_period_from"`

	// Current period end date and time for all current promotions
	CurrentPeriodTo time.Time `json:"current_period_to"`

	// If a promotion can be applied to the product at the expense of Ozon, this field is set to true
	OzonActionsExist bool `json:"ozon_actions_exist"`
}

type GetProductPriceInfoResultItemMarketingActionsAction struct {
	// Date and time when the seller's promotion starts
	DateFrom time.Time `json:"date_from"`

	// Date and time when the seller's promotion ends
	DateTo time.Time `json:"date_to"`

	// Discount on the seller's promotion
	DiscountValue string `json:"discount_value"`

	// Promotion name
	Title string `json:"title"`
}

type GetProductPriceInfoResultItemPrice struct {
	// If promos auto-application is enabled, the value is true
	AutoActionEnabled bool `json:"auto_action_enabled"`

	// Currency of your prices. It matches the currency set in the personal account settings
	CurrencyCode string `json:"currency_code"`

	// Product price including all promotion discounts. This value will be indicated on the Ozon storefront
	MarketingPrice string `json:"marketing_price"`

	// Product price with seller's promotions applied
	MarketingSellerPrice string `json:"marketing_seller_price"`

	// Minimum price for similar products on Ozon
	MinOzonPrice string `json:"min_ozon_price"`

	// Minimum product price with all promotions applied
	MinPrice string `json:"min_price"`

	// Price before discounts. Displayed strikethrough on the product description page
	OldPrice string `json:"old_price"`

	// Product price including discounts. This value is shown on the product description page
	Price string `json:"price"`

	// Retailer price
	RetailPrice string `json:"retail_price"`

	// Product VAT rate
	VAT string `json:"vat"`
}

type GetProductPriceInfoResultItemPriceIndexes struct {
	// Competitors' product price on other marketplaces
	ExternalIndexData GetProductPriceInfoResultItemPriceIndexesExternal `json:"external_index_data"`

	// Competitors' product price on Ozon
	OzonIndexData GetProductPriceInfoResultItemPriceIndexesOzon `json:"ozon_index_data"`

	// Resulting price index of the product
	PriceIndex string `json:"price_index"`

	// Price of your product on other marketplaces
	SelfMarketplaceIndexData GetProductPriceInfoResultItemPriceIndexesSelfMarketplace `json:"self_marketplaces_index_data"`
}

type GetProductPriceInfoResultItemPriceIndexesExternal struct {
	// Minimum competitors' product price on other marketplaces
	MinimalPrice string `json:"minimal_price"`

	// Price currency
	MinimalPriceCurrency string `json:"minimal_price_currency"`

	// Price index value
	PriceIndexValue float64 `json:"price_index_value"`
}

type GetProductPriceInfoResultItemPriceIndexesOzon struct {
	// Minimum competitors' product price on Ozon
	MinimalPrice string `json:"minimal_price"`

	// Price currency
	MinimalPriceCurrency string `json:"minimal_price_currency"`

	// Price index value
	PriceIndexValue float64 `json:"price_index_value"`
}

type GetProductPriceInfoResultItemPriceIndexesSelfMarketplace struct {
	// Minimum price of your product on other marketplaces
	MinimalPrice string `json:"minimal_price"`

	// Price currency
	MinimalPriceCurrency string `json:"minimal_price_currency"`

	// Price index value
	PriceIndexValue float64 `json:"price_index_value"`
}

// You can specify up to 1000 products in the request
//
// Check minimum and maximum commissions for FBO pipeline in your personal account.
// The `fbo_direct_flow_trans_max_amount` and `fbo_direct_flow_trans_min_amount` parameters
// from the method response are in development and return 0
func (c Products) GetProductPriceInfo(ctx context.Context, params *GetProductPriceInfoParams) (*GetProductPriceInfoResponse, error) {
	url := "/v4/product/info/prices"

	resp := &GetProductPriceInfoResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetMarkdownInfoParams struct {
	// Markdown products SKUs list
	DiscountedSKUs []string `json:"discounted_skus"`
}

type GetMarkdownInfoResponse struct {
	core.CommonResponse

	// Information about the markdown and the main product
	Items []GetMarkdownInfoItem `json:"items"`
}

type GetMarkdownInfoItem struct {
	// Comment on the damage reason
	CommentReasonDamaged string `json:"comment_reason_damaged"`

	// Product condition: new or used
	Condition string `json:"condition"`

	// Product condition on a 1 to 7 scale.
	//   - 1 — satisfactory,
	//   - 2 — good,
	//   - 3 — very good,
	//   - 4 — excellent,
	//   - 5–7 — like new
	ConditionEstimation string `json:"condition_estimation"`

	// Product defects
	Defects string `json:"defects"`

	// Markdown product SKU
	DiscountedSKU int64 `json:"discounted_sku"`

	// Mechanical damage description
	MechanicalDamage string `json:"mechanical_damage"`

	// Packaging damage description
	PackageDamage string `json:"package_damage"`

	// Indication of package integrity damage
	PackagingViolation string `json:"packaging_violation"`

	// Damage reason
	ReasonDamaged string `json:"reason_damaged"`

	// Indication of repaired product
	Repair string `json:"repair"`

	// Indication that the product is incomplete
	Shortage string `json:"shortage"`

	// Main products SKU
	SKU int64 `json:"sku"`

	// Indication that the product has a valid warranty
	WarrantyType string `json:"warranty_type"`
}

// Get information about the markdown and the main product by the markdown product SKU
//
// A method for getting information about the condition and defects of a markdown product by its SKU.
// The method also returns the SKU of the main product
func (c Products) GetMarkdownInfo(ctx context.Context, params *GetMarkdownInfoParams) (*GetMarkdownInfoResponse, error) {
	url := "/v1/product/info/discounted"

	resp := &GetMarkdownInfoResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type SetDiscountOnMarkdownProductParams struct {
	// Discount amount: from 3 to 99 percents
	Discount int32 `json:"discount"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

type SetDiscountOnMarkdownProductResponse struct {
	core.CommonResponse

	// Method result. true if the query was executed without errors
	Result bool `json:"result"`
}

// A method for setting the discount percentage on markdown products sold under the FBS scheme
func (c Products) SetDiscountOnMarkdownProduct(ctx context.Context, params *SetDiscountOnMarkdownProductParams) (*SetDiscountOnMarkdownProductResponse, error) {
	url := "/v1/product/update/discount"

	resp := &SetDiscountOnMarkdownProductResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type NumberOfSubsToProductAvailabilityParams struct {
	// List of SKUs, product identifiers in the Ozon system
	SKUS []int64 `json:"skus"`
}

type NumberOfSubsToProductAvailabilityResponse struct {
	core.CommonResponse

	// Method result
	Result []NumberOfSubsToProductAvailabilityResult `json:"result"`
}

type NumberOfSubsToProductAvailabilityResult struct {
	// Number of subscribed users
	Count int64 `json:"count"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

// You can pass multiple products in a request
func (c Products) NumberOfSubsToProductAvailability(ctx context.Context, params *NumberOfSubsToProductAvailabilityParams) (*NumberOfSubsToProductAvailabilityResponse, error) {
	url := "/v1/product/info/subscription"

	resp := &NumberOfSubsToProductAvailabilityResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateCharacteristicsParams struct {
	// Products and characteristics to be updated
	Items []UpdateCharacteristicsItem `json:"items"`
}

type UpdateCharacteristicsItem struct {
	// Product characteristics
	Attributes []UpdateCharacteristicsItemAttribute `json:"attributes"`

	// Product ID
	OfferId string `json:"offer_id"`
}

type UpdateCharacteristicsItemAttribute struct {
	// Identifier of the characteristic that supports nested properties.
	// Each of the nested characteristics can have multiple value variants
	ComplexId int64 `json:"complex_id"`

	// Characteristic identifier
	Id int64 `json:"id"`

	// Array of nested characteristic values
	Values []UpdateCharacteristicsItemValue `json:"values"`
}

type UpdateCharacteristicsItemValue struct {
	// Characteristic identifier in the dictionary
	DictionaryValueId int64 `json:"dictionary_value_id"`

	// Product characteristic value
	Value string `json:"value"`
}

type UpdateCharacteristicsResponse struct {
	core.CommonResponse

	// Products update task code.
	//
	// To check the update status, pass the received value to the `/v1/product/import/info` method
	TaskId int64 `json:"task_id"`
}

func (c Products) UpdateCharacteristics(ctx context.Context, params *UpdateCharacteristicsParams) (*UpdateCharacteristicsResponse, error) {
	url := "/v1/product/attributes/update"

	resp := &UpdateCharacteristicsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetRelatedSKUsParams struct {
	// List of SKUs
	SKUs []string `json:"sku"`
}

type GetRelatedSKUsResponse struct {
	core.CommonResponse

	// Related SKUs information
	Items []GetRelatedSKUsItem `json:"items"`

	// Errors
	Errors []GetRelatedSKUsError `json:"errors"`
}

type GetRelatedSKUsItem struct {
	// Product availability attribute by SKU
	Availability SKUAvailability `json:"availability"`

	// Date and time of deletion
	DeletedAt time.Time `json:"deleted_at"`

	// Delivery scheme
	DeliverySchema string `json:"delivery_schema"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

type GetRelatedSKUsError struct {
	// Error code
	Code string `json:"code"`

	// SKU, in which the error occurred
	SKU int `json:"sku"`

	// Error text
	Message string `json:"message"`
}

// Method for getting a single SKU based on the old SKU FBS and SKU FBO identifiers.
// The response will contain all SKUs related to the passed ones.
//
// The method can handle any SKU, even hidden or deleted.
//
// In one request, you can pass up to 200 SKUs.
func (c Products) GetRelatedSKUs(ctx context.Context, params *GetRelatedSKUsParams) (*GetRelatedSKUsResponse, error) {
	url := "/v1/product/related-sku/get"

	resp := &GetRelatedSKUsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
