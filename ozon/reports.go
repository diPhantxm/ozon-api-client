package ozon

import (
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
	ReportType string `json:"report_type"`
}

type GetReportsListResponse struct {
	core.CommonResponse

	// Method result
	Result struct {
		// Unique report identifier
		Code string `json:"code"`

		// Report creation date
		CreatedAt time.Time `json:"created_at"`

		// Error code when generating the report
		Error string `json:"error"`

		// Link to CSV file
		File string `json:"file"`

		// Array with the filters specified when the seller created the report
		Params struct {
		} `json:"params"`

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
	} `json:"result"`

	// Total number of reports
	Total int32 `json:"total"`
}

// Returns the list of reports that have been generated before
func (c Reports) GetList(params *GetReportsListParams) (*GetReportsListResponse, error) {
	url := "/v1/report/list"

	resp := &GetReportsListResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
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
	Result struct {
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
	} `json:"result"`
}

// Returns information about a created report by its identifier
func (c Reports) GetReportDetails(params *GetReportDetailsParams) (*GetReportDetailsResponse, error) {
	url := "/v1/report/list"

	resp := &GetReportDetailsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
