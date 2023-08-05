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

	// Proforma invoice link
	URL string `json:"url"`

	// Invoice HS-code. Pass a number up to 12 characters long
	HSCode string `json:"hs_code"`

	// Invoice date
	Date time.Time `json:"date"`

	// Invoice number. The number can contain letters and digits, maximum length is 50 characters
	Number string `json:"number"`

	// Cost stated in the invoice. The fractional part is separated by decimal point, up to two digits after the decimal poin
	Price float64 `json:"price"`

	// Invoice currency
	PriceCurrency InvoiceCurrency `json:"price_currency" default:"USD"`
}

type CreateUpdateProformaLinkResponse struct {
	core.CommonResponse

	// Method result
	Result bool `json:"result"`
}

// Create or edit proforma invoice link for VAT refund to Turkey sellers
func (c Invoices) CreateUpdate(ctx context.Context, params *CreateUpdateProformaLinkParams) (*CreateUpdateProformaLinkResponse, error) {
	url := "/v1/invoice/create-or-update"

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
	// Proforma invoice link
	FileURL string `json:"file_url"`
}

// Get a proforma invoice link
func (c Invoices) Get(ctx context.Context, params *GetProformaLinkParams) (*GetProformaLinkResponse, error) {
	url := "/v1/invoice/get"

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
