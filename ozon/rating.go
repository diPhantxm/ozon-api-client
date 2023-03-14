package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type GetCurrentSellerRatingInfoResponse struct {
	core.CommonResponse

	// Rating groups list
	Groups []struct {
		// Ratings group name
		GroupName string `json:"group_name"`

		// Ratings list
		Items []struct {

			// Rating change: the ratio of the previous value to the current one
			Change struct {
				// How the rating value has changed:
				//   - DIRECTION_UNKNOWN — unknown.
				//   - DIRECTION_NONE — has not changed.
				//   - DIRECTION_RISE — has increased.
				//   - DIRECTION_FALL — has dropped.
				Direction string `json:"direction"`

				// What the change means:
				//   - MEANING_UNKNOWN — unknown.
				//   - MEANING_NONE — neutral.
				//   - MEANING_GOOD — the indicator is improving, everything is good.
				//   - MEANING_BAD — the indicator is dropping, you should do something.
				Meaning string `json:"meaning"`
			} `json:"change"`

			// Current rating value
			CurrentValue float64 `json:"current_value"`

			// Rating name
			Name string `json:"name"`

			// Previous rating value
			PastValue float64 `json:"past_value"`

			// System rating name
			Rating string `json:"rating"`

			// What should be the rating value to be considered good:
			//   - UNKNOWN_DIRECTION — unknown.
			//   - NEUTRAL — doesn't matter.
			//   - HIGHER_IS_BETTER — the higher the better.
			//   - LOWER_IS_BETTER — the lower the better.
			RatingDirection string `json:"rating_direction"`

			// Rating status:
			//   - UNKNOWN_STATUS — unknown status.
			//   - OK — everything is OK.
			//   - WARNING — indicators require attention.
			//   - CRITICAL — critical rating
			Status string `json:"status"`

			// Value type:
			//   - UNKNOWN_VALUE — unknown,
			//   - INDEX,
			//   - PERCENT,
			//   - TIME,
			//   - RATIO — coefficient,
			//   - REVIEW_SCORE — score,
			//   - COUNT
			ValueType string `json:"value_type"`
		} `json:"items"`

		// An indication that the penalty points balance is exceeded
		PenaltyScoreExceeded bool `json:"penalty_score_exceeded"`

		// An indication that you participate in the Premium program
		Premium bool `json:"premium"`
	} `json:"groups"`
}

func (c Client) GetCurrentSellerRatingInfo() (*GetCurrentSellerRatingInfoResponse, error) {
	url := "/v1/rating/summary"

	resp := &GetCurrentSellerRatingInfoResponse{}

	response, err := c.client.Request(http.MethodPost, url, nil, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetSellerRatingInfoForPeriodParams struct {
	// Period start
	DateFrom time.Time `json:"date_from"`

	// Period end
	DateTo time.Time `json:"date_to"`

	// Filter by rating
	Ratings []string `json:"ratings"`

	// Indication that the response should contain information about Premium program penxalty points
	WithPremiumScores bool `json:"with_premium_scores"`
}

type GetSellerRatingInfoPeriodResponse struct {
	core.CommonResponse

	// Information on the Premium program penalty points
	PremiumScores []struct {
		// Rating name
		Rating string `json:"rating"`

		// Information on penalty points
		Scores []struct {
			// Date when the penalty points were received
			Date time.Time `json:"date"`

			// Rating value for which the penalty points were received
			RatingValue float64 `json:"rating_value"`

			// Number of received penalty points
			Value int32 `json:"value"`
		} `json:"scores"`
	} `json:"premium_scores"`

	// Information on the seller ratings
	Ratings []struct {
		// Rating threshold, after which sales will be blocked
		DangerThreshold float64 `json:"danger_threshold"`

		// Rating threshold for participation in the Premium program
		PremiumThreshold float64 `json:"premium_threshold"`

		// Rating system name
		Rating string `json:"rating"`

		// Rating values list
		Values []struct {
			// Rating calculation start date
			DateFrom time.Time `json:"date_from"`

			// Rating calculation end date
			DateTo time.Time `json:"date_to"`

			// Rating status
			Status struct {
				// Indication if the rating threshold for blocking is exceeded
				Danger bool `json:"danger"`

				// Indication whether the threshold for participation in the Premium program has been reached
				Premium bool `json:"premium"`

				// Indication of a warning that the threshold for blocking may be exceeded
				Warning bool `json:"warning"`
			} `json:"status"`

			// Rating value
			Value float64 `json:"value"`
		} `json:"values"`

		// Rating threshold, after which a warning about possible blocking appears
		WarningThreshold float64 `json:"warning_threshold"`
	} `json:"ratings"`
}

func (c Client) GetSellerRatingInfoForPeriod(params *GetSellerRatingInfoForPeriodParams) (*GetSellerRatingInfoPeriodResponse, error) {
	url := "/v1/rating/history"

	resp := &GetSellerRatingInfoPeriodResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
