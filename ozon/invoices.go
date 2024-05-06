package ozon

import (
	"context"
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Invoices struct {
	client *core.Client
}

type CreateUpdateProformaLinkParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Invoice link. Use the `v1/invoice/file/upload` method to create a link
	URL string `json:"url"`

	// Product HS-codes
	HSCodes []CreateUpdateProformaLinkHSCode `json:"hs_codes"`

	// Invoice date
	Date time.Time `json:"date"`

	// Invoice number. The number can contain letters and digits, maximum length is 50 characters
	Number string `json:"number"`

	// Cost stated in the invoice. The fractional part is separated by decimal point, up to two digits after the decimal point
	Price float64 `json:"price"`

	// Invoice currency
	PriceCurrency InvoiceCurrency `json:"price_currency" default:"USD"`
}

type CreateUpdateProformaLinkHSCode struct {
	// Product HS code
	Code string `json:"code"`

	// Product identifier in the Ozon system, SKU
	SKU string `json:"sku"`
}

type CreateUpdateProformaLinkResponse struct {
	core.CommonResponse

	// Method result
	Result bool `json:"result"`
}

// Create or edit an invoice for VAT refund to Turkey sellers
func (c Invoices) CreateUpdate(ctx context.Context, params *CreateUpdateProformaLinkParams) (*CreateUpdateProformaLinkResponse, error) {
	url := "/v2/invoice/create-or-update"

	resp := &CreateUpdateProformaLinkResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetProformaLinkParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`
}

type GetProformaLinkResponse struct {
	core.CommonResponse

	// Method result
	Result GetProformaLinkResult `json:"result"`
}

type GetProformaLinkResult struct {
	// Invoice uploading date
	Date time.Time `json:"date"`

	// Invoice link
	FileURL string `json:"file_url"`

	// Product HS-codes
	HSCodes []CreateUpdateProformaLinkHSCode `json:"hs_codes"`

	// Invoice number
	Number string `json:"number"`

	// Cost stated in the invoice.
	// The fractional part is separated by decimal point,
	// up to two digits after the decimal point.
	//
	// Example: 199.99
	Price float64 `json:"price"`

	// Invoice currency
	PriceCurrency InvoiceCurrency `json:"price_currency"`
}

// Get a proforma invoice link
func (c Invoices) Get(ctx context.Context, params *GetProformaLinkParams) (*GetProformaLinkResponse, error) {
	url := "/v2/invoice/get"

	resp := &GetProformaLinkResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type DeleteProformaLinkParams struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`
}

type DeleteProformaLinkResponse struct {
	core.CommonResponse

	// Method result
	Result bool `json:"result"`
}

func (c Invoices) Delete(ctx context.Context, params *DeleteProformaLinkParams) (*DeleteProformaLinkResponse, error) {
	url := "/v1/invoice/delete"

	resp := &DeleteProformaLinkResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UploadInvoiceParams struct {
	// Base64 encoded invoice
	Content string `json:"base64_content"`

	// Shipment number
	PostingNumber string `json:"posting_number"`
}

type UploadInvoiceResponse struct {
	core.CommonResponse

	// Link to invoice
	URL string `json:"url"`
}

// Available file types: JPEG and PDF. Maximum file size: 10 MB
func (c Invoices) Upload(ctx context.Context, params *UploadInvoiceParams) (*UploadInvoiceResponse, error) {
	url := "/v1/invoice/file/upload"

	resp := &UploadInvoiceResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
