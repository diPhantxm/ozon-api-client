package ozon

import (
	"context"
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Reports struct {
	client *core.Client
}

type GetReportsListParams struct {
	// Page number
	Page int32 `json:"page"`

	// The number of values on the page:
	//   - default value is 100,
	//   - maximum value is 1000
	PageSize int32 `json:"page_size"`

	// Default: "ALL"
	// Report type:
	//   - ALL — all reports,
	//   - SELLER_PRODUCTS — products report,,
	//   - SELLER_TRANSACTIONS — transactions report,
	//   - SELLER_PRODUCT_PRICES — product prices report,
	//   - SELLER_STOCK — stocks report,
	//   - SELLER_PRODUCT_MOVEMENT — products movement report,
	//   - SELLER_RETURNS — returns report,
	//   - SELLER_POSTINGS — shipments report,
	//   - SELLER_FINANCE — financial report
	ReportType string `json:"report_type" default:"ALL"`
}

type GetReportsListResponse struct {
	core.CommonResponse

	// Method result
	Result GetReportsListResult `json:"result"`
}

type GetReportsListResult struct {
	// Array with generated reports
	Reports []GetReportsListResultReport `json:"reports"`

	// Total number of reports
	Total int32 `json:"total"`
}

type GetReportsListResultReport struct {
	// Unique report identifier
	Code string `json:"code"`

	// Report creation date
	CreatedAt time.Time `json:"created_at"`

	// Error code when generating the report
	Error string `json:"error"`

	// Link to CSV file
	File string `json:"file"`

	// Array with the filters specified when the seller created the report
	Params map[string]string `json:"params"`

	// Report type:
	//   - SELLER_PRODUCTS — products report,
	//   - SELLER_TRANSACTIONS — transactions report,
	//   - SELLER_PRODUCT_PRICES — product prices report,
	//   - SELLER_STOCK — stocks report,
	//   - SELLER_PRODUCT_MOVEMENT — products movement report,
	//   - SELLER_RETURNS — returns report,
	//   - SELLER_POSTINGS — shipments report,
	//   - SELLER_FINANCE — financial report
	ReportType string `json:"report_type"`

	// Report generation status
	//   - `success`
	//   - `failed`
	Status string `json:"status"`
}

// Returns the list of reports that have been generated before
func (c Reports) GetList(ctx context.Context, params *GetReportsListParams) (*GetReportsListResponse, error) {
	url := "/v1/report/list"

	resp := &GetReportsListResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetReportDetailsParams struct {
	// Unique report identifier
	Code string `json:"code"`
}

type GetReportDetailsResponse struct {
	core.CommonResponse

	// Report details
	Result GetReportDetailResult `json:"result"`
}

type GetReportDetailResult struct {
	// Unique report identifier
	Code string `json:"code"`

	// Report creation date
	CreatedAt time.Time `json:"created_at"`

	// Error code when generating the report
	Error string `json:"error"`

	// Link to CSV file
	File string `json:"file"`

	// Array with the filters specified when the seller created the report
	Params map[string]string `json:"params"`

	// Report type:
	//   - SELLER_PRODUCTS — products report,
	//   - SELLER_TRANSACTIONS — transactions report,
	//   - SELLER_PRODUCT_PRICES — product prices report,
	//   - SELLER_STOCK — stocks report,
	//   - SELLER_PRODUCT_MOVEMENT — products movement report,
	//   - SELLER_RETURNS — returns report,
	//   - SELLER_POSTINGS — shipments report,
	//   - SELLER_FINANCE — financial report
	ReportType string `json:"report_type"`

	// Report generation status:
	//   - success
	//   - failed
	Status string `json:"status"`
}

// Returns information about a created report by its identifier
func (c Reports) GetReportDetails(ctx context.Context, params *GetReportDetailsParams) (*GetReportDetailsResponse, error) {
	url := "/v1/report/info"

	resp := &GetReportDetailsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetFinancialReportParams struct {
	// Report generation period
	Date GetFinancialReportDatePeriod `json:"date"`

	// Number of the page returned in the request
	Page int64 `json:"page"`

	// true, если нужно добавить дополнительные параметры в ответ
	WithDetails bool `json:"with_details"`

	// Number of items on the page
	PageSize int64 `json:"page_size"`
}

type GetFinancialReportDatePeriod struct {
	// Date from which the report is calculated
	From time.Time `json:"from"`

	// Date up to which the report is calculated
	To time.Time `json:"to"`
}

type GetFinancialReportResponse struct {
	core.CommonResponse

	// Method result
	Result GetFinancialResultResult `json:"result"`
}

type GetFinancialResultResult struct {
	// Reports list
	CashFlows []GetFinancialResultResultCashflow `json:"cash_flows"`

	// Detailed info
	Details GetFinancialResultResultDetail `json:"details"`

	// Number of pages with reports
	PageCount int64 `json:"page_count"`
}

type GetFinancialResultResultCashflow struct {
	// Period data
	Period GetFinancialResultResultCashflowPeriod `json:"period"`

	// Sum of sold products prices
	OrdersAmount float64 `json:"order_amount"`

	// Sum of returned products prices
	ReturnsAmount float64 `json:"returns_amount"`

	// Ozon sales commission
	CommissionAmount float64 `json:"commission_amount"`

	// Additional services cost
	ServicesAmount float64 `json:"services_amount"`

	// Logistic services cost
	ItemDeliveryAndReturnAmount float64 `json:"item_delivery_and_return_amount"`

	// Code of the currency used to calculate the commissions
	CurrencyCode string `json:"currency_code"`
}

type GetFinancialResultResultCashflowPeriod struct {
	// Period identifier
	Id int64 `json:"id"`

	// Period start
	Begin time.Time `json:"begin"`

	// Period end
	End time.Time `json:"end"`
}

type GetFinancialResultResultDetail struct {
	// Balance on the beginning of period
	BeginBalanceAmount float64 `json:"begin_balance_amount"`

	// Orders
	Delivery GetFinancialResultResultDetailDelivery `json:"delivery"`

	// Amount to be paid for the period
	InvoiceTransfer float64 `json:"invoice_transfer"`

	// Transfer under loan agreements
	Loan float64 `json:"loan"`

	// Paid for the period
	Payments []GetFinancialResultResultDetailPayment `json:"payments"`

	// Period data
	Period GetFinancialResultResultDetailPeriod `json:"period"`

	// Returns and cancellations
	Return GetFinancialResultResultDetailReturn `json:"return"`

	// rFBS transfers
	RFBS GetFinancialResultResultDetailRFBS `json:"rfbs"`

	// Services
	Services GetFinancialResultResultDetailService `json:"services"`

	// Compensation and other accruals
	Others GetFinancialResultResultDetailOthers `json:"others"`

	// Balance at the end of the period
	EndBalanceAmount float64 `json:"end_balance_amount"`
}

type GetFinancialResultResultDetailDelivery struct {
	// Total amount
	Total float64 `json:"total"`

	// Amount for which products were purchased, including commission fees
	Amount float64 `json:"amount"`

	// Processing and delivery fees
	DeliveryServices GetFinancialResultResultDetailDeliveryServices `json:"delivery_services"`
}

type GetFinancialResultResultDetailDeliveryServices struct {
	// Total amount
	Total float64 `json:"total"`

	// Details
	Items []GetFinancialResultResultDetailDeliveryServicesItem `json:"items"`
}

type GetFinancialResultResultDetailDeliveryServicesItem struct {
	// Operation name
	Name DetailsDeliveryItemName `json:"name"`

	// Amount by operation
	Price float64 `json:"price"`
}

type GetFinancialResultResultDetailPayment struct {
	// Currency
	CurrencyCode string `json:"currency_code"`

	// Payment amount
	Payment float64 `json:"payment"`
}

type GetFinancialResultResultDetailPeriod struct {
	// Period start
	Begin time.Time `json:"begin"`

	// Period end
	End time.Time `json:"end"`

	// Period identifier
	Id int64 `json:"id"`
}

type GetFinancialResultResultDetailReturn struct {
	// Total amount
	Total float64 `json:"total"`

	// Amount of returns received, including commission fees
	Amount float64 `json:"amount"`

	// Returns and cancellation fees
	ReturnServices GetFinancialResultResultDetailReturnServices `json:"return_services"`
}

type GetFinancialResultResultDetailReturnServices struct {
	// Total amount
	Total float64 `json:"total"`

	// Details
	Items []GetFinancialResultResultDetailReturnServicesItem `json:"items"`
}

type GetFinancialResultResultDetailReturnServicesItem struct {
	// Operation name
	Name DetailsReturnServiceName `json:"name"`

	// Amount by operation
	Price float64 `json:"price"`
}

type GetFinancialResultResultDetailRFBS struct {
	// Total amount
	Total float64 `json:"total"`

	// Transfers from customers
	TransferDelivery float64 `json:"transfer_delivery"`

	// Return of transfers to customers
	TransferDeliveryReturn float64 `json:"transfer_delivery_return"`

	// Compensation of delivery fees
	CompensationDeliveryReturn float64 `json:"compensation_delivery_return"`

	// Transfers of partial refunds to customers
	PartialCompensation float64 `json:"partial_compensation"`

	// Compensation of partial refunds
	PartialCompensationReturn float64 `json:"partial_compensation_return"`
}

type GetFinancialResultResultDetailService struct {
	// Total amount
	Total float64 `json:"total"`

	// Details
	Items []GetFinancialResultResultDetailServiceItem `json:"items"`
}

type GetFinancialResultResultDetailServiceItem struct {
	// Operation name
	Name DetailsServiceItemName `json:"name"`

	// Amount by operation
	Price float64 `json:"price"`
}

type GetFinancialResultResultDetailOthers struct {
	// Total amount
	Total float64 `json:"total"`

	// Details
	Items []GetFinancialResultResultDetailOthersItem `json:"items"`
}

type GetFinancialResultResultDetailOthersItem struct {
	// Operation name
	Name DetailsOtherItemName `json:"name"`

	// Amount by operation
	Price float64 `json:"price"`
}

// Returns information about a created report by its identifier
func (c Reports) GetFinancial(ctx context.Context, params *GetFinancialReportParams) (*GetFinancialReportResponse, error) {
	url := "/v1/finance/cash-flow-statement/list"

	resp := &GetFinancialReportResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductsReportParams struct {
	// Default: "DEFAULT"
	// Response language:
	//   - RU — Russian
	//   - EN — English
	Language string `json:"language" default:"DEFAULT"`

	// Product identifier in the seller's system
	OfferId []string `json:"offer_id"`

	// Search by record content, checks for availability
	Search string `json:"search"`

	// Product identifier in the Ozon system, SKU
	SKU []int64 `json:"sku"`

	// Default: "ALL"
	// Filter by product visibility
	Visibility string `json:"visibility" default:"ALL"`
}

type GetProductsReportResponse struct {
	core.CommonResponse

	// Method result
	Result GetProductsReportResult `json:"result"`
}

type GetProductsReportResult struct {
	// Unique report identifier
	Code string `json:"code"`
}

// Method for getting a report with products data. For example, Ozon ID, number of products, prices, status
func (c Reports) GetProducts(ctx context.Context, params *GetProductsReportParams) (*GetProductsReportResponse, error) {
	url := "/v1/report/products/create"

	resp := &GetProductsReportResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetStocksReportParams struct {
	// Default: "DEFAULT"
	// Response language:
	//   - RU — Russian
	//   - EN — English
	Language string `json:"language" default:"DEFAULT"`
}

type GetStocksReportResponse struct {
	core.CommonResponse

	// Method result
	Result GetStocksReportResult `json:"result"`
}

type GetStocksReportResult struct {
	// Unique report identifier
	Code string `json:"code"`
}

// Report with information about the number of available and reserved products in stock
func (c Reports) GetStocks(ctx context.Context, params *GetStocksReportParams) (*GetStocksReportResponse, error) {
	url := "/v1/report/stock/create"

	resp := &GetStocksReportResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProductsMovementReportParams struct {
	// Date from which the data will be in the report
	DateFrom time.Time `json:"date_from"`

	// Date up to which the data will be in the report
	DateTo time.Time `json:"date_to"`

	// Default: "DEFAULT"
	// Response language:
	//   - RU — Russian
	//   - EN — English
	Language string `json:"language" default:"DEFAULT"`
}

type GetProductsMovementReportResponse struct {
	core.CommonResponse

	// Method result
	Result GetProductsMovementReportResult `json:"result"`
}

type GetProductsMovementReportResult struct {
	// Unique report identifier
	Code string `json:"code"`
}

// Report with complete information on products, as well as the number of products with statuses:
//   - products with defects or in inventory,
//   - products in transit between the fulfillment centers,
//   - products in delivery,
//   - products to be sold
func (c Reports) GetProductsMovement(ctx context.Context, params *GetProductsMovementReportParams) (*GetProductsMovementReportResponse, error) {
	url := "/v1/report/products/movement/create"

	resp := &GetProductsMovementReportResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetReturnsReportParams struct {
	// Filter
	Filter GetReturnsReportsFilter `json:"filter"`

	// Default: "DEFAULT"
	// Response language:
	//   - RU — Russian
	//   - EN — English
	Language string `json:"language" default:"DEFAULT"`
}

type GetReturnsReportsFilter struct {
	// Order delivery scheme: fbs — delivery from seller's warehouse
	DeliverySchema string `json:"delivery_schema"`

	// Order identifier
	OrderId int64 `json:"order_id"`

	// Order status
	Status string `json:"status"`
}

type GetReturnsReportResponse struct {
	core.CommonResponse

	// Method result
	Result GetReturnReportResult `json:"result"`
}

type GetReturnReportResult struct {
	// Unique report identifier
	Code string `json:"code"`
}

// The report contains information about returned products that were accepted from the customer, ready for pickup, or delivered to the seller.
//
// The method is only suitable for orders shipped from the seller's warehouse
func (c Reports) GetReturns(ctx context.Context, params *GetReturnsReportParams) (*GetReturnsReportResponse, error) {
	url := "/v1/report/returns/create"

	resp := &GetReturnsReportResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetShipmentReportParams struct {
	// Filter
	Filter GetShipmentReportFilter `json:"filter"`

	// Default: "DEFAULT"
	// Response language:
	//   - RU — Russian
	//   - EN — English
	Language string `json:"language" default:"DEFAULT"`
}

type GetShipmentReportFilter struct {
	// Cancellation reason identifier
	CancelReasonId []int64 `json:"cancel_reason_id"`

	// Work scheme: FBO or FBS.
	//
	// To get an FBO scheme report, pass fbo in this parameter. For an FBS scheme report pass fbs
	DeliverySchema []string `json:"delivery_schema"`

	// Product identifier
	OfferId string `json:"offer_id"`

	// Order processing start date and time
	ProcessedAtFrom time.Time `json:"processed_at_from"`

	// Time when the order appeared in your personal account
	ProcessedAtTo time.Time `json:"processed_at_to"`

	// Product identifier in the Ozon system, SKU
	SKU []int64 `json:"sku"`

	// Status text
	StatusAlias []string `json:"status_alias"`

	// Numerical status
	Statuses []int64 `json:"statused"`

	// Product name
	Title string `json:"title"`
}

type GetShipmentReportResponse struct {
	core.CommonResponse

	// Method result
	Result GetShipmentReportResult `json:"result"`
}

type GetShipmentReportResult struct {
	// Unique report identifier
	Code string `json:"code"`
}

// Shipment report with orders details:
//   - order statuses
//   - processing start date
//   - order numbers
//   - shipment numbers
//   - shipment costs
//   - shipments contents
func (c Reports) GetShipment(ctx context.Context, params *GetShipmentReportParams) (*GetShipmentReportResponse, error) {
	url := "/v1/report/postings/create"

	resp := &GetShipmentReportResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type IssueOnDiscountedProductsResponse struct {
	core.CommonResponse

	// Unique report identifier
	Code string `json:"code"`
}

// Generates a report on discounted products in Ozon warehouses.
// For example, Ozon can discount a product due to damage when delivering.
//
// Returns report identifier. To get the report, send the identifier in the request body of a method `/v1/report/discounted/info`
func (c Reports) IssueOnDiscountedProducts(ctx context.Context) (*IssueOnDiscountedProductsResponse, error) {
	url := "/v1/report/discounted/create"

	resp := &IssueOnDiscountedProductsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ReportOnDiscountedProductsParams struct {
	// Unique report identifier
	Code string `json:"code"`
}

type ReportOnDiscountedProductsResponse struct {
	core.CommonResponse

	// Report information
	Report ReportonDiscountedProductsReport `json:"report"`
}

type ReportonDiscountedProductsReport struct {
	// Report creation date
	CreatedAt time.Time `json:"created_at"`

	// Link to report file
	File string `json:"file"`

	// Report status:
	//   - success — created
	//   - pending — waiting to be processed
	//   - processing — processed
	//   - failed — generation error
	Status string `json:"status"`

	// Report generation error code
	Error string `json:"error"`
}

// By report identifier, returns information about the report generated earlier
func (c Reports) ReportOnDiscountedProducts(ctx context.Context, params *ReportOnDiscountedProductsParams) (*ReportOnDiscountedProductsResponse, error) {
	url := "/v1/report/discounted/info"

	resp := &ReportOnDiscountedProductsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

// By report identifier, returns information about the report generated earlier
func (c Reports) ListReportsOnDiscountedProducts(ctx context.Context) (*ReportOnDiscountedProductsResponse, error) {
	url := "/v1/report/discounted/list"

	resp := &ReportOnDiscountedProductsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
