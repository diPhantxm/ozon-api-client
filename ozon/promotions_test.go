package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetAvailablePromotions(t *testing.T) {
	t.Parallel()

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().GetAvailablePromotions(ctx)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) > 0 {
				if resp.Result[0].Id == 0 {
					t.Errorf("Id cannot be 0")
				}
				if resp.Result[0].ActionType == "" {
					t.Errorf("Action type cannot be empty")
				}
			}
		}
	}
}

func TestAddToPromotion(t *testing.T) {
	t.Parallel()

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().AddToPromotion(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result.ProductIds) != len(test.params.Products) {
				t.Errorf("Length of products in response and request must be equal")
			}
		}
	}
}

func TestProductsAvailableForPromotion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ProductsAvailableForPromotionParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ProductsAvailableForPromotionParams{
				ActionId: 63337,
				Limit:    10,
				Offset:   0,
			},
			`{
				"result": {
				  "products": [
					{
					  "id": 226,
					  "price": 250,
					  "action_price": 0,
					  "max_action_price": 175,
					  "add_mode": "NOT_SET",
					  "stock": 0,
					  "min_stock": 0
					},
					{
					  "id": 1366,
					  "price": 2300,
					  "action_price": 630,
					  "max_action_price": 770,
					  "add_mode": "MANUAL",
					  "stock": 0,
					  "min_stock": 0
					}
				  ],
				  "total": 2
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ProductsAvailableForPromotionParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().ProductsAvailableForPromotion(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestProductsInPromotion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ProductsInPromotionParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ProductsInPromotionParams{
				ActionId: 66011,
				Limit:    10,
				Offset:   0,
			},
			`{
				"result": {
				  "products": [
					{
					  "id": 1383,
					  "price": 5503,
					  "action_price": 621,
					  "max_action_price": 3712.1,
					  "add_mode": "MANUAL",
					  "stock": 0,
					  "min_stock": 0
					}
				  ],
				  "total": 1
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ProductsInPromotionParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().ProductsInPromotion(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestRemoveProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *RemoveProductFromPromotionParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&RemoveProductFromPromotionParams{
				ActionId:   66011,
				ProductIds: []float64{14975},
			},
			`{
				"result": {
				  "product_ids": [
					14975
				  ],
				  "rejected": []
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&RemoveProductFromPromotionParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().RemoveProduct(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result.ProductIds) > 0 {
				if resp.Result.ProductIds[0] != test.params.ProductIds[0] {
					t.Errorf("Product ids in request and response are not equal")
				}
			}
		}
	}
}

func TestListHotSalePromotions(t *testing.T) {
	t.Parallel()

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
					"date_end": "string",
					"date_start": "string",
					"description": "string",
					"freeze_date": "string",
					"hotsale_id": 0,
					"is_participating": true,
					"title": "string"
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().ListHotSalePromotions(ctx)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestProductsAvailableForHotSalePromotion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ProductsAvailableForHotSalePromotionParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ProductsAvailableForHotSalePromotionParams{
				HotSaleId: 0,
				Limit:     0,
				Offset:    0,
			},
			`{
				"result": {
				  "products": [
					{
					  "action_price": 0,
					  "date_day_promo": "string",
					  "id": 0,
					  "is_active": true,
					  "max_action_price": 0,
					  "min_stock": 0,
					  "stock": 0
					}
				  ],
				  "total": 0
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ProductsAvailableForHotSalePromotionParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().ProductsAvailableForHotSalePromotion(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestAddProductsToHotSale(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *AddProductsToHotSaleParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&AddProductsToHotSaleParams{
				HotSaleId: 1234,
				Products: []AddProductsToHotSaleProduct{
					{
						ActionPrice: 12,
						ProductId:   111,
						Stock:       45,
					},
				},
			},
			`{
				"result": {
				  "rejected": [
					{
					  "product_id": 0,
					  "reason": "string"
					}
				  ]
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&AddProductsToHotSaleParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().AddProductsToHotSale(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestRemoveProductsToHotSale(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *RemoveProductsToHotSaleParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&RemoveProductsToHotSaleParams{
				HotSaleId:  12345,
				ProductIds: []float64{111},
			},
			`{
				"result": {
				  "rejected": [
					{
					  "product_id": 0,
					  "reason": "string"
					}
				  ]
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&RemoveProductsToHotSaleParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().RemoveProductsToHotSale(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestListDiscountRequests(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListDiscountRequestsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListDiscountRequestsParams{
				Status: "UNKNOWN",
				Page:   0,
				Limit:  100,
			},
			`{
				"result": [
				  {
					"id": 0,
					"created_at": "2019-08-24T14:15:22Z",
					"end_at": "2019-08-24T14:15:22Z",
					"edited_till": "2019-08-24T14:15:22Z",
					"status": "string",
					"customer_name": "string",
					"sku": 0,
					"user_comment": "string",
					"seller_comment": "string",
					"requested_price": 0,
					"approved_price": 0,
					"original_price": 0,
					"discount": 0,
					"discount_percent": 0,
					"base_price": 0,
					"min_auto_price": 0,
					"prev_task_id": 0,
					"is_damaged": true,
					"moderated_at": "2019-08-24T14:15:22Z",
					"approved_discount": 0,
					"approved_discount_percent": 0,
					"is_purchased": true,
					"is_auto_moderated": true,
					"offer_id": "string",
					"email": "string",
					"last_name": "string",
					"first_name": "string",
					"patronymic": "string",
					"approved_quantity_min": 0,
					"approved_quantity_max": 0,
					"requested_quantity_min": 0,
					"requested_quantity_max": 0,
					"requested_price_with_fee": 0,
					"approved_price_with_fee": 0,
					"approved_price_fee_percent": 0
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListDiscountRequestsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().ListDiscountRequests(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestApproveDiscountRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *DiscountRequestParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&DiscountRequestParams{
				Tasks: []DiscountRequestTask{
					{
						Id:                  123,
						ApprovedPrice:       11,
						SellerComment:       "string",
						ApprovedQuantityMin: 1,
						ApprovedQuantityMax: 2,
					},
				},
			},
			`{
				"result": {
				  "fail_details": [
					{
					  "task_id": 1234,
					  "error_for_user": "string"
					}
				  ],
				  "success_count": 1,
				  "fail_count": 1
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&DiscountRequestParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().ApproveDiscountRequest(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestDeclineDiscountRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *DiscountRequestParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&DiscountRequestParams{
				Tasks: []DiscountRequestTask{
					{
						Id:                  123,
						ApprovedPrice:       11,
						SellerComment:       "string",
						ApprovedQuantityMin: 1,
						ApprovedQuantityMax: 2,
					},
				},
			},
			`{
				"result": {
				  "fail_details": [
					{
					  "task_id": 1234,
					  "error_for_user": "string"
					}
				  ],
				  "success_count": 1,
				  "fail_count": 1
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&DiscountRequestParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Promotions().DeclineDiscountRequest(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
