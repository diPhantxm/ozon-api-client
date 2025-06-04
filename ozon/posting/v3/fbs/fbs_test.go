package fbs_test

import (
	"bytes"
	"context"
	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/andmetoo/ozon-api-client/ozon/posting/v3/fbs"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestList_Success(t *testing.T) {
	c := fbs.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v3/posting/fbs/list", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"dir":"ASC","filter":{"delivery_method_id":["123"],"is_quantum":false,"last_changed_status_date":{"from":"2023-11-03T11:47:39.878Z","to":"2023-12-04T11:47:39.878Z"},"order_id":0,"provider_id":["987"],"since":"2023-11-03T11:47:39.878Z","status":"awaiting_packaging","to":"2023-12-04T11:47:39.878Z","warehouse_id":["121212"]},"limit":0,"offset":0,"with":{"analytics_data":true,"barcodes":true,"financial_data":true,"translit":true}}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
  "result": {
    "postings": [
      {
        "posting_number": "05708065-0029-1",
        "order_id": 680420041,
        "order_number": "05708065-0029",
        "pickup_code_verified_at": "2025-01-17T10:59:26.614Z",
        "status": "awaiting_deliver",
        "substatus": "posting_awaiting_passport_data",
        "delivery_method": {
          "id": 21321684811000,
          "name": "Ozon Логистика самостоятельно, Красногорск",
          "warehouse_id": 21321684811000,
          "warehouse": "Стим Тойс Нахабино",
          "tpl_provider_id": 24,
          "tpl_provider": "Ozon Логистика"
        },
        "tracking_number": "",
        "tpl_integration_type": "ozon",
        "in_process_at": "2022-05-13T07:07:32Z",
        "shipment_date": "2022-05-13T10:00:00Z",
        "delivering_date": null,
        "optional": {
          "products_with_possible_mandatory_mark": [
            0
          ]
        },
        "cancellation": {
          "cancel_reason_id": 0,
          "cancel_reason": "",
          "cancellation_type": "",
          "cancelled_after_ship": false,
          "affect_cancellation_rating": false,
          "cancellation_initiator": ""
        },
        "customer": null,
        "products": [
          {
            "price": "1390.000000",
            "currency_code": "RUB",
            "is_blr_traceable": true,
            "offer_id": "205953",
            "name": "Электронный конструктор PinLab Позитроник",
            "sku": 358924380,
            "quantity": 1
          }
        ],
        "addressee": null,
        "barcodes": null,
        "analytics_data": null,
        "financial_data": null,
        "is_express": false,
        "quantum_id": 0
      }
    ],
    "has_next": true
  }
}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v3/posting/fbs",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.List(context.Background(), &fbs.ListRequest{
		Dir: "ASC",
		Filter: fbs.ListRequestFilter{
			DeliveryMethodID: []string{"123"},
			IsQuantum:        false,
			LastChangedStatusDate: fbs.ListRequestLastChangedStatusDate{
				From: time.Date(2023, 11, 3, 11, 47, 39, 878000000, time.UTC),
				To:   time.Date(2023, 12, 4, 11, 47, 39, 878000000, time.UTC),
			},
			OrderID:     0,
			ProviderID:  []string{"987"},
			Since:       time.Date(2023, 11, 3, 11, 47, 39, 878000000, time.UTC),
			Status:      "awaiting_packaging",
			To:          time.Date(2023, 12, 4, 11, 47, 39, 878000000, time.UTC),
			WarehouseID: []string{"121212"},
		},
		Limit:  0,
		Offset: 0,
		With: fbs.ListRequestWith{
			AnalyticsData: true,
			Barcodes:      true,
			FinancialData: true,
			Translit:      true,
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(
		t,
		&fbs.ListResponse{
			Result: fbs.ListResponseResult{
				Postings: []fbs.ListResponseResultPosting{
					{
						PostingNumber:        "05708065-0029-1",
						OrderID:              680420041,
						OrderNumber:          "05708065-0029",
						PickupCodeVerifiedAt: time.Date(2025, 1, 17, 10, 59, 26, 614000000, time.UTC),
						Status:               "awaiting_deliver",
						Substatus:            "posting_awaiting_passport_data",
						DeliveryMethod: fbs.ListResponseResultPostingDeliveryMethod{
							ID:            21321684811000,
							Name:          "Ozon Логистика самостоятельно, Красногорск",
							WarehouseID:   21321684811000,
							Warehouse:     "Стим Тойс Нахабино",
							TplProviderID: 24,
							TplProvider:   "Ozon Логистика",
						},
						TrackingNumber:     "",
						TplIntegrationType: "ozon",
						InProcessAt:        time.Date(2022, 5, 13, 7, 7, 32, 0, time.UTC),
						ShipmentDate:       time.Date(2022, 5, 13, 10, 0, 0, 0, time.UTC),
						DeliveringDate:     nil,
						Optional: fbs.ListResponseResultPostingOptional{
							ProductsWithPossibleMandatoryMark: []int{0},
						},
						Cancellation: fbs.ListResponseResultPostingCancellation{
							CancelReasonID:           0,
							CancelReason:             "",
							CancellationType:         "",
							CancelledAfterShip:       false,
							AffectCancellationRating: false,
							CancellationInitiator:    "",
						},
						Customer: nil,
						Products: []fbs.ListResponseResultPostingProduct{
							{
								Price:          "1390.000000",
								CurrencyCode:   "RUB",
								IsBlrTraceable: true,
								OfferID:        "205953",
								Name:           "Электронный конструктор PinLab Позитроник",
								SKU:            358924380,
								Quantity:       1,
							},
						},
						Addressee:     nil,
						Barcodes:      nil,
						AnalyticsData: nil,
						FinancialData: nil,
						IsExpress:     false,
						QuantumID:     0,
					},
				},
				HasNext: true,
			},
		},
		resp,
	)
}
