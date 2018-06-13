package models

type TransactionModel struct {
	TransactionID     *string           `json:"TransactionId,omitempty"` // String because of JS Number width. Better be on the safe side.
	TransactionType   *int64            `json:"TransactionType,omitempty"`
	UserID            *int64            `json:"UserId,omitempty"`
	VendorID          *int64            `json:"VendorId,omitempty"`
	PaymentTypeID     *int64            `json:"PaymentTypeId,omitempty"`
	Status            *int64            `json:"Status,omitempty"`
	CurrencyID        *int64            `json:"CurrencyId,omitempty"`
	Amount            *float64          `json:"Amount,omitempty"`
	BaseAmount        *float64          `json:"BaseAmount,omitempty"`
	BankMessage       *string           `json:"BankMessage,omitempty"`
	UserCardID        *int64            `json:"UserCardId,omitempty"`
	UserIP            *string           `json:"UserIp,omitempty"`
	Date              *IGCTime          `json:"Date,omitempty"`
	TransactionFee    *float64          `json:"TransactionFee,omitempty"`
	FeeAddedToAmount  *bool             `json:"FeeAddedToAmount,omitempty"`
	TrnReferenceID    *int64            `json:"TrnReferenceId,omitempty"`
	ProviderReference *string           `json:"ProviderReference,omitempty"`
	WalletTransID     *int64            `json:"WalletTransId,omitempty"`
	CardCountryID     *int64            `json:"CardCountryID,omitempty"`
	BonusCode         *string           `json:"BonusCode,omitempty"`
	GatewayID         *int64            `json:"GatewayId,omitempty"`
	TablePrefix       *string           `json:"TablePrefix,omitempty"`
	IDColumnsName     *string           `json:"IdColumnsName,omitempty"`
	PaymentType       *PaymentTypeModel `json:"PaymentType,omitempty"`
}
