package ozon

import (
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type GetAvailablePromotionsResponse struct {
	core.CommonResponse

	// Method result
	Result []struct {
		// Promotion identifier
		Id float64 `json:"id"`

		// Promotion name
		Title string `json:"title"`

		// Promotion type
		ActionType string `json:"action_type"`

		// Promotion description
		Description string `json:"description"`

		// Promotion start date
		DateStart string `json:"date_start"`

		// Promotion end date
		DateEnd string `json:"date_end"`

		// Promotion freeze date.
		//
		// If the field is filled, the seller can't increase prices, change the list of products, or decrease the number of product units in the promotion.
		//
		// The seller can lower prices and increase the product units number in the promotion
		FreezeDate string `json:"freeze_date"`

		// Number of products that can participate in the promotion
		PotentialProductsCount float64 `json:"potential_products_count"`

		// Number of products that participate in the promotion
		ParticipatingProductsCount float64 `json:"participating_products_count"`

		// Whether or not you participate in the promotion
		IsParticipating bool `json:"participating"`

		// Indication that customers need a promo code to participate in the promotion
		IsVoucherAction bool `json:"is_voucher_action"`

		// Number of blocked products
		BannedProductsCount float64 `json:"banned_products_count"`

		// Indication of the promotion is with the target audience
		WithTargeting bool `json:"with_targeting"`

		// Order amount
		OrderAmount float64 `json:"order_amount"`

		// Discount type
		DiscountType string `json:"discount_type"`

		// Discount size
		DiscountValue float64 `json:"discount_value"`
	} `json:"result"`
}

func (c Client) GetAvailablePromotions() (*GetAvailablePromotionsResponse, error) {
	url := "/v1/actions"

	resp := &GetAvailablePromotionsResponse{}

	response, err := c.client.Request(http.MethodGet, url, nil, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
