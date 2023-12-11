package ozon

import (
	"context"
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
	Result ListOfAccordanceTypesResult `json:"result"`
}

type ListOfAccordanceTypesResult struct {
	// Main accordance types
	Base []ListOfAccordanceTypesResultBase `json:"base"`

	// Main accordance types related to dangerous products
	Hazard []ListOfAccordanceTypesResultHazard `json:"hazard"`
}

type ListOfAccordanceTypesResultBase struct {
	// Accordance type code
	Code string `json:"code"`

	// Accordance type description
	Title string `json:"title"`
}

type ListOfAccordanceTypesResultHazard struct {
	// Accordance type code
	Code string `json:"code"`

	// Accordance type description
	Title string `json:"title"`
}

// List of accordance types (version 2)
func (c Certificates) ListOfAccordanceTypes(ctx context.Context) (*ListOfAccordanceTypesResponse, error) {
	url := "/v2/product/certificate/accordance-types/list"

	resp := &ListOfAccordanceTypesResponse{}

	response, err := c.client.Request(ctx, http.MethodGet, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type DirectoryOfDocumentTypesResponse struct {
	core.CommonResponse

	// List of certificate types and names
	Result []DirectoryOfDocumentTypesResult `json:"result"`
}

type DirectoryOfDocumentTypesResult struct {
	// Certificate name
	Name string `json:"name"`

	// Certificate type
	Value string `json:"value"`
}

// Directory of document types
func (c Certificates) DirectoryOfDocumentTypes(ctx context.Context) (*DirectoryOfDocumentTypesResponse, error) {
	url := "/v1/product/certificate/types"

	resp := &DirectoryOfDocumentTypesResponse{}

	response, err := c.client.Request(ctx, http.MethodGet, url, nil, resp, nil)
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
	Result ListOfCertifiedCategoriesResult `json:"result"`
}

type ListOfCertifiedCategoriesResult struct {
	// Certified categories details
	Certification []ListOfCertifiedCategoriesResultCert `json:"certification"`

	// Total number of categories
	Total int64 `json:"total"`
}

type ListOfCertifiedCategoriesResultCert struct {
	// Category name
	CategoryName string `json:"category_name"`

	// Indication of a mandatory category
	IsRequired bool `json:"is_required"`
}

// List of certified categories
func (c Certificates) ListOfCertifiedCategories(ctx context.Context, params *ListOfCertifiedCategoriesParams) (*ListOfCertifiedCategoriesResponse, error) {
	url := "/v1/product/certificate/types"

	resp := &ListOfCertifiedCategoriesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
func (c Certificates) LinkToProduct(ctx context.Context, params *LinkCertificateToProductParams) (*LinkCertificateToProductResponse, error) {
	url := "/v1/product/certificate/bind"

	resp := &LinkCertificateToProductResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Result DeleteCertificateResult `json:"result"`
}

type DeleteCertificateResult struct {
	// Indication that a certificate has been deleted:
	//   - true — deleted
	//   - false — not deleted
	IsDelete bool `json:"is_delete"`

	// Description of errors during certificate deletion
	ErrorMessage string `json:"error_message"`
}

// Delete certificate
func (c Certificates) Delete(ctx context.Context, params *DeleteCertificateParams) (*DeleteCertificateResponse, error) {
	url := "/v1/product/certificate/delete"

	resp := &DeleteCertificateResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Result GetCertificateInfoResult `json:"result"`
}

type GetCertificateInfoResult struct {
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
	RejectionReasonCode string `json:"rejection_reason_code"`

	// Moderator's comment
	VerificationComment string `json:"verification_comment"`

	// Issue date
	IssueDate time.Time `json:"issue_date"`

	// Expire date
	ExpireDate time.Time `json:"expire_date"`

	// Number of products associated with a certificate
	ProductsCount int32 `json:"products_count"`
}

// Certificate information
func (c Certificates) GetInfo(ctx context.Context, params *GetCertificateInfoParams) (*GetCertificateInfoResponse, error) {
	url := "/v1/product/certificate/info"

	resp := &GetCertificateInfoResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Result ListCertificatesResult `json:"result"`
}

type ListCertificatesResult struct {
	// Сertificate information
	Certificates []ListCertificatesResultCert `json:"certificates"`

	// Number of pages
	PageCount int32 `json:"page_count"`
}

type ListCertificatesResultCert struct {
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
	AccordanceTypecode string `json:"accordance_type_code"`

	// Certificate rejection reason
	RejectionReasonCode string `json:"rejection_reason_code"`

	// Moderator's comment
	VerificationComment string `json:"verification_comment"`

	// Issue date
	IssueDate time.Time `json:"issue_date"`

	// Expire date
	ExpireDate time.Time `json:"expire_date"`

	// Number of products associated with a certificate
	ProductsCount int32 `json:"products_count"`
}

// Certificates list
func (c Certificates) List(ctx context.Context, params *ListCertificatesParams) (*ListCertificatesResponse, error) {
	url := "/v1/product/certificate/list"

	resp := &ListCertificatesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ProductStatusesResponse struct {
	core.CommonResponse

	// Product statuses
	Result []ProductStatusesResult `json:"result"`
}

type ProductStatusesResult struct {
	// Product status code when linking it to the certificate
	Code string `json:"code"`

	// Status description
	Name string `json:"name"`
}

func (c Certificates) ProductStatuses(ctx context.Context) (*ProductStatusesResponse, error) {
	url := "/v1/product/certificate/list"

	resp := &ProductStatusesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
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
	Result ListProductsForCertificateResult `json:"result"`
}

type ListProductsForCertificateResult struct {
	// List of products
	Items []struct {
		// Product identifier
		ProductId int64 `json:"product_id"`

		// Status of the product processing when binding to a certificate
		ProductStatusCode string `json:"product_status_code"`
	} `json:"items"`

	// Number of products found
	Count int64 `json:"count"`
}

// A method for getting a list of possible statuses of products when binding them to a certificate
func (c Certificates) ListProductsForCertificate(ctx context.Context, params *ListProductsForCertificateParams) (*ListProductsForCertificateResponse, error) {
	url := "/v1/product/certificate/products/list"

	resp := &ListProductsForCertificateResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
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
	Result []UnlinkFromProductResult `json:"result"`
}

type UnlinkFromProductResult struct {
	// Error message when unbinding a product
	Error string `json:"error"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Indication that the product was unbound from a certificate:
	//   - true — it was unbound,
	//   - false — it is still bound
	Updated bool `json:"updated"`
}

// Unbind products from a certificate
func (c Certificates) UnlinkFromProduct(ctx context.Context, params *UnlinkFromProductParams) (*UnlinkFromProductResponse, error) {
	url := "/v1/product/certificate/unbind"

	resp := &UnlinkFromProductResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type PossibleRejectReasonsResponse struct {
	core.CommonResponse

	// Certificate rejection reasons
	Result []PossibleRejectReasonsResult `json:"result"`
}

type PossibleRejectReasonsResult struct {
	// Сode of a certificate rejection reason
	Code string `json:"code"`

	// Description of a certificate rejection reason
	Name string `json:"name"`
}

// Possible certificate rejection reasons
func (c Certificates) PossibleRejectReasons(ctx context.Context) (*PossibleRejectReasonsResponse, error) {
	url := "/v1/product/certificate/rejection_reasons/list"

	resp := &PossibleRejectReasonsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type PossibleStatusesResponse struct {
	core.CommonResponse

	// Possible certificate statuses
	Result []PossibleStatusesResult `json:"result"`
}

type PossibleStatusesResult struct {
	// Certificate status code
	Code string `json:"code"`

	// Status description
	Name string `json:"name"`
}

func (c Certificates) PossibleStatuses(ctx context.Context) (*PossibleStatusesResponse, error) {
	url := "/v1/product/certificate/status/list"

	resp := &PossibleStatusesResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
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
func (c Certificates) AddForProducts(ctx context.Context, params *AddCertificatesForProductsParams) (*AddCertificatesForProductsResponse, error) {
	url := "/v1/product/certificate/create"

	resp := &AddCertificatesForProductsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, map[string]string{
		"Content-Type": "multipart/form-data",
	})
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
