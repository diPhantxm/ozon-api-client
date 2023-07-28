package notifications

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type testData struct {
	raw    string
	object interface{}
}

func pingTest(t *testing.T) testData {
	return testData{
		object: &pingRequest{
			Common: Common{MessageType: PingType},
			Time:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
		},
		raw: `{
			"message_type": "TYPE_PING",
			"time": "2019-08-24T14:15:22Z"
		}`,
	}
}

func newPostingTest(t *testing.T) testData {
	return testData{
		object: &NewPosting{
			Common:        Common{MessageType: NewPostingType},
			PostingNumber: "24219509-0020-1",
			Products: []Product{
				{
					SKU:      147451959,
					Quantity: 2,
				},
			},
			InProccessAt: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-01-26T06:56:36.294Z"),
			WarehouseId:  18850503335000,
			SellerId:     15,
		},
		raw: `{
			"message_type": "TYPE_NEW_POSTING",
			"posting_number": "24219509-0020-1",
			"products": [
			  {
				"sku": 147451959,
				"quantity": 2
			  }
			],
			"in_process_at": "2021-01-26T06:56:36.294Z",
			"warehouse_id": 18850503335000,
			"seller_id": 15
		}`,
	}
}

func postingCancelledTest(t *testing.T) testData {
	return testData{
		object: &PostingCancelled{
			Common:        Common{MessageType: PostingCancelledType},
			PostingNumber: "24219509-0020-1",
			Products: []Product{
				{
					SKU:      147451959,
					Quantity: 1,
				},
			},
			OldState:         "posting_transferred_to_courier_service",
			NewState:         "posting_canceled",
			ChangedStateDate: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-01-26T06:56:36.294Z"),
			Reason:           Reason{Id: 1, Message: "message"},
			WarehouseId:      1,
			SellerId:         15,
		},
		raw: `{
			"message_type": "TYPE_POSTING_CANCELLED",
			"posting_number": "24219509-0020-1",
			"products": [
			  {
				"sku": 147451959,
				"quantity": 1
			  }
			],
			"old_state": "posting_transferred_to_courier_service",
			"new_state": "posting_canceled",
			"changed_state_date": "2021-01-26T06:56:36.294Z",
			"reason": {
			  "id": 1,
			  "message": "message"
			},
			"warehouse_id": 1,
			"seller_id": 15
		}`,
	}
}

func cutoffDateChangedTest(t *testing.T) testData {
	return testData{
		object: &CutoffDateChanged{
			Common:        Common{MessageType: CutoffDateChangedType},
			PostingNumber: "24219509-0020-2",
			NewCutoffDate: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-24T07:00:00Z"),
			OldCutoffDate: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-21T10:00:00Z"),
			WarehouseId:   1,
			SellerId:      15,
		},
		raw: `{
			"message_type": "TYPE_CUTOFF_DATE_CHANGED",
			"posting_number": "24219509-0020-2",
			"new_cutoff_date": "2021-11-24T07:00:00Z",
			"old_cutoff_date": "2021-11-21T10:00:00Z",
			"warehouse_id": 1,
			"seller_id": 15
		}`,
	}
}

func deliveryDateChangedTest(t *testing.T) testData {
	return testData{
		object: &DeliveryDateChanged{
			Common:               Common{MessageType: DeliveryDateChangedType},
			PostingNumber:        "24219509-0020-2",
			NewDeliveryDateBegin: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-24T07:00:00Z"),
			NewDeliveryDateEnd:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-24T16:00:00Z"),
			OldDeliveryDateBegin: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-21T10:00:00Z"),
			OldDeliveryDateEnd:   core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-11-21T19:00:00Z"),
			WarehouseId:          1,
			SellerId:             15,
		},
		raw: `{
			"message_type": "TYPE_DELIVERY_DATE_CHANGED",
			"posting_number": "24219509-0020-2",
			"new_delivery_date_begin": "2021-11-24T07:00:00Z",
			"new_delivery_date_end": "2021-11-24T16:00:00Z",
			"old_delivery_date_begin": "2021-11-21T10:00:00Z",
			"old_delivery_date_end": "2021-11-21T19:00:00Z",
			"warehouse_id": 1,
			"seller_id": 15
		}`,
	}
}

func priceIndexChangedTest(t *testing.T) testData {
	return testData{
		object: &PriceIndexChanged{
			Common:     Common{MessageType: PriceIndexChangedType},
			UpdatedAt:  core.TimeFromString(t, "2006-01-02T15:04:05Z", "2022-06-21T05:52:46.648533678Z"),
			SKU:        147451959,
			ProductId:  1234,
			PriceIndex: 5678,
			SellerId:   15,
		},
		raw: `{
			"seller_id": 15,
			"message_type": "TYPE_PRICE_INDEX_CHANGED",
			"updated_at":"2022-06-21T05:52:46.648533678Z",
			"sku": 147451959,
			"product_id": 1234,
			"price_index": 5678
		}`,
	}
}

func stocksChangedTest(t *testing.T) testData {
	return testData{
		object: &StocksChanged{
			Common: Common{MessageType: StocksChangedType},
			Items: []Item{
				{
					UpdatedAt: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-09-01T14:15:22Z"),
					SKU:       5678,
					ProductId: 1234,
					Stocks: []Stock{
						{
							WarehouseId: 10,
							Present:     50,
							Reserved:    5,
						},
					},
				},
			},
			SellerId: 15,
		},
		raw: `{
			"message_type": "TYPE_STOCKS_CHANGED",
			"seller_id": 15,
			"items": [
			  {
				"product_id": 1234,
				"sku": 5678,
				"updated_at": "2021-09-01T14:15:22Z",
				"stocks": [
				  {
					"warehouse_id": 10,
					"present": 50,
					"reserved": 5
				  }
				]
			  }
			]
		}`,
	}
}

func newMessageTest(t *testing.T) testData {
	return testData{
		object: &NewMessage{
			Common:    Common{MessageType: NewMessageType},
			ChatId:    "b646d975-0c9c-4872-9f41-8b1e57181063",
			ChatType:  "Buyer_Seller",
			MessageId: "3000000000817031942",
			CreatedAt: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2022-07-18T20:58:04.528Z"),
			User:      User{Id: "115568", Type: "Customer"},
			Data:      []string{"Message text"},
			SellerId:  7,
		},
		raw: `{  
			"message_type": "TYPE_NEW_MESSAGE",
			"chat_id": "b646d975-0c9c-4872-9f41-8b1e57181063",
			"chat_type": "Buyer_Seller",
			"message_id": "3000000000817031942",
			"created_at": "2022-07-18T20:58:04.528Z",
			"user": {
				"id": "115568",
				"type": "Customer"
			},
			"data": [
				"Message text"
			],  
			"seller_id": 7
		}`,
	}
}

func updateMessageTest(t *testing.T) testData {
	return testData{
		object: &UpdateMessage{
			NewMessage: NewMessage{
				Common:    Common{MessageType: UpdateMessageType},
				ChatId:    "b646d975-0c9c-4872-9f41-8b1e57181063",
				ChatType:  "Buyer_Seller",
				MessageId: "3000000000817031942",
				CreatedAt: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2022-07-18T20:58:04.528Z"),
				User: User{
					Id:   "115568",
					Type: "Сustomer",
				},
				Data:     []string{"Message text"},
				SellerId: 7,
			},
			UpdatedAt: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2022-07-18T20:59:04.528Z"),
		},
		raw: `{  
			"message_type": "TYPE_UPDATE_MESSAGE",
			"chat_id": "b646d975-0c9c-4872-9f41-8b1e57181063",
			"chat_type": "Buyer_Seller",
			"message_id": "3000000000817031942",
			"created_at": "2022-07-18T20:58:04.528Z",
			"updated_at": "2022-07-18T20:59:04.528Z",
			"user": {
				"id": "115568",
				"type": "Сustomer"
			},
			"data": [
				"Message text"
			], 
			"seller_id": 7
		}`,
	}
}

func createUpdateItemTest(t *testing.T) testData {
	return testData{
		object: &CreateOrUpdateItem{
			Common:    Common{MessageType: "TYPE_CREATE_OR_UPDATE_ITEM"},
			OfferId:   "1234",
			ProductId: 5678,
			IsError:   false,
			ChangedAt: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2022-09-01T14:15:22Z"),
			SellerId:  15,
		},
		raw: `{
			"message_type": "TYPE_CREATE_OR_UPDATE_ITEM",
			"seller_id": 15,
			"offer_id": "1234",
			"product_id": 5678,
			"is_error": false,
			"changed_at": "2022-09-01T14:15:22Z"
		}`,
	}
}

func stateChangedTest(t *testing.T) testData {
	return testData{
		object: &StateChanged{
			Common:           Common{MessageType: "TYPE_STATE_CHANGED"},
			PostingNumber:    "24219509-0020-2",
			NewState:         "posting_delivered",
			ChangedStateDate: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2021-02-02T15:07:46.765Z"),
			WarehouseId:      1,
			SellerId:         15,
		},
		raw: `{
			"message_type": "TYPE_STATE_CHANGED",
			"posting_number": "24219509-0020-2",
			"new_state": "posting_delivered",
			"changed_state_date": "2021-02-02T15:07:46.765Z",
			"warehouse_id": 1,
			"seller_id": 15
		}`,
	}
}

func messageReadTest(t *testing.T) testData {
	return testData{
		object: &MessageRead{
			LastReadMessageId: "3000000000817031942",
			NewMessage: NewMessage{
				Common:    Common{MessageType: "TYPE_MESSAGE_READ"},
				ChatId:    "b646d975-0c9c-4872-9f41-8b1e57181063",
				ChatType:  "Buyer_Seller",
				MessageId: "3000000000817031942",
				CreatedAt: core.TimeFromString(t, "2006-01-02T15:04:05Z", "2022-07-18T20:58:04.528Z"),
				User: User{
					Id:   "115568",
					Type: "Сustomer",
				},
				SellerId: 7,
			},
		},
		raw: `{  
			"message_type": "TYPE_MESSAGE_READ",
			"chat_id": "b646d975-0c9c-4872-9f41-8b1e57181063",
			"chat_type": "Buyer_Seller",
			"message_id": "3000000000817031942",
			"created_at": "2022-07-18T20:58:04.528Z",    
			"user": {
				"id": "115568",
				"type": "Сustomer"
			},
			"last_read_message_id": "3000000000817031942",
			"seller_id": 7
		}`,
	}
}

func chatClosedTest(t *testing.T) testData {
	return testData{
		object: &ChatClosed{
			Common:   Common{MessageType: ChatClosedType},
			ChatId:   "b646d975-0c9c-4872-9f41-8b1e57181063",
			ChatType: "Buyer_Seller",
			User:     User{Id: "115568", Type: "Customer"},
			SellerId: 7,
		},
		raw: `{  
			"message_type": "TYPE_CHAT_CLOSED",
			"chat_id": "b646d975-0c9c-4872-9f41-8b1e57181063",
			"chat_type": "Buyer_Seller",
			"user": {
				"id": "115568",
				"type": "Customer"
			},
			"seller_id": 7
		}`,
	}
}

func TestNotificationServer(t *testing.T) {
	testCases := []struct {
		request  testData
		response string
	}{
		{
			pingTest(t),
			`{
				"version": "1.0",
				"name": "Ozon Seller API"
			}`,
		},
		{
			newPostingTest(t),
			`{
				"result": true
			}`,
		},
		{
			postingCancelledTest(t),
			`{
				"result": true
			}`,
		},
		{
			stateChangedTest(t),
			`{
				"result": true
			}`,
		},
		{
			cutoffDateChangedTest(t),
			`{
				"result": true
			}`,
		},
		{
			deliveryDateChangedTest(t),
			`{
				"result": true
			}`,
		},
		{
			createUpdateItemTest(t),
			`{
				"result": true
			}`,
		},
		{
			priceIndexChangedTest(t),
			`{
				"result": true
			}`,
		},
		{
			stocksChangedTest(t),
			`{
				"result": true
			}`,
		},
		{
			newMessageTest(t),
			`{
				"result": true
			}`,
		},
		{
			updateMessageTest(t),
			`{
				"result": true
			}`,
		},
		{
			messageReadTest(t),
			`{
				"result": true
			}`,
		},
		{
			chatClosedTest(t),
			`{
				"result": true
			}`,
		},
	}

	port := getFreePort()

	client := http.Client{}
	server := NewNotificationServer(port)
	server.Register(NewPostingType, comparatorWith(newPostingTest(t).object))
	server.Register(PostingCancelledType, comparatorWith(postingCancelledTest(t).object))
	server.Register(StateChangedType, comparatorWith(stateChangedTest(t).object))
	server.Register(CutoffDateChangedType, comparatorWith(cutoffDateChangedTest(t).object))
	server.Register(DeliveryDateChangedType, comparatorWith(deliveryDateChangedTest(t).object))
	server.Register(CreateOrUpdateType, comparatorWith(createUpdateItemTest(t).object))
	server.Register(PriceIndexChangedType, comparatorWith(priceIndexChangedTest(t).object))
	server.Register(StocksChangedType, comparatorWith(stocksChangedTest(t).object))
	server.Register(NewMessageType, comparatorWith(newMessageTest(t).object))
	server.Register(UpdateMessageType, comparatorWith(updateMessageTest(t).object))
	server.Register(MessageReadType, comparatorWith(messageReadTest(t).object))
	server.Register(ChatClosedType, comparatorWith(chatClosedTest(t).object))
	go func() {
		if err := server.Run(); err != nil {
			t.Fatalf("notification server is down: %s", err)
		}
	}()

	// TODO: get rid of it
	// Needed to make sure server is running
	time.Sleep(3 * time.Second)

	for _, testCase := range testCases {
		httpResp, err := client.Post(fmt.Sprintf("http://0.0.0.0:%d/", port), "application/json", strings.NewReader(testCase.request.raw))
		if err != nil {
			t.Error(err)
			continue
		}

		gotJson, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			t.Error(err)
			continue
		}

		expected := map[string]interface{}{}
		got := map[string]interface{}{}
		err = json.Unmarshal(gotJson, &got)
		if err != nil {
			t.Error(err)
		}
		err = json.Unmarshal([]byte(testCase.response), &expected)
		if err != nil {
			t.Error(err)
		}

		if err := compare(expected, got); err != nil {
			t.Error(err)
			continue
		}
	}
}

func compare(expected map[string]interface{}, got map[string]interface{}) error {
	for k, v := range expected {
		if gotValue, ok := got[k]; !ok {
			return fmt.Errorf("key %s is expected to present", k)
		} else if !reflect.DeepEqual(gotValue, v) {
			return fmt.Errorf("key %s is not equal, got: %v, want: %v", k, gotValue, v)
		}
	}
	return nil
}

func getFreePort() int {
	listener, _ := net.Listen("tcp", ":0")
	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port
}

func comparatorWith(v1 interface{}) func(v2 interface{}) error {
	return func(v2 interface{}) error {
		if !reflect.DeepEqual(v1, v2) {
			return fmt.Errorf("objects are not equal:\n got: %#v,\n want: %#v", v2, v1)
		}
		return nil
	}
}
