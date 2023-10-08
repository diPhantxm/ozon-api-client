package ozon

import (
	"context"
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type Categories struct {
	client *core.Client
}

type GetProductTreeParams struct {
	// Response language
	Language Language `json:"language"`
}

type GetProductTreeResponse struct {
	core.CommonResponse

	// Categories list
	Result []GetProductTreeResult `json:"result"`
}

type GetProductTreeResult struct {
	// Category identifier
	CategoryId int64 `json:"category_id"`

	// Category name
	CategoryName string `json:"category_name"`

	// `true`, if you can't create products in the category. `false`, if you can
	Disabled bool `json:"disabled"`

	// Product type identifier
	TypeId int64 `json:"type_id"`

	// Product type name
	TypeName string `json:"type_name"`

	// Subcategory tree
	Children []GetProductTreeResult `json:"children"`
}

// Returns product categories in the tree view.
//
// New products can be created in the last level categories only.
// This means that you need to match these particular categories with the categories of your site.
// We don't create new categories by user request.
func (c *Categories) Tree(ctx context.Context, params *GetProductTreeParams) (*GetProductTreeResponse, error) {
	url := "/v1/description-category/tree"

	resp := &GetProductTreeResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetCategoryAttributesParams struct {
	// Category identifier
	CategoryId int64 `json:"category_id"`

	// Response language
	Language Language `json:"language"`

	// Product type identifier
	TypeId int64 `json:"type_id"`
}

type GetCategoryAttributesResponse struct {
	core.CommonResponse

	// Method result
	Result []GetCategoryAttributesResult `json:"result"`
}

type GetCategoryAttributesResult struct {
	// Characteristic description
	Description string `json:"description"`

	// Directory identifier
	DictionaryId int64 `json:"dictionary_id"`

	// Characteristics group identifier
	GroupId int64 `json:"group_id"`

	// Characteristics group name
	GroupName string `json:"group_name"`

	// Number of document generation task
	Id int64 `json:"id"`

	// Indicates that the attribute is aspect. An aspect attribute is a characteristic that distinguishes products of the same model.
	//
	// For example, clothes or shoes of the same model may have different colors and sizes. That is, color and size are aspect attributes.
	//
	// Values description:
	//
	// 	- `true`—the attribute is aspect and can't be changed after the products are delivered to the warehouse or sold from the seller's warehouse.
	// 	- `false`—the attribute is not aspect and can be changed at any time
	IsAspect bool `json:"is_aspect"`

	// Indicates that the characteristic is a set of values:
	//
	// 	- `true`—the characteristic is a set of values,
	// 	- `false`—the characteristic consists of a single value
	IsCollection bool `json:"is_collection"`

	// Indicates that the characteristic is mandatory:
	//
	// 	- `true`—a mandatory characteristic,
	// 	- `false`—an optional characteristic
	IsRequired bool `json:"is_required"`

	// Name
	Name string `json:"name"`

	// Characteristic type
	Type string `json:"type"`
}

// Getting characteristics for specified product category and type.
//
// If the dictionary_id value is 0, there is no directory.
// If the value is different, there are directories.
// Get them using the `/v1/description-category/attribute/values` method.
func (c *Categories) Attributes(ctx context.Context, params *GetCategoryAttributesParams) (*GetCategoryAttributesResponse, error) {
	url := "/v1/description-category/attribute"

	resp := &GetCategoryAttributesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetAttributeDictionaryParams struct {
	// Characteristics identifier
	AttributeId int64 `json:"attribute_id"`

	// Category identifier
	CategoryId int64 `json:"category_id"`

	// Response language
	Language Language `json:"language"`

	// Identifier of the directory to start the response with.
	// If `last_value_id` is 10, the response will contain directories starting from the 11th
	LastValueId int64 `json:"last_value_id"`

	// Number of values in the response:
	//
	// 	- maximum—5000,
	// 	- minimum—1.
	Limit int64 `json:"limit"`

	// Product type identifier
	TypeId int64 `json:"type_id"`
}

type GetAttributeDictionaryResponse struct {
	core.CommonResponse

	// Indication that only part of characteristic values was returned in the response:
	//
	// 	- true—make a request with a new last_value_id parameter value for getting the rest of characteristic values;
	// 	- false—all characteristic values were returned
	HasNext bool `json:"has_next"`

	// Characteristic values
	Result []GetAttributeDictionaryResult `json:"result"`
}

type GetAttributeDictionaryResult struct {
	// Characteristic value identifier
	Id int64 `json:"id"`

	// Additional description
	Info string `json:"info"`

	// Image link
	Picture string `json:"picture"`

	// Product characteristic value
	Value string `json:"value"`
}

// Returns characteristics value directory.
//
// To check if an attribute has a nested directory,
// use the `/v1/description-category/attribute` method.
func (c *Categories) AttributesDictionary(ctx context.Context, params *GetAttributeDictionaryParams) (*GetAttributeDictionaryResponse, error) {
	url := "/v1/description-category/attribute"

	resp := &GetAttributeDictionaryResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
