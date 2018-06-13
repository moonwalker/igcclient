package models

type WalletTransactionsModel struct {
	TransactionID                  *string  `json:"TransactionId,omitempty"`
	Date                           *IGCTime `json:"Date,omitempty"`
	TransactionType                *string  `json:"TransactionType,omitempty"`
	Description                    *string  `json:"Description,omitempty"`
	Currency                       *string  `json:"Currency,omitempty"`
	Amount                         *float64 `json:"Amount,omitempty"`
	BalanceBefore                  *float64 `json:"BalanceBefore,omitempty"`
	BalanceAfter                   *float64 `json:"BalanceAfter,omitempty"`
	StatusID                       *int64   `json:"StatusId,omitempty"`
	Status                         *string  `json:"Status,omitempty"`
	AllowWithdrawReversal          *bool    `json:"AllowWithdrawReversal,omitempty"`
	WithdrawReversalTimeoutMinutes *int64   `json:"WithdrawReversalTimeoutMinutes,omitempty"`
	PaymentTranID                  *string  `json:"PaymentTranId,omitempty"`
	TransactionTypeID              *int64   `json:"TransactionTypeId,omitempty"`
	ProviderReference              *string  `json:"ProviderReference,omitempty"`
	VendorID                       *int64   `json:"VendorId,omitempty"`
	PaymentTypeID                  *int64   `json:"PaymentTypeId,omitempty"`
	GatewayReference               *string  `json:"GatewayReference,omitempty"`
}
