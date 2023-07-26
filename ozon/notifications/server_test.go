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
)

func TestNotificationServer(t *testing.T) {
	testCases := []struct {
		request  string
		response string
	}{
		// PING
		{
			`{
				"message_type": "TYPE_PING",
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
				"result": true
			}`,
		},
	}

	port := getFreePort()

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

	// TODO: get rid of it
	// Needed to make sure server is running
	time.Sleep(3 * time.Second)

	for _, testCase := range testCases {
		httpResp, err := client.Post(fmt.Sprintf("http://0.0.0.0:%d/", port), "application/json", strings.NewReader(testCase.request))
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
