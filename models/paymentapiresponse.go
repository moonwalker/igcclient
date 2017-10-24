package models

type PaymentApiResponse struct {
	TransactionId       *int               `json:"TransactionId,omitempty"`
	StatusId            *int               `json:"StatusId,omitempty"`
	PaymentMethodId     *string            `json:"PaymentMethodId,omitempty"`
	PaymentTypeId       *string            `json:"PaymentTypeId,omitempty"`
	PaymentCategoryId   *string            `json:"PaymentCategoryId,omitempty"`
	VendorId            *string            `json:"VendorId,omitempty"`
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
	MainPageBackUrl     *string            `json:"MainPageBackUrl,omitempty"`
	BackUrl             *string            `json:"BackUrl,omitempty"`
	ProviderUrl         *string            `json:"ProviderUrl,omitempty"`
	ProviderFrameWidth  *int               `json:"ProviderFrameWidth,omitempty"`
	ProviderFrameHeight *int               `json:"ProviderFrameHeight,omitempty"`
	ProviderFrameTop    *int               `json:"ProviderFrameTop,omitempty"`
	ExtraDetails        *map[string]string `json:"ExtraDetails,omitempty"`
}
