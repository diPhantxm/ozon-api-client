package info_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/andmetoo/ozon-api-client/ozon/product/v5/info"

	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestPrices_Success(t *testing.T) {
	c := info.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v4/product/info/prices", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"cursor":"string","filter":{"offer_id":["356792"],"product_id":["243686911"],"visibility":"ALL"},"limit":100}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"cursor": "string",
							"items": [
								{
									"acquiring": 0,
									"commissions": {
										"fbo_deliv_to_customer_amount": 14.75,
										"fbo_direct_flow_trans_max_amount": 46.5,
										"fbo_direct_flow_trans_min_amount": 31,
										"fbo_return_flow_amount": 50,
										"fbs_deliv_to_customer_amount": 60,
										"fbs_direct_flow_trans_max_amount": 61.5,
										"fbs_direct_flow_trans_min_amount": 41,
										"fbs_first_mile_max_amount": 25,
										"fbs_first_mile_min_amount": 0,
										"fbs_return_flow_amount": 40,
										"sales_percent_fbo": 15,
										"sales_percent_fbs": 0
									},
									"marketing_actions": {
										"actions": [
											{
												"date_from": "2024-12-13T06:49:37.591Z",
												"date_to": "2024-12-13T06:49:37.591Z",
												"title": "string",
												"value": 0
											}
										],
										"current_period_from": "2024-12-13T06:49:37.591Z",
										"current_period_to": "2024-12-13T06:49:37.591Z",
										"ozon_actions_exist": true
									},
									"offer_id": "356792",
									"price": {
										"auto_action_enabled": true,
										"currency_code": "RUB",
										"marketing_price": 0,
										"marketing_seller_price": 0,
										"min_price": 0,
										"old_price": 579,
										"price": 499,
										"retail_price": 0,
										"vat": 0.2
									},
									"price_indexes": {
										"color_index": "WITHOUT_INDEX",
										"external_index_data": {
											"min_price": 0,
											"min_price_currency": "string",
											"price_index_value": 0
										},
										"ozon_index_data": {
											"min_price": 0,
											"min_price_currency": "string",
											"price_index_value": 0
										},
										"self_marketplaces_index_data": {
											"min_price": 0,
											"min_price_currency": "string",
											"price_index_value": 0
										}
									},
									"product_id": 243686911,
									"volume_weight": 0
								}
							],
							"total": 0
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v4/product/info",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Prices(context.Background(), &info.PricesRequest{
		Cursor: "string",
		Filter: info.PricesRequestFilter{
			OfferID: []string{
				"356792",
			},
			ProductID: []string{
				"243686911",
			},
			Visibility: info.PricesRequestFilterVisibilityALL,
		},
		Limit: 100,
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &info.PricesResponse{
		Items: []info.PricesResponseItem{
			{
				Acquiring: 0,
				Commissions: info.PricesResponseItemCommissions{
					FBODelivToCustomerAmount:    14.75,
					FBODirectFlowTransMaxAmount: 46.5,
					FBODirectFlowTransMinAmount: 31,
					FBOReturnFlowAmount:         50,
					FBSDelivToCustomerAmount:    60,
					FBSDirectFlowTransMaxAmount: 61.5,
					FBSDirectFlowTransMinAmount: 41,
					FBSFirstMileMaxAmount:       25,
					FBSFirstMileMinAmount:       0,
					FBSReturnFlowAmount:         40,
					SalesPercentFBO:             15,
					SalesPercentFBS:             0,
				},
				MarketingActions: info.PricesResponseItemMarketingActions{
					Actions: []info.PricesResponseItemMarketingActionsAction{
						{
							DateFrom: "2024-12-13T06:49:37.591Z",
							DateTo:   "2024-12-13T06:49:37.591Z",
							Title:    "string",
							Value:    0,
						},
					},
					CurrentPeriodFrom: "2024-12-13T06:49:37.591Z",
					CurrentPeriodTo:   "2024-12-13T06:49:37.591Z",
					OzonActionsExist:  true,
				},
				OfferID: "356792",
				Price: info.PricesResponseItemPrice{
					AutoActionEnabled:    true,
					CurrencyCode:         info.PricesResponseItemPriceCurrencyCodeRUB,
					MarketingPrice:       0,
					MarketingSellerPrice: 0,
					MinPrice:             0,
					OldPrice:             579,
					Price:                499,
					RetailPrice:          0,
					Vat:                  0.2,
				},
				PriceIndexes: info.PricesResponseItemPriceIndexes{
					ColorIndex: info.PricesResponseItemPriceIndexesColorIndexWITHOUTINDEX,
					ExternalIndexData: info.PricesResponseItemPriceIndexesExternalIndexData{
						MinPrice:         0.0,
						MinPriceCurrency: "string",
						PriceIndexValue:  0,
					},
					OzonIndexData: info.PricesResponseItemPriceIndexesOzonIndexData{
						MinPrice:         0.0,
						MinPriceCurrency: "string",
						PriceIndexValue:  0,
					},
					SelfMarketplacesIndexData: info.PricesResponseItemPriceIndexesSelfMarketplacesIndexData{
						MinPrice:         0.0,
						MinPriceCurrency: "string",
						PriceIndexValue:  0,
					},
				},
				ProductID:    243686911,
				VolumeWeight: 0,
			},
		},
		Total:  0,
		Cursor: "string",
	}, resp)
}
