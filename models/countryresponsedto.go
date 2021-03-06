package models

type CountryResponseDTO struct {
	CountryID         *int64   `json:"CountryId,omitempty"`
	Name              *string  `json:"Name,omitempty"`
	AlphaCode2        *string  `json:"AlphaCode2,omitempty"`
	AlphaCode3        *string  `json:"AlphaCode3,omitempty"`
	NumericCode       *int64   `json:"NumericCode,omitempty"`
	PhoneCode         *string  `json:"PhoneCode,omitempty"`
	CurrencyID        *int64   `json:"CurrencyId,omitempty"`
	MinimumSignupAge  *int64   `json:"MinimumSignupAge,omitempty"`
	IsLive            *bool    `json:"IsLive,omitempty"`
	Created           *IGCTime `json:"Created,omitempty"`
	IsExcluded        *bool    `json:"IsExcluded,omitempty"`
	IsFeatured        *bool    `json:"IsFeatured,omitempty"`
	DefaultLanguageID *int64   `json:"DefaultLanguageId,omitempty"`
}
