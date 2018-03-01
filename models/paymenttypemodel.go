package models

type PaymentTypeModel struct {
	PaymentTypeID           *int64   `json:"PaymentTypeID,omitempty"`
	Name                    *string  `json:"Name,omitempty"`
	Enabled                 *bool    `json:"Enabled,omitempty"`
	Created                 *IGCTime `json:"Created,omitempty"`
	MinAmount               *float64 `json:"MinAmount,omitempty"`
	MaxAmount               *float64 `json:"MaxAmount,omitempty"`
	ProcessingTime          *string  `json:"ProcessingTime,omitempty"`
	IsFeePercentage         *bool    `json:"IsFeePercentage,omitempty"`
	Fee                     *float64 `json:"Fee,omitempty"`
	FeeAddedOnRequest       *bool    `json:"FeeAddedOnRequest,omitempty"`
	MinFee                  *float64 `json:"MinFee,omitempty"`
	MaxFee                  *float64 `json:"MaxFee,omitempty"`
	MinWithdrawAmount       *float64 `json:"MinWithdrawAmount,omitempty"`
	MaxWithdrawAmount       *float64 `json:"MaxWithdrawAmount,omitempty"`
	WithdrawProcessingTime  *string  `json:"WithdrawProcessingTime,omitempty"`
	IsWithdrawFeePercentage *bool    `json:"IsWithdrawFeePercentage,omitempty"`
	WithdrawFee             *float64 `json:"WithdrawFee,omitempty"`
	WithdrawMinFee          *float64 `json:"WithdrawMinFee,omitempty"`
	WithdrawMaxFee          *float64 `json:"WithdrawMaxFee,omitempty"`
	CaptureAfterMinutes     *int64   `json:"CaptureAfterMinutes,omitempty"`
	TwoStepWithdraw         *bool    `json:"TwoStepWithdraw,omitempty"`
	CardProcessing          *bool    `json:"CardProcessing,omitempty"`
	BankTransfer            *bool    `json:"BankTransfer,omitempty"`
}
