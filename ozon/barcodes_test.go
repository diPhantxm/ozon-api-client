package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGenerateBarcodes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GenerateBarcodesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GenerateBarcodesParams{
				ProductIds: []int64{123456789},
			},
			`{
				"errors": [
				  {
					"code": "code 200",
					"error": "no error",
					"barcode": "456",
					"product_id": 123456789
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GenerateBarcodesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Barcodes().Generate(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Errors) != 0 {
				if resp.Errors[0].ProductId != test.params.ProductIds[0] {
					t.Errorf("Product ids are not equal")
				}
			}
		}
	}
}

func TestBindBarcodes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *BindBarcodesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&BindBarcodesParams{
				Barcodes: []BindBarcode{
					{
						Barcode: "some barcode",
						SKU:     123456789,
					},
				},
			},
			`{
				"errors": [
				  {
					"code": "code 200",
					"error": "no error",
					"barcode": "some barcode",
					"sku": 123456789
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&BindBarcodesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Barcodes().Bind(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Errors) != 0 {
				if resp.Errors[0].Barcode != test.params.Barcodes[0].Barcode {
					t.Errorf("Barcodes are not equal")
				}
				if resp.Errors[0].SKU != test.params.Barcodes[0].SKU {
					t.Errorf("Barcodes are not equal")
				}
			}
		}
	}
}
