package models

type UserLimitResponse struct {
	Id          *int64   `json:"Id,omitempty"`          //The ID of the user limit
	LimitType   *int64   `json:"LimitType,omitempty"`   //The Limit Type
	Duration    *int64   `json:"Duration,omitempty"`    //The Duration of the limit
	UserId      *int64   `json:"UserId,omitempty"`      //The User ID of the limit
	CreatedDate *IGCTime `json:"CreatedDate,omitempty"` //The date the limit was created
	StartDate   *IGCTime `json:"StartDate,omitempty"`   //The date the limit starts
	EndDate     *IGCTime `json:"EndDate,omitempty"`     //The Date the limit ends
	Amount      *float64 `json:"Amount,omitempty"`      //The amount of the limit
	AmountHR    *string  `json:"Amount_HR,omitempty"`
	AmountDAY   *string  `json:"Amount_DAY,omitempty"`
}
