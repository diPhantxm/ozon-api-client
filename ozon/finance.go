package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Finance struct {
	client *core.Client
}

type ReportOnSoldProductsParams struct {
	// Time period in the `YYYY-MM` format
	Date string `json:"date"`
}

type ReportOnSoldProductsResponse struct {
	core.CommonResponse

	// Query result
	Result []struct {
		// Report title page
		Header []struct {
			// Report ID
			Id string `json:"num"`

			// Report generation date
			DocDate string `json:"doc_date"`

			// Date of the offer agreement
			ContractDate string `json:"contract_date"`

			// Offer agreement number
			ContractNum string `json:"contract_num"`

			// Currency of your prices
			CurrencyCode string `json:"currency_code"`

			// Amount to accrue
			DocAmount float64 `json:"doc_amount"`

			// Amount to accrue with VAT
			VATAmount float64 `json:"vat_amount"`

			// Payer's TIN
			PayerINN string `json:"payer_inn"`

			// Payer's Tax Registration Reason Code (KPP)
			PayerKPP string `json:"payer_kpp"`

			// Payer's name
			PayerName string `json:"payer_name"`

			// Recipient's TIN
			RecipientINN string `json:"rcv_inn"`

			// Recipient's Tax Registration Reason Code (KPP)
			RecipientKPP string `json:"rcv_kpp"`

			// Recipient's name
			RecipientName string `json:"rcv_name"`

			// Period start in the report
			StartDate string `json:"start_date"`

			// Period end in the report
			StopDate string `json:"stop_date"`
		} `json:"header"`

		// Report table
		Rows []struct {
			// Row number
			RowNumber int32 `json:"row_number"`

			// Product ID
			ProductId int64 `json:"product_id"`

			// Product name
			ProductName string `json:"product_name"`

			// Product barcode
			Barcode string `json:"barcode"`

			// Product identifier in the seller's system
			OfferId string `json:"offer_id"`

			// Sales commission by category
			CommissionPercent float64 `json:"commission_percent"`

			// Seller's price with their discount
			Price float64 `json:"price"`

			// Selling price: the price at which the customer purchased the product. For sold products
			PriceSale float64 `json:"price_sale"`

			// Sold for amount.
			//
			// Sold products cost considering the quantity and regional coefficients. Calculation is made by the sale_amount price
			SaleAmount float64 `json:"sale_amount"`

			// Commission for sold products, including discounts and extra charges
			SaleCommission float64 `json:"sale_commission"`

			// Extra charge at the expense of Ozon.
			//
			// Amount that Ozon will compensate the seller if the Ozon discount is greater than or equal to the sales commission
			SaleDiscount float64 `json:"sale_discount"`

			// Total accrual for the products sold.
			//
			// Amount after deduction of sales commission, application of discounts and extra charges
			SalePriceSeller float64 `json:"sale_price_seller"`

			// Quantity of products sold at the price_sale price
			SaleQuantity int32 `json:"sale_qty"`

			// Price at which the customer purchased the product. For returned products
			ReturnSale float64 `json:"return_sale"`

			// Cost of returned products, taking into account the quantity and regional coefficients.
			// Calculation is carried out at the return_sale price
			ReturnAmount float64 `json:"return_amount"`

			// Commission including the quantity of products, discounts and extra charges.
			// Ozon compensates it for the returned products
			ReturnCommission float64 `json:"return_commission"`

			// Extra charge at the expense of Ozon.
			//
			// Amount of the discount at the expense of Ozon on returned products.
			// Ozon will compensate it to the seller if the Ozon discount is greater than or equal to the sales commission
			ReturnDiscount float64 `json:"return_discount"`

			// Amount charged to the seller for returned products after deducing sales commissions, applying discounts and extra charges
			ReturnPriceSeller float64 `json:"return_price_seller"`

			// Quantity of returned products
			ReturnQuantity int32 `json:"return_qty"`
		} `json:"rows"`
	} `json:"result"`
}

// Returns information on products sold and returned within a month. Canceled or non-purchased products are not included.
//
// Report is returned no later than the 5th day of the next month
func (c Finance) ReportOnSoldProducts(params *ReportOnSoldProductsParams) (*ReportOnSoldProductsResponse, error) {
	url := "/v1/finance/realization"

	resp := &ReportOnSoldProductsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetTotalTransactionsSumParams struct {
	// Filter by date
	Date GetTotalTransactionsSumDate `json:"date"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Transaction type:
	//
	//   - all — all,
	//   - orders — orders,
	//   - returns — returns and cancellations,
	//   - services — service fees,
	//   - compensation — compensation,
	//   - transferDelivery — delivery cost,
	//   - other — other
	TransactionType string `json:"transaction_type"`
}

type GetTotalTransactionsSumDate struct {
	// Period start.
	//
	// Format: YYYY-MM-DDTHH:mm:ss.sssZ.
	// Example: 2019-11-25T10:43:06.51
	From time.Time `json:"from"`

	// Period end.
	//
	// Format: YYYY-MM-DDTHH:mm:ss.sssZ.
	// Example: 2019-11-25T10:43:06.51
	To time.Time `json:"to"`
}

type GetTotalTransactionsSumResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Total cost of products and returns for specified period
		AccrualsForSale float64 `json:"accruals_for_sale"`

		// Compensations
		CompensationAmount float64 `json:"compensatino_amount"`

		// Charges for delivery and returns when working under rFBS scheme
		MoneyTransfer float64 `json:"money_transfer"`

		// Other accurals
		OthersAmount float64 `json:"others_amount"`

		// Cost of shipment processing, orders packaging, pipeline and last mile services, and delivery cost before the new commissions and rates applied from February 1, 2021.
		//
		// Pipeline is delivery of products from one cluster to another.
		//
		// Last mile is products delivery to the pick-up point, parcle terminal, or by courier
		ProcessingAndDelivery float64 `json:"processing_and_delivery"`

		// Cost of reverse pipeline, returned, canceled and unredeemed orders processing, and return cost before the new commissions and rates applied from February 1, 2021.
		//
		// Pipeline is delivery of products from one cluster to another.
		//
		// Last mile is products delivery to the pick-up point, parcle terminal, or by courier
		RefundsAndCancellations float64 `json:"refunds_and_cancellations"`

		// The commission withheld when the product was sold and refunded when the product was returned
		SaleCommission float64 `json:"sale_commission"`

		// The additional services cost that are not directly related to deliveries and returns.
		// For example, promotion or product placement
		ServicesAmount float64 `json:"services_amount"`
	} `json:"result"`
}

// Returns total sums for transactions for specified period
func (c Finance) GetTotalTransactionsSum(params *GetTotalTransactionsSumParams) (*GetTotalTransactionsSumResponse, error) {
	url := "/v3/finance/transaction/totals"

	resp := &GetTotalTransactionsSumResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListTransactionsParams struct{
	// Filter
	Filter ListTransactionsFilter `json:"filter"`

	// Number of the page returned in the request
	Page int64 `json:"page"`

	// Number of items on the page
	PageSize int64 `json:"page_size"`
}

type ListTransactionsFilter struct{
	// Filter by date
	Date ListTransactionsFilterDate `json:"date"`

	// Operation type
	OperationType string `json:"operation_type"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Transaction type
	TransactionType string `json:"transaction_type"`
}

type ListTransactionsFilterDate struct{
	// Period start.
	// 
	// Format: YYYY-MM-DDTHH:mm:ss.sssZ.
	// Example: 2019-11-25T10:43:06.51
	From time.Time `json:"from"`

	// Period end.
	// 
	// Format: YYYY-MM-DDTHH:mm:ss.sssZ.
	// Example: 2019-11-25T10:43:06.51
	To time.Time `json:"to"`
}

type ListTransactionsResponse struct{
	core.CommonResponse

	// Method result
	Result struct{
		// Transactions infromation
		Operations []struct{
			// Cost of the products with seller's discounts applied
			AccrualsForSale float64 `json:"accruals_for_sale"`

			// Total transaction sum
			Amount float64 `json:"amount"`

			// Delivery cost for charges by rates that were in effect until February 1, 2021, and for charges for bulky products
			DeliveryCharge float64 `json:"delivery_charge"`

			// Product information
			Items []struct{
				// Product name
				Name string `json:"name"`

				// Product identifier in the Ozon system, SKU
				SKU int64 `json:"sku"`
			} `json:"items"`

			// Operation date
			OperationDate string `json:"operation_date"`

			// Operation identifier
			OperationId int64 `json:"operation_id"`

			// Operation type
			OperationType string `json:"operation_type"`

			// Operation type name
			OperationTypeName string `json:"operation_type_name"`

			// Shipment information
			Posting struct{
				// Delivery scheme:
				//   - FBO — delivery to Ozon warehouse
				//   - FBS — delivery from seller's warehouse
				//   - RFBS — delivery service of seller's choice
				//   - Crossborder — delivery from abroad
				DeliverySchema string `json:"delivery_schema"`

				// Date the product was accepted for processing
				OrderDate string `json:"order_date"`

				// Shipment number
				PostingNumber string `json:"posting_number"`

				// Warehouse identifier
				WarehouseId int64 `json:"warehouse_id"`
			} `json:"posting"`

			// Returns and cancellation cost for charges by rates that were in effect until February 1, 2021, and for charges for bulky products
			ReturnDeliveryCharge float64 `json:"return_delivery_charge"`

			// Sales commission or sales commission refund
			SaleCommission float64 `json:"sale_commission"`

			// Additional services
			Services []struct{
				// Service name
				Name string `json:"name"`

				// Price
				Price float64 `json:"price"`
			} `json:"services"`

			// Transaction type
			Type string `json:"type"`
		} `json:"operations"`

		// Number of pages
		PageCount int64 `json:"page_count"`

		// Number of products
		RowCount int64 `json:"row_count"`
	} `json:"result"`
}

// Returns detailed information on all accruals. The maximum period for which you can get information in one request is 1 month.
// 
// If you don't specify the posting_number in request, the response contains all shipments for the specified period or shipments of a certain type
func (c Finance) ListTransactions(params *ListTransactionsParams) (*ListTransactionsResponse, error) {
	url := "/v3/finance/transaction/list"

	resp := &ListTransactionsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
