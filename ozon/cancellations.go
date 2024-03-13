package ozon

import (
	"context"
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Cancellations struct {
	client *core.Client
}

type GetCancellationInfoParams struct {
	// Cancellation request identifier
	CancellationId int64 `json:"cancellation_id"`
}

type GetCancellationInfoResponse struct {
	core.CommonResponse

	// Method result
	Result CancellationInfo `json:"result"`
}

type CancellationInfo struct {
	// Cancellation request identifier
	CancellationId int64 `json:"cancellation_id"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Cancellation reason
	CancellationReason CancellationInfoReason `json:"cancellation_reason"`

	// Cancellation request creation date
	CancelledAt time.Time `json:"cancelled_at"`

	// Comment to cancellation submitted by cancellation initiator
	CancellationReasonMessage string `json:"cancellation_reason_message"`

	// Delivery service integration type
	TPLIntegrationType TPLIntegrationType `json:"tpl_integration_type"`

	// Cancellation request status
	State CancellationInfoState `json:"state"`

	// Cancellation initiator
	CancellationInitiator string `json:"cancellation_initiator"`

	// Order creation date
	OrderDate time.Time `json:"order_date"`

	// Comment submitted on the cancellation request approval or rejection
	ApproveComment string `json:"approve_comment"`

	// Cancellation request approval or rejection date
	ApproveDate time.Time `json:"approve_date"`

	// Date after which the request will be automatically approved
	AutoApproveDate time.Time `json:"auto_approve_date"`
}

type CancellationInfoReason struct {
	// Cancellation reason identifier
	Id int64 `json:"id"`

	// Cancellation reason name
	Name string `json:"name"`
}

type CancellationInfoState struct {
	// Status identifier
	Id int64 `json:"id"`

	// Status name
	Name string `json:"name"`

	// Request status
	State string `json:"state"`
}

// Method for getting information about a rFBS cancellation request
func (c Cancellations) GetInfo(ctx context.Context, params *GetCancellationInfoParams) (*GetCancellationInfoResponse, error) {
	url := "/v1/conditional-cancellation/get"

	resp := &GetCancellationInfoResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ListCancellationsParams struct {
	// Filters
	Filter *ListCancellationsFilter `json:"filter,omitempty"`

	// Number of cancellation requests in the response
	Limit int32 `json:"limit,omitempty"`

	// Number of elements that will be skipped in the response.
	// For example, if offset=10, the response will start with the 11th element found
	Offset int32 `json:"offset,omitempty"`

	// Additional information
	With *ListCancellationWith `json:"with,omitempty"`
}

type ListCancellationsFilter struct {
	// Filter by cancellation initiator
	CancellationInitiator []string `json:"cancellation_initiator"`

	// Filter by shipment number.
	//
	// Optional parameter. You can pass several values here
	PostingNumber string `json:"posting_number"`

	// Filter by cancellation request status
	State string `json:"state"`
}

type ListCancellationWith struct {
	// Indication that the counter of requests in different statuses should be displayed in the response
	Counters bool `json:"counters"`
}

type ListCancellationsResponse struct {
	core.CommonResponse

	// Cancellation requests list
	Result []CancellationInfo `json:"result"`

	// The total number of requests by the specified filters
	Total int32 `json:"total"`

	// Counter of requests in different statuses
	Counters ListCancellationResponseCounters `json:"counters"`
}

type ListCancellationResponseCounters struct {
	// Number of requests for approval
	OnApproval int64 `json:"on_approval"`

	// Number of approved requests
	Approved int64 `json:"approved"`

	// Number of rejected requests
	Rejected int64 `json:"rejected"`
}

// Method for getting a list of rFBS cancellation requests
func (c Cancellations) List(ctx context.Context, params *ListCancellationsParams) (*ListCancellationsResponse, error) {
	url := "/v1/conditional-cancellation/list"

	resp := &ListCancellationsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ApproveRejectCancellationsParams struct {
	// Cancellation request identifier
	CancellationId int64 `json:"cancellation_id"`

	// Comment
	Comment string `json:"comment,omitempty"`
}

type ApproveRejectCancellationsResponse struct {
	core.CommonResponse
}

// The method allows to approve an rFBS cancellation request in the ON_APPROVAL status.
// The order will be canceled and the money will be returned to the customer
func (c Cancellations) Approve(ctx context.Context, params *ApproveRejectCancellationsParams) (*ApproveRejectCancellationsResponse, error) {
	url := "/v1/conditional-cancellation/approve"

	resp := &ApproveRejectCancellationsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

// The method allows to reject an rFBS cancellation request in the ON_APPROVAL status. Explain your decision in the comment parameter.
//
// The order will remain in the same status and must be delivered to the customer
func (c Cancellations) Reject(ctx context.Context, params *ApproveRejectCancellationsParams) (*ApproveRejectCancellationsResponse, error) {
	url := "/v1/conditional-cancellation/reject"

	resp := &ApproveRejectCancellationsResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
