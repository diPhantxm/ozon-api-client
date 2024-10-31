package ozon

import (
	"context"
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Rating struct {
	client *core.Client
}

type GetCurrentSellerRatingInfoResponse struct {
	core.CommonResponse

	// Rating groups list
	Groups []GetCurrentSellerRatingInfoGroup `json:"groups"`

	// Localization index details.
	// If you had no sales in the last 14 days,
	// the parameter fields will be empty
	LocalizationIndex []LocalizationIndex `json:"localization_index"`

	// An indication that the penalty points balance is exceeded
	PenaltyScoreExceeded bool `json:"penalty_score_exceeded"`

	// An indication that you participate in the Premium program
	Premium bool `json:"premium"`
}

type LocalizationIndex struct {
	// Date of localization index calculation
	CalculationDate time.Time `json:"calculation_date"`

	// Localization index value
	LocalizationPercentage int32 `json:"localization_percentage"`
}

type GetCurrentSellerRatingInfoGroup struct {
	// Ratings group name
	GroupName string `json:"group_name"`

	// Ratings list
	Items []GetCurrentSellerRatingInfoGroupItem `json:"items"`
}

type GetCurrentSellerRatingInfoGroupItem struct {

	// Rating change: the ratio of the previous value to the current one
	Change GetCurrentSellerRatingInfoGroupItemChange `json:"change"`

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
}

type GetCurrentSellerRatingInfoGroupItemChange struct {
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
}

func (c Rating) GetCurrentSellerRatingInfo(ctx context.Context) (*GetCurrentSellerRatingInfoResponse, error) {
	url := "/v1/rating/summary"

	resp := &GetCurrentSellerRatingInfoResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, nil, resp, nil)
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
	PremiumScores []GetSellerRatingInfoPeriodPremiumScores `json:"premium_scores"`

	// Information on the seller ratings
	Ratings []GetSellerRatingInfoPeriodRating `json:"ratings"`
}

type GetSellerRatingInfoPeriodPremiumScores struct {
	// Rating name
	Rating string `json:"rating"`

	// Information on penalty points
	Scores []GetSellerRatingInfoPeriodPremiumScore `json:"scores"`
}

type GetSellerRatingInfoPeriodPremiumScore struct {
	// Date when the penalty points were received
	Date time.Time `json:"date"`

	// Rating value for which the penalty points were received
	RatingValue float64 `json:"rating_value"`

	// Number of received penalty points
	Value int32 `json:"value"`
}

type GetSellerRatingInfoPeriodRating struct {
	// Rating threshold, after which sales will be blocked
	DangerThreshold float64 `json:"danger_threshold"`

	// Rating threshold for participation in the Premium program
	PremiumThreshold float64 `json:"premium_threshold"`

	// Rating system name
	Rating string `json:"rating"`

	// Rating values list
	Values []GetSellerRatingInfoPeriodRatingValue `json:"values"`

	// Rating threshold, after which a warning about possible blocking appears
	WarningThreshold float64 `json:"warning_threshold"`
}

type GetSellerRatingInfoPeriodRatingValue struct {
	// Rating calculation start date
	DateFrom time.Time `json:"date_from"`

	// Rating calculation end date
	DateTo time.Time `json:"date_to"`

	// Rating status
	Status GetSellerRatingInfoPeriodRatingValueStatus `json:"status"`

	// Rating value
	Value float64 `json:"value"`
}

type GetSellerRatingInfoPeriodRatingValueStatus struct {
	// Indication if the rating threshold for blocking is exceeded
	Danger bool `json:"danger"`

	// Indication whether the threshold for participation in the Premium program has been reached
	Premium bool `json:"premium"`

	// Indication of a warning that the threshold for blocking may be exceeded
	Warning bool `json:"warning"`
}

func (c Rating) GetSellerRatingInfoForPeriod(ctx context.Context, params *GetSellerRatingInfoForPeriodParams) (*GetSellerRatingInfoPeriodResponse, error) {
	url := "/v1/rating/history"

	resp := &GetSellerRatingInfoPeriodResponse{}

	response, err := c.client.Request(ctx, http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
