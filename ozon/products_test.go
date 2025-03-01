package ozon

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetStocksInfo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetStocksInfoParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetStocksInfoParams{
				Limit:  100,
				Cursor: "",
				Filter: GetStocksInfoFilter{
					OfferId:    []string{"136834"},
					ProductId:  []int64{214887921},
					Visibility: "ALL",
				},
			},
			`{
				"cursor": "string",
				"items": [
				  {
					"offer_id": "string",
					"product_id": 123,
					"stocks": [
					  {
						"present": 0,
						"reserved": 0,
						"shipment_type": "SHIPMENT_TYPE_GENERAL",
						"sku": 0,
						"type": "string"
					  }
					]
				  }
				],
				"total": 0
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetStocksInfoParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().GetStocksInfo(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetStocksInfoResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Items) > int(test.params.Limit) {
				t.Errorf("Amount of items in response cannot be bigger than limit")
			}
			if len(resp.Items) > 0 {
				if resp.Items[0].ProductId == 0 {
					t.Errorf("Product id cannot be 0")
				}
				if resp.Items[0].OfferId == "" {
					t.Errorf("Offer id cannot be empty")
				}
			}
		}
	}
}

func TestUpdateStocks(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdateStocksParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdateStocksParams{
				Stocks: []UpdateStocksStock{
					{
						OfferId:   "PG-2404С1",
						ProductId: 55946,
						Stock:     4,
					},
				},
			},
			`{
				"result": [
				  {
					"product_id": 55946,
					"offer_id": "PG-2404С1",
					"updated": true,
					"errors": []
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&UpdateStocksParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().UpdateStocks(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &UpdateStocksResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) != len(test.params.Stocks) {
				t.Errorf("Length of stocks in request and response are not equal")
			}
			if len(resp.Result) > 0 {
				if resp.Result[0].OfferId != test.params.Stocks[0].OfferId {
					t.Errorf("Offer ids in request and response are not equal")
				}
				if resp.Result[0].ProductId != test.params.Stocks[0].ProductId {
					t.Errorf("Product ids in request and response are not equal")
				}
			}
		}
	}
}

func TestStocksInSellersWarehouse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *StocksInSellersWarehouseParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&StocksInSellersWarehouseParams{
				SKU: []string{"123"},
			},
			`{
				"result": [
				  {
					"sku": 12,
					"present": 34,
					"product_id": 548761,
					"reserved": 5,
					"warehouse_id": 156778,
					"warehouse_name": "something"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&StocksInSellersWarehouseParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().StocksInSellersWarehouse(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &StocksInSellersWarehouseResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) != len(test.params.SKU) {
				t.Errorf("Length of skus in request and response must be equal")
			}
			if len(resp.Result) > 0 {
				if fmt.Sprint(resp.Result[0].SKU) == test.params.SKU[0] {
					t.Errorf("fbs sku in request and response are not equal")
				}
			}
		}
	}
}

func TestUpdatePrices(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdatePricesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdatePricesParams{
				Prices: []UpdatePricesPrice{
					{
						AutoActionEnabled:    "UNKNOWN",
						CurrencyCode:         "RUB",
						MinPrice:             "800",
						OldPrice:             "0",
						Price:                "1448",
						ProductId:            1386,
						PriceStrategyEnabled: PriceStrategyUnknown,
					},
				},
			},
			`{
				"result": [
				  {
					"product_id": 1386,
					"offer_id": "PH8865",
					"updated": true,
					"errors": []
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&UpdatePricesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().UpdatePrices(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &UpdatePricesResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) != len(test.params.Prices) {
				t.Errorf("Length of prices in request and response are not equal")
			}
			if len(resp.Result) > 0 {
				if resp.Result[0].ProductId != test.params.Prices[0].ProductId {
					t.Errorf("Product ids in request and response are not equal")
				}
			}
		}
	}
}

func TestUpdateQuantityStockProducts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdateQuantityStockProductsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdateQuantityStockProductsParams{
				Stocks: []UpdateQuantityStockProductsStock{
					{
						OfferId:     "PH11042",
						ProductId:   118597312,
						Stock:       100,
						WarehouseId: 22142605386000,
					},
				},
			},
			`{
				"result": [
				  {
					"warehouse_id": 22142605386000,
					"product_id": 118597312,
					"offer_id": "PH11042",
					"updated": true,
					"errors": []
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&UpdateQuantityStockProductsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().UpdateQuantityStockProducts(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &UpdateQuantityStockProductsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) != len(test.params.Stocks) {
				t.Errorf("Length of stocks in request and response are not equal")
			}
			if len(resp.Result) > 0 {
				if resp.Result[0].Offerid != test.params.Stocks[0].OfferId {
					t.Errorf("Offer ids in request and response are not equal")
				}
				if resp.Result[0].ProductId != test.params.Stocks[0].ProductId {
					t.Errorf("Product ids in request and response are not equal")
				}
				if resp.Result[0].WarehouseId != test.params.Stocks[0].WarehouseId {
					t.Errorf("Warehouse ids in request and response are not equal")
				}
			}
		}
	}
}

func TestCreateOrUpdateProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreateOrUpdateProductParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreateOrUpdateProductParams{
				Items: []CreateOrUpdateProductItem{
					{
						Attributes: []CreateOrUpdateAttribute{
							{
								ComplexId: 0,
								Id:        5076,
								Values: []CreateOrUpdateAttributeValue{
									{
										DictionaryValueId: 971082156,
										Value:             "Стойка для акустической системы",
									},
								},
							},
							{
								ComplexId: 0,
								Id:        9048,
								Values: []CreateOrUpdateAttributeValue{
									{
										Value: "Комплект защитных плёнок для X3 NFC. Темный хлопок",
									},
								},
							},
							{
								ComplexId: 0,
								Id:        8229,
								Values: []CreateOrUpdateAttributeValue{
									{
										DictionaryValueId: 95911,
										Value:             "Комплект защитных плёнок для X3 NFC. Темный хлопок",
									},
								},
							},
							{
								ComplexId: 0,
								Id:        85,
								Values: []CreateOrUpdateAttributeValue{
									{
										DictionaryValueId: 5060050,
										Value:             "Samsung",
									},
								},
							},
							{
								ComplexId: 0,
								Id:        10096,
								Values: []CreateOrUpdateAttributeValue{
									{
										DictionaryValueId: 61576,
										Value:             "серый",
									},
								},
							},
						},
						Barcode:               "112772873170",
						DescriptionCategoryId: 17033876,
						CurrencyCode:          "RUB",
						Depth:                 10,
						DimensionUnit:         "mm",
						Height:                250,
						Name:                  "Комплект защитных плёнок для X3 NFC. Темный хлопок",
						OfferId:               "143210608",
						OldPrice:              "1100",
						Price:                 "1000",
						VAT:                   "0.1",
						Weight:                100,
						WeightUnit:            "g",
						Width:                 150,
					},
				},
			},
			`{
				"result": {
				  "task_id": 172549793
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CreateOrUpdateProductParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().CreateOrUpdateProduct(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CreateOrUpdateProductResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.TaskId == 0 {
				t.Errorf("Task id cannot be 0")
			}
		}
	}
}

func TestGetListOfProducts(t *testing.T) {
	t.Parallel()

	testTimeout := 5 * time.Second

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetListOfProductsParams
		response   string
	}{
		// Test OK
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetListOfProductsParams{
				Filter: GetListOfProductsFilter{
					OfferId:    []string{"136748"},
					ProductId:  []int64{223681945},
					Visibility: "ALL",
				},
				LastId: "",
				Limit:  100,
			},
			`{
				"result": {
				  "items": [
					{
					  "product_id": 223681945,
					  "offer_id": "136748",
					  "has_fbo_stocks": false,
					  "has_fbs_stocks": true,
					  "archived": false,
					  "is_discounted": true,
					  "quants": [
						{
						  "warehouse_id": 123,
						  "quantity": 50,
						  "reserved": 10
						}
					  ]
					}
				  ],
				  "total": 1,
				  "last_id": "bnVсbA=="
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetListOfProductsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
		defer cancel()

		resp, err := c.Products().GetListOfProducts(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetListOfProductsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result.Items) != int(resp.Result.Total) {
				t.Errorf("Length of items is not equal total")
			}
			if resp.Result.Total > int32(test.params.Limit) {
				t.Errorf("Length of items is bigger than limit")
			}
			if len(resp.Result.Items) > 0 {
				if resp.Result.Items[0].OfferId == "" {
					t.Errorf("Offer id cannot be empty")
				}
				if resp.Result.Items[0].ProductId == 0 {
					t.Errorf("Product id cannot be 0")
				}
				// Optional: check we successfully parse quants
				if len(resp.Result.Items[0].Quants) == 0 {
					t.Errorf("Expected some quants, got none")
				}
			}
		}
	}
}

func TestGetProductsRatingBySKU(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProductsRatingBySKUParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductsRatingBySKUParams{
				SKUs: []int64{179737222},
			},
			`{
				"products": [
				  {
					"sku": 179737222,
					"rating": 42.5,
					"groups": [
					  {
						"key": "media",
						"name": "Медиа",
						"rating": 70,
						"weight": 25,
						"conditions": [
						  {
							"key": "media_images_2",
							"description": "Добавлено 2 изображения",
							"fulfilled": true,
							"cost": 50
						  },
						  {
							"key": "media_images_3",
							"description": "Добавлено 3 изображения и более",
							"fulfilled": true,
							"cost": 20
						  },
						  {
							"key": "media_image_3d",
							"description": "Добавлено 3D-изображение",
							"fulfilled": false,
							"cost": 15
						  },
						  {
							"key": "media_video",
							"description": "Добавлено видео",
							"fulfilled": false,
							"cost": 15
						  }
						],
						"improve_attributes": [
						  {
							"id": 4074,
							"name": "Код ролика на YouTube"
						  },
						  {
							"id": 4080,
							"name": "3D-изображение"
						  }
						],
						"improve_at_least": 2
					  },
					  {
						"key": "important_attributes",
						"name": "Важные атрибуты",
						"rating": 50,
						"weight": 30,
						"conditions": [
						  {
							"key": "important_2",
							"description": "Заполнено 2 атрибута и более",
							"fulfilled": true,
							"cost": 50
						  },
						  {
							"key": "important_50_percent",
							"description": "Заполнено более 50% атрибутов",
							"fulfilled": false,
							"cost": 25
						  },
						  {
							"key": "important_70_percent",
							"description": "Заполнено более 70% атрибутов",
							"fulfilled": false,
							"cost": 25
						  }
						],
						"improve_attributes": [
						  {
							"id": 4385,
							"name": "Гарантийный срок"
						  },
						  {
							"id": 4389,
							"name": "Страна-изготовитель"
						  },
						  {
							"id": 8513,
							"name": "Количество в упаковке, шт"
						  },
						  {
							"id": 8590,
							"name": "Макс. диагональ, дюймы"
						  },
						  {
							"id": 8591,
							"name": "Мин. диагональ, дюймы"
						  },
						  {
							"id": 9336,
							"name": "Модель браслета/умных часов"
						  },
						  {
							"id": 11046,
							"name": "Покрытие"
						  },
						  {
							"id": 11047,
							"name": "Прозрачность покрытия"
						  },
						  {
							"id": 11048,
							"name": "Дополнительные свойства покрытия"
						  },
						  {
							"id": 11049,
							"name": "Вид стекла"
						  },
						  {
							"id": 11603,
							"name": "Размер циферблата"
						  }
						],
						"improve_at_least": 6
					  },
					  {
						"key": "other_attributes",
						"name": "Остальные атрибуты",
						"rating": 0,
						"weight": 25,
						"conditions": [
						  {
							"key": "other_2",
							"description": "Заполнено 2 атрибута и более",
							"fulfilled": false,
							"cost": 50
						  },
						  {
							"key": "other_50_percent",
							"description": "Заполнено более 50% атрибутов",
							"fulfilled": false,
							"cost": 50
						  }
						],
						"improve_attributes": [
						  {
							"id": 4382,
							"name": "Размеры, мм"
						  }
						],
						"improve_at_least": 1
					  },
					  {
						"key": "text",
						"name": "Текстовое описание",
						"rating": 50,
						"weight": 20,
						"conditions": [
						  {
							"key": "text_annotation_100_chars",
							"description": "Аннотация более 100 знаков",
							"fulfilled": true,
							"cost": 25
						  },
						  {
							"key": "text_annotation_500_chars",
							"description": "Аннотация более 500 знаков",
							"fulfilled": true,
							"cost": 25
						  },
						  {
							"key": "text_rich",
							"description": "Заполнен Rich-контент",
							"fulfilled": false,
							"cost": 100
						  }
						],
						"improve_attributes": [
						  {
							"id": 11254,
							"name": "Rich-контент JSON"
						  }
						],
						"improve_at_least": 1
					  }
					]
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetProductsRatingBySKUParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().GetProductsRatingBySKU(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetProductsRatingBySKUResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Products) != len(test.params.SKUs) {
				t.Errorf("Length of products in response is not equal length of skus in request")
			}
			if len(resp.Products) > 0 {
				if resp.Products[0].SKU != test.params.SKUs[0] {
					t.Errorf("SKU in request and response are not equal")
				}
			}
		}
	}
}

func TestGetProductImportStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProductImportStatusParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductImportStatusParams{
				TaskId: 172549793,
			},
			`{
				"result": {
				  "items": [
					{
					  "offer_id": "143210608",
					  "product_id": 137285792,
					  "status": "imported",
					  "errors": []
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
			&GetProductImportStatusParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().GetProductImportStatus(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetProductImportStatusResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result.Items) > 0 {
				if resp.Result.Items[0].ProductId == 0 {
					t.Errorf("Product id cannot be 0")
				}
				if resp.Result.Items[0].OfferId == "" {
					t.Errorf("Offer id cannot be empty")
				}
			}
		}
	}
}

func TestCreateProductByOzonID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreateProductByOzonIDParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreateProductByOzonIDParams{
				Items: []CreateProductsByOzonIDItem{
					{
						Name:         "string",
						OfferId:      "91132",
						OldPrice:     "2590",
						Price:        "2300",
						CurrencyCode: "RUB",
						SKU:          298789742,
						VAT:          "0.1",
					},
				},
			},
			`{
				"result": {
				  "task_id": 176594213,
				  "unmatched_sku_list": []
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CreateProductByOzonIDParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().CreateProductByOzonID(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CreateProductByOzonIDResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestUpdateProductImages(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdateProductImagesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdateProductImagesParams{
				ColorImage: "string",
				Images:     []string{"string"},
				Images360:  []string{"string"},
				ProductId:  12345,
			},
			`{
				"result": {
				  "pictures": [
					{
					  "is_360": true,
					  "is_color": true,
					  "is_primary": true,
					  "product_id": 12345,
					  "state": "string",
					  "url": "string"
					},
					{
						"is_360": false,
						"is_color": true,
						"is_primary": true,
						"product_id": 12345,
						"state": "string",
						"url": "string"
					  }
				  ]
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&UpdateProductImagesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().UpdateProductImages(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ProductInfoResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result.Pictures) != len(test.params.Images)+len(test.params.Images360) {
				t.Errorf("Amount of pictures in request and response are not equal")
			}
			if len(resp.Result.Pictures) > 0 {
				if resp.Result.Pictures[0].ProductId != test.params.ProductId {
					t.Errorf("Product ids in request and response are not equal")
				}
			}
		}
	}
}

func TestCheckImageUploadingStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CheckImageUploadingStatusParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CheckImageUploadingStatusParams{
				ProductId: []int64{123456},
			},
			`{
				"items": [
				  {
					"product_id": 123456,
					"primary_photo": [
					  "string"
					],
					"photo": [
					  "string"
					],
					"color_photo": [
					  "string"
					],
					"photo_360": [
					  "string"
					]
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CheckImageUploadingStatusParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().CheckImageUploadingStatus(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CheckImageUploadingStatusResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Items) > 0 {
				if resp.Items[0].ProductId != test.params.ProductId[0] {
					t.Errorf("Product ids in request and response are not equal")
				}
			}
		}
	}
}

func TestListProductsByIDs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListProductsByIDsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListProductsByIDsParams{
				OfferId: []string{"010", "23"},
			},
			`{
				"items": [
				  {
					"barcodes": [
					  "string"
					],
					"color_image": [
					  "string"
					],
					"commissions": [
					  {
						"delivery_amount": 0,
						"percent": 0,
						"return_amount": 0,
						"sale_schema": "string",
						"value": 0
					  }
					],
					"created_at": "2019-08-24T14:15:22Z",
					"currency_code": "string",
					"description_category_id": 0,
					"discounted_fbo_stocks": 0,
					"errors": [
					  {
						"attribute_id": 0,
						"code": "string",
						"field": "string",
						"level": "ERROR_LEVEL_UNSPECIFIED",
						"state": "string",
						"texts": {
						  "attribute_name": "string",
						  "description": "string",
						  "hint_code": "string",
						  "message": "string",
						  "params": [
							{
							  "name": "string",
							  "value": "string"
							}
						  ],
						  "short_description": "string"
						}
					  }
					],
					"has_discounted_fbo_item": true,
					"id": 0,
					"images": [
					  "string"
					],
					"images360": [
					  "string"
					],
					"is_archived": true,
					"is_autoarchived": true,
					"is_discounted": true,
					"is_kgt": true,
					"is_prepayment_allowed": true,
					"is_super": true,
					"marketing_price": "string",
					"min_price": "string",
					"model_info": {
					  "count": 0,
					  "model_id": 0
					},
					"name": "string",
					"offer_id": "string",
					"old_price": "string",
					"price": "string",
					"price_indexes": {
					  "color_index": "COLOR_INDEX_UNSPECIFIED",
					  "external_index_data": {
						"minimal_price": "string",
						"minimal_price_currency": "string",
						"price_index_value": 0
					  },
					  "ozon_index_data": {
						"minimal_price": "string",
						"minimal_price_currency": "string",
						"price_index_value": 0
					  },
					  "self_marketplaces_index_data": {
						"minimal_price": "string",
						"minimal_price_currency": "string",
						"price_index_value": 0
					  }
					},
					"primary_image": [
					  "string"
					],
					"sources": [
					  {
						"created_at": "2019-08-24T14:15:22Z",
						"quant_code": "string",
						"shipment_type": "SHIPMENT_TYPE_UNSPECIFIED",
						"sku": 0,
						"source": "string"
					  }
					],
					"statuses": {
					  "is_created": true,
					  "moderate_status": "string",
					  "status": "string",
					  "status_description": "string",
					  "status_failed": "string",
					  "status_name": "string",
					  "status_tooltip": "string",
					  "status_updated_at": "2019-08-24T14:15:22Z",
					  "validation_status": "string"
					},
					"stocks": {
					  "has_stock": true,
					  "stocks": [
						{
						  "present": 0,
						  "reserved": 0,
						  "sku": 0,
						  "source": "string"
						}
					  ]
					},
					"type_id": 0,
					"updated_at": "2019-08-24T14:15:22Z",
					"vat": "string",
					"visibility_details": {
					  "has_price": true,
					  "has_stock": true
					},
					"volume_weight": 0
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListProductsByIDsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().ListProductsByIDs(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ListProductsByIDsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetDescriptionOfProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetDescriptionOfProductParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetDescriptionOfProductParams{
				Filter: GetDescriptionOfProductFilter{
					ProductId:  []int64{213761435},
					Visibility: "ALL",
				},
				LastId: "okVsfA==«",
				SortBy: "ASC",
				Limit:  100,
			},
			`{
				"result": [
				  {
					"id": 213761435,
					"barcode": "",
					"description_category_id": 17038062,
					"name": "Пленка защитная для Xiaomi Redmi Note 10 Pro 5G",
					"offer_id": "21470",
					"height": 10,
					"depth": 210,
					"width": 140,
					"dimension_unit": "mm",
					"weight": 50,
					"weight_unit": "g",
					"images": [
					  {
						"file_name": "https://cdn1.ozone.ru/s3/multimedia-f/6190456071.jpg",
						"default": true,
						"index": 0
					  },
					  {
						"file_name": "https://cdn1.ozone.ru/s3/multimedia-7/6190456099.jpg",
						"default": false,
						"index": 1
					  },
					  {
						"file_name": "https://cdn1.ozone.ru/s3/multimedia-9/6190456065.jpg",
						"default": false,
						"index": 2
					  }
					],
					"images360": [],
					"pdf_list": [],
					"attributes": [
					  {
						"attribute_id": 5219,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 970718176,
							"value": "универсальный"
						  }
						]
					  },
					  {
						"attribute_id": 11051,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 970736931,
							"value": "Прозрачный"
						  }
						]
					  },
					  {
						"attribute_id": 10100,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 0,
							"value": "false"
						  }
						]
					  },
					  {
						"attribute_id": 11794,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 970860783,
							"value": "safe"
						  }
						]
					  },
					  {
						"attribute_id": 9048,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 0,
							"value": "Пленка защитная для Xiaomi Redmi Note 10 Pro 5G"
						  }
						]
					  },
					  {
						"attribute_id": 5076,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 39638,
							"value": "Xiaomi"
						  }
						]
					  },
					  {
						"attribute_id": 9024,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 0,
							"value": "21470"
						  }
						]
					  },
					  {
						"attribute_id": 10015,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 0,
							"value": "false"
						  }
						]
					  },
					  {
						"attribute_id": 85,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 971034861,
							"value": "Brand"
						  }
						]
					  },
					  {
						"attribute_id": 9461,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 349824787,
							"value": "Защитная пленка для смартфона"
						  }
						]
					  },
					  {
						"attribute_id": 4180,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 0,
							"value": "Пленка защитная для Xiaomi Redmi Note 10 Pro 5G"
						  }
						]
					  },
					  {
						"attribute_id": 4191,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 0,
							"value": "Пленка предназначена для модели Xiaomi Redmi Note 10 Pro 5G. Защитная гидрогелевая пленка обеспечит защиту вашего смартфона от царапин, пыли, сколов и потертостей."
						  }
						]
					  },
					  {
						"attribute_id": 8229,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 91521,
							"value": "Защитная пленка"
						  }
						]
					  }
					],
					"complex_attributes": [],
					"color_image": ""
				  }
				],
				"total": 1,
				"last_id": "onVsfA=="
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetDescriptionOfProductParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().GetDescriptionOfProduct(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetDescriptionOfProductResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) != len(test.params.Filter.ProductId) && len(resp.Result) != len(test.params.Filter.OfferId) {
				t.Errorf("Amount of products in request and response are not equal")
			}
		}
	}
}

func TestGetDescriptionOfProductV4(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetDescriptionOfProductsParams
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetDescriptionOfProductsParams{
				Filter: GetDescriptionOfProductsFilter{
					ProductId:  []string{"330186294"},
					Visibility: "ALL",
				},
				Limit:         100,
				SortDirection: "ASC",
			},
			`{
				"result": [
				  {
					"id": 330186294,
					"barcode": "OZN653473453",
					"name": "PC ЮКОМС Ryzen 7 5700G ...",
					"offer_id": "ju-cas2-r5700g-bl",
					"height": 360,
					"depth": 420,
					"width": 220,
					"dimension_unit": "mm",
					"weight": 4500,
					"weight_unit": "g",
					"description_category_id": 17028619,
					"type_id": 91476,
					"primary_image": "https://cdn1.ozone.ru/s3/multimedia-1-3/7084786431.jpg",
					"model_info": {
					  "model_id": 379410772,
					  "count": 126
					},
					"images": [
					  "https://cdn1.ozone.ru/s3/multimedia-1-0/7084786428.jpg",
					  "https://cdn1.ozone.ru/s3/multimedia-1-k/7084786304.jpg"
					],
					"pdf_list": [],
					"attributes": [
					  {
						"id": 85,
						"complex_id": 0,
						"values": [
						  {
							"dictionary_value_id": 971195426,
							"value": "ЮКОМС"
						  }
						]
					  }
					],
					"complex_attributes": [],
					"color_image": ""
				  }
				],
				"total": 1,
				"last_id": ""
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetDescriptionOfProductsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(
			test.statusCode,
			test.response,
			test.headers,
		))

		ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
		defer cancel()

		resp, err := c.Products().GetDescriptionOfProducts(ctx, test.params)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			continue
		}

		compareJsonResponse(t, test.response, &GetDescriptionOfProductsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("wrong status code: got: %d, want: %d", resp.StatusCode, test.statusCode)
		}

		if test.statusCode == http.StatusOK {
			if len(resp.Result) == 0 {
				t.Error("expected non-empty result in success case")
			}
		}
	}
}

func TestGetProductDescription(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProductDescriptionParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductDescriptionParams{
				OfferId:   "5",
				ProductId: 73453843,
			},
			`{
				"result": {
				  "id": 73453843,
				  "offer_id": "5",
				  "name": "Онлайн курс по дрессировке собак \"Воспитанная собака за 3 недели\"",
				  "description": "Экспресс-курс - это сокращённый вариант курса \"Собака: инструкция по применению\", дающий базовый минимум знаний, навыков, умений. Это оптимальный вариант для совершения первых шагов по воспитанию!<br/><br/>Что дает Экспресс-курс:<ul><li>Контакт с собакой </li></ul>К концу экспресс-курса дрессировки вы получаете воспитанного друга и соратника, который ориентируется на вас в любой ситуации.<ul><li>Уверенность в безопасности</li></ul>Благополучие собаки: больше не будет срывов с поводка, преследования кошек, попыток съесть что-либо на улице и т. д.<ul><li>Комфортная жизнь</li></ul>Принципиально другой уровень общения, без раздражения, криков и недовольства поведением животного."
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetProductDescriptionParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().GetProductDescription(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetProductDescriptionResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.Id != test.params.ProductId {
				t.Errorf("Product ids in request and response are not equal")
			}
			if resp.Result.OfferId != test.params.OfferId {
				t.Errorf("Offer ids in request and response are not equal")
			}
		}
	}
}

func TestGetProductRangeLimit(t *testing.T) {
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
				"daily_create": {
				  "limit": 0,
				  "reset_at": "2019-08-24T14:15:22Z",
				  "usage": 0
				},
				"daily_update": {
				  "limit": 0,
				  "reset_at": "2019-08-24T14:15:22Z",
				  "usage": 0
				},
				"total": {
				  "limit": 0,
				  "usage": 0
				}
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
		resp, err := c.Products().GetProductRangeLimit(ctx)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetProductRangeLimitResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestChangeProductIDs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ChangeProductIDsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ChangeProductIDsParams{
				UpdateOfferId: []ChangeProductIDsUpdateOffer{
					{
						NewOfferId: "new id",
						OfferId:    "old id",
					},
				},
			},
			`{
				"errors": [
				  {
					"message": "string",
					"offer_id": "string"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ChangeProductIDsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().ChangeProductIDs(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ChangeProductIDsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestArchiveProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ArchiveProductParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ArchiveProductParams{
				ProductId: []int64{125529926},
			},
			`{
				"result": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ArchiveProductParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().ArchiveProduct(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ArchiveProductResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestUnarchiveProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ArchiveProductParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ArchiveProductParams{
				ProductId: []int64{125529926},
			},
			`{
				"result": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ArchiveProductParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().UnarchiveProduct(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ArchiveProductResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestRemoveProductWithoutSKU(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *RemoveProductWithoutSKUParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&RemoveProductWithoutSKUParams{
				Products: []RemoveProductWithoutSKUProduct{
					{
						OfferId: "033",
					},
				},
			},
			`{
				"status": [
				  {
					"offer_id": "033",
					"is_deleted": true,
					"error": ""
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&RemoveProductWithoutSKUParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().RemoveProductWithoutSKU(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &RemoveProductWithoutSKUResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Status) > 0 {
				if resp.Status[0].OfferId != test.params.Products[0].OfferId {
					t.Errorf("Offer ids in request and response are not equal")
				}
			}
		}
	}
}

func TestUploadActivationCodes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UploadActivationCodesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UploadActivationCodesParams{
				DigitalCodes: []string{"764282654334"},
				ProductId:    73160317,
			},
			`{
				"result": {
				  "task_id": 172549811
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&UploadActivationCodesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().UploadActivationCodes(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &UploadActivationCodesResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestStatusOfUploadingActivationCodes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *StatusOfUploadingActivationCodesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&StatusOfUploadingActivationCodesParams{
				TaskId: 178574231,
			},
			`{
				"result": {
				  "status": "imported"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&StatusOfUploadingActivationCodesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().StatusOfUploadingActivationCodes(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &StatusOfUploadingActivationCodesResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetProductPriceInfo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProductPriceInfoParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductPriceInfoParams{
				Filter: GetProductPriceInfoFilter{
					OfferId:    []string{"356792"},
					ProductId:  []int64{243686911},
					Visibility: "ALL",
				},
			},
			`{
				"result": {
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
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetProductPriceInfoParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().GetProductPriceInfo(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetProductPriceInfoResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetMarkdownInfo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetMarkdownInfoParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetMarkdownInfoParams{
				DiscountedSKUs: []string{"635548518"},
			},
			`{
				"items": [
				  {
					"discounted_sku": 635548518,
					"sku": 320067758,
					"condition_estimation": "4",
					"packaging_violation": "",
					"warranty_type": "",
					"reason_damaged": "Механическое повреждение",
					"comment_reason_damaged": "повреждена заводская упаковка",
					"defects": "",
					"mechanical_damage": "",
					"package_damage": "",
					"shortage": "",
					"repair": "",
					"condition": ""
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetMarkdownInfoParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().GetMarkdownInfo(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			compareJsonResponse(t, test.response, &GetMarkdownInfoResponse{})

			if len(resp.Items) > 0 {
				if fmt.Sprint(resp.Items[0].DiscountedSKU) != test.params.DiscountedSKUs[0] {
					t.Errorf("SKUs in reqest and resonse are not equal")
				}
			}
		}
	}
}

func TestSetDiscountOnMarkdownProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *SetDiscountOnMarkdownProductParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&SetDiscountOnMarkdownProductParams{
				Discount:  0,
				ProductId: 0,
			},
			`{
				"result": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&SetDiscountOnMarkdownProductParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().SetDiscountOnMarkdownProduct(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		compareJsonResponse(t, test.response, &SetDiscountOnMarkdownProductResponse{})
	}
}

func TestNumberOfSubsToProductAvailability(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *NumberOfSubsToProductAvailabilityParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&NumberOfSubsToProductAvailabilityParams{
				SKUS: []int64{1234},
			},
			`{
				"result": [
				  {
					"count": 2,
					"sku": 1234
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&NumberOfSubsToProductAvailabilityParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().NumberOfSubsToProductAvailability(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		compareJsonResponse(t, test.response, &NumberOfSubsToProductAvailabilityResponse{})

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) != len(test.params.SKUS) {
				t.Errorf("Length of SKUS in request and response are not equal")
			}

			if len(resp.Result) > 0 {
				if resp.Result[0].SKU != test.params.SKUS[0] {
					t.Errorf("SKU in request and response are not equal")
				}
			}
		}
	}
}

func TestUpdateCharacteristics(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdateCharacteristicsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdateCharacteristicsParams{
				Items: []UpdateCharacteristicsItem{
					{
						Attributes: []UpdateCharacteristicsItemAttribute{
							{
								ComplexId: 0,
								Id:        0,
								Values: []UpdateCharacteristicsItemValue{
									{
										DictionaryValueId: 0,
										Value:             "string",
									},
								},
							},
						},
						OfferId: "string",
					},
				},
			},
			`{
				"task_id": 0
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&UpdateCharacteristicsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().UpdateCharacteristics(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		compareJsonResponse(t, test.response, &UpdateCharacteristicsResponse{})
	}
}

func TestGetRelatedSKUs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetRelatedSKUsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetRelatedSKUsParams{
				SKUs: []string{"321", "322"},
			},
			`{
				"items": [
				  {
					"availability": "HIDDEN",
					"deleted_at": "2019-08-24T14:15:22Z",
					"delivery_schema": "fbs",
					"product_id": 123,
					"sku": 321
				  }
				],
				"errors": [
				  {
					"code": "test_code",
					"sku": 322,
					"message": "test_message"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetRelatedSKUsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().GetRelatedSKUs(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		compareJsonResponse(t, test.response, &GetRelatedSKUsResponse{})

		if len(resp.Errors)+len(resp.Items) != len(test.params.SKUs) {
			t.Errorf("expected equal length of skus in request and response")
		}
	}
}

func TestEconomyInfo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetEconomyInfoParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetEconomyInfoParams{
				QuantCode: []string{"321", "322"},
			},
			`{
				"items": [
				  {
					"offer_id": "string",
					"product_id": 0,
					"quant_info": {
					  "quants": [
						{
						  "barcodes_extended": [
							{
							  "barcode": "string",
							  "error": "string",
							  "status": "string"
							}
						  ],
						  "dimensions": {
							"depth": 0,
							"height": 0,
							"weight": 0,
							"width": 0
						  },
						  "marketing_price": {
							"price": "string",
							"seller_price": "string"
						  },
						  "min_price": "string",
						  "old_price": "string",
						  "price": "string",
						  "quant_code": "string",
						  "quant_sice": 0,
						  "shipment_type": "string",
						  "sku": 0,
						  "statuses": {
							"state_description": "string",
							"state_name": "string",
							"state_sys_name": "string",
							"state_tooltip": "string"
						  }
						}
					  ]
					}
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetEconomyInfoParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().EconomyInfo(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		compareJsonResponse(t, test.response, &GetEconomyInfoResponse{})
	}
}

func TestListEconomy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListEconomyProductsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListEconomyProductsParams{
				Cursor:     "string",
				Limit:      10,
				Visibility: "ALL",
			},
			`{
				"cursor": "string",
				"products": [
				  {
					"offer_id": "string",
					"product_id": 0,
					"quants": [
					  {
						"quant_code": "string",
						"quant_size": 0
					  }
					]
				  }
				],
				"total_items": 0
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListEconomyProductsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().ListEconomy(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		compareJsonResponse(t, test.response, &ListEconomyProductsResponse{})
	}
}

func TestUpdatePriceRelevanceTimer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdatePriceRelevanceTimerParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdatePriceRelevanceTimerParams{
				ProductIds: []string{"string"},
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&UpdatePriceRelevanceTimerParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().UpdatePriceRelevanceTimer(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		compareJsonResponse(t, test.response, &UpdatePriceRelevanceTimerResponse{})
	}
}

func TestStatusPriceRelevanceTimer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *StatusPriceRelevanceTimerParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&StatusPriceRelevanceTimerParams{
				ProductIds: []string{"string"},
			},
			`{
				"statuses": [
				  {
					"expired_at": "2019-08-24T14:15:22Z",
					"min_price_for_auto_actions_enabled": true,
					"product_id": 0
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&StatusPriceRelevanceTimerParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Products().StatusPriceRelevanceTimer(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		compareJsonResponse(t, test.response, &StatusPriceRelevanceTimerResponse{})
	}
}
