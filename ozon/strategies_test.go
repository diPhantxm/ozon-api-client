package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestListCompetitors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListCompetitorsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListCompetitorsParams{
				Page:  1,
				Limit: 20,
			},
			`{
				"competitor": [
				  {
					"competitor_name": "string",
					"competitor_id": 0
				  }
				],
				"total": 0
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListCompetitorsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().ListCompetitors(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &ListCompetitorsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestListStrategies(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListStrategiesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListStrategiesParams{
				Page:  1,
				Limit: 20,
			},
			`{
				"strategies": [
				  {
					"strategy_id": "string",
					"strategy_name": "string",
					"type": "string",
					"update_type": "string",
					"updated_at": "string",
					"products_count": 0,
					"competitors_count": 0,
					"enabled": true
				  }
				],
				"total": 0
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListStrategiesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().List(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &ListStrategiesResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestCreateStrategy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreateStrategyParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreateStrategyParams{
				StrategyName: "New strategy",
				Competitors: []CreateStrategyCompetitor{
					{
						CompetitorId: 1008426,
						Coefficient:  1,
					},
					{
						CompetitorId: 204,
						Coefficient:  1,
					},
					{
						CompetitorId: 91,
						Coefficient:  1,
					},
					{
						CompetitorId: 48,
						Coefficient:  1,
					},
				},
			},
			`{
				"result": {
				  "id": "4f3a1d4c-5833-4f04-b69b-495cbc1f6f1c"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CreateStrategyParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().Create(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &CreateStrategyResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestInfoStrategy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *InfoStrategyParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&InfoStrategyParams{
				StrategyId: "1",
			},
			`{
				"result": {
				  "strategy_name": "test1",
				  "enabled": true,
				  "update_type": "strategyItemsListChanged",
				  "type": "COMP_PRICE",
				  "competitors": [
					{
					  "competitor_id": 204,
					  "coefficient": 1
					},
					{
					  "competitor_id": 1008426,
					  "coefficient": 1
					}
				  ]
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&InfoStrategyParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().Info(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &InfoStrategyResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestUpdateStrategy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdateStrategyParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdateStrategyParams{
				StrategyId:   "a3de1826-9c54-40f1-bb6d-1a9e2638b058",
				StrategyName: "New Strategy",
				Competitors: []CreateStrategyCompetitor{
					{
						CompetitorId: 1008426,
						Coefficient:  1,
					},
					{
						CompetitorId: 204,
						Coefficient:  1,
					},
					{
						CompetitorId: 91,
						Coefficient:  1,
					},
					{
						CompetitorId: 48,
						Coefficient:  1,
					},
					{
						CompetitorId: 45,
						Coefficient:  1,
					},
				},
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&UpdateStrategyParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().Update(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &UpdateStrategyResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestAddProductsToStrategy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *AddProductsToStrategyParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&AddProductsToStrategyParams{
				ProductId:  []int64{29209},
				StrategyId: "e29114f0-177d-4160-8d06-2bc528470dda",
			},
			`{
				"result": {
				  "failed_product_count": 0
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&AddProductsToStrategyParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().AddProducts(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &AddProductsToStrategyResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetStrategiesByProductIds(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetStrategiesByProductIdsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetStrategiesByProductIdsParams{
				ProductId: []int64{29209},
			},
			`{
				"result": {
				  "products_info": [
					{
					  "product_id": 29209,
					  "strategy_id": "b7cd30e6-5667-424d-b105-fbec30a52477"
					}
				  ]
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetStrategiesByProductIdsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().GetByProductIds(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &GetStrategiesByProductIdsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result.ProductsInfo) != len(test.params.ProductId) {
				t.Errorf("Length of product ids in request and response are not equal")
			}

			if len(resp.Result.ProductsInfo) > 0 {
				if resp.Result.ProductsInfo[0].ProductId != test.params.ProductId[0] {
					t.Errorf("Product ids in request and response are not equal")
				}
			}
		}
	}
}

func TestListProductsInStrategy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListProductsInStrategyParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListProductsInStrategyParams{
				StrategyId: "string",
			},
			`{
				"result": {
				  "product_id": [
					"string"
				  ]
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListProductsInStrategyParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().ListProducts(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &ListProductsInStrategyResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestGetCompetitorPrice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetCompetitorPriceParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetCompetitorPriceParams{
				ProductId: 0,
			},
			`{
				"result": {
				  "strategy_id": "string",
				  "is_enabled": true,
				  "strategy_product_price": 0,
				  "price_downloaded_at": "2022-11-17T15:33:53.936Z",
				  "strategy_competitor_id": 0,
				  "strategy_competitor_product_url": "string"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetCompetitorPriceParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().GetCompetitorPrice(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &GetCompetitorPriceResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestRemoveProductsFromStrategy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *RemoveProductsFromStrategyParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&RemoveProductsFromStrategyParams{
				ProductId: []int64{0},
			},
			`{
				"result": {
				  "failed_product_count": 0
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&RemoveProductsFromStrategyParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().RemoveProducts(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &RemoveProductsFromStrategyResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestChangeStrategyStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ChangeStrategyStatusParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ChangeStrategyStatusParams{
				Enabled:    true,
				StrategyId: "c7516438-7124-4e2c-85d3-ccd92b6b9b65",
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ChangeStrategyStatusParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().ChangeStatus(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &ChangeStatusToResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestRemoveStrategy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *RemoveStrategyParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&RemoveStrategyParams{
				StrategyId: "strategy",
			},
			`{}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&RemoveStrategyParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Strategies().Remove(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		compareJsonResponse(t, test.response, &RemoveStrategyResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
