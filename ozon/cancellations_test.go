package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetCancellationInfo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetCancellationInfoParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetCancellationInfoParams{
				CancellationId: 90066344,
			},
			`{
				"result": {
				  "cancellation_id": 90066344,
				  "posting_number": "47134289-0029-1",
				  "cancellation_reason": {
					"id": 508,
					"name": "Покупатель отменил заказ"
				  },
				  "cancelled_at": "2022-04-07T06:37:26.871105Z",
				  "cancellation_reason_message": "Изменение пункта выдачи заказа.",
				  "tpl_integration_type": "ThirdPartyTracking",
				  "state": {
					"id": 2,
					"name": "Подтверждена",
					"state": "APPROVED"
				  },
				  "cancellation_initiator": "CLIENT",
				  "order_date": "2022-04-06T17:17:24.517Z",
				  "approve_comment": "",
				  "approve_date": "2022-04-07T07:52:45.971824Z",
				  "auto_approve_date": "2022-04-09T06:37:26.871105Z"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetCancellationInfoParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Cancellations().GetInfo(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.CancellationId != test.params.CancellationId {
				t.Errorf("Cancellation ids in request and response are not equal")
			}
		}
	}
}

func TestListCancellations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListCancellationsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListCancellationsParams{
				Filter: ListCancellationsFilter{
					CancellationInitiator: []string{"CLIENT"},
					State:                 "ALL",
				},
				Limit:  2,
				Offset: 0,
				With: ListCancellationWith{
					Counters: true,
				},
			},
			`{
				"result": [
				  {
					"cancellation_id": 50186754,
					"posting_number": "41267064-0032-1",
					"cancellation_reason": {
					  "id": 508,
					  "name": "Покупатель отменил заказ"
					},
					"cancelled_at": "2021-09-03T07:17:12.116114Z",
					"cancellation_reason_message": "",
					"tpl_integration_type": "ThirdPartyTracking",
					"state": {
					  "id": 2,
					  "name": "Подтверждена",
					  "state": "APPROVED"
					},
					"cancellation_initiator": "CLIENT",
					"order_date": "2021-09-03T07:04:53.220Z",
					"approve_comment": "",
					"approve_date": "2021-09-03T09:13:12.614200Z",
					"auto_approve_date": "2021-09-06T07:17:12.116114Z"
				  },
				  {
					"cancellation_id": 51956491,
					"posting_number": "14094410-0018-1",
					"cancellation_reason": {
					  "id": 507,
					  "name": "Покупатель передумал"
					},
					"cancelled_at": "2021-09-13T15:03:25.155827Z",
					"cancellation_reason_message": "",
					"tpl_integration_type": "ThirdPartyTracking",
					"state": {
					  "id": 5,
					  "name": "Автоматически отменена",
					  "state": "REJECTED"
					},
					"cancellation_initiator": "CLIENT",
					"order_date": "2021-09-13T07:48:50.143Z",
					"approve_comment": "",
					"approve_date": null,
					"auto_approve_date": "2021-09-16T15:03:25.155827Z"
				  }
				],
				"total": 19,
				"counters": {
				  "on_approval": 0,
				  "approved": 14,
				  "rejected": 5
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListCancellationsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Cancellations().List(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestApproveCancellations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ApproveRejectCancellationsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ApproveRejectCancellationsParams{
				CancellationId: 74393917,
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ApproveRejectCancellationsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Cancellations().Approve(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestRejectCancellations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ApproveRejectCancellationsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ApproveRejectCancellationsParams{
				CancellationId: 74393917,
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ApproveRejectCancellationsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Cancellations().Reject(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
