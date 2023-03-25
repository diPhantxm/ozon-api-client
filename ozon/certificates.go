package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Certificates struct {
	client *core.Client
}

type ListOfAccordanceTypesResponse struct {
	core.CommonResponse

	// Accordance types
	Result struct {
		// Main accordance types
		Base []struct {
			// Accordance type code
			Code string `json:"code"`

			// Accordance type description
			Title string `json:"title"`
		} `json:"base"`

		// Main accordance types related to dangerous products
		Hazard []struct {
			// Accordance type code
			Code string `json:"code"`

			// Accordance type description
			Title string `json:"title"`
		} `json:"hazard"`
	} `json:"result"`
}

// List of accordance types (version 2)
func (c Certificates) ListOfAccordanceTypes() (*ListOfAccordanceTypesResponse, error) {
	url := "/v2/product/certificate/accordance-types/list"

	resp := &ListOfAccordanceTypesResponse{}

	response, err := c.client.Request(http.MethodGet, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type DirectoryOfDocumentTypesResponse struct {
	core.CommonResponse

	// List of certificate types and names
	Result []struct {
		// Certificate name
		Name string `json:"name"`

		// Certificate type
		Value string `json:"value"`
	} `json:"result"`
}

// Directory of document types
func (c Certificates) DirectoryOfDocumentTypes() (*DirectoryOfDocumentTypesResponse, error) {
	url := "/v1/product/certificate/types"

	resp := &DirectoryOfDocumentTypesResponse{}

	response, err := c.client.Request(http.MethodGet, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListOfCertifiedCategoriesParams struct {
	// Number of the page returned in the query
	Page int32 `json:"page"`

	// Number of elements on the page
	PageSize int32 `json:"page_size"`
}

type ListOfCertifiedCategoriesResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Certified categories details
		Certification []struct {
			// Category name
			CategoryName string `json:"category_name"`

			// Indication of a mandatory category
			IsRequired bool `json:"is_required"`
		} `json:"certification"`

		// Total number of categories
		Total int64 `json:"total"`
	} `json:"reult"`
}

// List of certified categories
func (c Certificates) ListOfCertifiedCategories(params *ListOfCertifiedCategoriesParams) (*ListOfCertifiedCategoriesResponse, error) {
	url := "/v1/product/certificate/types"

	resp := &ListOfCertifiedCategoriesResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type LinkCertificateToProductParams struct {
	// Certificate identifier that was assigned when it was uploaded
	CertificateId int64 `json:"certificate_id"`

	// An array of product identifiers that this certificate applies to
	ProductId []int64 `json:"product_id"`
}

type LinkCertificateToProductResponse struct {
	core.CommonResponse

	// The result of processing the request. true if the request was executed without errors
	Result bool `json:"result"`
}

// Link the certificate to the product
func (c Certificates) LinkToProduct(params *LinkCertificateToProductParams) (*LinkCertificateToProductResponse, error) {
	url := "/v1/product/certificate/bind"

	resp := &LinkCertificateToProductResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type DeleteCertificateParams struct {
	// Certificate identifier
	CertificateId int32 `json:"certificate_id"`
}

type DeleteCertificateResponse struct {
	core.CommonResponse

	// Result of deleting the certificate
	Result struct {
		// Indication that a certificate has been deleted:
		//   - true — deleted
		//   - false — not deleted
		IsDelete bool `json:"is_delete"`

		// Description of errors during certificate deletion
		ErrorMessage string `json:"error_message"`
	} `json:"result"`
}

// Delete certificate
func (c Certificates) Delete(params *DeleteCertificateParams) (*DeleteCertificateResponse, error) {
	url := "/v1/product/certificate/delete"

	resp := &DeleteCertificateResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetCertificateInfoParams struct {
	// Certificate identifier
	CertificateNumber string `json:"certificate_number"`
}

type GetCertificateInfoResponse struct {
	core.CommonResponse

	// Certificate information
	Result struct {
		// Identifier
		CertificateId int32 `json:"certificate_id"`

		// Number
		CertificateNumber string `json:"certificate_number"`

		// Name
		CertificateName string `json:"certificate_name"`

		// Type
		TypeCode string `json:"type_code"`

		// Status
		StatusCode string `json:"status_code"`

		// Accordance type
		AccordanceTypeCode string `json:"accordance_type_code"`

		// Certificate rejection reason
		RejectionReasonCode string `json:"rejectio_reason_code"`

		// Moderator's comment
		VerificationComment string `json:"verification_comment"`

		// Issue date
		IssueDate time.Time `json:"issue_date"`

		// Expire date
		ExpireDate time.Time `json:"expire_date"`

		// Number of products associated with a certificate
		ProductsCount int32 `json:"products_count"`
	} `json:"result"`
}

// Certificate information
func (c Certificates) GetInfo(params *GetCertificateInfoParams) (*GetCertificateInfoResponse, error) {
	url := "/v1/product/certificate/info"

	resp := &GetCertificateInfoResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListCertificatesParams struct {
	// Product identifier associated with the certificate.
	// Pass the parameter if you need certificates that certain products are associated with
	OfferId string `json:"offer_id"`

	// Certificate status. Pass the parameter if you need certificates with a certain status
	Status string `json:"status"`

	// Certificate type. Pass the parameter if you need certificates with a certain type
	Type string `json:"type"`

	// Page from which the list should be displayed. The minimum value is 1
	Page int32 `json:"page"`

	// Number of objects on the page. The value is from 1 to 1000
	PageSize int32 `json:"page_size"`
}

type ListCertificatesResponse struct {
	core.CommonResponse

	// Certificates
	Result struct {
		// Сertificate information
		Certificates []struct {
			// Identifier
			CertificateId int32 `json:"certificate_id"`

			// Number
			CertificateNumber string `json:"certificate_number"`

			// Name
			CertificateName string `json:"certificate_name"`

			// Type
			TypeCode string `json:"type"`

			// Status
			StatusCode string `json:"status_code"`

			// Accordance type
			AccordanceTypecode string `json:"accordance_type_code"`

			// Certificate rejection reason
			RejectionReasonCode string `json:"rejectio_reason_code"`

			// Moderator's comment
			VerificationComment string `json:"verification_comment"`

			// Issue date
			IssueDate time.Time `json:"issue_data"`

			// Expire date
			ExpireDate time.Time `json:"expire_date"`

			// Number of products associated with a certificate
			ProductsCount int32 `json:"products_count"`
		} `json:"certificates"`

		// Number of pages
		PageCount int32 `json:"page_count"`
	} `json:"result"`
}

// Certificates list
func (c Certificates) List(params *ListCertificatesParams) (*ListCertificatesResponse, error) {
	url := "/v1/product/certificate/list"

	resp := &ListCertificatesResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ProductStatusesResponse struct {
	core.CommonResponse

	// Product statuses
	Result []struct {
		// Product status code when linking it to the certificate
		Code string `json:"code"`

		// Status description
		Name string `json:"name"`
	} `json:"result"`
}

func (c Certificates) ProductStatuses() (*ProductStatusesResponse, error) {
	url := "/v1/product/certificate/list"

	resp := &ProductStatusesResponse{}

	response, err := c.client.Request(http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListProductsForCertificateParams struct {
	// Certificate identifier
	CertificateId int32 `json:"certificate_id"`

	// Status of the product verification when binding to a certificate
	ProductStatusCode string `json:"product_status_code"`

	// Page from which the list should be displayed. The minimum value is 1
	Page int32 `json:"page"`

	// Number of objects on the page. The value is from 1 to 1000
	PageSize int32 `json:"page_size"`
}

type ListProductsForCertificateResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// List of products
		Items []struct {
			// Product identifier
			ProductId int64 `json:"product_id"`

			// Status of the product processing when binding to a certificate
			ProductStatusCode string `json:"product_status_code"`
		} `json:"items"`

		// Number of products found
		Count int64 `json:"count"`
	} `json:"result"`
}

// A method for getting a list of possible statuses of products when binding them to a certificate
func (c Certificates) ListProductsForCertificate(params *ListProductsForCertificateParams) (*ListProductsForCertificateResponse, error) {
	url := "/v1/product/certificate/products/list"

	resp := &ListProductsForCertificateResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UnlinkFromProductParams struct {
	// Certificate identifier
	CertificateId int32 `json:"certificate_id"`

	// List of product identifiers that you want to unbind from a certificate
	ProductId []int64 `json:"product_id"`
}

type UnlinkFromProductResponse struct {
	core.CommonResponse

	// Method result
	Result []struct {
		// Error message when unbinding a product
		Error string `json:"error"`

		// Product identifier
		ProductId int64 `json:"product_id"`

		// Indication that the product was unbound from a certificate:
		//   - true — it was unbound,
		//   - false — it is still bound
		Updated bool `json:"updated"`
	} `json:"result"`
}

// Unbind products from a certificate
func (c Certificates) UnlinkFromProduct(params *UnlinkFromProductParams) (*UnlinkFromProductResponse, error) {
	url := "/v1/product/certificate/unbind"

	resp := &UnlinkFromProductResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type PossibleRejectReasonsResponse struct {
	core.CommonResponse

	// Certificate rejection reasons
	Result []struct {
		// Сode of a certificate rejection reason
		Code string `json:"code"`

		// Description of a certificate rejection reason
		Name string `json:"name"`
	} `json:"result"`
}

// Possible certificate rejection reasons
func (c Certificates) PossibleRejectReasons() (*PossibleRejectReasonsResponse, error) {
	url := "/v1/product/certificate/rejection_reasons/list"

	resp := &PossibleRejectReasonsResponse{}

	response, err := c.client.Request(http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type PossibleStatusesResponse struct {
	core.CommonResponse

	// Possible certificate statuses
	Result []struct {
		// Certificate status code
		Code string `json:"code"`

		// Status description
		Name string `json:"name"`
	} `json:"result"`
}

func (c Certificates) PossibleStatuses() (*PossibleStatusesResponse, error) {
	url := "/v1/product/certificate/status/list"

	resp := &PossibleStatusesResponse{}

	response, err := c.client.Request(http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type AddCertificatesForProductsParams struct {
	// Array of certificates for the product. Valid extensions are jpg, jpeg, png, pdf
	Files []byte `json:"files"`

	// Certificate name. No more than 100 characters
	Name string `json:"name"`

	// Certificate number. No more than 100 characters
	Number string `json:"number"`

	// Certificate type. To get the list of types, use the GET `/v1/product/certificate/types` method
	TypeCode string `json:"type_code"`

	// Accordance type. To get the list of types, use the GET `/v1/product/certificate/accordance-types` method
	AccordanceTypeCode string `json:"accordance_type_code"`

	// Issue date of the certificate
	IssueDate time.Time `json:"issue_date"`

	// Expiration date of the certificate. Can be empty for permanent certificates
	ExpireDate time.Time `json:"expire_date"`
}

type AddCertificatesForProductsResponse struct {
	core.CommonResponse

	Id int `json:"id"`
}

// Adding certificates for products
func (c Certificates) AddForProducts(params *AddCertificatesForProductsParams) (*AddCertificatesForProductsResponse, error) {
	url := "/v1/product/certificate/create"

	resp := &AddCertificatesForProductsResponse{}

	response, err := c.client.Request(http.MethodPost, url, nil, resp, map[string]string{
		"Content-Type": "multipart/form-data",
	})
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
