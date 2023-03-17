package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestReportOnSoldProducts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode   int
		headers      map[string]string
		params       *ReportOnSoldProductsParams
		response     string
		errorMessage string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ReportOnSoldProductsParams{
				Date: "2022-09",
			},
			`{
				"result": [
					{
						"header": [
							{
								"doc_date": "2022-09-22",
								"num": "string",
								"start_date": "2022-09-02",
								"stop_date": "2022-09-22",
								"contract_date": "2022-09-02",
								"contract_num": "string",
								"payer_name": "string",
								"payer_inn": "string",
								"payer_kpp": "string",
								"rcv_name": "string",
								"rcv_inn": "string",
								"rcv_kpp": "string",
								"doc_amount": 1,
								"vat_amount": 1,
								"currency_code": "string"
							}
						],
						"rows": [
							{
								"row_number": 0,
								"product_id": 0,
								"product_name": "string",
								"offer_id": "string",
								"barcode": "string",
								"price": 0,
								"commission_percent": 0,
								"price_sale": 0,
								"sale_qty": 0,
								"sale_amount": 0,
								"sale_discount": 0,
								"sale_commission": 0,
								"sale_price_seller": 0,
								"return_sale": 0,
								"return_qty": 0,
								"return_amount": 0,
								"return_discount": 0,
								"return_commission": 0,
								"return_price_seller": 0
							}
						]
					}
				]
			}`,
			"",
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ReportOnSoldProductsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
			"Client-Id and Api-Key headers are required",
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Finance().ReportOnSoldProducts(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode != http.StatusOK {
			if resp.Message != test.errorMessage {
				t.Errorf("got wrong error message: got: %s, expected: %s", resp.Message, test.errorMessage)
			}
		}
	}
}

func TestGetTotalTransactionsSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode   int
		headers      map[string]string
		params       *GetTotalTransactionsSumParams
		response     string
		errorMessage string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetTotalTransactionsSumParams{
				Date: GetTotalTransactionsSumDate{
					From: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-01T00:00:00.000Z"),
					To:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-02T00:00:00.000Z"),
				},
				TransactionType: "ALL",
			},
			`{
				"result": {
				  "accruals_for_sale": 96647.58,
				  "sale_commission": -11456.65,
				  "processing_and_delivery": -24405.68,
				  "refunds_and_cancellations": -330,
				  "services_amount": -1307.57,
				  "compensation_amount": 0,
				  "money_transfer": 0,
				  "others_amount": 113.05
				}
			}`,
			"",
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetTotalTransactionsSumParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
			"Client-Id and Api-Key headers are required",
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Finance().GetTotalTransactionsSum(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode != http.StatusOK {
			if resp.Message != test.errorMessage {
				t.Errorf("got wrong error message: got: %s, expected: %s", resp.Message, test.errorMessage)
			}
		}
	}
}
