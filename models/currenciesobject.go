package models

type CurrenciesObject struct {
	CurrencyId     *int     `json:"CurrencyId,omitempty"`
	Code           *string  `json:"Code,omitempty"`
	Description    *string  `json:"Description,omitempty"`
	Multiplier     *int     `json:"Multiplier,omitempty"`
	Format         *string  `json:"Format,omitempty"`
	ExchangeRate   *float64 `json:"ExchangeRate,omitempty"`
	MaxBetLimit    *float64 `json:"MaxBetLimit,omitempty"`
	CurrencySymbol *string  `json:"CurrencySymbol,omitempty"`
}
