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
		// Test Ok
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
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Promotions().GetAvailablePromotions()
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestAddToPromotion(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *AddProductToPromotionParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&AddProductToPromotionParams{
				ActionId: 60564,
				Products: []AddProductToPromotionProduct{
					{
						ProductId:   1389,
						ActionPrice: 356,
						Stock:       10,
					},
				},
			},
			`{
			"result": {
				"product_ids": [
						1389
					],
					"rejected": []
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&AddProductToPromotionParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Promotions().AddToPromotion(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
