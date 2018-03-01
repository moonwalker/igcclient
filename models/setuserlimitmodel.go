package models

type SetUserLimitModel struct {
	UserId      *int64         `json:"UserId,omitempty"`        // The User ID for the limit
	Type        *LimitType     `json:"Type,omitempty"`          // The Limit Type
	Duration    *LimitDuration `json:"LimitDuration,omitempty"` // The Limit Duration
	LimitAmount *float64       `json:"LimitAmount,omitempty"`   // The Limit Amount
}
