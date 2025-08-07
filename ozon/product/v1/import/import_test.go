package _import_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	_import "github.com/andmetoo/ozon-api-client/ozon/product/v1/import"

	"github.com/andmetoo/ozon-api-client/internal/auth"
	"github.com/andmetoo/ozon-api-client/internal/test"
	"github.com/stretchr/testify/require"
)

func TestInfo_Success(t *testing.T) {
	c := _import.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/import/info", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"task_id":"172549793"}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": {
								"items": [
									{
										"offer_id": "143210608",
										"product_id": 137285792,
										"status": "imported",
										"errors": [ ]
									}
								],
								"total": 1
							}
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/import",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Info(context.Background(), &_import.InfoRequest{
		TaskID: "172549793",
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &_import.InfoResponse{
		Result: _import.InfoResponseResult{
			Items: []_import.InfoResponseResultItem{
				{
					OfferID:   "143210608",
					ProductID: 137285792,
					Status:    _import.InfoResponseResultItemStatusImported,
					Errors:    make([]_import.InfoResponseResultItemError, 0),
				},
			},
			Total: 1,
		},
	}, resp)
}

func TestStocks_Success(t *testing.T) {
	c := _import.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/import/stocks", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"stocks":[{"offer_id":"PG-2404С1","product_id":55946,"stock":4}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": [
								{
									"product_id": 55946,
									"offer_id": "PG-2404С1",
									"updated": true,
									"errors": []
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/import",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Stocks(context.Background(), &_import.StocksRequest{
		Stocks: []_import.StocksRequestStock{
			{
				OfferID:   "PG-2404С1",
				ProductID: 55946,
				Stock:     4,
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &_import.StocksResponse{
		Result: []_import.StocksResponseResult{
			{
				ProductID: 55946,
				OfferID:   "PG-2404С1",
				Updated:   true,
				Errors:    make([]_import.StocksResponseResultError, 0),
			},
		},
	}, resp)
}

func TestPrices_Success(t *testing.T) {
	c := _import.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-seller.ozon.ru/v1/product/import/prices", test.FullURL(r))
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"prices":[{"auto_action_enabled":"UNKNOWN","currency_code":"RUB","min_price":"800","net_price":"100","offer_id":"","old_price":"0","price":"1448","price_strategy_enabled":"UNKNOWN","product_id":1386}]}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"result": [
								{
									"product_id": 1386,
									"offer_id": "PH8865",
									"updated": true,
									"errors": []
								}
							]
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-seller.ozon.ru/v1/product/import",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.Prices(context.Background(), &_import.PricesRequest{
		Prices: []_import.PricesRequestPrice{
			{
				AutoActionEnabled:    _import.PricesRequestPriceAutoActionEnabledUNKNOWN,
				CurrencyCode:         _import.PricesRequestPriceCurrencyCodeRUB,
				MinPrice:             "800",
				NetPrice:             "100",
				OfferID:              "",
				OldPrice:             "0",
				Price:                "1448",
				PriceStrategyEnabled: _import.PricesRequestPricePriceStrategyEnabledUNKNOWN,
				ProductID:            1386,
			},
		},
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &_import.PricesResponse{
		Result: []_import.PricesResponseResult{
			{
				ProductID: 1386,
				OfferID:   "PH8865",
				Updated:   true,
				Errors:    make([]_import.PricesResponseResultError, 0),
			},
		},
	}, resp)
}
