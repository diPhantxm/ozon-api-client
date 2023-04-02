package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Promotions struct {
	client *core.Client
}

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

// A method for getting a list of promotions that you can participate in
func (c Promotions) GetAvailablePromotions() (*GetAvailablePromotionsResponse, error) {
	url := "/v1/actions"

	resp := &GetAvailablePromotionsResponse{}

	response, err := c.client.Request(http.MethodGet, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type AddProductToPromotionParams struct {
	// Promotion identifier
	ActionId float64 `json:"action_id"`

	// Products list
	Products []AddProductToPromotionProduct `json:"products"`
}

type AddProductToPromotionProduct struct {
	// Product identifier
	ProductId float64 `json:"produt_id"`

	// Promotional product price
	ActionPrice float64 `json:"action_price"`

	// Number of product units in a stock discount type promotion
	Stock float64 `json:"stock"`
}

type AddProductToPromotionResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// List of product identifiers that were added to the promotion
		ProductIds []float64 `json:"product_ids"`

		// List of products that weren't added to the promotion
		Rejected []struct {
			// Product identifier
			ProductId float64 `json:"product_id"`

			// Reason why the product wasn't added to the promotion
			Reason string `json:"reason"`
		} `json:"rejected"`
	} `json:"result"`
}

// A method for adding products to an available promotion
func (c Promotions) AddToPromotion(params *AddProductToPromotionParams) (*AddProductToPromotionResponse, error) {
	url := "/v1/actions/products/activate"

	resp := &AddProductToPromotionResponse{}

	response, err := c.client.Request(http.MethodGet, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ProductsAvailableForPromotionParams struct {
	// Promotion identifier
	ActionId float64 `json:"action_id"`

	// Number of values in the response. The default value is 100
	Limit float64 `json:"limit"`

	// Number of elements that will be skipped in the response.
	// For example, if offset=10, the response will start with the 11th element found
	Offset float64 `json:"offset"`
}

type ProductsAvailableForPromotionResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Products list
		Products []PromotionProduct `json:"products"`

		// Total number of products that can participate in the promotion
		Total float64 `json:"total"`
	} `json:"result"`
}

type PromotionProduct struct {
	// Product identifier
	Id float64 `json:"id"`

	// Current product price without a discount
	Price float64 `json:"price"`

	// Promotional product price
	ActionPrice float64 `json:"action_price"`

	// Maximum possible promotional product price
	MaxActionType float64 `json:"max_action_type"`

	// Type of adding a product to the promotion: automatically or manually by the seller
	AddMode string `json:"add_mode"`

	// Minimum number of product units in a stock discount type promotion
	MinStock float64 `json:"min_stock"`

	// Number of product units in a stock discount type promotion
	Stock float64 `json:"stock"`
}

// A method for getting a list of products that can participate in the promotion by the promotion identifier
func (c Promotions) ProductsAvailableForPromotion(params *ProductsAvailableForPromotionParams) (*ProductsAvailableForPromotionResponse, error) {
	url := "/v1/actions/candidates"

	resp := &ProductsAvailableForPromotionResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ProductsInPromotionParams struct {
	// Promotion identifier
	ActionId float64 `json:"action_id"`

	// Number of values in the response. The default value is 100
	Limit float64 `json:"limit"`

	// Number of elements that will be skipped in the response. For example, if offset=10, the response will start with the 11th element found
	Offset float64 `json:"offset"`
}

type ProductsInPromotionResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Products list
		Products []PromotionProduct `json:"products"`

		// Total number of products that can participate in the promotion
		Total float64 `json:"total"`
	} `json:"reuslt"`
}

// A method for getting the list of products participating in the promotion by its identifier
func (c Promotions) ProductsInPromotion(params *ProductsInPromotionParams) (*ProductsInPromotionResponse, error) {
	url := "/v1/actions/products"

	resp := &ProductsInPromotionResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type RemoveProductFromPromotionParams struct {
	// Promotion identifier
	ActionId float64 `json:"action_id"`

	// List of products identifiers
	ProductIds []float64 `json:"product_ids"`
}

type RemoveProductFromPromotionResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// List of product identifiers that were removed from the promotion
		ProductIds []float64 `json:"product_ids"`

		// List of product identifiers that weren't removed from the promotion
		Rejected []struct {
			// Product identifier
			ProductId float64 `json:"product_id"`

			// Reason why the product wasn't removed from the promotion
			Reason string `json:"reason"`
		} `json:"rejected"`
	} `json:"result"`
}

// A method for removing products from the promotion
func (c Promotions) RemoveProduct(params *RemoveProductFromPromotionParams) (*RemoveProductFromPromotionResponse, error) {
	url := "/v1/actions/products/deactivate"

	resp := &RemoveProductFromPromotionResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListHotSalePromotionsResponse struct {
	core.CommonResponse

	// Method result
	Result []struct {
		// Promotion end date
		DateEnd string `json:"date_end"`

		// Promotion start date
		DateStart string `json:"date_start"`

		// Promotion description
		Description string `json:"description"`

		// Promotion freeze date.
		//
		// If the field is filled, the seller can't increase prices, change the list of products,
		// or decrease the number of product units in the promotion.
		//
		// The seller can lower prices and increase the product units number in the promotion
		FreezeDate string `json:"freeze_date"`

		// Hot Sale promotion identifier
		HotsaleId float64 `json:"hotsale_id"`

		// Indication that you participate in this promotion
		IsParticipating bool `json:"is_participating"`

		// Promotion name
		Title string `json:"title"`
	} `json:"result"`
}

// List of available Hot Sale promotions
func (c Promotions) ListHotSalePromotions() (*ListHotSalePromotionsResponse, error) {
	url := "/v1/actions/hotsales/list"

	resp := &ListHotSalePromotionsResponse{}

	response, err := c.client.Request(http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ProductsAvailableForHotSalePromotionParams struct {
	// Hot Sale promotion identifier
	HotSaleId float64 `json:"hotsale_id"`

	// Number of elements in the response. Default value is 100
	Limit float64 `json:"limit"`

	// Number of elements that will be skipped in the response. For example, if offset=10, the response will start with the 11th element found
	Offset float64 `json:"offset"`
}

type ProductsAvailableForHotSalePromotionResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Products list
		Products []struct {
			// Promotional product price
			ActionPrice float64 `json:"action_price"`

			// Date when the product participates in the promotion in the YYYY-MM-DD format
			DateDayPromo string `json:"date_day_promo"`

			// Product identifier
			Id float64 `json:"id"`

			// Indication that product participates in the promotion
			IsActive bool `json:"is_active"`

			// Maximum possible promotional price of the product
			MaxActionPrice float64 `json:"max_action_type"`

			// Minimum number of product units in a stock discount type promotion
			MinStock float64 `json:"min_stock"`

			// Number of product units in a stock discount type promotion
			Stock float64 `json:"stock"`
		} `json:"products"`

		// Total number of products that are available for the promotion
		Total float64 `json:"total"`
	} `json:"result"`
}

// Method for getting a list of products that can participate or are already participating in the Hot Sale promotion
func (c Promotions) ProductsAvailableForHotSalePromotion(params *ProductsAvailableForHotSalePromotionParams) (*ProductsAvailableForHotSalePromotionResponse, error) {
	url := "/v1/actions/hotsales/products"

	resp := &ProductsAvailableForHotSalePromotionResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type AddProductsToHotSaleParams struct {
	// Hot Sale promotion identifier
	HotSaleId float64 `json:"hotsale_id"`

	// Products to be added to the promotion. The maximum number in one request is 100
	Products []AddProductsToHotSaleProduct `json:"products"`
}

type AddProductsToHotSaleProduct struct {
	// Promotional product price
	ActionPrice float64 `json:"action_price"`

	// Product identifier
	ProductId float64 `json:"product_id"`

	// Number of product units in a stock discount type promotion
	Stock float64 `json:"stock"`
}

type ProductsToHotSaleResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// List of products that haven't been added to the promotion
		Rejected []struct {
			//Product identifier
			ProductId float64 `json:"product_id"`

			// Reason why the product hasn't been added to the promotion
			Reason string `json:"reason"`
		} `json:"rejected"`
	} `json:"result"`
}

func (c Promotions) AddProductsToHotSale(params *AddProductsToHotSaleParams) (*ProductsToHotSaleResponse, error) {
	url := "/v1/actions/hotsales/activate"

	resp := &ProductsToHotSaleResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type RemoveProductsToHotSaleParams struct {
	// Hot Sale promotion identifier
	HotSaleId float64 `json:"hotsale_id"`

	// List of products identifiers. Maximum number of values in one request is 100
	ProductIds []float64 `json:"product_ids"`
}

// Remove product from the Hot Sale promotion
func (c Promotions) RemoveProductsToHotSale(params *RemoveProductsToHotSaleParams) (*ProductsToHotSaleResponse, error) {
	url := "/v1/actions/hotsales/activate"

	resp := &ProductsToHotSaleResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListDiscountRequestsParams struct {
	// Discount request status
	Status ListDiscountRequestsStatus `json:"status" default:"UNKNOWN"`

	// Page number from which you want to download the list of discount requests
	Page uint64 `json:"page"`

	// The maximum number of requests on a page
	Limit uint64 `json:"limit"`
}

type ListDiscountRequestsResponse struct {
	core.CommonResponse

	// List of requests
	Result []struct {
		// Request ID
		Id uint64 `json:"id"`

		// Request created date
		CreatedAt time.Time `json:"created_at"`

		// End time of the request
		EndAt time.Time `json:"end_at"`

		// Time to change the decision
		EditedTill time.Time `json:"edited_till"`

		// Request status
		Status string `json:"status"`

		// Customer's name
		CustomerName string `json:"customer_name"`

		// Product identifier in the Ozon system, SKU
		SKU uint64 `json:"sku"`

		// Customer's comment on the request
		UserComment string `json:"user_comment"`

		// Seller's comment on the request
		SellerComment string `json:"seller_comment"`

		// Requested price
		RequestedPrice float64 `json:"requested_price"`

		// Approved price
		ApprovedPrice float64 `json:"approved_price"`

		// Product price before all discounts
		OriginalPrice float64 `json:"original_price"`

		// Discount in rubles
		Discount float64 `json:"discount"`

		// Discount percentage
		DiscountPercent float64 `json:"discount_percent"`

		// Base price at which a product is selling on Ozon, if not eligible for a promotion
		BasePrice float64 `json:"base_price"`

		// The minimum price after auto-application of discounts and promotions
		MinAutoPrice float64 `json:"min_auto_price"`

		// ID of the previous customer request for this product
		PrevTaskId uint64 `json:"prev_task_id"`

		// If product is damaged — true
		IsDamaged bool `json:"is_damaged"`

		// Moderation date: review, approval or decline of the request
		ModeratedAt time.Time `json:"moderated_at"`

		// Discount in rubles approved by the seller. Pass the value 0 if the seller did not approve the request
		ApprovedDiscount float64 `json:"approved_discount"`

		// Discount percentage approved by the seller. Pass the value 0 if the seller did not approve the request
		ApprovedDiscountPercent float64 `json:"approved_discount_percent"`

		// Whether the customer has purchased the product. true if purchased
		IsPurchased bool `json:"is_purchased"`

		// Whether the request was moderated automatically. true if moderation was automatic
		IsAutoModerated bool `json:"is_auto_moderated"`

		// Product identifier in the seller's system
		OfferId string `json:"offer_id"`

		// Email of the user who processed the request
		Email string `json:"email"`

		// Last name of the user who processed the request
		LastName string `json:"last_name"`

		// First name of the user who processed the request
		FirstName string `json:"first_name"`

		// Patronymic of the user who processed the request
		Patronymic string `json:"patronymic"`

		// Approved minimum quantity of products
		ApprovedQuantityMin uint64 `json:"approved_quantity_min"`

		// Approved maximum quantity of products
		ApprovedQuantityMax uint64 `json:"approved_quantity_max"`

		// Requested minimum number of products
		RequestedQuantityMin uint64 `json:"requested_quantity_min"`

		// Requested maximum number of products
		RequestedQuantityMax uint64 `json:"requested_quantity_max"`

		// Requested price with fee
		RequestedPriceWithFee float64 `json:"requested_price_with_fee"`

		// Approved price with fee
		ApprovedPriceWithFee float64 `json:"approved_price_with_fee"`

		// Approved price fee percent
		ApprovedPriceFeePercent float64 `json:"approved_price_fee_percent"`
	} `json:"result"`
}

// Method for getting a list of products that customers want to buy with discount
func (c Promotions) ListDiscountRequests(params *ListDiscountRequestsParams) (*ListDiscountRequestsResponse, error) {
	url := "/v1/actions/discounts-task/list"

	resp := &ListDiscountRequestsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type DiscountRequestParams struct {
	// List of discount requests
	Tasks []DiscountRequestTask `json:"tasks"`
}

type DiscountRequestTask struct {
	// Request ID
	Id uint64 `json:"id"`

	// Approved price
	ApprovedPrice float64 `json:"approved_price"`

	// Seller's comment on the request
	SellerComment string `json:"seller_comment"`

	// Approved minimum quantity of products
	ApprovedQuantityMin uint64 `json:"approved_quantity_min"`

	// Approved maximum quantity of products
	ApprovedQuantityMax uint64 `json:"approved_quantity_max"`
}

type DiscountRequestResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Errors when creating a request
		FailDetails []struct {
			// Request ID
			TaskId uint64 `json:"task_id"`

			// Error message
			ErrorForUser string `json:"error_for_user"`
		} `json:"fail_details"`

		// The number of requests with a successful status change
		SuccessCount int32 `json:"success_count"`

		// The number of requests that failed to change their status
		FailCount int32 `json:"fail_count"`
	} `json:"result"`
}

// You can approve applications in statuses:
//   - NEW — new
//   - SEEN — viewed
func (c Promotions) ApproveDiscountRequest(params *DiscountRequestParams) (*DiscountRequestResponse, error) {
	url := "/v1/actions/discounts-task/approve"

	resp := &DiscountRequestResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

// You can decline applications in statuses:
//   - NEW—new
//   - SEEN—viewed
func (c Promotions) DeclineDiscountRequest(params *DiscountRequestParams) (*DiscountRequestResponse, error) {
	url := "/v1/actions/discounts-task/decline"

	resp := &DiscountRequestResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
