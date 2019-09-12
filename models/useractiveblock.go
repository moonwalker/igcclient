package models

type UserActiveBlock struct {
	UserBlockID         *int64   `json:"UserBlockId"`
	BlockReasonID       *int64   `json:"BlockReasonId"`
	BlockReason         *string  `json:"BlockReason"`
	BlockTypeID         *int64   `json:"BlockTypeId"`
	BlockType           *string  `json:"BlockType"`
	Comment             *string  `json:"Comment"`
	StartDate           *IGCTime `json:"StartDate"`
	EndDate             *IGCTime `json:"EndDate"`
	CooldownEndDate     *IGCTime `json:"CooldownEndDate"`
	ExpireAutomatically *bool    `json:"ExpireAutomatically"`
	CanBeCancelled      *bool    `json:"CanBeCancelled"`
	CreatedDate         *IGCTime `json:"CreatedDate"`
	CreatedBy           *int64   `json:"CreatedBy"`
}
