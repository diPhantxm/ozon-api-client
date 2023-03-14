package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetAvailablePromotions(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			`{
				"result": [
				  {
					"id": 71342,
					"title": "test voucher #2",
					"date_start": "2021-11-22T09:46:38Z",
					"date_end": "2021-11-30T20:59:59Z",
					"potential_products_count": 0,
					"is_participating": true,
					"participating_products_count": 5,
					"description": "",
					"action_type": "DISCOUNT",
					"banned_products_count": 0,
					"with_targeting": false,
					"discount_type": "UNKNOWN",
					"discount_value": 0,
					"order_amount": 0,
					"freeze_date": "",
					"is_voucher_action": true
				  }
				]
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.GetAvailablePromotions()
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
