package ozon

import (
	"context"
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type Strategies struct {
	client *core.Client
}

type ListCompetitorsParams struct {
	// Page number from which you want to download the list of competitors.
	// The minimum value is 1
	Page int64 `json:"page"`

	// Maximum number of competitors on the page. Allowed values: 1–50
	Limit int64 `json:"limit"`
}

type ListCompetitorsResponse struct {
	core.CommonResponse

	// List of competitors
	Competitor []ListCompetitorsCompetitor `json:"competitor"`

	// Total number of competitors
	Total int32 `json:"total"`
}

type ListCompetitorsCompetitor struct {
	// Competitor's name
	Name string `json:"competitor_name"`

	// Competitor identifier
	Id int64 `json:"competitor_id"`
}

// Method for getting a list of competitors—sellers with similar products in other online stores and marketplaces
func (c Strategies) ListCompetitors(ctx context.Context, params *ListCompetitorsParams) (*ListCompetitorsResponse, error) {
	url := "/v1/pricing-strategy/competitors/list"

	resp := &ListCompetitorsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListStrategiesParams struct {
	// Page number from which you want to download the list of competitors.
	// The minimum value is 1
	Page int64 `json:"page"`

	// Maximum number of competitors on the page. Allowed values: 1–50
	Limit int64 `json:"limit"`
}

type ListStrategiesResponse struct {
	core.CommonResponse

	// List of strategies
	Strategies []ListStrategiesStrategy `json:"strategies"`

	// Total number of strategies
	Total int32 `json:"total"`
}

type ListStrategiesStrategy struct {
	// Strategy identifier
	Id string `json:"strategy_id"`

	// Strategy name
	Name string `json:"strategy_name"`

	// Strategy type
	Type StrategyType `json:"type"`

	// Type of the last strategy change
	UpdateType StrategyUpdateType `json:"update_type"`

	// Date of last change
	UpdatedAt string `json:"updated_at"`

	// Number of products in the strategy
	ProductsCount int64 `json:"products_count"`

	// Number of selected competitors
	CompetitorsCount int64 `json:"competitors_count"`

	// Strategy status
	Enabled bool `json:"enabled"`
}

func (c Strategies) List(ctx context.Context, params *ListStrategiesParams) (*ListStrategiesResponse, error) {
	url := "/v1/pricing-strategy/list"

	resp := &ListStrategiesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CreateStrategyParams struct {
	// List of competitors
	Competitors []CreateStrategyCompetitor `json:"competitors"`

	// Strategy name
	StrategyName string `json:"strategy_name"`
}

type CreateStrategyCompetitor struct {
	// Coefficient by which the minimum price among competitors will be multiplied.
	// The allowed range is from 0.5 to 1.2
	Coefficient float32 `json:"coefficient"`

	// Competitor identifier
	CompetitorId int64 `json:"competitor_id"`
}

type CreateStrategyResponse struct {
	core.CommonResponse

	// Method result
	Result CreateStrategyResult `json:"result"`
}

type CreateStrategyResult struct {
	// Strategy identifier
	StrategyId string `json:"strategy_id"`
}

func (c Strategies) Create(ctx context.Context, params *CreateStrategyParams) (*CreateStrategyResponse, error) {
	url := "/v1/pricing-strategy/create"

	resp := &CreateStrategyResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type InfoStrategyParams struct {
	// Strategy identifier
	StrategyId string `json:"strategy_id"`
}

type InfoStrategyResponse struct {
	core.CommonResponse

	// Method result
	Result InfoStrategyResult `json:"result"`
}

type InfoStrategyResult struct {
	// List of competitors
	Competitors []CreateStrategyCompetitor `json:"competitors"`

	// Strategy status
	Enabled bool `json:"enabled"`

	// Strategy name
	Name string `json:"name"`

	// Strategy type
	Type StrategyType `json:"type"`

	// Type of the last strategy change
	UpdateType StrategyUpdateType `json:"update_type"`
}

func (c Strategies) Info(ctx context.Context, params *InfoStrategyParams) (*InfoStrategyResponse, error) {
	url := "/v1/pricing-strategy/info"

	resp := &InfoStrategyResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateStrategyParams struct {
	// List of competitors
	Competitors []CreateStrategyCompetitor `json:"competitors"`

	// Product identifier
	StrategyId string `json:"strategy_id"`

	// Strategy name
	StrategyName string `json:"strategy_name"`
}

type UpdateStrategyResponse struct {
	core.CommonResponse
}

func (c Strategies) Update(ctx context.Context, params *UpdateStrategyParams) (*UpdateStrategyResponse, error) {
	url := "/v1/pricing-strategy/update"

	resp := &UpdateStrategyResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type AddProductsToStrategyParams struct {
	// List of product identifiers. The maximum number is 50
	ProductId []int64 `json:"product_id"`

	// Product identifier
	StrategyId string `json:"strategy_id"`
}

type AddProductsToStrategyResponse struct {
	core.CommonResponse

	// Method result
	Result AddProductsToStrategyResult `json:"result"`
}

type AddProductsToStrategyResult struct {
	// Products with errors
	Errors []AddProductsToStrategyResultError `json:"errors"`

	// Number of products with errors
	FailedProductCount int32 `json:"failed_product_count"`
}

type AddProductsToStrategyResultError struct {
	// Error code
	Code string `json:"code"`

	// Error message
	Error string `json:"error"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

func (c Strategies) AddProducts(ctx context.Context, params *AddProductsToStrategyParams) (*AddProductsToStrategyResponse, error) {
	url := "/v1/pricing-strategy/products/add"

	resp := &AddProductsToStrategyResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetStrategiesByProductIdsParams struct {
	// List of product identifiers. The maximum number is 50
	ProductId []int64 `json:"product_id"`
}

type GetStrategiesByProductIdsResponse struct {
	core.CommonResponse

	// Method result
	Result GetStrategiesByProductIdsResult `json:"result"`
}

type GetStrategiesByProductIdsResult struct {
	// Product information
	ProductsInfo []GetStrategiesByProductIdsResultProductInfo `json:"products_info"`
}

type GetStrategiesByProductIdsResultProductInfo struct {
	// Product identifier
	ProductId int64 `json:"product_id"`

	// Strategy identifier to which the product is bounded
	StrategyId string `json:"strategy_id"`
}

func (c Strategies) GetByProductIds(ctx context.Context, params *GetStrategiesByProductIdsParams) (*GetStrategiesByProductIdsResponse, error) {
	url := "/v1/pricing-strategy/strategy-ids-by-product-ids"

	resp := &GetStrategiesByProductIdsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListProductsInStrategyParams struct {
	// Strategy identifier
	StrategyId string `json:"strategy_id"`
}

type ListProductsInStrategyResponse struct {
	core.CommonResponse

	// Method result
	Result ListProductsInStrategyResult `json:"result"`
}

type ListProductsInStrategyResult struct {
	// Product identifier
	ProductId []string `json:"product_id"`
}

func (c Strategies) ListProducts(ctx context.Context, params *ListProductsInStrategyParams) (*ListProductsInStrategyResponse, error) {
	url := "/v1/pricing-strategy/products/list"

	resp := &ListProductsInStrategyResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetCompetitorPriceParams struct {
	// Product identifier
	ProductId int64 `json:"product_id"`
}

type GetCompetitorPriceResponse struct {
	core.CommonResponse

	// Method result
	Result GetCompetitorPriceResult `json:"result"`
}

type GetCompetitorPriceResult struct {
	// Product identifier
	StrategyId string `json:"strategy_id"`

	// true if the product is in the pricing strategy
	IsEnabled bool `json:"is_enabled"`

	// Price of product in the strategy
	StrategyProductPrice int32 `json:"strategy_product_price"`

	// Price setting date
	PriceDownloadedAt string `json:"price_downloaded_at"`

	// Competitor identifier
	StrategyCompetitorId int64 `json:"strategy_competitor_id"`

	// Link to a competitor's product
	StrategyCompetitorProductURL string `json:"strategy_competitor_product_url"`
}

func (c Strategies) GetCompetitorPrice(ctx context.Context, params *GetCompetitorPriceParams) (*GetCompetitorPriceResponse, error) {
	url := "/v1/pricing-strategy/product/info"

	resp := &GetCompetitorPriceResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type RemoveProductsFromStrategyParams struct {
	// List of product identifiers. The maximum number is 50
	ProductId []int64 `json:"product_id"`
}

type RemoveProductsFromStrategyResponse struct {
	core.CommonResponse

	// Method result
	Result RemoveProductsFromStrategyResult `json:"result"`
}

type RemoveProductsFromStrategyResult struct {
	// Number of products with errors
	FailedProductCount int32 `json:"failed_product_count"`
}

func (c Strategies) RemoveProducts(ctx context.Context, params *RemoveProductsFromStrategyParams) (*RemoveProductsFromStrategyResponse, error) {
	url := "/v1/pricing-strategy/products/delete"

	resp := &RemoveProductsFromStrategyResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ChangeStrategyStatusParams struct {
	// Strategy status
	Enabled bool `json:"enabled"`

	// Product identifier
	StrategyId string `json:"strategy_id"`
}

type ChangeStrategyStatusResponse struct {
	core.CommonResponse
}

func (c Strategies) ChangeStatus(ctx context.Context, params *ChangeStrategyStatusParams) (*ChangeStrategyStatusResponse, error) {
	url := "/v1/pricing-strategy/status"

	resp := &ChangeStrategyStatusResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type RemoveStrategyParams struct {
	// Strategy identifier
	StrategyId string `json:"strategy_id"`
}

type RemoveStrategyResponse struct {
	core.CommonResponse
}

func (c Strategies) Remove(ctx context.Context, params *RemoveStrategyParams) (*RemoveStrategyResponse, error) {
	url := "/v1/pricing-strategy/delete"

	resp := &RemoveStrategyResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
