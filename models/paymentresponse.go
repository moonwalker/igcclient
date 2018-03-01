package models

type PaymentResponse struct {
	Status            *TransactionStatus `json:"Status,omitempty"`
	RedirectUrl       *string            `json:"RedirectUrl,omitempty"`
	ErrorMsg          *string            `json:"ErrorMsg,omitempty"`
	BankMsg           *string            `json:"BankMsg,omitempty"`
	TransactionId     *int64             `json:"TransactionId,omitempty"`
	ProviderResponse  *string            `json:"ProviderResponse,omitempty"`
	WalletAccount     *string            `json:"WalletAccount,omitempty"`
	ProviderReference *string            `json:"ProviderReference,omitempty"`
	GatewayReference  *string            `json:"GatewayReference,omitempty"`
	UserCardId        *int64             `json:"UserCardId,omitempty"`
	FraudScore        *int64             `json:"FraudScore,omitempty"`
	ExtraDetails      *map[string]string `json:"ExtraDetails,omitempty"`
	CardCountryID     *int64             `json:"CardCountryID,omitempty"`
	PaymentTypeId     *int64             `json:"PaymentTypeId,omitempty"`
	VendorId          *int64             `json:"VendorId,omitempty"`
	GatewayId         *int64             `json:"GatewayId,omitempty"`
	ActionId          *string            `json:"ActionId,omitempty"`
}
