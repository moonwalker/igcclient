package models

type CurrenciesObject struct {
	CurrencyID     *int64   `json:"CurrencyId,omitempty"`
	Code           *string  `json:"Code,omitempty"`
	Description    *string  `json:"Description,omitempty"`
	Multiplier     *int64   `json:"Multiplier,omitempty"`
	Format         *string  `json:"Format,omitempty"`
	ExchangeRate   *float64 `json:"ExchangeRate,omitempty"`
	MaxBetLimit    *float64 `json:"MaxBetLimit,omitempty"`
	CurrencySymbol *string  `json:"CurrencySymbol,omitempty"`
}
