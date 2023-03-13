package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetAnalyticsData(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetAnalyticsDataParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetAnalyticsDataParams{
				DateFrom:  core.TimeFromString(t, "2006-01-02", "2020-09-01"),
				DateTo:    core.TimeFromString(t, "2006-01-02", "2021-10-15"),
				Dimension: []string{"sku", "day"},
				Metrics:   []string{"hits_view_search"},
				Sort: []GetAnalyticsDataSort{
					{
						Key:   "hits_view_search",
						Order: "DESC",
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
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.GetAnalyticsData(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetStocksOnWarehouses(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetStocksOnWarehousesParams
		response   string
	}{
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
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.GetStocksOnWarehouses(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
