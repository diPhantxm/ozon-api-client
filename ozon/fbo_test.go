package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetFBOShipmentsList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetFBOShipmentsListParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetFBOShipmentsListParams{
				Direction: "ASC",
				Filter: GetFBOShipmentsListFilter{
					Since:  core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-09-01T00:00:00.000Z"),
					Status: "awaiting_packaging",
					To:     core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-17T10:44:12.828Z"),
				},
				Limit:    5,
				Offset:   0,
				Translit: true,
				With: GetFBOShipmentsListWith{
					AnalyticsData: true,
					FinancialData: true,
				},
			},
			`{
				"result": [
				  {
					"order_id": 354680487,
					"order_number": "16965409-0014",
					"posting_number": "16965409-0014-1",
					"status": "delivered",
					"cancel_reason_id": 0,
					"created_at": "2021-09-01T00:23:45.607000Z",
					"in_process_at": "2021-09-01T00:25:30.120000Z",
					"products": [
					  {
						"sku": 160249683,
						"name": "Так говорил Омар Хайям. Жизнеописание. Афоризмы и рубайят. Классика в словах и картинках",
						"quantity": 1,
						"offer_id": "978-5-906864-56-7",
						"price": "81.00",
						"digital_codes": [],
						"currency_code": "RUB"
					  }
					],
					"analytics_data": {
					  "region": "РОСТОВСКАЯ ОБЛАСТЬ",
					  "city": "Ростов-на-Дону",
					  "delivery_type": "PVZ",
					  "is_premium": false,
					  "payment_type_group_name": "Карты оплаты",
					  "warehouse_id": 17717042026000,
					  "warehouse_name": "РОСТОВ-НА-ДОНУ_РФЦ",
					  "is_legal": false
					},
					"financial_data": {
					  "products": [
						{
						  "commission_amount": 12.15,
						  "commission_percent": 15,
						  "payout": 68.85,
						  "product_id": 160249683,
						  "currency_code": "RUB",
						  "old_price": 115,
						  "price": 81,
						  "total_discount_value": 34,
						  "total_discount_percent": 29.57,
						  "actions": [
							"Системная виртуальная скидка селлера"
						  ],
						  "picking": null,
						  "quantity": 0,
						  "client_price": "",
						  "item_services": {
							"marketplace_service_item_fulfillment": -31.5,
							"marketplace_service_item_pickup": 0,
							"marketplace_service_item_dropoff_pvz": 0,
							"marketplace_service_item_dropoff_sc": 0,
							"marketplace_service_item_dropoff_ff": 0,
							"marketplace_service_item_direct_flow_trans": -5,
							"marketplace_service_item_return_flow_trans": 0,
							"marketplace_service_item_deliv_to_customer": -20,
							"marketplace_service_item_return_not_deliv_to_customer": 0,
							"marketplace_service_item_return_part_goods_customer": 0,
							"marketplace_service_item_return_after_deliv_to_customer": 0
						  }
						}
					  ],
					  "posting_services": {
						"marketplace_service_item_fulfillment": 0,
						"marketplace_service_item_pickup": 0,
						"marketplace_service_item_dropoff_pvz": 0,
						"marketplace_service_item_dropoff_sc": 0,
						"marketplace_service_item_dropoff_ff": 0,
						"marketplace_service_item_direct_flow_trans": 0,
						"marketplace_service_item_return_flow_trans": 0,
						"marketplace_service_item_deliv_to_customer": 0,
						"marketplace_service_item_return_not_deliv_to_customer": 0,
						"marketplace_service_item_return_part_goods_customer": 0,
						"marketplace_service_item_return_after_deliv_to_customer": 0
					  }
					},
					"additional_data": []
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetFBOShipmentsListParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().GetShipmentsList(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &GetFBSShipmentsListResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetShipmentDetails(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetShipmentDetailsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetShipmentDetailsParams{
				PostingNumber: "50520644-0012-7",
				Translit:      true,
				With: GetShipmentDetailsWith{
					AnalyticsData: true,
					FinancialData: true,
				},
			},
			`{
				"result": {
				  "order_id": 354679434,
				  "order_number": "50520644-0012",
				  "posting_number": "50520644-0012-7",
				  "status": "delivered",
				  "cancel_reason_id": 0,
				  "created_at": "2021-09-01T00:34:56.563Z",
				  "in_process_at": "2021-09-01T00:34:56.103Z",
				  "products": [
					{
					  "sku": 254665483,
					  "name": "Мочалка натуральная из люфы с деревянной ручкой",
					  "quantity": 1,
					  "offer_id": "PS1033",
					  "price": "137.00",
					  "digital_codes": [],
					  "currency_code": "RUB"
					}
				  ],
				  "analytics_data": {
					"region": "МОСКВА",
					"city": "Москва",
					"delivery_type": "Courier",
					"is_premium": false,
					"payment_type_group_name": "Карты оплаты",
					"warehouse_id": 15431806189000,
					"warehouse_name": "ХОРУГВИНО_РФЦ",
					"is_legal": false
				  },
				  "financial_data": {
					"products": [
					  {
						"commission_amount": 13.7,
						"commission_percent": 10,
						"payout": 123.3,
						"product_id": 254665483,
						"currency_code": "RUB",
						"old_price": 198,
						"price": 137,
						"total_discount_value": 61,
						"total_discount_percent": 30.81,
						"actions": [
						  "Системная виртуальная скидка селлера"
						],
						"picking": null,
						"quantity": 0,
						"client_price": "",
						"item_services": {
						  "marketplace_service_item_fulfillment": -31.5,
						  "marketplace_service_item_pickup": 0,
						  "marketplace_service_item_dropoff_pvz": 0,
						  "marketplace_service_item_dropoff_sc": 0,
						  "marketplace_service_item_dropoff_ff": 0,
						  "marketplace_service_item_direct_flow_trans": -5,
						  "marketplace_service_item_return_flow_trans": 0,
						  "marketplace_service_item_deliv_to_customer": -20,
						  "marketplace_service_item_return_not_deliv_to_customer": 0,
						  "marketplace_service_item_return_part_goods_customer": 0,
						  "marketplace_service_item_return_after_deliv_to_customer": 0
						}
					  }
					],
					"posting_services": {
					  "marketplace_service_item_fulfillment": 0,
					  "marketplace_service_item_pickup": 0,
					  "marketplace_service_item_dropoff_pvz": 0,
					  "marketplace_service_item_dropoff_sc": 0,
					  "marketplace_service_item_dropoff_ff": 0,
					  "marketplace_service_item_direct_flow_trans": 0,
					  "marketplace_service_item_return_flow_trans": 0,
					  "marketplace_service_item_deliv_to_customer": 0,
					  "marketplace_service_item_return_not_deliv_to_customer": 0,
					  "marketplace_service_item_return_part_goods_customer": 0,
					  "marketplace_service_item_return_after_deliv_to_customer": 0
					}
				  },
				  "additional_data": []
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetShipmentDetailsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().GetShipmentDetails(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &GetShipmentDetailsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.PostingNumber != test.params.PostingNumber {
				t.Errorf("Posting numbers in request and response are not equal")
			}
		}
	}
}

func TestListSupplyRequests(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListSupplyRequestsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListSupplyRequestsParams{
				Page:     0,
				PageSize: 0,
				States:   []SupplyRequestState{AcceptanceAtStorageWarehouse},
			},
			`{
				"has_next": true,
				"supply_orders": [
				  {
					"created_at": "string",
					"local_timeslot": {
					  "from": "string",
					  "to": "string"
					},
					"preferred_supply_date_from": "string",
					"preferred_supply_date_to": "string",
					"state": "string",
					"supply_order_id": 0,
					"supply_order_number": "string",
					"supply_warehouse": {
					  "address": "string",
					  "name": "string",
					  "warehouse_id": 0
					},
					"time_left_to_prepare_supply": 0,
					"time_left_to_select_supply_variant": 0,
					"total_items_count": 0,
					"total_quantity": 0
				  }
				],
				"total_supply_orders_count": 0
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListSupplyRequestsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().ListSupplyRequests(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &ListSupplyRequestsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetSupplyRequestInfo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetSupplyRequestInfoParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetSupplyRequestInfoParams{
				SupplyOrderId: 0,
			},
			`{
				"created_at": "string",
				"local_timeslot": {
				  "from": "string",
				  "to": "string"
				},
				"preferred_supply_date_from": "string",
				"preferred_supply_date_to": "string",
				"seller_warehouse": {
				  "address": "string",
				  "name": "string",
				  "warehouse_id": 0
				},
				"state": "string",
				"supply_order_id": 0,
				"supply_order_number": "string",
				"supply_warehouse": {
				  "address": "string",
				  "name": "string",
				  "warehouse_id": 0
				},
				"time_left_to_prepare_supply": 0,
				"time_left_to_select_supply_variant": 0,
				"total_items_count": 0,
				"total_quantity": 0,
				"vehicle_info": {
				  "driver_name": "string",
				  "driver_phone": "string",
				  "vehicle_model": "string",
				  "vehicle_number": "string"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetSupplyRequestInfoParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().GetSupplyRequestInfo(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &GetSupplyRequestInfoResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestListProductsInSupplyRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListProductsInSupplyRequestParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListProductsInSupplyRequestParams{
				Page:          0,
				PageSize:      0,
				SupplyOrderId: 0,
			},
			`{
				"has_next": true,
				"items": [
				  {
					"icon_path": "string",
					"name": "string",
					"offer_id": "string",
					"quantity": 0,
					"sku": 0
				  }
				],
				"total_items_count": 0
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListProductsInSupplyRequestParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().ListProductsInSupplyRequest(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &ListProductsInSupplyRequestResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetWarehouseWorkload(t *testing.T) {
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
					"schedule": {
					  "capacity": [
						{
						  "start": "2019-08-24T14:15:22Z",
						  "end": "2019-08-24T14:15:22Z",
						  "value": 0
						}
					  ],
					  "date": "2019-08-24T14:15:22Z"
					},
					"warehouse": {
					  "id": "string",
					  "name": "string"
					}
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
		resp, err := c.FBO().GetWarehouseWorkload(ctx)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &GetWarehouseWorkloadResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
