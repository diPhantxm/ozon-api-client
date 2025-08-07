package _import

type InfoRequest struct {
	TaskID string `json:"task_id"`
}

type InfoResponseResultItemErrorOptionalDescriptionElements struct {
	PropertyName string `json:"property name*"`
}

type InfoResponseResultItemError struct {
	Code                        string                                                 `json:"code"`
	Message                     string                                                 `json:"message"`
	State                       string                                                 `json:"state"`
	Level                       string                                                 `json:"level"`
	Description                 string                                                 `json:"description"`
	Field                       string                                                 `json:"field"`
	AttributeID                 int64                                                  `json:"attribute_id"`
	AttributeName               string                                                 `json:"attribute_name"`
	OptionalDescriptionElements InfoResponseResultItemErrorOptionalDescriptionElements `json:"optional_description_elements"`
}

type InfoResponseResultItem struct {
	OfferID   string                        `json:"offer_id"`
	ProductID int64                         `json:"product_id"`
	Status    InfoResponseResultItemStatus  `json:"status"`
	Errors    []InfoResponseResultItemError `json:"errors"`
}

type InfoResponseResult struct {
	Items []InfoResponseResultItem `json:"items"`
	Total int32                    `json:"total"`
}

type InfoResponse struct {
	Result InfoResponseResult `json:"result"`
}

type StocksRequestStock struct {
	OfferID   string `json:"offer_id"`
	ProductID int64  `json:"product_id"`
	Stock     int64  `json:"stock"`
}

type StocksRequest struct {
	Stocks []StocksRequestStock `json:"stocks"`
}

type StocksResponseResultError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type StocksResponseResult struct {
	Errors    []StocksResponseResultError `json:"errors"`
	OfferID   string                      `json:"offer_id"`
	ProductID int64                       `json:"product_id"`
	Updated   bool                        `json:"updated"`
}

type StocksResponse struct {
	Result []StocksResponseResult
}

type PricesRequestPrice struct {
	AutoActionEnabled    PricesRequestPriceAutoActionEnabled    `json:"auto_action_enabled"`
	CurrencyCode         PricesRequestPriceCurrencyCode         `json:"currency_code"`
	MinPrice             string                                 `json:"min_price"`
	NetPrice             string                                 `json:"net_price"`
	OfferID              string                                 `json:"offer_id"`
	OldPrice             string                                 `json:"old_price"`
	Price                string                                 `json:"price"`
	PriceStrategyEnabled PricesRequestPricePriceStrategyEnabled `json:"price_strategy_enabled"`
	ProductID            int64                                  `json:"product_id"`
}

type PricesRequest struct {
	Prices []PricesRequestPrice `json:"prices"`
}

type PricesResponseResultError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PricesResponseResult struct {
	Errors    []PricesResponseResultError `json:"errors"`
	OfferID   string                      `json:"offer_id"`
	ProductID int64                       `json:"product_id"`
	Updated   bool                        `json:"updated"`
}

type PricesResponse struct {
	Result []PricesResponseResult `json:"result"`
}
