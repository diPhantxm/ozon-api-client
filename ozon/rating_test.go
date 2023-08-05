package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetCurrentRatingInfo(t *testing.T) {
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
				"groups": [
				  {
					"group_name": "my-group",
					"items": [
					  {
						"change": {
						  "direction": "DIRECTION_RISE",
						  "meaning": "MEANING_GOOD"
						},
						"current_value": 15.2,
						"name": "my-rating-name",
						"past_value": 15.1,
						"rating": "my-rating",
						"rating_direction": "HIGHER_IS_BETTER",
						"status": "OK",
						"value_type": "INDEX"
					  }
					]
				  }
				],
				"penalty_score_exceeded": true,
				"premium": true
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
		resp, err := c.Rating().GetCurrentSellerRatingInfo(ctx)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Groups) > 0 {
				if len(resp.Groups[0].Items) == 0 {
					t.Errorf("Length of items in a group cannot be 0")
				}
			}
		}
	}
}

func TestGetRatingInfoForPeriod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetSellerRatingInfoForPeriodParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetSellerRatingInfoForPeriodParams{
				DateFrom:          core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
				DateTo:            core.TimeFromString(t, "2006-01-02T15:04:05Z", "2019-08-24T14:15:22Z"),
				Ratings:           []string{"rating_order_accuracy", "rating_shipment_delay"},
				WithPremiumScores: true,
			},
			`{
				"premium_scores": [
				  {
					"rating": "string",
					"scores": [
					  {
						"date": "2019-08-24T14:15:22Z",
						"rating_value": 5,
						"value": 5
					  }
					]
				  }
				],
				"ratings": [
				  {
					"danger_threshold": 1,
					"premium_threshold": 2,
					"rating": "my-rating-system",
					"values": [
					  {
						"date_from": "2019-08-24T14:15:22Z",
						"date_to": "2019-08-24T14:15:22Z",
						"status": {
						  "danger": true,
						  "premium": true,
						  "warning": true
						},
						"value": 12
					  }
					],
					"warning_threshold": 3
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetSellerRatingInfoForPeriodParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Rating().GetSellerRatingInfoForPeriod(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Ratings) > 0 {
				if resp.Ratings[0].Rating == "" {
					t.Errorf("Rating system cannot be empty")
				}
			}
		}
	}
}
