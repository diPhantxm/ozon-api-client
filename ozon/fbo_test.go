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
					Since:  core.NewTimeFormat(core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-09-01T00:00:00.000Z"), "2006-01-02T15:04:05Z"),
					Status: "awaiting_packaging",
					To:     core.NewTimeFormat(core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-17T10:44:12.828Z"), "2006-01-02T15:04:05Z"),
				},
				Limit:    5,
				Offset:   0,
				Translit: true,
				With: &GetFBOShipmentsListWith{
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
					"created_at": "2021-09-01T00:23:45.607Z",
					"in_process_at": "2021-09-01T00:25:30.12Z",
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
			continue
		}

		compareJsonResponse(t, test.response, &GetFBOShipmentsListResponse{})

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
				With: &GetShipmentDetailsWith{
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
			continue
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
				Filter: &ListSupplyRequestsFilter{
					States: []string{"ORDER_STATE_DATA_FILLING"},
				},
				Paging: &ListSupplyRequestsPaging{
					FromOrderId: 0,
					Limit:       0,
				},
			},
			`{
				"last_supply_order_id": 0,
				"supply_order_id": [
				  "string"
				]
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
			continue
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
				OrderIds: []string{"string"},
			},
			`{
				"orders": [
				  {
					"creation_date": "string",
					"creation_flow": "string",
					"data_filling_deadline_utc": "2019-08-24T14:15:22Z",
					"dropoff_warehouse_id": 0,
					"state": "ORDER_STATE_UNSPECIFIED",
					"supplies": [
					  {
						"bundle_id": "string",
						"storage_warehouse_id": 0,
						"supply_id": 0
					  }
					],
					"supply_order_id": 0,
					"supply_order_number": "string",
					"timeslot": [
					  {
						"can_not_set_reasons": [
						  "string"
						],
						"can_set": true,
						"is_required": true,
						"value": {
						  "timeslot": [
							{
							  "from": "2019-08-24T14:15:22Z",
							  "to": "2019-08-24T14:15:22Z"
							}
						  ],
						  "timezone_info": [
							{
							  "iana_name": "string",
							  "offset": "string"
							}
						  ]
						}
					  }
					],
					"vehicle": [
					  {
						"can_not_set_reasons": [
						  "string"
						],
						"can_set": true,
						"is_required": true,
						"value": [
						  {
							"driver_name": "string",
							"driver_phone": "string",
							"vehicle_model": "string",
							"vehicle_number": "string"
						  }
						]
					  }
					]
				  }
				],
				"warehouses": [
				  {
					"address": "string",
					"name": "string",
					"warehouse_id": 0
				  }
				]
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
			continue
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
			continue
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
			continue
		}

		compareJsonResponse(t, test.response, &GetWarehouseWorkloadResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetSupplyOrdersByStatus(t *testing.T) {
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
				"items": [
				  {
					"count": 0,
					"order_state": "ORDER_STATE_UNSPECIFIED"
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
		resp, err := c.FBO().GetSupplyOrdersByStatus(ctx)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetSupplyOrdersByStatusResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetSupplyTimeslots(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetSupplyTimeslotsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetSupplyTimeslotsParams{
				SupplyOrderId: 0,
			},
			`{
				"timeslots": [
				  {
					"from": "2019-08-24T14:15:22Z",
					"to": "2019-08-24T14:15:22Z"
				  }
				],
				"timezone": [
				  {
					"iana_name": "string",
					"offset": "string"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetSupplyTimeslotsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().GetSupplyTimeslots(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetSupplyTimeslotsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestUpdateSupplyTimeslot(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdateSupplyTimeslotParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdateSupplyTimeslotParams{
				SupplyOrderId: 0,
				Timeslot: SupplyTimeslotValueTimeslot{
					From: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
					To:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
				},
			},
			`{
				"errors": [
				  "UPDATE_TIMESLOT_ERROR_UNSPECIFIED"
				],
				"operation_id": "string"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&UpdateSupplyTimeslotParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().UpdateSupplyTimeslot(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &UpdateSupplyTimeslotResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetSupplyTimeslotStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetSupplyTimeslotStatusParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetSupplyTimeslotStatusParams{
				OperationId: "string",
			},
			`{
				"errors": [
				  "UPDATE_TIMESLOT_ERROR_UNSPECIFIED"
				],
				"status": "STATUS_UNSPECIFIED"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetSupplyTimeslotStatusParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().GetSupplyTimeslotStatus(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetSupplyTimeslotStatusResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestCreatePass(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreatePassParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreatePassParams{
				SupplyOrderId: 123,
				Vehicle: GetSupplyRequestInfoVehicle{
					DriverName:    "string",
					DriverPhone:   "string",
					VehicleModel:  "string",
					VehicleNumber: "string",
				},
			},
			`{
				"error_reasons": [
				  "SET_VEHICLE_ERROR_UNSPECIFIED"
				],
				"operation_id": "string"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CreatePassParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().CreatePass(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CreatePassResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetPass(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetPassParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetPassParams{
				OperationId: "string",
			},
			`{
				"errors": [
				  "SET_VEHICLE_ERROR_UNSPECIFIED"
				],
				"result": "Unknown"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetPassParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().GetPass(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetPassResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetSupplyContent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetSupplyContentParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetSupplyContentParams{
				BundleIds: []string{"string"},
				IsAsc:     true,
				Limit:     0,
				Query:     "string",
				SortField: "UNSPECIFIED",
			},
			`{
				"items": [
				  {
					"icon_path": "string",
					"sku": 0,
					"name": "string",
					"quantity": 0,
					"barcode": "string",
					"product_id": 0,
					"quant": 0,
					"is_quant_editable": true,
					"volume_in_litres": 0,
					"total_volume_in_litres": 0,
					"contractor_item_code": "string",
					"sfbo_attribute": "ITEM_SFBO_ATTRIBUTE_UNSPECIFIED",
					"shipment_type": "BUNDLE_ITEM_SHIPMENT_TYPE_UNSPECIFIED"
				  }
				],
				"total_count": 0,
				"has_next": true,
				"last_id": "string"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetSupplyContentParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().GetSupplyContent(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetSupplyContentResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestCreateSupplyDraft(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreateSupplyDraftParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreateSupplyDraftParams{
				ClusterIds:         []string{"string"},
				DropoffWarehouseId: 0,
				Items: []CreateSupplyDraftItem{
					{
						Quantity: 1,
						SKU:      11,
					},
				},
				Type: "CREATE_TYPE_CROSSDOCK",
			},
			`{
				"operation_id": "string"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CreateSupplyDraftParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().CreateSupplyDraft(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CreateSupplyDraftResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetSupplyDraftInfo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetSupplyDraftInfoParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetSupplyDraftInfoParams{
				OperationId: "string",
			},
			`{
				"clusters": [
				  {
					"cluster_id": 0,
					"cluster_name": "string",
					"warehouses": [
					  {
						"address": "string",
						"bundle_ids": [
						  {
							"bundle_id": "string",
							"is_docless": true
						  }
						],
						"name": "string",
						"restricted_bundle_id": "string",
						"status": {
						  "invalid_reason": "WAREHOUSE_SCORING_INVALID_REASON_UNSPECIFIED",
						  "is_available": true,
						  "state": "WAREHOUSE_SCORING_STATUS_FULL_AVAILABLE"
						},
						"supply_warehouse": {
						  "address": "string",
						  "name": "string",
						  "warehouse_id": 0
						},
						"total_rank": 0,
						"total_score": 0,
						"travel_time_days": 0,
						"warehouse_id": 0
					  }
					]
				  }
				],
				"draft_id": 0,
				"errors": [
				  {
					"error_message": "string",
					"items_validation": [
					  {
						"reasons": [
						  "string"
						],
						"sku": 0
					  }
					],
					"unknown_cluster_ids": [
					  "string"
					]
				  }
				],
				"status": "CALCULATION_STATUS_FAILED"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetSupplyDraftInfoParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().GetSupplyDraftInfo(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetSupplyDraftInfoResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestCreateSupplyFromDraft(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreateSupplyFromDraftParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreateSupplyFromDraftParams{
				DraftId: 11,
				Timeslot: CreateSupplyFromDraftTimeslot{
					FromInTimezone: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
					ToInTimezone:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
				},
				WarehouseId: 45,
			},
			`{
				"operation_id": "string"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CreateSupplyFromDraftParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBO().CreateSupplyFromDraft(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CreateSupplyFromDraftResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
