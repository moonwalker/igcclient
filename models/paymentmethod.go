package models

type PaymentMethod struct {
	PaymentMethodID             *int64   `json:"PaymentMethodID,omitempty"`
	PaymentTypeID               *int64   `json:"PaymentTypeID,omitempty"`
	Image                       *string  `json:"Image,omitempty"`
	IsDeposit                   *bool    `json:"IsDeposit,omitempty"`
	IsLive                      *bool    `json:"IsLive,omitempty"`
	Created                     *IGCTime `json:"Created,omitempty"`
	LastEdited                  *IGCTime `json:"LastEdited,omitempty"`
	CreatedBy                   *int64   `json:"CreatedBy,omitempty"`
	LastEditedBy                *int64   `json:"LastEditedBy,omitempty"`
	Sort                        *int64   `json:"Sort,omitempty"`
	Description                 *string  `json:"Description,omitempty"`
	CallExternalPaymentFunction *bool    `json:"CallExternalPaymentFunction,omitempty"`
	ReturnJson                  *bool    `json:"ReturnJson,omitempty"`
	ExternalPaymentFrameWidth   *int64   `json:"ExternalPaymentFrameWidth,omitempty"`
	ExternalPaymentFrameHeight  *int64   `json:"ExternalPaymentFrameHeight,omitempty"`
	ExternalPaymentFrameTop     *int64   `json:"ExternalPaymentFrameTop,omitempty"`
	PaymentMethodCode           *string  `json:"PaymentMethodCode,omitempty"`
	InstantPayment              *bool    `json:"InstantPayment,omitempty"`
	RedirectToFullScreen        *bool    `json:"RedirectToFullScreen,omitempty"`
	Name                        *string  `json:"Name,omitempty"`
	Enabled                     *bool    `json:"Enabled,omitempty"`
	MinAmount                   *float64 `json:"MinAmount,omitempty"`
	MaxAmount                   *float64 `json:"MaxAmount,omitempty"`
	ProcessingTime              *string  `json:"ProcessingTime,omitempty"`
	IsFeePercentage             *bool    `json:"IsFeePercentage,omitempty"`
	Fee                         *float64 `json:"Fee,omitempty"`
	FeeAddedOnRequest           *bool    `json:"FeeAddedOnRequest,omitempty"`
	MinFee                      *float64 `json:"MinFee,omitempty"`
	MaxFee                      *float64 `json:"MaxFee,omitempty"`
	MinWithdrawAmount           *float64 `json:"MinWithdrawAmount,omitempty"`
	MaxWithdrawAmount           *float64 `json:"MaxWithdrawAmount,omitempty"`
	WithdrawProcessingTime      *string  `json:"WithdrawProcessingTime,omitempty"`
	IsWithdrawFeePercentage     *bool    `json:"IsWithdrawFeePercentage,omitempty"`
	WithdrawFee                 *float64 `json:"WithdrawFee,omitempty"`
	WithdrawMinFee              *float64 `json:"WithdrawMinFee,omitempty"`
	WithdrawMaxFee              *float64 `json:"WithdrawMaxFee,omitempty"`
	CaptureAfterMinutes         *int64   `json:"CaptureAfterMinutes,omitempty"`
	TwoStepWithdraw             *bool    `json:"TwoStepWithdraw,omitempty"`
	CardProcessing              *bool    `json:"CardProcessing,omitempty"`
	BankTransfer                *bool    `json:"BankTransfer,omitempty"`
	AutoShowInWithdrawal        *bool    `json:"AutoShowInWithdrawal,omitempty"`
}
