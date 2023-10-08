package ozon

import (
	"context"
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type Barcodes struct {
	client *core.Client
}

type GenerateBarcodesParams struct {
	// List of products for which you want to generate barcodes
	ProductIds []int64 `json:"product_ids"`
}

type GenerateBarcodesResponse struct {
	core.CommonResponse

	Errors []GenerateBarcodesError `json:"errors"`
}

type GenerateBarcodesError struct {
	// Error code
	Code string `json:"code"`

	// Error details
	Error string `json:"error"`

	// Barcode that is failed to generate
	Barcode string `json:"barcode"`

	// Product identifier for which the barcode generation failed
	ProductId int64 `json:"product_id"`
}

// If a product doesn't have a barcode, you can create it using this method. If a barcode already exists,
// but it isn't specified in your account, you can bind it using the `/v1/barcode/add` method.
//
// You can't generate barcodes for more than 100 products per request.
// You can use the method no more than 20 times per minute.
func (b *Barcodes) Generate(ctx context.Context, params *GenerateBarcodesParams) (*GenerateBarcodesResponse, error) {
	url := "/v1/barcode/generate"

	resp := &GenerateBarcodesResponse{}

	response, err := b.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type BindBarcodesParams struct {
	// List of barcodes and products
	Barcodes []BindBarcode `json:"barcodes"`
}

type BindBarcode struct {
	// Barcode. Maximum 100 characters
	Barcode string `json:"barcode"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

type BindBarcodesResponse struct {
	core.CommonResponse

	// Errors while binding barcodes
	Errors []BindBarcodesError `json:"errors"`
}

type BindBarcodesError struct {
	// Error code
	Code string `json:"code"`

	// Error details
	Error string `json:"error"`

	// Barcode that is failed to generate
	Barcode string `json:"barcode"`

	// SKU of the product for which the barcode binding failed
	SKU int64 `json:"sku"`
}

// If a product has a barcode that isn't specified in your account,
// bind it using this method. If a product doesn't have a barcode,
// you can create it using the `/v1/barcode/generate` method.
//
// You can't bind barcodes to more than 100 products per request.
// Each product can have up to 100 barcodes.
// You can use the method no more than 20 times per minute.
func (b *Barcodes) Bind(ctx context.Context, params *BindBarcodesParams) (*BindBarcodesResponse, error) {
	url := "/v1/barcode/add"

	resp := &BindBarcodesResponse{}

	response, err := b.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
