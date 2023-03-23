package ozon

import (
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type Brands struct {
	client *core.Client
}

type ListCertifiedBrandsParams struct {
	// Number of the page returned in the request
	Page int32 `json:"page"`

	// Number of elements on the page
	PageSize int32 `json:"page_size"`
}

type ListCertifiedBrandsResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Certified brands details
		BrandCertification []struct {
			// Brand name
			BrandName string `json:"brand_name"`

			// Indication that the certificate is required:
			//   - true if the certificate is required;
			//   - false if not
			HasCertificate bool `json:"has_certificate"`
		} `json:"brand_certification"`

		// Total number of brands
		Total int64 `json:"total"`
	} `json:"result"`
}

// List of certified brands
func (c Brands) List(params *ListCertifiedBrandsParams) (*ListCertifiedBrandsResponse, error) {
	url := "/v1/brand/company-certification/list"

	resp := &ListCertifiedBrandsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
