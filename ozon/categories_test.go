package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetProductTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetProductTreeParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetProductTreeParams{
				CategoryId: 17034410,
			},
			`{
				"result": [
				  {
					"category_id": 17034410,
					"title": "Развивающие игрушки",
					"children": []
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetProductTreeParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Categories().Tree(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) > 0 {
				if resp.Result[0].CategoryId != test.params.CategoryId {
					t.Errorf("First category ids in request and response are not equal")
				}
			}
		}
	}
}

func TestGetCategoryAttributes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetCategoryAttributesParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetCategoryAttributesParams{
				CategoryId: []int64{17034410},
			},
			`{
				"result": [
				  {
					"category_id": 17034410,
					"attributes": [
					  {
						"id": 85,
						"name": "Бренд",
						"description": "Укажите наименование бренда, под которым произведен товар. Если товар не имеет бренда, используйте значение \"Нет бренда\"",
						"type": "String",
						"is_collection": false,
						"is_required": true,
						"group_id": 0,
						"group_name": "",
						"dictionary_id": 28732849
					  }
					]
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetCategoryAttributesParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Categories().Attributes(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) != len(test.params.CategoryId) {
				t.Errorf("Length of categories in request and response are not equal")
			}
			if len(resp.Result) > 0 {
				if resp.Result[0].CategoryId != test.params.CategoryId[0] {
					t.Errorf("Category ids in request and response are not equal")
				}
			}
		}
	}
}

func TestGetAttributeDictionary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *GetAttributeDictionaryParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&GetAttributeDictionaryParams{
				AttributeId: 10096,
				CategoryId:  17028968,
				LastValueId: 0,
				Limit:       3,
			},
			`{
				"result": [
				  {
					"id": 61571,
					"value": "белый",
					"info": "",
					"picture": ""
				  },
				  {
					"id": 61572,
					"value": "прозрачный",
					"info": "",
					"picture": ""
				  },
				  {
					"id": 61573,
					"value": "бежевый",
					"info": "",
					"picture": ""
				  }
				],
				"has_next": true
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&GetAttributeDictionaryParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Categories().AttributesDictionary(ctx, test.params)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Result) > int(test.params.Limit) {
				t.Errorf("Length of response result is bigger than limit")
			}
		}
	}
}
