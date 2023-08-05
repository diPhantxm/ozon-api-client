package ozon

import (
	"context"
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
					"company_id": 123456789,
					"current_place_name": "my-place",
					"dst_place_name": "that-place",
					"id": 123456789,
					"is_opened": true,
					"posting_number": "some number",
					"return_reason_name": "ripped",
					"returned_to_ozon_moment": "2019-08-24T14:15:22Z",
					"sku": 123456789,
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().GetFBOReturns(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Returns) > 0 {
				if resp.Returns[0].Id == 0 {
					t.Errorf("Id cannot be 0")
				}
				if resp.Returns[0].CompanyId == 0 {
					t.Errorf("Company id cannot be 0")
				}
				if resp.Returns[0].SKU == 0 {
					t.Errorf("SKU cannot be 0")
				}
			}
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
				LastId: 0,
			},
			`{
				"last_id": 0,
				"returns": [
				  {
					"accepted_from_customer_moment": "string",
					"clearing_id": 23,
					"commission": 21,
					"commission_percent": 0,
					"exemplar_id": 42,
					"id": 123,
					"is_moving": true,
					"is_opened": true,
					"last_free_waiting_day": "string",
					"place_id": 122,
					"moving_to_place_name": "string",
					"picking_amount": 0,
					"posting_number": "string",
					"picking_tag": "string",
					"price": 0,
					"price_without_commission": 0,
					"product_id": 2222,
					"product_name": "string",
					"quantity": 0,
					"return_barcode": "string",
					"return_clearing_id": 0,
					"return_date": "string",
					"return_reason_name": "string",
					"waiting_for_seller_date_time": "string",
					"returned_to_seller_date_time": "string",
					"waiting_for_seller_days": 0,
					"returns_keeping_cost": 0,
					"sku": 33332,
					"status": "string"
				  }
				]
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().GetFBSReturns(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Returns) > 0 {
				if resp.Returns[0].Id == 0 {
					t.Errorf("Id cannot be 0")
				}
				if resp.Returns[0].ProductId == 0 {
					t.Errorf("Product id cannot be 0")
				}
				if resp.Returns[0].SKU == 0 {
					t.Errorf("SKU cannot be 0")
				}
				if resp.Returns[0].Status == "" {
					t.Errorf("Status cannot be empty")
				}
			}
		}
	}
}
