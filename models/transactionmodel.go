package models

type TransactionModel struct {
	TransactionId     *string           `json:"TransactionId,omitempty"` // String because of JS Number width. Better be on the safe side.
	TransactionType   *int              `json:"TransactionType,omitempty"`
	UserId            *int              `json:"UserId,omitempty"`
	VendorId          *int              `json:"VendorId,omitempty"`
	PaymentTypeId     *int              `json:"PaymentTypeId,omitempty"`
	Status            *int              `json:"Status,omitempty"`
	CurrencyId        *int              `json:"CurrencyId,omitempty"`
	Amount            *float64          `json:"Amount,omitempty"`
	BaseAmount        *float64          `json:"BaseAmount,omitempty"`
	BankMessage       *string           `json:"BankMessage,omitempty"`
	UserCardId        *int              `json:"UserCardId,omitempty"`
	UserIp            *string           `json:"UserIp,omitempty"`
	Date              *IGCTime          `json:"Date,omitempty"`
	TransactionFee    *float64          `json:"TransactionFee,omitempty"`
	FeeAddedToAmount  *bool             `json:"FeeAddedToAmount,omitempty"`
	TrnReferenceId    *int              `json:"TrnReferenceId,omitempty"`
	ProviderReference *string           `json:"ProviderReference,omitempty"`
	WalletTransId     *int              `json:"WalletTransId,omitempty"`
	CardCountryID     *int              `json:"CardCountryID,omitempty"`
	BonusCode         *string           `json:"BonusCode,omitempty"`
	GatewayId         *int              `json:"GatewayId,omitempty"`
	TablePrefix       *string           `json:"TablePrefix,omitempty"`
	IdColumnsName     *string           `json:"IdColumnsName,omitempty"`
	PaymentType       *PaymentTypeModel `json:"PaymentType,omitempty"`
}
