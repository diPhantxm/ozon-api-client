package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetFBOReturns(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetFBOReturnsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetFBOReturnsParams{
				Filter: GetFBOReturnsFilter{
					PostingNumber: "some number",
				},
				LastId: 123,
				Limit:  100,
			},
			`{
				"last_id": 0,
				"returns": [
				  {
					"accepted_from_customer_moment": "2019-08-24T14:15:22Z",
					"company_id": 0,
					"current_place_name": "my-place",
					"dst_place_name": "that-place",
					"id": 0,
					"is_opened": true,
					"posting_number": "some number",
					"return_reason_name": "ripped",
					"returned_to_ozon_moment": "2019-08-24T14:15:22Z",
					"sku": 0,
					"status_name": "delivering"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetFBOReturnsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Returns().GetFBOReturns(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetFBSReturns(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetFBSReturnsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetFBSReturnsParams{
				Filter: GetFBSReturnsFilter{
					PostingNumber: []string{"07402477-0022-2"},
					Status:        "returned_to_seller",
				},
				Limit:  1000,
				Offset: 0,
			},
			`{
				"result": {
				  "returns": [
					{
					  "id": 19166541735000,
					  "clearing_id": 19166541725000,
					  "posting_number": "07402477-0022-2",
					  "product_id": 172423678,
					  "sku": 172423678,
					  "status": "returned_to_seller",
					  "returns_keeping_cost": 0,
					  "return_reason_name": "5.12 Заказ более не актуален: долгие сроки доставки",
					  "return_date": "2020-08-12T17:27:50+00:00",
					  "quantity": 1,
					  "product_name": "Кофе ароматизированный \"Лесной орех\" 250 гр",
					  "price": 294,
					  "waiting_for_seller_date_time": "2020-08-16T02:50:35+00:00",
					  "returned_to_seller_date_time": "2020-08-21T10:07:13+00:00",
					  "last_free_waiting_day": "2020-08-19T23:59:59+00:00",
					  "is_opened": false,
					  "place_id": 0,
					  "commission_percent": 0,
					  "commission": 0,
					  "price_without_commission": 0,
					  "is_moving": false,
					  "moving_to_place_name": "МОСКВА_ХАБ",
					  "waiting_for_seller_days": 2,
					  "picking_amount": null,
					  "accepted_from_customer_moment": null,
					  "picking_tag": null
					}
				  ],
				  "count": 1
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetFBSReturnsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Returns().GetFBSReturns(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
