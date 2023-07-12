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
				Filter: ListOfShipmentCertificatesFilter{
					DateFrom:        "2021-08-04",
					DateTo:          "2022-08-04",
					IntegrationType: "ozon",
					Status:          []ShipmentCertificateFilterStatus{ShitmentCertificateFilterFormed},
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

func TestRescheduleShipmentDeliveryDate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *RescheduleShipmentDeliveryDateParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&RescheduleShipmentDeliveryDateParams{
				PostingNumber: "23281294-0063-2",
				NewTimeslot: RescheduleShipmentDeliveryDateTimeslot{
					From: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2023-03-25T08:51:56.932Z"),
					To:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2023-03-25T08:51:56.932Z"),
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
			&RescheduleShipmentDeliveryDateParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().RescheduleShipmentDeliveryDate(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestDateAvailableForDeliverySchedule(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *DateAvailableForDeliveryScheduleParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&DateAvailableForDeliveryScheduleParams{
				PostingNumber: "23281294-0063-2",
			},
			`{
				"available_change_count": 0,
				"delivery_interval": {
				  "begin": "2019-08-24T14:15:22Z",
				  "end": "2019-08-24T14:15:22Z"
				},
				"remaining_change_count": 0
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&DateAvailableForDeliveryScheduleParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().DateAvailableForDeliverySchedule(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestListManufactoruingCountries(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListManufacturingCountriesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListManufacturingCountriesParams{
				NameSearch: "some name",
			},
			`{
				"result": [
				  {
					"name": "Алжир",
					"country_iso_code": "DZ"
				  },
				  {
					"name": "Ангилья",
					"country_iso_code": "AI"
				  },
				  {
					"name": "Виргинские Острова (Великобритания)",
					"country_iso_code": "VG"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListManufacturingCountriesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().ListManufacturingCountries(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) > 0 {
				if resp.Result[0].Name == "" {
					t.Errorf("Name cannot be empty")
				}
				if resp.Result[0].CountriISOCode == "" {
					t.Errorf("ISO code cannot be empty")
				}
			}
		}
	}
}

func TestSetManufacturingCountry(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *SetManufacturingCountryParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&SetManufacturingCountryParams{
				PostingNumber:  "57195475-0050-3",
				ProductId:      180550365,
				CountryISOCode: "NO",
			},
			`{
				"product_id": 180550365,
				"is_gtd_needed": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&SetManufacturingCountryParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().SetManufacturingCountry(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.ProductId != test.params.ProductId {
				t.Errorf("Product ids in request and response are not equal")
			}
		}
	}
}

func TestPartialPackOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *PartialPackOrderParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&PartialPackOrderParams{
				PostingNumber: "48173252-0034-4",
				Products: []PartialPackOrderProduct{
					{
						ExemplarInfo: []FBSProductExemplar{
							{
								MandatoryMark: "mark",
								GTD:           "gtd",
								IsGTDAbsest:   true,
							},
						},
						ProductId: 247508873,
						Quantity:  1,
					},
				},
			},
			`{
				"result": [
				  "48173252-0034-9"
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&PartialPackOrderParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().PartialPackOrder(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestAvailableFreightsList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *AvailableFreightsListParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&AvailableFreightsListParams{
				DeliveryMethodId: 0,
				DepartureDate:    core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
			},
			`{
				"result": [
				  {
					"carriage_id": 0,
					"carriage_postings_count": 0,
					"carriage_status": "string",
					"cutoff_at": "2019-08-24T14:15:22Z",
					"delivery_method_id": 0,
					"delivery_method_name": "string",
					"errors": [
					  {
						"code": "string",
						"status": "string"
					  }
					],
					"first_mile_type": "string",
					"has_entrusted_acceptance": true,
					"mandatory_postings_count": 0,
					"mandatory_packaged_count": 0,
					"tpl_provider_icon_url": "string",
					"tpl_provider_name": "string",
					"warehouse_city": "string",
					"warehouse_id": 0,
					"warehouse_name": "string",
					"warehouse_timezone": "string"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&AvailableFreightsListParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().AvailableFreightsList(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGenerateAct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GenerateActParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GenerateActParams{
				Id: 123,
			},
			`{
				"id": 421,
				"status": "string"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GenerateActParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().GenerateAct(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetDigitalAct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetDigitalActParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetDigitalActParams{
				Id:      900000250859000,
				DocType: "act_of_acceptance",
			},
			`{
				"content": "string",
				"name": "string",
				"type": "string"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetDigitalActParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().GetDigitalAct(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Content == "" {
				t.Errorf("Content cannot be empty")
			}
		}
	}
}

func TestPackageUnitLabels(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *PackageUnitLabelsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&PackageUnitLabelsParams{
				Id: 295662811,
			},
			`{
				"content": "string",
				"name": "string",
				"type": "string"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&PackageUnitLabelsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().PackageUnitLabel(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Content == "" {
				t.Errorf("Content cannot be empty")
			}
		}
	}
}

func TestOpenDisputeOverShipment(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *OpenDisputeOverShipmentParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&OpenDisputeOverShipmentParams{
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
			&OpenDisputeOverShipmentParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().OpenDisputeOverShipment(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestShipmentCancellationReasons(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ShipmentCancellationReasonsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ShipmentCancellationReasonsParams{
				RelatedPostingNumbers: []string{"73837363-0010-3"},
			},
			`{
				"result": [
				  {
					"posting_number": "73837363-0010-3",
					"reasons": [
					  {
						"id": 352,
						"title": "The goods ran out at the seller's warehouse",
						"type_id": "seller"
					  },
					  {
						"id": 400,
						"title": "Only defective goods remained",
						"type_id": "seller"
					  },
					  {
						"id": 402,
						"title": "Other (seller's fault)",
						"type_id": "seller"
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
			&ShipmentCancellationReasonsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().ShipmentCancellationReasons(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) != len(test.params.RelatedPostingNumbers) {
				t.Errorf("Length of postings numbers in request and response are not equal")
			}
			if len(resp.Result) > 0 {
				if resp.Result[0].PostingNumber != test.params.RelatedPostingNumbers[0] {
					t.Errorf("Posting number in request and response are not equal")
				}
			}
		}
	}
}

func TestShipmentsCancellationReasons(t *testing.T) {
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
					"id": 352,
					"title": "The goods ran out at the seller's warehouse",
					"type_id": "seller",
					"is_available_for_cancellation": true
				  },
				  {
					"id": 400,
					"title": "Only defective goods remained",
					"type_id": "seller",
					"is_available_for_cancellation": true
				  },
				  {
					"id": 401,
					"title": "Seller rejects arbitration",
					"type_id": "seller",
					"is_available_for_cancellation": false
				  },
				  {
					"id": 402,
					"title": "Other (seller's fault)",
					"type_id": "seller",
					"is_available_for_cancellation": true
				  },
				  {
					"id": 665,
					"title": "The buyer did not pick up the order",
					"type_id": "seller",
					"is_available_for_cancellation": false
				  },
				  {
					"id": 666,
					"title": "Return from the delivery service: there is no delivery to the specified region",
					"type_id": "seller",
					"is_available_for_cancellation": false
				  },
				  {
					"id": 667,
					"title": "Order lost by delivery service",
					"type_id": "seller",
					"is_available_for_cancellation": false
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

		resp, err := c.FBS().ShipmentsCancellationReasons()
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
				if resp.Result[0].TypeId == "" {
					t.Errorf("Type id cannot be empty")
				}
			}
		}
	}
}

func TestAddWeightForBulkProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *AddWeightForBulkProductParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&AddWeightForBulkProductParams{
				Items: AddWeightForBulkProductItem{
					SKU:        1231428352,
					WeightReal: []float64{0.3},
				},
				PostingNumber: "33920158-0006-1",
			},
			`{
				"result": "33920158-0006-1"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&AddWeightForBulkProductParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().AddWeightForBulkProduct(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result != test.params.PostingNumber {
				t.Errorf("Posting numbers in request and response are not equal")
			}
		}
	}
}

func TestCancelSending(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CancelSendingParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CancelSendingParams{
				PostingNumber:       "33920113-1231-1",
				CancelReasonId:      352,
				CancelReasonMessage: "Product is out of stock",
				Items: []CancelSendingItem{
					{
						Quantity: 5,
						SKU:      150587396,
					},
				},
			},
			`{
				"result": ""
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CancelSendingParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().CancelSending(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestListShipmentInCertificate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListShipmentInCertificateParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListShipmentInCertificateParams{
				Id: 900000250859000,
			},
			`{
				"result": [
				  {
					"id": 0,
					"multi_box_qty": 0,
					"posting_number": "string",
					"status": "string",
					"seller_error": "string",
					"updated_at": "2019-08-24T14:15:22Z",
					"created_at": "2019-08-24T14:15:22Z",
					"products": [
					  {
						"name": "string",
						"offer_id": "string",
						"price": "string",
						"quantity": 0,
						"sku": 0
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
			&ListShipmentInCertificateParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().ListShipmentInCertificate(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestSpecifyNumberOfBoxes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *SpecifyNumberOfBoxesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&SpecifyNumberOfBoxesParams{
				PostingNumber: "string",
			},
			`{
				"result": {
				  "result": true
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&SpecifyNumberOfBoxesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().SpecifyNumberOfBoxes(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestStatusOfAct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *StatusOfActParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&StatusOfActParams{
				Id: 900000250859000,
			},
			`{
				"result": {
				  "result": true
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&StatusOfActParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().StatusOfAct(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestETGBCustomsDeclarations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ETGBCustomsDeclarationsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ETGBCustomsDeclarationsParams{
				Date: ETGBCustomsDeclarationsDate{
					From: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2023-02-13T12:13:16.818Z"),
					To:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2023-02-13T12:13:16.818Z"),
				},
			},
			`{
				"result": [
				  {
					"posting_number": "string",
					"etgb": {
					  "number": "string",
					  "date": "string",
					  "url": "string"
					}
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ETGBCustomsDeclarationsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().ETGBCustomsDeclarations(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestBarcodeFromProductShipment(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *BarcodeFromProductShipmentParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&BarcodeFromProductShipmentParams{
				Id: 295662811,
			},
			`{
				"content": "https://cdn.ozone.ru/s3/ozon-disk-api/techdoc/seller-api/barcode_1684849346.png",
				"name": "barcode-test",
				"type": "PNG"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&BarcodeFromProductShipmentParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().BarcodeFromProductShipment(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Content == "" {
				t.Errorf("content cannot be empty")
			}
			if resp.Type == "" {
				t.Error("type cannot be empty")
			}
			if resp.Name == "" {
				t.Error("name cannot be empty")
			}
		}
	}
}

func TestBarcodeValueFromProductShipment(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *BarcodeValueFromProductShipmentParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&BarcodeValueFromProductShipmentParams{
				Id: 295662811,
			},
			`{
				"result": "%303%24276481394"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&BarcodeValueFromProductShipmentParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.FBS().BarcodeValueFromProductShipment(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result == "" {
				t.Errorf("result cannot be empty")
			}
		}
	}
}
