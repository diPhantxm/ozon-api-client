package ozon

import (
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type Brands struct {
	client *core.Client
}

type GetCertefiedBrandsListParams struct {
	// Number of the page returned in the request
	Page int `json:"page"`

	// Number of elements on the page
	PageSize int `json:"page_size"`
}

type GetCertefiedBrandsListResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Certified brands details
		BrandCertification []struct {
			// Brand ID
			BrandID int `json:"brand_id"`

			// Brand name
			BrandName string `json:"brand_name"`

			// Indication that the certificate is required
			HasCertificate bool `json:"has_certificate"`
		} `json:"brand_certification"`

		// Total number of brands
		Total int `json:"total"`
	} `json:"result"`
}

// Get a list of certified brands
func (c Brands) GetCertefiedBrandsList(params *GetCertefiedBrandsListParams) (*GetCertefiedBrandsListResponse, error) {
	url := "/v1/brand/company-certification/list"

	resp := &GetCertefiedBrandsListResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
