package models

type PaymentResponse struct {
	Status            *TransactionStatus `json:"Status,omitempty"`
	RedirectURL       *string            `json:"RedirectUrl,omitempty"`
	ErrorMsg          *string            `json:"ErrorMsg,omitempty"`
	BankMsg           *string            `json:"BankMsg,omitempty"`
	TransactionID     *int64             `json:"TransactionId,omitempty"`
	ProviderResponse  *string            `json:"ProviderResponse,omitempty"`
	WalletAccount     *string            `json:"WalletAccount,omitempty"`
	ProviderReference *string            `json:"ProviderReference,omitempty"`
	GatewayReference  *string            `json:"GatewayReference,omitempty"`
	UserCardID        *int64             `json:"UserCardId,omitempty"`
	FraudScore        *int64             `json:"FraudScore,omitempty"`
	ExtraDetails      *map[string]string `json:"ExtraDetails,omitempty"`
	CardCountryID     *int64             `json:"CardCountryID,omitempty"`
	PaymentTypeID     *int64             `json:"PaymentTypeId,omitempty"`
	VendorID          *int64             `json:"VendorId,omitempty"`
	GatewayID         *int64             `json:"GatewayId,omitempty"`
	ActionID          *string            `json:"ActionId,omitempty"`
}
