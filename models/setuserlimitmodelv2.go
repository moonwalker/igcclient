package models

type SetUserLimitModelV2 struct {
	UserID      *int64         `json:"UserId,omitempty"`        // The User ID for the limit
	Type        *LimitType     `json:"Type,omitempty"`          // The Limit Type
	Duration    *LimitDuration `json:"LimitDuration,omitempty"` // The Limit Duration
	LimitAmount *float64       `json:"LimitAmount,omitempty"`   // The Limit Amount
	StartDate   *IGCTime       `json:"StartDate,omitempty"`     // By Default, Current UTC datetime is limit start date can be set for future as well
	EndDate     *IGCTime       `json:"EndDate,omitempt"`        // for a non-recurrent limit, End date is required, NULL for recurrent limits
}
