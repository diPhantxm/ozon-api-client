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
	// Category identifier
	CategoryId int64 `json:"category_id"`

	// Response language
	Language Language `json:"language" default:"DEFAULT"`
}

type GetProductTreeResponse struct {
	core.CommonResponse

	// Category list
	Result []GetProductTreeResult `json:"result"`
}

type GetProductTreeResult struct {
	// Category identifier
	CategoryId int64 `json:"category_id"`

	// Subcategory tree
	Children []GetProductTreeResponse `json:"children"`

	// Category name
	Title string `json:"title"`
}

// Returns product categories in the tree view.
// New products can be created in the last level categories only.
// This means that you need to match these particular categories with the categories of your site.
// It is not possible to create categories by user request
func (c Categories) Tree(ctx context.Context, params *GetProductTreeParams) (*GetProductTreeResponse, error) {
	url := "/v2/category/tree"

	resp := &GetProductTreeResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetCategoryAttributesParams struct {
	// Filter by characteristics
	AttributeType AttributeType `json:"attribute_type" default:"ALL"`

	// Category identifier
	CategoryId []int64 `json:"category_id"`

	// Response language
	Language Language `json:"language" default:"DEFAULT"`
}

type GetCategoryAttributesResponse struct {
	core.CommonResponse

	// Method result
	Result []GetCategoryAttributesResult `json:"result"`
}

type GetCategoryAttributesResult struct {
	// Array of product characteristics
	Attributes []GetCategoryAttributesResultAttribute `json:"attributes"`

	// Category identifier
	CategoryId int64 `json:"category_id"`
}

type GetCategoryAttributesResultAttribute struct {
	// Indication that the dictionary attribute values depend on the category:
	//   - true — the attribute has its own set of values for each category.
	//   - false — the attribute has the same set of values for all categories
	CategoryDependent bool `json:"category_dependent"`

	// Characteristic description
	Description string `json:"description"`

	// Directory identifier
	DictionaryId int64 `json:"dictionary_id"`

	// Characteristics group identifier
	GroupId int64 `json:"group_id"`

	// Characteristics group name
	GroupName string `json:"group_name"`

	// Document generation task number
	Id int64 `json:"id"`

	// Indicates that the attribute is aspect. An aspect attribute is a characteristic that distinguishes products of the same model.
	//
	// For example, clothes and shoes of the same model may have different colors and sizes. That is, color and size are aspect attributes.
	//
	// Values description:
	//   - true — the attribute is aspect and cannot be changed after the products are delivered to the warehouse or sold from the seller's warehouse.
	//   - false — the attribute is not aspect and can be changed at any time
	IsAspect bool `json:"is_aspect"`

	// Indicates that the characteristic is a set of values:
	//   - true — the characteristic is a set of values,
	//   - false — the characteristic consists of a single value
	IsCollection bool `json:"is_collection"`

	// Indicates that the characteristic is mandatory:
	//   - true — a mandatory characteristic,
	//   - false — you can leave the characteristic out
	IsRequired bool `json:"is_required"`

	// Name
	Name string `json:"name"`

	// Characteristic type
	Type string `json:"type"`
}

// Getting characteristics for specified product category.
//
// Pass up to 20 category identifiers in the `category_id` list.
//
// You can check whether the attribute has a nested directory by the `dictionary_id` parameter.
// The 0 value means there is no directory. If the value is different, then there are directories.
// You can get them using the `/v2/category/attribute/values` method
func (c Categories) Attributes(ctx context.Context, params *GetCategoryAttributesParams) (*GetCategoryAttributesResponse, error) {
	url := "/v3/category/attribute"

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
	// The default language is Russian
	Language Language `json:"language" default:"DEFAULT"`

	LastValueId int64 `json:"last_value_id"`

	// Number of values in the response:
	//   - maximum — 5000
	//   - minimum — 1
	Limit int64 `json:"limit"`
}

type GetAttributeDictionaryResponse struct {
	core.CommonResponse

	HasNext bool `json:"has_next"`

	// Method result
	Result []GetAttributeDictionaryResult `json:"result"`
}

type GetAttributeDictionaryResult struct {
	Id      int64  `json:"id"`
	Info    string `json:"info"`
	Picture string `json:"picture"`

	// Product characteristic value
	Value string `json:"value"`
}

// You can use the `/v3/category/attribute` method to check if an attribute has a nested directory.
// If there are directories, get them using this method
func (c Categories) AttributesDictionary(ctx context.Context, params *GetAttributeDictionaryParams) (*GetAttributeDictionaryResponse, error) {
	url := "/v2/category/attribute/values"

	resp := &GetAttributeDictionaryResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
