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
				Filter: &GetFBOReturnsFilter{
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
			continue
		}

		compareJsonResponse(t, test.response, &GetFBOReturnsResponse{})

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
				Filter: &GetFBSReturnsFilter{
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
			continue
		}

		compareJsonResponse(t, test.response, &GetFBSReturnsResponse{})

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

func TestGetRFBSReturns(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetRFBSReturnsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetRFBSReturnsParams{
				LastId: 999,
				Limit:  555,
				Filter: &GetRFBSReturnsFilter{
					OfferId:       "123",
					PostingNumber: "111",
					GroupState:    []RFBSReturnsGroupState{RFBSReturnsGroupStateAll},
					CreatedAt: GetRFBSReturnsFilterCreatedAt{
						From: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
						To:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
					},
				},
			},
			`{
				"returns": {
				  "client_name": "string",
				  "created_at": "2019-08-24T14:15:22Z",
				  "order_number": "string",
				  "posting_number": "111",
				  "product": {
					"name": "string",
					"offer_id": "123",
					"currency_code": "string",
					"price": "string",
					"sku": 123
				  },
				  "return_id": 0,
				  "return_number": "string",
				  "state": {
					"group_state": "All",
					"money_return_state_name": "string",
					"state": "string",
					"state_name": "string"
				  }
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetRFBSReturnsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().GetRFBSReturns(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetRFBSReturnsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Returns.Product.OfferId != test.params.Filter.OfferId {
				t.Errorf("expected offer ID %s, but got: %s", test.params.Filter.OfferId, resp.Returns.Product.OfferId)
			}
			if resp.Returns.PostingNumber != test.params.Filter.PostingNumber {
				t.Errorf("expected posting number %s, but got: %s", test.params.Filter.PostingNumber, resp.Returns.PostingNumber)
			}
			if resp.Returns.State.GroupState != test.params.Filter.GroupState[0] {
				t.Errorf("expected group state %s, but got: %s", test.params.Filter.GroupState[0], resp.Returns.State.GroupState)
			}
		}
	}
}

func TestGetRFBSReturn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetRFBSReturnParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetRFBSReturnParams{
				ReturnId: 123,
			},
			`{
				"returns": {
				  "available_actions": [
					{
					  "id": 0,
					  "name": "string"
					}
				  ],
				  "client_name": "string",
				  "client_photo": [
					"string"
				  ],
				  "client_return_method_type": {
					"id": 0,
					"name": "string"
				  },
				  "comment": "string",
				  "created_at": "2019-08-24T14:15:22Z",
				  "order_number": "string",
				  "posting_number": "string",
				  "product": {
					"name": "string",
					"offer_id": "string",
					"currency_code": "string",
					"price": "string",
					"sku": 0
				  },
				  "rejection_comment": "string",
				  "rejection_reason": [
					{
					  "hint": "string",
					  "id": 0,
					  "is_comment_required": true,
					  "name": "string"
					}
				  ],
				  "return_method_description": "string",
				  "return_number": "string",
				  "return_reason": {
					"id": 0,
					"is_defect": true,
					"name": "string"
				  },
				  "ru_post_tracking_number": "string",
				  "state": {
					"state": "string",
					"state_name": "string"
				  },
				  "warehouse_id": 0
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetRFBSReturnParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().GetRFBSReturn(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetRFBSReturnResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestRejectRFBSReturn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *RejectRFBSReturnParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&RejectRFBSReturnParams{
				ReturnId:          123,
				Comment:           "No comment",
				RejectionReasonId: 111,
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&RejectRFBSReturnParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().RejectRFBSReturn(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &RejectRFBSReturnResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestCompensateRFBSreturn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CompensateRFBSReturnParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CompensateRFBSReturnParams{
				ReturnId:           123,
				CompensationAmount: "11",
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CompensateRFBSReturnParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().CompensateRFBSReturn(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CompensateRFBSReturnResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestApproveRFBSReturn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ApproveRFBSReturnParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ApproveRFBSReturnParams{
				ReturnId:                123,
				ReturnMethodDescription: "some description",
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ApproveRFBSReturnParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().ApproveRFBSReturn(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ApproveRFBSReturnResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestReceiveRFBSReturn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ReceiveRFBSReturnParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ReceiveRFBSReturnParams{
				ReturnId: 123,
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ReceiveRFBSReturnParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().ReceiveRFBSReturn(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ReceiveRFBSReturnResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestRefundRFBS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *RefundRFBSParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&RefundRFBSParams{
				ReturnId:         123,
				ReturnForBackWay: 111,
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&RefundRFBSParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().RefundRFBS(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &RefundRFBSResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestIsGiveoutEnabled(t *testing.T) {
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
				"enabled": true
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
		resp, err := c.Returns().IsGiveoutEnabled(ctx)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &IsGiveoutEnabledResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetGiveoutPDF(t *testing.T) {
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
				"content_type": "application/pdf",
				"file_name": "string",
				"file_content": "string"
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
		resp, err := c.Returns().GetGiveoutPDF(ctx)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetGiveoutResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetGiveoutPNG(t *testing.T) {
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
				"content_type": "image/png",
				"file_name": "string",
				"file_content": "string"
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
		resp, err := c.Returns().GetGiveoutPNG(ctx)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetGiveoutResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetGiveoutBarcode(t *testing.T) {
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
				"barcode": "string"
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
		resp, err := c.Returns().GetGiveoutBarcode(ctx)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetGiveoutBarcodeResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestResetGiveoutBarcode(t *testing.T) {
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
				"content_type": "image/png",
				"file_name": "string",
				"file_content": "string"
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
		resp, err := c.Returns().ResetGiveoutBarcode(ctx)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetGiveoutResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetGiveoutList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetGiveoutListParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetGiveoutListParams{
				LastId: 0,
				Limit:  0,
			},
			`{
				"giveouts": [
				  {
					"approved_articles_count": 0,
					"created_at": "2019-08-24T14:15:22Z",
					"giveout_id": 0,
					"giveout_status": "string",
					"total_articles_count": 0,
					"warehouse_address": "string",
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
			&GetGiveoutListParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().GetGiveoutList(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetGiveoutListResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetGiveoutInfo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetGiveoutInfoParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetGiveoutInfoParams{
				GiveoutId: 11,
			},
			`{
				"articles": [
				  {
					"approved": true,
					"delivery_schema": "string",
					"name": "string",
					"seller_id": 0
				  }
				],
				"giveout_id": 11,
				"giveout_status": "string",
				"warehouse_address": "string",
				"warehouse_name": "string"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetGiveoutInfoParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().GetGiveoutInfo(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetGiveoutInfoResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
		if resp.GiveoutId != test.params.GiveoutId {
			t.Errorf("expected giveout id to be equal: got: %d, expected: %d", resp.GiveoutId, test.params.GiveoutId)
		}
	}
}

func TestFBSQuantity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetFBSQuantityReturnsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetFBSQuantityReturnsParams{
				Filter: GetFBSQuantityReturnsFilter{
					PlaceId: 1,
				},
				Pagination: GetFBSQuantityReturnsPagination{
					LastId: 2,
					Limit:  3,
				},
			},
			`{
				"company_id": 0,
				"drop_off_points": [
				  {
					"address": "string",
					"id": 0,
					"name": "string",
					"pass_info": {
					  "count": 0,
					  "is_required": true
					},
					"place_id": 0,
					"returns_count": 0,
					"warehouses_ids": [
					  "string"
					]
				  }
				],
				"has_next": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetFBSQuantityReturnsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().FBSQuantity(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetFBSQuantityReturnsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestListReturns(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListReturnsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListReturnsParams{
				Filter: &ListReturnsFilter{
					LogisticReturnDate: &GetFBSReturnsFilterTimeRange{
						TimeFrom: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
						TimeTo:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
					},
					StorageTarifficationDate: &GetFBSReturnsFilterTimeRange{
						TimeFrom: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
						TimeTo:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
					},
					VisualStatusChangeMoment: &GetFBSReturnsFilterTimeRange{
						TimeFrom: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
						TimeTo:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
					},
					WarehouseId:  911,
					ReturnSchema: "FBO",
					ProductName:  "string",
				},
				Limit:  500,
				LastId: 0,
			},
			`{
				"returns": [
				  {
					"exemplars": [
					  {
						"id": 1019562967545956
					  }
					],
					"id": 1000015552,
					"company_id": 3058,
					"return_reason_name": "Customer refused on receipt: not satisfied with the quality of the product",
					"type": "FullReturn",
					"schema": "Fbs",
					"order_id": 24540784250,
					"order_number": "58544282-0057",
					"place": {
					  "id": 23869688194000,
					  "name": "СЦ_Львовский_Возвраты",
					  "address": "Россия, обл. Московская, г. Подольск, промышленная зона Львовский, ул. Московская, д. 69, стр. 5"
					},
					"target_place": {
					  "id": 23869688194000,
					  "name": "СЦ_Львовский_Возвраты",
					  "address": "Россия, обл. Московская, г. Подольск, промышленная зона Львовский, ул. Московская, д. 69, стр. 5"
					},
					"storage": {
					  "sum": {
						"currency_code": "RUB",
						"price": 1231
					  },
					  "tariffication_first_date": "2024-07-30T06:15:48.998146Z",
					  "tariffication_start_date": "2024-07-29T06:15:48.998146Z",
					  "arrived_moment": "2024-07-29T06:15:48.998146Z",
					  "days": 0,
					  "utilization_sum": {
						"currency_code": "RUB",
						"price": 1231
					  },
					  "utilization_forecast_date": "2024-07-29T06:15:48.998146Z"
					},
					"product": {
					  "sku": 1100526203,
					  "offer_id": "81451",
					  "name": "Кукла Дотти Плачущий младенец Cry Babies Dressy Dotty",
					  "price": {
						"currency_code": "RUB",
						"price": 3318
					  },
					  "price_without_commission": {
						"currency_code": "RUB",
						"price": 3318
					  },
					  "commission_percent": 1.2,
					  "commission": {
						"currency_code": "RUB",
						"price": 2312
					  }
					},
					"logistic": {
					  "technical_return_moment": "2024-07-29T06:15:48.998146Z",
					  "final_moment": "2024-07-29T06:15:48.998146Z",
					  "cancelled_with_compensation_moment": "2024-07-29T06:15:48.998146Z",
					  "return_date": "2024-07-29T06:15:48.998146Z",
					  "barcode": "ii5275210303"
					},
					"visual": {
					  "status": {
						"id": 3,
						"display_name": "At the pick-up point",
						"sys_name": "ArrivedAtReturnPlace"
					  },
					  "change_moment": "2024-07-29T06:15:48.998146Z"
					},
					"additional_info": {
					  "is_opened": true,
					  "is_super_econom": false
					},
					"source_id": 90426223,
					"posting_number": "58544282-0057-1",
					"clearing_id": 21190893156000,
					"return_clearing_id": null
				  }
				],
				"has_next": false
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListReturnsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Returns().List(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ListReturnsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
