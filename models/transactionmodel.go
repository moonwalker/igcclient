package models

type TransactionModel struct {
	TransactionId     *string           `json:"TransactionId,omitempty"` // String because of JS Number width. Better be on the safe side.
	TransactionType   *int64            `json:"TransactionType,omitempty"`
	UserId            *int64            `json:"UserId,omitempty"`
	VendorId          *int64            `json:"VendorId,omitempty"`
	PaymentTypeId     *int64            `json:"PaymentTypeId,omitempty"`
	Status            *int64            `json:"Status,omitempty"`
	CurrencyId        *int64            `json:"CurrencyId,omitempty"`
	Amount            *float64          `json:"Amount,omitempty"`
	BaseAmount        *float64          `json:"BaseAmount,omitempty"`
	BankMessage       *string           `json:"BankMessage,omitempty"`
	UserCardId        *int64            `json:"UserCardId,omitempty"`
	UserIp            *string           `json:"UserIp,omitempty"`
	Date              *IGCTime          `json:"Date,omitempty"`
	TransactionFee    *float64          `json:"TransactionFee,omitempty"`
	FeeAddedToAmount  *bool             `json:"FeeAddedToAmount,omitempty"`
	TrnReferenceId    *int64            `json:"TrnReferenceId,omitempty"`
	ProviderReference *string           `json:"ProviderReference,omitempty"`
	WalletTransId     *int64            `json:"WalletTransId,omitempty"`
	CardCountryID     *int64            `json:"CardCountryID,omitempty"`
	BonusCode         *string           `json:"BonusCode,omitempty"`
	GatewayId         *int64            `json:"GatewayId,omitempty"`
	TablePrefix       *string           `json:"TablePrefix,omitempty"`
	IdColumnsName     *string           `json:"IdColumnsName,omitempty"`
	PaymentType       *PaymentTypeModel `json:"PaymentType,omitempty"`
}
