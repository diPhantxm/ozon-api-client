package ozon

import (
	"net/http"

	core "github.com/diphantxm/ozon-api-client"
)

type GetListOfWarehousesResponse struct {
	core.CommonResponse

	Result []struct {
		// Trusted acceptance attribute. `true` if trusted acceptance is enabled in the warehouse
		HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`

		// Indication that the warehouse works under the rFBS scheme:
		//   - true — the warehouse works under the rFBS scheme;
		//   - false — the warehouse does not work under the rFBS scheme.
		IsRFBS bool `json:"is_rfbs"`

		// Warehouse name
		Name string `json:"name"`

		// Warehouse identifier
		WarehouseId int64 `json:"warehouse_id"`

		// Possibility to print an acceptance certificate in advance. `true` if printing in advance is possible
		CanPrintActInAdvance bool `json:"can_print_act_in_advance"`

		// FBS first mile
		FirstMileType struct {
			// DropOff point identifier
			DropoffPointId string `json:"dropoff_point_id"`

			// DropOff timeslot identifier
			DropoffTimeslotId int64 `json:"dropoff_timeslot_id"`

			// Indication that the warehouse settings are being updated
			FirstMileIsChanging bool `json:"first_mile_is_changing"`

			// First mile type:
			//
			// Enum: "DropOff" "Pickup"
			//   - DropOff
			//   - Pickup
			FirstMileType string `json:"first_mile_type"`
		} `json:"first_mile_type"`

		// Indication if there is a limit on the minimum number of orders. `true` if there is such a limit
		HasPostingsLimit bool `json:"has_postings_limit"`

		// Indication that the warehouse is not working due to quarantine
		IsKarantin bool `json:"is_karantin"`

		// Indication that the warehouse accepts bulky products
		IsKGT bool `json:"is_kgt"`

		// Indication that warehouse schedule can be changed
		IsTimetableEditable bool `json:"is_timetable_editable"`

		// Minimum limit value: the number of orders that can be brought in one shipment
		MinPostingsLimit int32 `json:"min_postings_limit"`

		// Limit value. -1 if there is no limit
		PostingsLimit int32 `json:"postings_limit"`

		// Number of warehouse working days
		MinWorkingDays int64 `json:"min_working_days"`

		// Warehouse status
		Status string `json:"status"`

		// Warehouse working days
		WorkingDays []string `json:"working_days"`
	} `json:"resulCommonResponse"`
}

// You do not need to specify any parameters in the request. Your company will be identified by the Client ID
func (c Client) GetListOfWarehouses() (*GetListOfWarehousesResponse, error) {
	url := "/v1/warehouse/list"

	resp := &GetListOfWarehousesResponse{}

	response, err := c.client.Request(http.MethodPost, url, nil, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
