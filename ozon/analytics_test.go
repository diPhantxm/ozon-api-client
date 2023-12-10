package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetAnalyticsData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetAnalyticsDataParams
		response   string
	}{
		// Test ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetAnalyticsDataParams{
				DateFrom:  core.TimeFromString(t, "2006-01-02", "2020-09-01"),
				DateTo:    core.TimeFromString(t, "2006-01-02", "2021-10-15"),
				Dimension: []GetAnalyticsDataDimension{SKUDimension, DayDimension},
				Metrics:   []GetAnalyticsDataFilterMetric{HistViewPDP},
				Sort: []GetAnalyticsDataSort{
					{
						Key:   HistViewPDP,
						Order: Descending,
					},
				},
				Limit:  1000,
				Offset: 0,
			},
			`{
				"result": {
				  "data": [],
				  "totals": [
					0
				  ]
				},
				"timestamp": "2021-11-25 15:19:21"
			}`,
		},
		{
			// Test No Client-Id or Api-Key
			http.StatusUnauthorized,
			map[string]string{},
			&GetAnalyticsDataParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Analytics().GetAnalyticsData(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &GetAnalyticsDataResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetStocksOnWarehouses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetStocksOnWarehousesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetStocksOnWarehousesParams{
				Limit:         1000,
				Offset:        0,
				WarehouseType: "ALL",
			},
			`{
				"result": {
				  "rows": [
					{
					  "free_to_sell_amount": 15,
					  "item_code": "my-code",
					  "item_name": "my-name",
					  "promised_amount": 12,
					  "reserved_amount": 11,
					  "sku": 12345,
					  "warehouse_name": "my-warehouse"
					}
				  ]
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetStocksOnWarehousesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Analytics().GetStocksOnWarehouses(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &GetStocksOnWarehousesResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result.Rows) > int(test.params.Limit) {
				t.Errorf("Length of rows is bigger than limit")
			}
		}
	}
}
