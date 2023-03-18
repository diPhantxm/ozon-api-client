package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Analytics struct {
	client *core.Client
}

type GetAnalyticsDataParams struct {
	// Date from which the data will be in the report
	DateFrom time.Time `json:"date_from"`

	// Date up to which the data will be in the report
	DateTo time.Time `json:"date_to"`

	// Items Enum: "unknownDimension" "sku" "spu" "day" "week" "month" "year" "category1" "category2" "category3" "category4" "brand" "modelID"
	Dimension []string `json:"dimension"`

	// Filters
	Filters []struct {
		// Sorting parameter. You can pass any attribute from the `dimension` and `metric` parameters except the `brand` attribute
		Key string `json:"key"`

		// Comparison operation
		//
		// Enum: "EQ" "GT" "GTE" "LT" "LTE"
		Operation string `json:"operation"`

		// Value for comparison
		Value string `json:"value"`
	} `json:"filters"`

	// Number of items in the respones:
	//   - maximum is 1000,
	//   - minimum is 1.
	Limit int64 `json:"limit"`

	// Specify up to 14 metrics. If there are more, you will get an error with the InvalidArgument code
	//
	// Items Enum: "unknown_metric" "hits_view_search" "hits_view_pdp" "hits_view" "hits_tocart_search" "hits_tocart_pdp" "hits_tocart" "session_view_search"
	// "session_view_pdp" "session_view" "conv_tocart_search" "conv_tocart_pdp" "conv_tocart" "revenue" "returns" "cancellations" "ordered_units" "delivered_units"
	// "adv_view_pdp" "adv_view_search_category" "adv_view_all" "adv_sum_all" "position_category" "postings" "postings_premium"
	Metrics []string `json:"metrics"`

	// Number of elements that will be skipped in the response. For example, if `offset=10`, the response will start with the 11th element found
	Offset int64 `json:"offset"`

	// Report sorting settings
	Sort []GetAnalyticsDataSort `json:"sort"`
}

// Report sorting settings
type GetAnalyticsDataSort struct {
	// Metric by which the method result will be sorted
	Key string `json:"key"`

	// Sorting type
	//   - ASC — in ascending order,
	//   - DESC — in descending order.
	Order string `json:"order"`
}

type GetAnalyticsDataResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Data array
		Data []struct {
			// Data grouping in the report
			Dimensions []struct {
				// Identifier
				Id string `json:"id"`

				// Name
				Name string `json:"name"`
			} `json:"dimensions"`

			// Metric values list
			Metrics []float64 `json:"metrics"`
		} `json:"data"`

		// Total and average metrics values
		Totals []float64 `json:"totals"`
	} `json:"result"`

	// Report creation time
	Timestamp string `json:"timestamp"`
}

// Specify the period and metrics that are required. The response will contain analytical data grouped by the `dimensions` parameter.
func (c Analytics) GetAnalyticsData(params *GetAnalyticsDataParams) (*GetAnalyticsDataResponse, error) {
	url := "/v1/analytics/data"

	resp := &GetAnalyticsDataResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
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

	// Warehouse type filter:
	//   - EXPRESS_DARK_STORE — Ozon warehouses with Fresh delivery.
	//   - NOT_EXPRESS_DARK_STORE — Ozon warehouses without Fresh delivery.
	//   - ALL — all Ozon warehouses.
	WarehouseType string `json:"warehouse_type"`
}

type GetStocksOnWarehousesResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Information about products and stocks
		Rows []struct {
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
		} `json:"rows"`
	} `json:"result"`
}

// Report on stocks and products movement at Ozon warehouses
func (c Analytics) GetStocksOnWarehouses(params *GetStocksOnWarehousesParams) (*GetStocksOnWarehousesResponse, error) {
	url := "/v2/analytics/stock_on_warehouses"

	resp := &GetStocksOnWarehousesResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
