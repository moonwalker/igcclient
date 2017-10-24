package models

type PaymentResponse struct {
	Status            *TransactionStatus `json:"Status,omitempty"`
	RedirectUrl       *string            `json:"RedirectUrl,omitempty"`
	ErrorMsg          *string            `json:"ErrorMsg,omitempty"`
	BankMsg           *string            `json:"BankMsg,omitempty"`
	TransactionId     *int               `json:"TransactionId,omitempty"`
	ProviderResponse  *string            `json:"ProviderResponse,omitempty"`
	WalletAccount     *string            `json:"WalletAccount,omitempty"`
	ProviderReference *string            `json:"ProviderReference,omitempty"`
	GatewayReference  *string            `json:"GatewayReference,omitempty"`
	UserCardId        *int               `json:"UserCardId,omitempty"`
	FraudScore        *int               `json:"FraudScore,omitempty"`
	ExtraDetails      *map[string]string `json:"ExtraDetails,omitempty"`
	CardCountryID     *int               `json:"CardCountryID,omitempty"`
	PaymentTypeId     *int               `json:"PaymentTypeId,omitempty"`
	VendorId          *int               `json:"VendorId,omitempty"`
	GatewayId         *int               `json:"GatewayId,omitempty"`
	ActionId          *string            `json:"ActionId,omitempty"`
}
