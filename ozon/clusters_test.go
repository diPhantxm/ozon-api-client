package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestListClusters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListClustersParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListClustersParams{
				ClusterIds:  []string{"string"},
				ClusterType: "CLUSTER_TYPE_UNKNOWN",
			},
			`{
				"clusters": [
				  {
					"id": 0,
					"logistic_clusters": [
					  {
						"is_archived": true,
						"warehouses": [
						  {
							"name": "string",
							"type": "FULL_FILLMENT",
							"warehouse_id": 0
						  }
						]
					  }
					],
					"name": "string",
					"type": "CLUSTER_TYPE_UNKNOWN"
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListClustersParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Clusters().List(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ListClustersResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
