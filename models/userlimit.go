package models

type UserLimit struct {
	Id          *int           `json:"Id,omitempty"`          // The ID of the user limit
	LimitType   *LimitType     `json:"LimitType,omitempty"`   // The Limit Type
	Duration    *LimitDuration `json:"Duration,omitempty"`    // The Duration of the limit
	UserId      *int           `json:"UserId,omitempty"`      // The User ID of the limit
	CreatedDate *IGCTime       `json:"CreatedDate,omitempty"` // The date the limit was created
	StartDate   *IGCTime       `json:"StartDate,omitempty"`   // The date the limit starts
	EndDate     *IGCTime       `json:"EndDate,omitempty"`     // The date the limit ends
	Amount      *float64       `json:"Amount,omitempty"`      // The amount of the limit
}
