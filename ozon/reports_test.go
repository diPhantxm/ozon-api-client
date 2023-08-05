package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetReportsListParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetReportsListParams{
				ReportType: "ALL",
				PageSize:   1000,
			},
			`{
				"result": {
				  "reports": [
					{
					  "code": "cd3f2c76-2b9a-4e77-a5a9-7ab52678b3bf",
					  "status": "success",
					  "error": "",
					  "file": "https://minio-production.s3.s.o3.ru:8000/sc-temporary/89/0e/890ef6e360a6396f.csv",
					  "report_type": "seller_products",
					  "params": {
						"visibility": "3"
					  },
					  "created_at": "2019-02-06T12:09:47.258062Z"
					},
					{
					  "code": "c39f5fe4-c00b-4e95-a487-6ad34c1e34a3",
					  "status": "success",
					  "error": "",
					  "file": "https://minio-production.s3.s.o3.ru:8000/reports/a7/48/a7481a083873e164.csv",
					  "report_type": "seller_products",
					  "params": {
						"visibility": "3"
					  },
					  "created_at": "2019-02-15T08:34:32.267178Z"
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
			&GetReportsListParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Reports().GetList(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if int(resp.Result.Total) != len(resp.Result.Reports) {
			t.Errorf("Amount of reports (%d) is not equal to total (%d)", len(resp.Result.Reports), resp.Result.Total)
		}
		if len(resp.Result.Reports) > 0 {
			if resp.Result.Reports[0].Status == "" {
				t.Errorf("Status must be 'success' or 'failed'")
			}
		}
	}
}

func TestGetReportDetails(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetReportDetailsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetReportDetailsParams{},
			`{
				"result": {
				  "code": "257bf213-ca57-405c-8edf-41d2ce22decf",
				  "status": "success",
				  "error": "",
				  "file": "https://storage.yandexcloud.net/ozon.reports/95/c1/95c1ae93320294cb.csv",
				  "report_type": "seller_products",
				  "params": {},
				  "created_at": "2021-11-25T14:54:55.688260Z"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetReportDetailsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Reports().GetReportDetails(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.Status == "" {
				t.Errorf("Status must be 'success' or 'failed'")
			}
		}
	}
}

func TestGetFinancialReport(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetFinancialReportParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetFinancialReportParams{
				Date: GetFinancialReportDatePeriod{
					From: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2022-01-01T00:00:00.000Z"),
					To:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2022-12-31T00:00:00.000Z"),
				},
				WithDetails: true,
				Page:        1,
				PageSize:    1,
			},
			`{
				"result": {
				  "cash_flows": [
					{
					  "commission_amount": 1437,
					  "currency_code": "string",
					  "item_delivery_and_return_amount": 1991,
					  "orders_amount": 1000,
					  "period": {
						"begin": "2023-04-03T09:12:10.239Z",
						"end": "2023-04-03T09:12:10.239Z",
						"id": 11567022278500
					  },
					  "returns_amount": -3000,
					  "services_amount": 8471.28
					}
				  ],
				  "details": {
					"period": {
					  "begin": "2023-04-03T09:12:10.239Z",
					  "end": "2023-04-03T09:12:10.239Z",
					  "id": 11567022278500
					},
					"payments": [
					  {
						"payment": 0,
						"currency_code": "string"
					  }
					],
					"begin_balance_amount": 0,
					"delivery": {
					  "total": 0,
					  "amount": 0,
					  "delivery_services": {
						"total": 0,
						"items": [
						  {
							"name": "string",
							"price": 0
						  }
						]
					  }
					},
					"return": {
					  "total": 0,
					  "amount": 0,
					  "return_services": {
						"total": 0,
						"items": [
						  {
							"name": "string",
							"price": 0
						  }
						]
					  }
					},
					"loan": 0,
					"invoice_transfer": 0,
					"rfbs": {
					  "total": 0,
					  "transfer_delivery": 0,
					  "transfer_delivery_return": 0,
					  "compensation_delivery_return": 0,
					  "partial_compensation": 0,
					  "partial_compensation_return": 0
					},
					"services": {
					  "total": 0,
					  "items": [
						{
						  "name": "string",
						  "price": 0
						}
					  ]
					},
					"others": {
					  "total": 0,
					  "items": [
						{
						  "name": "string",
						  "price": 0
						}
					  ]
					},
					"end_balance_amount": 0
				  }
				},
				"page_count": 15
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetFinancialReportParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Reports().GetFinancial(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result.CashFlows) > 0 {
				if resp.Result.CashFlows[0].CurrencyCode == "" {
					t.Errorf("Currency Code cannot be empty")
				}
			}
		}
	}
}

func TestGetProductsReport(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProductsReportParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductsReportParams{},
			`{
				"result": {
				  "code": "d55f4517-8347-4e24-9d93-d6e736c1c07c"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetProductsReportParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Reports().GetProducts(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.Code == "" {
				t.Errorf("Code cannot be empty")
			}
		}
	}
}

func TestGetStocksReport(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetStocksReportParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetStocksReportParams{},
			`{
				"result": {
				  "code": "d55f4517-8347-4e24-9d93-d6e736c1c07c"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetStocksReportParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Reports().GetStocks(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.Code == "" {
				t.Errorf("Code cannot be empty")
			}
		}
	}
}

func TestGetProductsMovementReport(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProductsMovementReportParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductsMovementReportParams{
				DateFrom: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2020-08-01T14:15:22Z"),
				DateTo:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2020-08-01T14:15:22Z"),
			},
			`{
				"result": {
				  "code": "h56f4917-1346-4e64-9d90-с6e736c1e07c"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetProductsMovementReportParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Reports().GetProductsMovement(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.Code == "" {
				t.Errorf("Code cannot be empty")
			}
		}
	}
}

func TestGetReturnsReport(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetReturnsReportParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetReturnsReportParams{
				Filter: GetReturnsReportsFilter{
					DeliverySchema: "fbs",
				},
			},
			`{
				"result": {
				  "code": "d55f4517-8347-4e24-9d93-d6e736c1c07c"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetReturnsReportParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Reports().GetReturns(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.Code == "" {
				t.Errorf("Code cannot be empty")
			}
		}
	}
}

func TestGetShipmentReport(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetShipmentReportParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetShipmentReportParams{
				Filter: GetShipmentReportFilter{
					DeliverySchema:  []string{"fbs", "fbo", "crossborder"},
					ProcessedAtFrom: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-09-02T17:10:54.861Z"),
					ProcessedAtTo:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-02T17:10:54.861Z"),
				},
			},
			`{
				"result": {
				  "code": "d55f4517-8347-4e24-9d93-d6e736c1c07c"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetShipmentReportParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Reports().GetShipment(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.Code == "" {
				t.Errorf("Code cannot be empty")
			}
		}
	}
}

func TestIssueOnDiscountedProducts(t *testing.T) {
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
				"code": "d55f4517-8347-4e24-9d93-d6e736c1c07c"
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
		resp, err := c.Reports().IssueOnDiscountedProducts(ctx)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Code == "" {
				t.Errorf("Code cannot be empty")
			}
		}
	}
}

func TestReportOnDiscountedProducts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ReportOnDiscountedProductsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ReportOnDiscountedProductsParams{
				Code: "d55f4517-8347-4e24-9d93-d6e736c1c07c",
			},
			`{
				"report": {
				  "created_at": "2022-10-04T10:07:08.146Z",
				  "error": "string",
				  "file": "string",
				  "status": "string"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ReportOnDiscountedProductsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Reports().ReportOnDiscountedProducts(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestListReportsOnDiscountedProducts(t *testing.T) {
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
				"reports": [
				  {
					"created_at": "2022-10-04T10:07:08.146Z",
					"error": "string",
					"file": "string",
					"status": "string"
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
		resp, err := c.Reports().ListReportsOnDiscountedProducts(ctx)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
