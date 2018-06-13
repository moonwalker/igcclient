package models

type PaymentApiResponse struct {
	TransactionID       *int64             `json:"TransactionId,omitempty"`
	StatusID            *int64             `json:"StatusId,omitempty"`
	PaymentMethodID     *string            `json:"PaymentMethodId,omitempty"`
	PaymentTypeID       *string            `json:"PaymentTypeId,omitempty"`
	PaymentCategoryID   *string            `json:"PaymentCategoryId,omitempty"`
	VendorID            *string            `json:"VendorId,omitempty"`
	DeviceType          *string            `json:"DeviceType,omitempty"`
	TransactionType     *string            `json:"TransactionType,omitempty"`
	Currency            *string            `json:"Currency,omitempty"`
	Amount              *string            `json:"Amount,omitempty"`
	Fee                 *string            `json:"Fee,omitempty"`
	TransactionDate     *string            `json:"TransactionDate,omitempty"`
	ErrorMsg            *string            `json:"ErrorMsg,omitempty"`
	BankMsg             *string            `json:"BankMsg,omitempty"`
	ProviderReference   *string            `json:"ProviderReference,omitempty"`
	GatewayReference    *string            `json:"GatewayReference,omitempty"`
	MainPageBackURL     *string            `json:"MainPageBackUrl,omitempty"`
	BackURL             *string            `json:"BackUrl,omitempty"`
	ProviderURL         *string            `json:"ProviderUrl,omitempty"`
	ProviderFrameWidth  *int64             `json:"ProviderFrameWidth,omitempty"`
	ProviderFrameHeight *int64             `json:"ProviderFrameHeight,omitempty"`
	ProviderFrameTop    *int64             `json:"ProviderFrameTop,omitempty"`
	ExtraDetails        *map[string]string `json:"ExtraDetails,omitempty"`
}
