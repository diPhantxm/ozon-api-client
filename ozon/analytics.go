package ozon

import (
	"context"
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type Analytics struct {
	client *core.Client
}

type GetAnalyticsDataParams struct {
	// Date from which the data will be in the report
	DateFrom *core.TimeFormat `json:"date_from"`

	// Date up to which the data will be in the report
	DateTo *core.TimeFormat `json:"date_to"`

	// Items Enum: "unknownDimension" "sku" "spu" "day" "week" "month" "year" "category1" "category2" "category3" "category4" "brand" "modelID"
	// Data grouping available to all sellers:
	//   - unknownDimension—unknown,
	//   - sku—product identifier,
	//   - spu—product identifier,
	//   - day—day,
	//   - week—week,
	//   - month—month.
	// Data grouping available to sellers with Premium subscription:
	//   - year—year,
	//   - category1—first level category,
	//   - category2—second level category,
	//   - category3—third level category,
	//   - category4—fourth level category,
	//   - brand—brand,
	//   - modelID—model.
	Dimension []GetAnalyticsDataDimension `json:"dimension"`

	// Filters
	Filters []GetAnalyticsDataFilter `json:"filters"`

	// Number of items in the respones:
	//   - maximum is 1000,
	//   - minimum is 1.
	Limit int64 `json:"limit"`

	// Specify up to 14 metrics. If there are more, you will get an error with the InvalidArgument code
	// The list of metrics for which the report will be generated.
	//
	// Metrics available to all sellers:
	//
	// 	 - revenue—ordered amount,
	//   - ordered_units—ordered products.
	// Metrics available to sellers with Premium subscription:
	//   - unknown_metric—unknown metric,
	//   - hits_view_search—impressions in search and category,
	//   - hits_view_pdp—impressions on the product description page,
	//   - hits_view—total impressions,
	//   - hits_tocart_search—added to cart from search or category,
	//   - hits_tocart_pdp—added to cart from the product description page,
	//   - hits_tocart—added to cart total,
	//   - session_view_search—sessions with impressions in search or category,
	//   - session_view_pdp—sessions with impressions on the product description page,
	//   - session_view—sessions total,
	//   - conv_tocart_search—conversion to cart from search or category,
	//   - conv_tocart_pdp—conversion to cart from a product description page,
	//   - conv_tocart—total conversion to cart,
	//   - returns—returned products,
	//   - cancellations—canceled products,
	//   - delivered_units—delivered products,
	//   - position_category—position in search and category.
	Metrics []GetAnalyticsDataFilterMetric `json:"metrics"`

	// Number of elements that will be skipped in the response. For example, if `offset=10`, the response will start with the 11th element found
	Offset int64 `json:"offset"`

	// Report sorting settings
	Sort []GetAnalyticsDataSort `json:"sort"`
}

type GetAnalyticsDataFilter struct {
	// Sorting parameter. You can pass any attribute from the `dimension` and `metric` parameters except the `brand` attribute
	Key string `json:"key"`

	// Comparison operation
	Operation GetAnalyticsDataFilterOperation `json:"operation"`

	// Value for comparison
	Value string `json:"value"`
}

// Report sorting settings
type GetAnalyticsDataSort struct {
	// Metric by which the method result will be sorted
	Key GetAnalyticsDataFilterMetric `json:"key"`

	// Sorting type
	Order Order `json:"order"`
}

type GetAnalyticsDataResponse struct {
	core.CommonResponse

	// Method result
	Result GetAnalyticsDataResult `json:"result"`

	// Report creation time
	Timestamp string `json:"timestamp"`
}

type GetAnalyticsDataResult struct {
	// Data array
	Data []GetAnalyticsDataResultData `json:"data"`

	// Total and average metrics values
	Totals []float64 `json:"totals"`
}

type GetAnalyticsDataResultData struct {
	// Data grouping in the report
	Dimensions []GetAnalyticsDataResultDimension `json:"dimensions"`

	// Metric values list
	Metrics []float64 `json:"metrics"`
}

type GetAnalyticsDataResultDimension struct {
	// Product SKU
	Id string `json:"id"`

	// Name
	Name string `json:"name"`
}

// Specify the period and metrics that are required. The response will contain analytical data grouped by the `dimensions` parameter.
//
// There are restrictions for sellers without Premium subscription:
//
//   - data is available for the last 3 months,
//   - some of the data grouping methods and metrics aren't available.
//
// There are no restrictions for sellers with Premium subscription
func (c Analytics) GetAnalyticsData(ctx context.Context, params *GetAnalyticsDataParams) (*GetAnalyticsDataResponse, error) {
	url := "/v1/analytics/data"

	resp := &GetAnalyticsDataResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetStocksOnWarehousesParams struct {
	// Number of values per page.
	//
	// Default is 100
	Limit int64 `json:"limit" default:"100"`

	// Number of elements that will be skipped in the response. For example, if `offset=10`, the response will start with the 11th element found
	Offset int64 `json:"offset"`

	// Warehouse type filter
	WarehouseType WarehouseType `json:"warehouse_type" default:"ALL"`
}

type GetStocksOnWarehousesResponse struct {
	core.CommonResponse

	// Method result
	Result GetStocksOnWarehousesResult `json:"result"`
}

type GetStocksOnWarehousesResult struct {
	// Information about products and stocks
	Rows []GetStocksOnWarehousesResultRow `json:"rows"`
}

type GetStocksOnWarehousesResultRow struct {
	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Product identifier in the seller's system
	ItemCode string `json:"item_code"`

	// Product name in the Ozon system
	ItemName string `json:"item_name"`

	// Product amount available for sale on Ozon
	FreeToSellAmount int64 `json:"free_to_sell_amount"`

	// Product amount specified for confirmed future supplies
	PromisedAmount int64 `json:"promised_amount"`

	// Product amount reserved for purchase, returns, and transportation between warehouses
	ReservedAmount int64 `json:"reserved_amount"`

	// Name of the warehouse where the products are stored
	WarehouseName string `json:"warehouse_name"`
}

// Method for getting a report on leftover stocks and products movement at Ozon warehouses
func (c Analytics) GetStocksOnWarehouses(ctx context.Context, params *GetStocksOnWarehousesParams) (*GetStocksOnWarehousesResponse, error) {
	url := "/v2/analytics/stock_on_warehouses"

	resp := &GetStocksOnWarehousesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductTurnoverParams struct {
	// Number of values in the response
	Limit int64 `json:"limit"`

	// Number of elements to skip in the response.
	//
	// For example, if offset = 10, the response starts with the 11th element found
	Offset int32 `json:"offset"`

	// Product identifiers in the Ozon system, SKU
	SKU []string `json:"sku"`
}

type GetProductTurnoverResponse struct {
	core.CommonResponse

	// Products
	Items []ProductTurnoverItem `json:"items"`
}

type ProductTurnoverItem struct {
	// Average daily number of product items sold over the last 60 days
	Ads float64 `json:"ads"`

	// Product stock, pcs
	CurrentStock int64 `json:"current_stock"`

	// Number of days the stock will last based on your average daily sales
	IDC float64 `json:"idc"`

	// Product stock level
	IDCGrade string `json:"idc_grade"`
}

// Use the method to get the product turnover rate and the number of days the current stock will last.
//
// If you request a list of products by sku, the limit and offset parameters are optional.
func (c Analytics) GetProductTurnover(ctx context.Context, params *GetProductTurnoverParams) (*GetProductTurnoverResponse, error) {
	url := "/v1/analytics/turnover/stocks"

	resp := &GetProductTurnoverResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetStockManagementParams struct {
	// GetStockManagementFilter
	Filter GetStockManagementFilter `json:"filter"`

	// Number of values in the response
	Limit int32 `json:"limit,omitempty"`

	// Number of elements to skip in the response
	Offset int32 `json:"offset,omitempty"`
}

type GetStockManagementFilter struct {
	// Product identifiers in the Ozon system, SKU
	SKUs []string `json:"skus"`

	// The type of item in stock
	StockTypes string `json:"stock_types"`

	// Warehouse identifiers
	WarehouseIds []string `json:"warehouse_ids"`
}

type GetStockManagementResponse struct {
	core.CommonResponse

	// Products
	Items []StockItem `json:"items"`
}

type StockItem struct {
	// Stock of defective products, pcs
	DefectCount int64 `json:"defect_stock_count"`

	// Stock of near-expired products, pcs
	ExpiringCount int64 `json:"expiring_stock_count"`

	// Product name
	ProductName string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Stock of valid products
	ValidCount int64 `json:"valid_stock_count"`

	// Stock of products that waiting for documents
	WaitingDocsCount int64 `json:"waitingdocs_stock_count"`

	// Warehouse name
	WarehouseName string `json:"warehouse_name"`
}

// Use the method to find out how many product items are left in stock
func (c Analytics) Stock(ctx context.Context, params *GetStockManagementParams) (*GetStockManagementResponse, error) {
	url := "/v1/analytics/manage/stocks"

	resp := &GetStockManagementResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductQueriesParams struct {
	// Date when analytics generation starts
	DateFrom string `json:"date_from"`

	//Date when analytics generation ends
	DateTo string `json:"date_to"`

	// Number of page returned in the request
	Page int32 `json:"page"`

	// Number of items on the pag
	PageSize int32 `json:"page_size"`

	// List of SKUs—product identifiers in the Ozon system.
	// Analytics on requests is returned for them.
	// Maximum value is 1,000 SKUs
	SKUs []string `json:"skus"`

	// Parameter by which products are sorted
	SortBy string `json:"sort_by"`

	// Sorting direction
	SortDir string `json:"sort_dir"`
}

type GetProductQueriesResponse struct {
	core.CommonResponse

	// Period for which the analytics is generated
	AnalyticsPeriod AnalyticsPeriod `json:"analytics_period"`

	// Product list
	Items []GetProductQueriesItem `json:"items"`

	// Number of pages
	PageCount int64 `json:"page_count"`

	// Total number of queries
	Total int64 `json:"total"`
}

type AnalyticsPeriod struct {
	// Date when analytics generation starts
	DateFrom string `json:"date_from"`

	// Date when analytics generation ends
	DateTo string `json:"date_to"`
}

type GetProductQueriesItem struct {
	// Category name
	Category string `json:"category"`

	// Currency
	Currency string `json:"currency"`

	// Sales by queries
	GMV float64 `json:"gmv"`

	// Product name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Average product position. Available only with the Premium or Premium Plus subscription, otherwise the field returns empty
	Position float64 `json:"position"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Number of customers who searched for your product on Ozon
	UniqueSearchUsers int64 `json:"unique_search_users"`

	// Number of customers who have seen your product on Ozon.
	// Available only with the Premium or Premium Plus subscription,
	// otherwise the field returns empty
	UniqueViewUsers int64 `json:"unique_view_users"`

	// Conversion from product views.
	// Available only with the Premium or Premium Plus subscription,
	// otherwise the field returns empty
	ViewConversion float64 `json:"view_conversion"`
}

// Use the method to get data about your product queries.
// Full analytics is available with the Premium and Premium Plus subscription.
// Without subscription, you can see a part of the metrics.
// The method is similar to the Products in Search → Queries for my product tab in your personal account.
//
// You can view analytics by queries for certain dates.
// To do this, specify the interval in the date_from and date_to fields.
// Data for the last month are available in any interval except for
// three days from the current date because these days the calculation is performed.
// Analytics for dates later than a month ago is available only with
// the Premium and Premium Plus subscription, and only by weeks.
// Specify the date_from parameter in the request
func (c Analytics) GetProductQueries(ctx context.Context, params *GetProductQueriesParams) (*GetProductQueriesResponse, error) {
	url := "/v1/analytics/product-queries"

	resp := &GetProductQueriesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
