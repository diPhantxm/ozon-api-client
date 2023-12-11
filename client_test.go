package core

import (
	"context"
	"net/http"
	"testing"
	"time"
)

const (
	testTimeout = 5 * time.Second
)

type TestRequestRequest struct {
	FirstField  string `json:"first_field"`
	SecondField int32  `json:"second_field"`
}

type TestRequestResponse struct {
	FirstField  string `json:"first_header"`
	SecondField struct {
		Key int32 `json:"key"`
	} `json:"second_header"`
	ThirdField []struct {
		FirstElement string `json:"first_element"`
	} `json:"third_header"`
}

func TestRequest(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *TestRequestRequest
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{
				"first_header":  "first-value",
				"second_header": "second-value",
			},
			&TestRequestRequest{
				FirstField:  "test",
				SecondField: 123,
			},
			`{
				"first_field": "first-value",
				"second_field": {
					"key": 12
				},
				"third_field": [
					{
						"first_element": "first-element-value"
					}
				]
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(NewMockHttpHandler(test.statusCode, test.response, test.headers))

		respStruct := &TestRequestResponse{}
		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Request(ctx, http.MethodPost, "/", test.params, respStruct, nil)

		if err != nil {
			t.Error(err)
			continue
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
