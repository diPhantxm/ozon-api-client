package ozon

import (
	"net/http"

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
}

type CreateUpdateProformaLinkResponse struct {
	core.CommonResponse

	// Method result
	Result bool `json:"result"`
}

// Create or edit proforma invoice link for VAT refund to Turkey sellers
func (c Invoices) CreateUpdate(params *CreateUpdateProformaLinkParams) (*CreateUpdateProformaLinkResponse, error) {
	url := "/v1/invoice/create-or-update"

	resp := &CreateUpdateProformaLinkResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
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
	Result struct {
		// Proforma invoice link
		FileURL string `json:"file_url"`
	} `json:"result"`
}

// Get a proforma invoice link
func (c Invoices) Get(params *GetProformaLinkParams) (*GetProformaLinkResponse, error) {
	url := "/v1/invoice/get"

	resp := &GetProformaLinkResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
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

func (c Invoices) Delete(params *DeleteProformaLinkParams) (*DeleteProformaLinkResponse, error) {
	url := "/v1/invoice/delete"

	resp := &DeleteProformaLinkResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
