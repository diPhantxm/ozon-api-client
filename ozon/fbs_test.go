package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestListUnprocessedShipments(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListUnprocessedShipmentsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListUnprocessedShipmentsParams{
				Direction: "ASC",
				Filter: ListUnprocessedShipmentsFilter{
					CutoffFrom: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-08-24T14:15:22Z"),
					CutoffTo:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-08-31T14:15:22Z"),
					Status:     "awaiting_packaging",
				},
				Limit: 100,
				With: ListUnprocessedShipmentsWith{
					AnalyticsData: true,
					Barcodes:      true,
					FinancialData: true,
					Translit:      true,
				},
			},
			`{
				"result": {
				  "postings": [
					{
					  "posting_number": "23713478-0018-3",
					  "order_id": 559293114,
					  "order_number": "33713378-0051",
					  "status": "awaiting_packaging",
					  "delivery_method": {
						"id": 15110442724000,
						"name": "Ozon Логистика курьеру, Москва",
						"warehouse_id": 15110442724000,
						"warehouse": "Склад на Ленина",
						"tpl_provider_id": 24,
						"tpl_provider": "Ozon Логистика"
					  },
					  "tracking_number": "",
					  "tpl_integration_type": "ozon",
					  "in_process_at": "2021-08-25T10:48:38Z",
					  "shipment_date": "2021-08-26T10:00:00Z",
					  "delivering_date": null,
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
						  "currency_code": "RUB",
						  "price": "1259",
						  "offer_id": "УТ-0001365",
						  "name": "Мяч, цвет: черный, 5 кг",
						  "sku": 140048123,
						  "quantity": 1,
						  "mandatory_mark": []
						}
					  ],
					  "addressee": null,
					  "barcodes": {
						"upper_barcode": "%101%806044518",
						"lower_barcode": "23024930500000"
					  },
					  "analytics_data": {
						"region": "Санкт-Петербург",
						"city": "Санкт-Петербург",
						"delivery_type": "PVZ",
						"is_premium": false,
						"payment_type_group_name": "Карты оплаты",
						"warehouse_id": 15110442724000,
						"warehouse": "Склад на Ленина",
						"tpl_provider_id": 24,
						"tpl_provider": "Ozon Логистика",
						"delivery_date_begin": "2022-08-28T14:00:00Z",
						"delivery_date_end": "2022-08-28T18:00:00Z",
						"is_legal": false
					  },
					  "financial_data": {
						"products": [
						  {
							"commission_amount": 0,
							"commission_percent": 0,
							"payout": 0,
							"product_id": 140048123,
							"old_price": 1888,
							"price": 1259,
							"total_discount_value": 629,
							"total_discount_percent": 33.32,
							"actions": [
							  "Системная виртуальная скидка селлера"
							],
							"picking": null,
							"quantity": 1,
							"client_price": "",
							"item_services": {
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
					  "is_express": false,
					  "requirements": {
						"products_requiring_gtd": [],
						"products_requiring_country": []
					  }
					}
				  ],
				  "count": 55
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListUnprocessedShipmentsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().ListUnprocessedShipments(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetFBSShipmentsList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetFBSShipmentsListParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetFBSShipmentsListParams{
				Direction: "ASC",
				Filter: GetFBSShipmentsListFilter{
					Since:  core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-01T00:00:00.000Z"),
					To:     core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-12-01T23:59:59.000Z"),
					Status: "awaiting_packaging",
				},
				Limit:  100,
				Offset: 0,
				With: GetFBSShipmentsListWith{
					AnalyticsData: true,
					FinancialData: true,
					Translit:      true,
				},
			},
			`{
				"result": {
				  "postings": [
					{
					  "posting_number": "05708065-0029-1",
					  "order_id": 680420041,
					  "order_number": "05708065-0029",
					  "status": "awaiting_deliver",
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
						  "currency_code": "RUB",
						  "price": "1390.000000",
						  "offer_id": "205953",
						  "name": " Электронный конструктор PinLab Позитроник",
						  "sku": 358924380,
						  "quantity": 1,
						  "mandatory_mark": []
						}
					  ],
					  "addressee": null,
					  "barcodes": null,
					  "analytics_data": null,
					  "financial_data": null,
					  "is_express": false,
					  "requirements": {
						"products_requiring_gtd": [],
						"products_requiring_country": [],
						"products_requiring_mandatory_mark": []
					  }
					}
				  ],
				  "has_next": true
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetFBSShipmentsListParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().GetFBSShipmentsList(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestPackOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *PackOrderParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&PackOrderParams{
				Packages: []PackOrderPackage{
					{
						Products: []PackOrderPackageProduct{
							{
								ProductId: 185479045,
								Quantity:  1,
							},
						},
					},
				},
				PostingNumber: "89491381-0072-1",
				With: PackOrderWith{
					AdditionalData: true,
				},
			},
			`{
				"result": [
				  "89491381-0072-1"
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&PackOrderParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().PackOrder(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if len(resp.Result) != len(test.params.Packages) {
			t.Errorf("Length of packages in request and response are not equal")
		}
	}
}

func TestValidateLabelingCodes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ValidateLabelingCodesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ValidateLabelingCodesParams{
				PostingNumber: "23281294-0063-2",
				Products: []ValidateLabelingCodesProduct{
					{
						Exemplars: []ValidateLabelingCodesExemplar{
							{
								GTD:           "",
								MandatoryMark: "010290000151642731tVMohkbfFgunB",
							},
						},
						ProductId: 476925391,
					},
				},
			},
			`{
				"result": {
				  "products": [
					{
					  "product_id": 476925391,
					  "exemplars": [
						{
						  "mandatory_mark": "010290000151642731tVMohkbfFgunB",
						  "gtd": "",
						  "valid": true,
						  "errors": []
						}
					  ],
					  "valid": true,
					  "error": ""
					}
				  ]
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ValidateLabelingCodesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().ValidateLabelingCodes(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result.Products) != len(test.params.Products) {
				t.Errorf("Length of products in request and response are not equal")
			}
			if len(resp.Result.Products) > 0 {
				if resp.Result.Products[0].ProductId != test.params.Products[0].ProductId {
					t.Errorf("Product ids in request and response are not equal")
				}
			}
		}
	}
}

func TestGetShipmentDataByBarcode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetShipmentDataByBarcodeParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetShipmentDataByBarcodeParams{
				Barcode: "20325804886000",
			},
			`{
				"result": {
				  "posting_number": "57195475-0050-3",
				  "order_id": 438764970,
				  "order_number": "57195475-0050",
				  "status": "awaiting_packaging",
				  "delivery_method": {
					"id": 18114520187000,
					"name": "Ozon Логистика самостоятельно, Москва",
					"warehouse_id": 18114520187000,
					"warehouse": "Москва основной",
					"tpl_provider_id": 24,
					"tpl_provider": "Ozon Логистика"
				  },
				  "tracking_number": "",
				  "tpl_integration_type": "ozon",
				  "in_process_at": "2021-11-20T09:14:16Z",
				  "shipment_date": "2021-11-23T10:00:00Z",
				  "delivering_date": null,
				  "provider_status": "",
				  "delivery_price": "",
				  "cancellation": {
					"cancel_reason_id": 0,
					"cancel_reason": "",
					"cancellation_type": "",
					"cancelled_after_ship": false,
					"affect_cancellation_rating": false,
					"cancellation_initiator": ""
				  },
				  "customer": null,
				  "addressee": null,
				  "products": [
					{
					  "price": "279.0000",
					  "offer_id": "250-7898-1",
					  "name": "Кофе ароматизированный \"Шоколадный апельсин\" 250 гр",
					  "sku": 180550365,
					  "quantity": 1,
					  "mandatory_mark": [],
					  "dimensions": {
						"height": "40.00",
						"length": "240.00",
						"weight": "260",
						"width": "140.00"
					  }
					}
				  ],
				  "barcodes": null,
				  "analytics_data": null,
				  "financial_data": null,
				  "additional_data": [],
				  "is_express": false,
				  "requirements": {
					"products_requiring_gtd": [],
					"products_requiring_country": []
				  },
				  "product_exemplars": null
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetShipmentDataByBarcodeParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().GetShipmentDataByBarcode(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.OrderId == 0 {
				t.Errorf("Order id cannot be 0")
			}
			if resp.Result.Status == "" {
				t.Errorf("Status cannot be empty")
			}
		}
	}
}
