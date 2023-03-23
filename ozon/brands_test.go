package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetCertefiedBrandsList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetCertefiedBrandsListParams
		response   string
	}{
		// Test ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetCertefiedBrandsListParams{
				Page:     1,
				PageSize: 100,
			},
			`{
				"result": {
					"brand_certification": [
						{
							"brand_id": 135476863,
							"brand_name": "Sea of Spa",
							"has_certificate": false
						}
					],
					"total": 1
				}
			}`,
		},
		{
			// Test No Client-Id or Api-Key
			http.StatusUnauthorized,
			map[string]string{},
			&GetCertefiedBrandsListParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.Brands().GetCertefiedBrandsList(test.params)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}

}
