package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestCreateUpdateProformaLink(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreateUpdateProformaLinkParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreateUpdateProformaLinkParams{
				PostingNumber: "posting number",
				URL:           "link",
			},
			`{
				"result": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CreateUpdateProformaLinkParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Invoices().CreateUpdate(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetProformaLink(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProformaLinkParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProformaLinkParams{
				PostingNumber: "posting number",
			},
			`{
				"result": {
				  "file_url": "string"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetProformaLinkParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Invoices().Get(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestDeleteProformaLink(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *DeleteProformaLinkParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&DeleteProformaLinkParams{
				PostingNumber: "posting number",
			},
			`{
				"result": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&DeleteProformaLinkParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Invoices().Delete(test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
