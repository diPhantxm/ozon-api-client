package ozon

import (
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type Strategies struct {
	client *core.Client
}

type ListCompetitorsParams struct {
	Page int64 `json:"page"`

	Limit int64 `json:"limit"`
}

type ListCompetitorsResponse struct {
	core.CommonResponse

	Competitors []struct {
		Name string `json:"name"`

		Id int64 `json:"id"`
	} `json:"competitors"`

	Total int32 `json:"total"`
}

func (c Strategies) ListCompetitors(params *ListCompetitorsParams) (*ListCompetitorsResponse, error) {
	url := "/v1/pricing-strategy/competitors/list"

	resp := &ListCompetitorsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListStrategiesParams struct {
	Page int64 `json:"page"`

	Limit int64 `json:"limit"`
}

type ListStrategiesResponse struct {
	core.CommonResponse

	Strategies []struct {
		Id string `json:"id"`

		Name string `json:"name"`

		Type StrategyType `json:"type"`

		UpdateType StrategyUpdateType `json:"update_type"`

		UpdatedAt string `json:"updated_at"`

		ProductsCount int64 `json:"products_count"`

		CompetitorsCount int64 `json:"competitors_count"`

		Enabled bool `json:"enabled"`
	} `json:"strategies"`

	Total int32 `json:"total"`
}

func (c Strategies) List(params *ListStrategiesParams) (*ListStrategiesResponse, error) {
	url := "/v1/pricing-strategy/list"

	resp := &ListStrategiesResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CreateStrategyParams struct {
	Competitors []CreateStrategyCompetitor `json:"competitors"`

	StrategyName string `json:"strategy_name"`
}

type CreateStrategyCompetitor struct {
	Coefficient float32 `json:"coefficient"`

	CompetitorId int64 `json:"competitor_id"`
}

type CreateStrategyResponse struct {
	core.CommonResponse

	Result struct {
		StrategyId string `json:"strategy_id"`
	} `json:"result"`
}

func (c Strategies) Create(params *CreateStrategyParams) (*CreateStrategyResponse, error) {
	url := "/v1/pricing-strategy/create"

	resp := &CreateStrategyResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type InfoStrategyParams struct {
	StrategyId string `json:"strategy_id"`
}

type InfoStrategyResponse struct {
	core.CommonResponse

	Result struct {
		Competitors []CreateStrategyCompetitor `json:"competitors"`

		Enabled bool `json:"enabled"`

		Name string `json:"name"`

		Type StrategyType `json:"type"`

		UpdateType StrategyUpdateType `json:"update_type"`
	} `json:"result"`
}

func (c Strategies) Info(params *InfoStrategyParams) (*InfoStrategyResponse, error) {
	url := "/v1/pricing-strategy/info"

	resp := &InfoStrategyResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateStrategyParams struct {
	Competitors []CreateStrategyCompetitor `json:"competitors"`

	StrategyId string `json:"strategy_id"`

	StrategyName string `json:"strategy_name"`
}

type UpdateStrategyResponse struct {
	core.CommonResponse
}

func (c Strategies) Update(params *UpdateStrategyParams) (*UpdateStrategyResponse, error) {
	url := "/v1/pricing-strategy/update"

	resp := &UpdateStrategyResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type AddProductsToStrategyParams struct {
	ProductId []int64 `json:"product_id"`

	StrategyId string `json:"strategy_id"`
}

type AddProductsToStrategyResponse struct {
	core.CommonResponse

	Result struct {
		Errors []struct {
			Code string `json:"code"`

			Error string `json:"error"`

			ProductId int64 `json:"product_id"`
		} `json:"errors"`

		FailedProductCount int32 `json:"failed_product_count"`
	} `json:"result"`
}

func (c Strategies) AddProducts(params *AddProductsToStrategyParams) (*AddProductsToStrategyResponse, error) {
	url := "/v1/pricing-strategy/products/add"

	resp := &AddProductsToStrategyResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetStrategiesByProductIdsParams struct {
	ProductId []int64 `json:"product_id"`
}

type GetStrategiesByProductIdsResponse struct {
	core.CommonResponse

	Result struct {
		ProductsInfo []struct {
			ProductId int64 `json:"product_id"`

			StrategyId string `json:"strategy_id"`
		} `json:"products_info"`
	} `json:"result"`
}

func (c Strategies) GetByProductIds(params *GetStrategiesByProductIdsParams) (*GetStrategiesByProductIdsResponse, error) {
	url := "/v1/pricing-strategy/strategy-ids-by-product-ids"

	resp := &GetStrategiesByProductIdsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListProductsInStrategyParams struct {
	StrategyId string `json:"strategy_id"`
}

type ListProductsInStrategyResponse struct {
	core.CommonResponse

	Result struct {
		ProductId []string `json:"product_id"`
	} `json:"result"`
}

func (c Strategies) ListProducts(params *ListProductsInStrategyParams) (*ListProductsInStrategyResponse, error) {
	url := "/v1/pricing-strategy/products/list"

	resp := &ListProductsInStrategyResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetCompetitorPriceParams struct {
	ProductId int64 `json:"product_id"`
}

type GetCompetitorPriceResponse struct {
	core.CommonResponse

	Result struct {
		StrategyId string `json:"strategy_id"`

		IsEnabled bool `json:"is_enabled"`

		StrategyProductPrice int32 `json:"strategy_product_price"`

		PriceDownloadedAt string `json:"price_downloaded_at"`

		StrategyCompetitorId int64 `json:"strategy_competitor_id"`

		StrategyCompetitorProductURL string `json:"strategy_competitor_product_url"`
	} `json:"result"`
}

func (c Strategies) GetCompetitorPrice(params *GetCompetitorPriceParams) (*GetCompetitorPriceResponse, error) {
	url := "/v1/pricing-strategy/product/info"

	resp := &GetCompetitorPriceResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type RemoveProductsFromStrategyParams struct {
	ProductId []int64 `json:"product_id"`
}

type RemoveProductsFromStrategyResponse struct {
	core.CommonResponse

	Result struct {
		FailedProductCount int32 `json:"failed_product_count"`
	} `json:"result"`
}

func (c Strategies) RemoveProducts(params *RemoveProductsFromStrategyParams) (*RemoveProductsFromStrategyResponse, error) {
	url := "/v1/pricing-strategy/products/delete"

	resp := &RemoveProductsFromStrategyResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ChangeStrategyStatusParams struct {
	Enabled bool `json:"enabled"`

	StrategyId string `json:"strategy_id"`
}

type ChangeStrategyStatusResponse struct {
	core.CommonResponse
}

func (c Strategies) ChangeStatus(params *ChangeStrategyStatusParams) (*ChangeStrategyStatusResponse, error) {
	url := "/v1/pricing-strategy/status"

	resp := &ChangeStrategyStatusResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type RemoveStrategyParams struct {
	StrategyId string `json:"strategy_id"`
}

type RemoveStrategyResponse struct {
	core.CommonResponse
}

func (c Strategies) Remove(params *RemoveStrategyParams) (*RemoveStrategyResponse, error) {
	url := "/v1/pricing-strategy/delete"

	resp := &RemoveStrategyResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
