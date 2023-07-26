package notifications

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	port = 5000
)

func TestNotificationServer(t *testing.T) {
	testCases := []struct {
		request  string
		response string
	}{
		// PING
		{
			`{
				"message_type": "string",
				"time": "2019-08-24T14:15:22Z"
			}`,
			`{
				"version": "1.0",
				"name": "Ozon Seller API"
			}`,
		},
		// REGISTERED HANDLER
		{
			`{  
				"message_type": "TYPE_CHAT_CLOSED",
				"chat_id": "b646d975-0c9c-4872-9f41-8b1e57181063",
				"chat_type": "Buyer_Seller",
				"user": {
					"id": "115568",
					"type": "Сustomer"
				},
				"seller_id": "7"
			}`,
			`{
				"result": true
			}`,
		},
		// UNREGISTERED HANDLER
		{
			`{  
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
				"seller_id": "7"
			}`,
			`{
				"result": false
			}`,
		},
	}

	client := http.Client{}
	server := NewNotificationServer(port)
	server.Register(ChatClosedType, func(req interface{}) error {
		_, ok := req.(*ChatClosed)
		if !ok {
			return fmt.Errorf("req is not of ChatClosed type")
		}
		return nil
	})
	go func() {
		if err := server.Run(); err != nil {
			t.Fatalf("notification server is down: %s", err)
		}
	}()

	for _, testCase := range testCases {
		httpResp, err := client.Post(fmt.Sprintf("http://127.0.0.1:%d/", port), "application/json", strings.NewReader(testCase.request))
		if err != nil {
			t.Error(err)
			continue
		}

		gotJson, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			t.Error(err)
			continue
		}

		expected := map[string]string{}
		got := map[string]string{}
		json.Unmarshal(gotJson, got)
		json.Unmarshal([]byte(testCase.response), expected)

		if err := compare(expected, got); err != nil {
			t.Error(err)
			continue
		}
	}
}

func compare(expected map[string]string, got map[string]string) error {
	for k, v := range expected {
		if gotValue, ok := got[k]; !ok {
			return fmt.Errorf("key %s is expected to present", k)
		} else if gotValue != v {
			return fmt.Errorf("key %s is not equal, got: %s, want: %s", k, gotValue, v)
		}
	}
	return nil
}
