package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetList(t *testing.T) {
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

		resp, err := c.Reports().GetList(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetReportDetails(t *testing.T) {
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

		resp, err := c.Reports().GetReportDetails(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetFinancial(t *testing.T) {
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
				Page:     1,
				PageSize: 1,
			},
			`{
				"result": {
				  "cash_flows": [
					{
					  "period": {
						"id": 11567022278500,
						"begin": "2022-08-01T00:00:00Z",
						"end": "2022-08-15T00:00:00Z"
					  },
					  "orders_amount": 1000,
					  "returns_amount": -3000,
					  "commission_amount": 1437,
					  "services_amount": 8471.28,
					  "item_delivery_and_return_amount": 1991,
					  "currency_code": "RUB"
					}
				  ],
				  "page_count": 15
				}
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

		resp, err := c.Reports().GetFinancial(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
