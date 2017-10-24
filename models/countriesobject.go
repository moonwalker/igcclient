package models

type CountriesObject struct {
	CountryID         *int     `json:"CountryID,omitempty"`
	Name              *string  `json:"Name,omitempty"`
	AlphaCode2        *string  `json:"AlphaCode2,omitempty"`
	AlphaCode3        *string  `json:"AlphaCode3,omitempty"`
	NumericCode       *int     `json:"NumericCode,omitempty"`
	PhoneCode         *string  `json:"PhoneCode,omitempty"`
	CurrencyId        *int     `json:"CurrencyId,omitempty"`
	MinimumSignupAge  *int     `json:"MinimumSignupAge,omitempty"`
	IsLive            *bool    `json:"IsLive,omitempty"`
	Created           *IGCTime `json:"Created,omitempty"`
	IsExcluded        *bool    `json:"IsExcluded,omitempty"`
	IsFeatured        *bool    `json:"IsFeatured,omitempty"`
	KycLimit          *float64 `json:"KycLimit,omitempty"`
	DefaultLanguageId *int     `json:"DefaultLanguageId,omitempty"`
}
