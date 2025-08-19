package info

type PricesRequestFilter struct {
	OfferID    []string                      `json:"offer_id"`
	ProductID  []string                      `json:"product_id"`
	Visibility PricesRequestFilterVisibility `json:"visibility"`
}

type PricesRequest struct {
	Cursor string              `json:"cursor"`
	Filter PricesRequestFilter `json:"filter"`
	Limit  int32               `json:"limit"`
}

type PricesResponseItemCommissions struct {
	FBODelivToCustomerAmount    float64 `json:"fbo_deliv_to_customer_amount"`
	FBODirectFlowTransMaxAmount float64 `json:"fbo_direct_flow_trans_max_amount"`
	FBODirectFlowTransMinAmount float64 `json:"fbo_direct_flow_trans_min_amount"`
	FBOReturnFlowAmount         float64 `json:"fbo_return_flow_amount"`
	FBSDelivToCustomerAmount    float64 `json:"fbs_deliv_to_customer_amount"`
	FBSDirectFlowTransMaxAmount float64 `json:"fbs_direct_flow_trans_max_amount"`
	FBSDirectFlowTransMinAmount float64 `json:"fbs_direct_flow_trans_min_amount"`
	FBSFirstMileMaxAmount       float64 `json:"fbs_first_mile_max_amount"`
	FBSFirstMileMinAmount       float64 `json:"fbs_first_mile_min_amount"`
	FBSReturnFlowAmount         float64 `json:"fbs_return_flow_amount"`
	SalesPercentFBO             float64 `json:"sales_percent_fbo"`
	SalesPercentFBS             float64 `json:"sales_percent_fbs"`
}

type PricesResponseItemMarketingActionsAction struct {
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
	Title    string `json:"title"`
	Value    int    `json:"value"`
}

type PricesResponseItemMarketingActions struct {
	Actions           []PricesResponseItemMarketingActionsAction `json:"actions"`
	CurrentPeriodFrom string                                     `json:"current_period_from"`
	CurrentPeriodTo   string                                     `json:"current_period_to"`
	OzonActionsExist  bool                                       `json:"ozon_actions_exist"`
}

type PricesResponseItemPrice struct {
	AutoActionEnabled    bool                                `json:"auto_action_enabled"`
	CurrencyCode         PricesResponseItemPriceCurrencyCode `json:"currency_code"`
	MarketingPrice       float64                             `json:"marketing_price"`
	MarketingSellerPrice float64                             `json:"marketing_seller_price"`
	MinPrice             float64                             `json:"min_price"`
	OldPrice             float64                             `json:"old_price"`
	Price                float64                             `json:"price"`
	RetailPrice          float64                             `json:"retail_price"`
	Vat                  float64                             `json:"vat"`
}

type PricesResponseItemPriceIndexesExternalIndexData struct {
	MinPrice         float64 `json:"min_price"`
	MinPriceCurrency string  `json:"min_price_currency"`
	PriceIndexValue  float64 `json:"price_index_value"`
}

type PricesResponseItemPriceIndexesOzonIndexData struct {
	MinPrice         float64 `json:"min_price"`
	MinPriceCurrency string  `json:"min_price_currency"`
	PriceIndexValue  float64 `json:"price_index_value"`
}

type PricesResponseItemPriceIndexesSelfMarketplacesIndexData struct {
	MinPrice         float64 `json:"min_price"`
	MinPriceCurrency string  `json:"min_price_currency"`
	PriceIndexValue  float64 `json:"price_index_value"`
}

type PricesResponseItemPriceIndexes struct {
	ColorIndex                PricesResponseItemPriceIndexesColorIndex                `json:"color_index"`
	ExternalIndexData         PricesResponseItemPriceIndexesExternalIndexData         `json:"external_index_data"`
	OzonIndexData             PricesResponseItemPriceIndexesOzonIndexData             `json:"ozon_index_data"`
	SelfMarketplacesIndexData PricesResponseItemPriceIndexesSelfMarketplacesIndexData `json:"self_marketplaces_index_data"`
}

type PricesResponseItem struct {
	Acquiring        float64                            `json:"acquiring"`
	Commissions      PricesResponseItemCommissions      `json:"commissions"`
	MarketingActions PricesResponseItemMarketingActions `json:"marketing_actions"`
	OfferID          string                             `json:"offer_id"`
	Price            PricesResponseItemPrice            `json:"price"`
	PriceIndexes     PricesResponseItemPriceIndexes     `json:"price_indexes"`
	ProductID        int64                              `json:"product_id"`
	VolumeWeight     float64                            `json:"volume_weight"`
}

type PricesResponse struct {
	Cursor string               `json:"cursor"`
	Items  []PricesResponseItem `json:"items"`
	Total  int32                `json:"total"`
}
