package ozon

import (
	"context"
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().ListUnprocessedShipments(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ListUnprocessedShipmentsResponse{})

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
					Since:            core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-01T00:00:00.000Z"),
					To:               core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-12-01T23:59:59.000Z"),
					Status:           "awaiting_packaging",
					WarehouseId:      []int64{123},
					ProviderId:       []int64{223},
					FBPFilter:        FBPFilterAll,
					DeliveryMethodId: []int64{456},
					OrderId:          0,
				},
				Limit:  0,
				Offset: 0,
				With: GetFBSShipmentsListWith{
					AnalyticsData: true,
					FinancialData: true,
					Translit:      true,
					Barcodes:      true,
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().GetFBSShipmentsList(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetFBSShipmentsListResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().PackOrder(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &PackOrderResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().ValidateLabelingCodes(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ValidateLabelingCodesResponse{})

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
				  "in_process_at": "2021-11-20T09:14:16Z",
				  "shipment_date": "2021-11-23T10:00:00Z",
				  "products": [
					{
					  "price": "279.0000",
					  "offer_id": "250-7898-1",
					  "name": "Кофе ароматизированный \"Шоколадный апельсин\" 250 гр",
					  "sku": 180550365,
					  "quantity": 1,
					  "mandatory_mark": []
					}
				  ],
				  "barcodes": null,
				  "analytics_data": null,
				  "financial_data": null
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().GetShipmentDataByBarcode(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetShipmentDataByBarcodeResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().GetShipmentDataByIdentifier(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetShipmentDataByIdentifierResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().AddTrackingNumbers(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &AddTrackingNumbersResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().ListOfShipmentCertificates(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ListOfShipmentCertificatesResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().SignShipmentCertificate(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &SignShipmentCertificateResponse{})

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

		deliveringctx, _ := context.WithTimeout(context.Background(), testTimeout)
		deliveringResp, err := c.FBS().ChangeStatusToDelivering(deliveringctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ChangeStatusToResponse{})

		lastMilectx, _ := context.WithTimeout(context.Background(), testTimeout)
		lastMileResp, err := c.FBS().ChangeStatusToLastMile(lastMilectx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		deliveredctx, _ := context.WithTimeout(context.Background(), testTimeout)
		deliveredResp, err := c.FBS().ChangeStatusToDelivered(deliveredctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		sendBySellerctx, _ := context.WithTimeout(context.Background(), testTimeout)
		sendBySellerResp, err := c.FBS().ChangeStatusToSendBySeller(sendBySellerctx, test.params)
		if err != nil {
			t.Error(err)
			continue
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().PassShipmentToShipping(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &PassShipmentToShippingResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().CancelShipment(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CancelShipmentResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().CreateAct(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CreateActResponse{})

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
				  "status": "completed",
				  "file_url": "https://cdn1.ozone.ru/s3/sc-temporary/e6/0c/e60cdfd7aed78c2b44d134504fbd591d.pdf"
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().GetLabeling(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetLabelingResponse{})

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
				"content_type": "application/pdf",
				"file_name": "ticket-170660-2023-07-13T13:17:06Z.pdf",
				"file_content": "%PDF-1.7\n%âãÏÓ\n53 0 obj\n<</MarkInfo<</Marked true/Type/MarkInfo>>/Pages 9 0 R/StructTreeRoot 10 0 R/Type/Catalog>>\nendobj\n8 0 obj\n<</Filter/FlateDecode/Length 2888>>\nstream\nxå[[ݶ\u0011~?¿BÏ\u0005Bs\u001c^\u0000Àwí5ú\u0010 m\u0016Èsà¦)\n;hÒ\u0014èÏïG\u0014)<{äµ] ]?¬¬oIÎ}¤F±óϤñï\u001bÕü×X­´OÏï?^~¹$<ø¨È9q\u0013Y\u0012åñì§_¼|ÿégü\t+\u0012\u001bxª}Æxҿ¿¼_º¼xg¦þ5OkuÌ3ýíògüûå\"Ni\u0016C\u0001°\u000fA9g'r¢\"\u0013YóĪ\u0018NÑ{\u001dÕóZ¬\\Ô\""
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().PrintLabeling(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &PrintLabelingResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Content == "" {
				t.Error("content cannot be empty")
			}
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().CreateTaskForGeneratingLabel(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CreateTaskForGeneratingLabelResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().GetDropOffPointRestrictions(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetDropOffPointRestrictionsResponse{})

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
				MultiBoxQuantity: 0,
				PostingNumber:    "1234",
				Products: []CheckProductItemsDataProduct{
					{
						Exemplars: []CheckProductItemsDataProductExemplar{
							{
								ExemplarId:    1,
								GTD:           "string",
								IsGTDAbsent:   true,
								IsRNTPAbsent:  true,
								MandatoryMark: "string",
								RNTP:          "string",
								JWUIN:         "string",
							},
						},
						IsGTDNeeded:           true,
						IsMandatoryMarkNeeded: true,
						IsRNTPNeeded:          true,
						ProductId:             22,
						Quantity:              11,
					},
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().CheckProductItemsData(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CheckProductItemsDataResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().GetProductItemsCheckStatuses(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetProductItemsCheckStatusesResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().RescheduleShipmentDeliveryDate(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &RescheduleShipmentDeliveryDateResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().DateAvailableForDeliverySchedule(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &DateAvailableForDeliveryScheduleResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().ListManufacturingCountries(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ListManufacturingCountriesResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().SetManufacturingCountry(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &SetManufacturingCountryResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().PartialPackOrder(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &PartialPackOrderResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().AvailableFreightsList(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &AvailableFreightsListResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().GenerateAct(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GenerateActResponse{})

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
				"content_type": "application/pdf",
				"file_name": "20816409_act_of_mismatch.pdf",
				"file_content": "%PDF-1.4\n%ÓôÌá\n1 0 obj\n<<\n/Creator(Chromium)\n/Producer(PDFsharp 1.50.5147 \\([www.pdfsharp.com|http://www.pdfsharp.com/]\\) \\(Original: Skia/PDF m103\\))\n/CreationDate(D:20230625092529+00'00')\n/ModDate(D:20230625092529+00'00')\n>>\nendobj\n2 0 obj\n<<\n/Type/Page\n/Resources\n<<\n/ProcSet[/PDF/Text/ImageB/ImageC/ImageI]\n/ExtGState\n<<\n/G3 3 0 R\n/G8 8 0 R\n>>\n/XObject\n<<\n/X6 6 0 R\n/X7 7 0 R\n>>\n/Font\n<<\n/F4 4 0 R\n/F5 5 0 R\n>>\n>>\n/MediaBox[0 0 594.96 841.92]\n/Contents 9 0 R\n/StructParents 0\n/Parent 13 0 R\n/Group\n<<\n/CS/DeviceRGB\n/S/Transparency\n>>\n>>\nendobj\n3 0 obj\n<<\n/ca 1\n/BM/Normal\n>>\nendobj\n4 0 obj\n<<\n/Type/Font\n/Subtype/Type0\n/BaseFont/AAAAAA+LiberationSans\n/Encoding/Identity-H\n/DescendantFonts[160 0 R]\n/ToUnicode 161 0 R\n>>\nendobj\n5 0 obj\n<<\n/Type/Font\n/Subtype/Type0\n/BaseFont/BAAAAA+LiberationSans-Bold\n/Encoding/Identity-H\n/DescendantFonts[164 0"
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().GetDigitalAct(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetDigitalActResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().PackageUnitLabel(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &PackageUnitLabelsResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().OpenDisputeOverShipment(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &OpenDisputeOverShipmentResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().ShipmentCancellationReasons(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ShipmentCancellationReasonsResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().ShipmentsCancellationReasons(ctx)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ShipmentsCancellationReasonsResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().AddWeightForBulkProduct(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &AddWeightForBulkProductResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().CancelSending(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CancelSendingResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().ListShipmentInCertificate(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ListShipmentInCertificateResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().SpecifyNumberOfBoxes(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &SpecifyNumberOfBoxesResponse{})

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
				  "added_to_act": [
					"true"
				  ],
				  "removed_from_act": [
					"false"
				  ],
				  "status": "ready"
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().StatusOfAct(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &StatusOfActResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().ETGBCustomsDeclarations(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ETGBCustomsDeclarationsResponse{})

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
				"content_type": "image/png",
				"file_name": "20913984_barcode.png",
				"file_content": "PNG\r\n\u001a\n\u0000\u0000\u0000\rIHDR\u0000\u0000\u0003\u0010\u0000\u0000\u0000\u0010\u0000\u0000\u0000\u0000íZ\u000e'\u0000\u0000\u0002pIDATxìÕÁJ\u00031\u0014@Q+þÿ/×E\u0017\u000e¼\u0010u¡-ç¬$£Éˌp?î÷·§t» }ýü¸Ãcåz¹2wOWû\\Ϛ뫧×Ùö;ì|rÇýßîç¼úî{§¬N?í7oìv¸®µ¹Ãùû¹¾ÿÏ9ÿî?a¸ºéê7O&߿É9çÉ\u000eÏáý¯\u0007\u0000à\u0012\b\u0000@\u0000\u0004\u0002$\u0010\u0000$\u0000 \t\u0004\u0000I \u0000H\u0002\u0001@\u0012\b\u0000@\u0000\u0004\u0002$\u0010\u0000$\u0000 \t\u0004\u0000I \u0000H\u0002\u0001@\u0012\b\u0000@\u0000\u0004\u0002$\u0010\u0000$\u0000 \t\u0004\u0000I \u0000H\u0002\u0001@\u0012\b\u0000@\u0000\u0004\u0002$\u0010\u0000$\u0000 \t\u0004\u0000I \u0000H\u0002\u0001@\u0012\b\u0000@\u0000\u0004\u0002$\u0010\u0000"
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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().BarcodeFromProductShipment(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &BarcodeFromProductShipmentResponse{})

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

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().BarcodeValueFromProductShipment(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &BarcodeValueFromProductShipmentResponse{})

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

func TestGetActPDF(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetActPDFParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetActPDFParams{
				Id: 22435521842000,
			},
			`{
				"content_type": "application/pdf",
				"file_name": "20928233.pdf",
				"file_content": "binarystring"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetActPDFParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.FBS().GetActPDF(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &GetActPDFResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.FileContent == "" {
				t.Errorf("result cannot be empty")
			}
		}
	}
}
