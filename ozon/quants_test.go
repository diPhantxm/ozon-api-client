package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestListQuants(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListQuantsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListQuantsParams{
				Cursor: "string",
				Filter: ListQuantsFilter{
					InvQuantIds:        []string{"string"},
					DestinationPlaceId: 123,
					OfferId:            "string",
					SKUName:            "string",
					Statuses:           []string{"unknown"},
					WarehouseId:        456,
				},
				Limit:   10,
				Sort:    "string",
				SortDir: "string",
			},
			`{
				"result": {
				  "cursor": "string",
				  "has_next": true,
				  "quants": [
					{
					  "available_actions": [
						"string"
					  ],
					  "awaiting_stock_due_date": "2019-08-24T14:15:22Z",
					  "cancel_reason": {
						"cancel_reason_id": 0,
						"cancel_reason_name": "string",
						"responsible": "string"
					  },
					  "company_id": 0,
					  "created_at": "2019-08-24T14:15:22Z",
					  "current_postings_count": 0,
					  "cutoff": "2019-08-24T14:15:22Z",
					  "delivery_method_name": "string",
					  "destination_place_id": 0,
					  "destination_place_name": "string",
					  "filling_percent": 0,
					  "first_posting_cancellation_date": "2019-08-24T14:15:22Z",
					  "id": 0,
					  "inv_quant_id": 0,
					  "last_status_change_at": "2019-08-24T14:15:22Z",
					  "offer_id": "string",
					  "products_price": 0,
					  "quantum_start_date": "2019-08-24T14:15:22Z",
					  "sku": 0,
					  "sku_name": "string",
					  "status": "unknown",
					  "target_postings_count": 0,
					  "tpl_provider_name": "string",
					  "type": "string",
					  "warehouse_id": 0,
					  "warehouse_name": "string"
					}
				  ]
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListQuantsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Quants().List(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ListQuantsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetQuant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetQuantParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetQuantParams{
				QuantId: 456,
			},
			`{
				"result": [
				  {
					"available_actions": [
					  "string"
					],
					"awaiting_stock_due_date": "2019-08-24T14:15:22Z",
					"cancel_reason": {
					  "cancel_reason_id": 0,
					  "cancel_reason_name": "string",
					  "responsible": "string"
					},
					"current_postings_count": 0,
					"cutoff": "2019-08-24T14:15:22Z",
					"delivery_method_name": "string",
					"destination_place_id": 0,
					"destination_place_name": "string",
					"filling_percent": 0,
					"first_posting_cancellation_date": "2019-08-24T14:15:22Z",
					"id": 0,
					"inv_quant_id": 0,
					"offer_id": "string",
					"postings": [
					  {
						"cancel_reason": {
						  "cancel_reason_id": 0,
						  "cancel_reason_name": "string",
						  "responsible": "string"
						},
						"posting_number": "string",
						"products_price": 0,
						"status_alias": "string",
						"status_id": 0
					  }
					],
					"product_picture_url": "string",
					"products_price": 0,
					"quantum_start_date": "2019-08-24T14:15:22Z",
					"sku": 0,
					"sku_name": "string",
					"status": "unknown",
					"target_postings_count": 0,
					"tpl_provider_name": "string",
					"type": "string",
					"warehouse_id": 0,
					"warehouse_name": "string"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetQuantParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Quants().Get(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetQuantResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestShipQuant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ShipQuantParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ShipQuantParams{
				QuantId: 456,
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ShipQuantParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Quants().Ship(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ShipQuantResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestStatusQuant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *StatusQuantParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&StatusQuantParams{
				QuantId: 456,
			},
			`{
				"status": "unknown"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&StatusQuantParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Quants().Status(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &StatusQuantResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
