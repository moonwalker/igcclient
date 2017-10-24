package models

type WalletTransactionsModel struct {
	TransactionId                  *string  `json:"TransactionId,omitempty"`
	Date                           *IGCTime `json:"Date,omitempty"`
	TransactionType                *string  `json:"TransactionType,omitempty"`
	Description                    *string  `json:"Description,omitempty"`
	Currency                       *string  `json:"Currency,omitempty"`
	Amount                         *float64 `json:"Amount,omitempty"`
	BalanceBefore                  *float64 `json:"BalanceBefore,omitempty"`
	BalanceAfter                   *float64 `json:"BalanceAfter,omitempty"`
	StatusId                       *int     `json:"StatusId,omitempty"`
	Status                         *string  `json:"Status,omitempty"`
	AllowWithdrawReversal          *bool    `json:"AllowWithdrawReversal,omitempty"`
	WithdrawReversalTimeoutMinutes *int     `json:"WithdrawReversalTimeoutMinutes,omitempty"`
	PaymentTranId                  *string  `json:"PaymentTranId,omitempty"`
	TransactionTypeId              *int     `json:"TransactionTypeId,omitempty"`
	ProviderReference              *string  `json:"ProviderReference,omitempty"`
	VendorId                       *int     `json:"VendorId,omitempty"`
	PaymentTypeId                  *int     `json:"PaymentTypeId,omitempty"`
	GatewayReference               *string  `json:"GatewayReference,omitempty"`
}
