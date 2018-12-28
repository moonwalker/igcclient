package models

type UserLimitResponseDTO struct {
	Duration  *LimitDuration `json:"Duration"`
	LimitType *LimitType     `json:"LimitType"`
	ID        *int64         `json:"Id"`
	Amount    *float64       `json:"Amount"`
}
