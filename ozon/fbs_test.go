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

func TestGetShipmentDataByIdentifier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetShipmentDataByIdentifierParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetShipmentDataByIdentifierParams{
				PostingNumber: "57195475-0050-3",
				With:          GetShipmentDataByIdentifierWith{},
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
					  "currency_code": "RUB",
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
			&GetShipmentDataByIdentifierParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().GetShipmentDataByIdentifier(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.PostingNumber != test.params.PostingNumber {
				t.Errorf("Posting numbers in request and response are not equal")
			}
			if resp.Result.OrderId == 0 {
				t.Errorf("Order id cannot be 0")
			}
			if resp.Result.Status == "" {
				t.Errorf("Status cannot be empty")
			}
			if resp.Result.TPLIntegrationType == "" {
				t.Errorf("TPL integration type cannot be empty")
			}
		}
	}
}

func TestAddTrackingNumbers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *AddTrackingNumbersParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&AddTrackingNumbersParams{
				TrackingNumbers: []FBSTrackingNumbersParams{
					{
						PostingNumber:  "48173252-0033-2",
						TrackingNumber: "123123123",
					},
				},
			},
			`{
				"result": [
				  {
					"error": [],
					"posting_number": "48173252-0033-2",
					"result": true
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&AddTrackingNumbersParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().AddTrackingNumbers(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) != len(test.params.TrackingNumbers) {
				t.Errorf("Length of tracking numbers in reqeust and response are not equal")
			}
			if len(resp.Result) > 0 {
				if resp.Result[0].PostingNumber != test.params.TrackingNumbers[0].PostingNumber {
					t.Errorf("Posting numbers in request and response are not equal")
				}
			}
		}
	}
}

func TestListOfShipmentCertificates(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListOfShipmentCertificatesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListOfShipmentCertificatesParams{
				Limit: 100,
				Filter: ListOfShipmentCertificates{
					DateFrom:        "2021-08-04",
					DateTo:          "2022-08-04",
					IntegrationType: "ozon",
					Status:          []string{"delivered"},
				},
			},
			`{
				"result": [
				  {
					"id": 1234,
					"delivery_method_id": 1234,
					"delivery_method_name": "string",
					"integration_type": "string",
					"containers_count": 2,
					"status": "string",
					"departure_date": "string",
					"created_at": "2019-08-24T14:15:22Z",
					"updated_at": "2019-08-24T14:15:22Z",
					"act_type": "string",
					"is_partial": true,
					"has_postings_for_next_carriage": true,
					"partial_num": 0,
					"related_docs": {
					  "act_of_acceptance": {
						"created_at": "2019-08-24T14:15:22Z",
						"document_status": "string"
					  },
					  "act_of_mismatch": {
						"created_at": "2019-08-24T14:15:22Z",
						"document_status": "string"
					  },
					  "act_of_excess": {
						"created_at": "2019-08-24T14:15:22Z",
						"document_status": "string"
					  }
					}
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListOfShipmentCertificatesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().ListOfShipmentCertificates(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) > 0 {
				if resp.Result[0].Id == 0 {
					t.Errorf("Id cannot be 0")
				}
				if resp.Result[0].Status == "" {
					t.Errorf("Status cannot be empty")
				}
				if resp.Result[0].DeliveryMethodId == 0 {
					t.Errorf("Delivery method id cannot be 0")
				}
			}
		}
	}
}

func TestSignShipmentCertificate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *SignShipmentCertificateParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&SignShipmentCertificateParams{
				Id:      900000250859000,
				DocType: "act_of_mismatch",
			},
			`{
				"result": "string"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&SignShipmentCertificateParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().SignShipmentCertificate(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
		if resp.StatusCode == http.StatusOK {
			if resp.Result == "" {
				t.Errorf("Result cannot be empty")
			}
		}
	}
}

func TestChangeStatusTo(t *testing.T) {
	t.Parallel()

	type test struct {
		statusCode int
		headers    map[string]string
		params     *ChangeStatusToParams
		response   string
	}

	assertResponse := func(t *testing.T, test *test, resp *ChangeStatusToResponse) {
		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) != len(test.params.PostingNumber) {
				t.Errorf("Length of posting numbers in request and response are not equal")
			}
			if len(resp.Result) > 0 {
				if resp.Result[0].PostingNumber != test.params.PostingNumber[0] {
					t.Errorf("Posting numbers in request and response are not equal")
				}
			}
		}
	}

	tests := []test{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ChangeStatusToParams{
				PostingNumber: []string{"48173252-0033-2"},
			},
			`{
				"result": [
				  {
					"error": [],
					"posting_number": "48173252-0033-2",
					"result": true
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ChangeStatusToParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		deliveringResp, err := c.FBS().ChangeStatusToDelivering(test.params)
		if err != nil {
			t.Error(err)
		}

		lastMileResp, err := c.FBS().ChangeStatusToLastMile(test.params)
		if err != nil {
			t.Error(err)
		}

		deliveredResp, err := c.FBS().ChangeStatusToDelivered(test.params)
		if err != nil {
			t.Error(err)
		}

		sendBySellerResp, err := c.FBS().ChangeStatusToSendBySeller(test.params)
		if err != nil {
			t.Error(err)
		}

		assertResponse(t, &test, deliveringResp)
		assertResponse(t, &test, lastMileResp)
		assertResponse(t, &test, deliveredResp)
		assertResponse(t, &test, sendBySellerResp)
	}
}

func TestPassShipmentToShipping(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *PassShipmentToShippingParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&PassShipmentToShippingParams{
				PostingNumber: []string{"33920143-1195-1"},
			},
			`{
				"result": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&PassShipmentToShippingParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().PassShipmentToShipping(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestCancelShipment(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CancelShipmentParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CancelShipmentParams{
				CancelReasonId:      352,
				CancelReasonMessage: "Product is out of stock",
				PostingNumber:       "33920113-1231-1",
			},
			`{
				"result": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CancelShipmentParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().CancelShipment(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestCreateAct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreateActParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreateActParams{
				ContainersCount:  1,
				DeliveryMethodId: 230039077005,
				DepartureDate:    core.TimeFromString(t, "2006-01-02T15:04:05Z", "2022-06-10T11:42:06.444Z"),
			},
			`{
				"result": {
				  "id": 5819327210249
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CreateActParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().CreateAct(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetLabeling(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetLabelingParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetLabelingParams{
				TaskId: 158,
			},
			`{
				"result": {
					"error": "24",
					"file_url": "some-url",
					"status": "completed"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetLabelingParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().GetLabeling(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
		if resp.StatusCode == http.StatusOK {
			if resp.Result.Status == "" {
				t.Errorf("Status cannot be empty")
			}
		}
	}
}

func TestPrintLabeling(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *PrintLabelingParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&PrintLabelingParams{
				PostingNumber: []string{"48173252-0034-4"},
			},
			`{
				"content": "https://cdn1.ozone.ru/s3/ozon-disk-api/c4a11c8b748033daf6cdd44aca7ed4c492e55d6f4810f13feae4792afa7934191647255705"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&PrintLabelingParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().PrintLabeling(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestCreateTaskForGeneratingLabel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreateTaskForGeneratingLabelParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreateTaskForGeneratingLabelParams{
				PostingNumber: []string{"48173252-0034-4"},
			},
			`{
				"result": {
				  "task_id": 5819327210249
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CreateTaskForGeneratingLabelParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().CreateTaskForGeneratingLabel(test.params)
		if err != nil {
			t.Error(err)
		}

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

func TestGetDropOffPointRestrictions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetDropOffPointRestrictionsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetDropOffPointRestrictionsParams{
				PostingNumber: "48173252-0034-4",
			},
			`{
				"result": {
				  "posting_number": "48173252-0034-4",
				  "max_posting_weight": 0,
				  "min_posting_weight": 0,
				  "width": 0,
				  "length": 0,
				  "height": 0,
				  "max_posting_price": 0,
				  "min_posting_price": 0
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetDropOffPointRestrictionsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().GetDropOffPointRestrictions(test.params)
		if err != nil {
			t.Error(err)
		}

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

func TestCheckProductItemsData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CheckProductItemsDataParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CheckProductItemsDataParams{
				PostingNumber: "48173252-0034-4",
				Products: CheckProductItemsDataProduct{
					Exemplars: []FBSProductExemplar{
						{
							IsGTDAbsest:   true,
							MandatoryMark: "010290000151642731tVMohkbfFgunB",
						},
					},
					ProductId: 476925391,
				},
			},
			`{
				"result": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CheckProductItemsDataParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().CheckproductItemsData(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetProductItemsCheckStatuses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProductItemsCheckStatusesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductItemsCheckStatusesParams{
				PostingNumber: "23281294-0063-2",
			},
			`{
				"posting_number": "23281294-0063-2",
				"products": [
				  {
					"product_id": 476925391,
					"exemplars": [
					  {
						"mandatory_mark": "010290000151642731tVMohkbfFgunB",
						"gtd": "",
						"is_gtd_absent": true,
						"mandatory_mark_check_status": "passed",
						"mandatory_mark_error_codes": [],
						"gtd_check_status": "passed",
						"gtd_error_codes": []
					  }
					]
				  }
				],
				"status": "ship_available"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetProductItemsCheckStatusesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().GetProductItemsCheckStatuses(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
		if resp.StatusCode == http.StatusOK {
			if resp.PostingNumber != test.params.PostingNumber {
				t.Errorf("Posting numbers in request and response are not equal")
			}
			if resp.Status == "" {
				t.Errorf("Status cannot be empty")
			}
			if len(resp.Products) > 0 {
				if resp.Products[0].ProductId == 0 {
					t.Errorf("Product id cannot be 0")
				}
				if len(resp.Products[0].Exemplars) > 0 {
					if resp.Products[0].Exemplars[0].MandatoryMark == "" {
						t.Errorf("Mandatory mark cannot be empty")
					}
				}
			}
		}
	}
}
