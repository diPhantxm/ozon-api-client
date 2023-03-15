package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetListOfWarehouses(t *testing.T) {
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
					"warehouse_id": 15588127982000,
					"name": "Proffi (Панорама Групп)",
					"is_rfbs": false
				  },
				  {
					"warehouse_id": 22142605386000,
					"name": "Склад на производственной",
					"is_rfbs": true
				  },
				  {
					"warehouse_id": 22208673494000,
					"name": "Тест 37349",
					"is_rfbs": true
				  },
				  {
					"warehouse_id": 22240462819000,
					"name": "Тест12",
					"is_rfbs": true
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

		resp, err := c.GetListOfWarehouses()
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetListOfDeliveryMethods(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetListOfDeliveryMethodsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetListOfDeliveryMethodsParams{
				Filter: GetListOfDeliveryMethodsFilter{
					WarehouseId: 15588127982000,
				},
				Limit:  100,
				Offset: 0,
			},
			`{
				"result": [
				  {
					"id": 15588127982000,
					"company_id": 1,
					"name": "Ozon Логистика курьеру, Есипово",
					"status": "ACTIVE",
					"cutoff": "13:00",
					"provider_id": 24,
					"template_id": 0,
					"warehouse_id": 15588127982000,
					"created_at": "2019-04-04T15:22:31.048202Z",
					"updated_at": "2021-08-15T10:21:44.854209Z"
				  }
				],
				"has_next": false
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetListOfDeliveryMethodsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.GetListOfDeliveryMethods(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
