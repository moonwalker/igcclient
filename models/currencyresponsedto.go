package models

type CurrencyResponseDTO struct {
	CurrencyID        *int64   `json:"CurrencyId,omitempty"`
	Code              *string  `json:"Code,omitempty"`
	Description       *string  `json:"Description,omitempty"`
	CurrencySymbol    *string  `json:"CurrencySymbol,omitempty"`
	Format            *string  `json:"Format,omitempty"`
	ExchangeRate      *float64 `json:"ExchangeRate,omitempty"`
	Multiplier        *int64   `json:"Multiplier,omitempty"`
	MaxBetLimit       *float64 `json:"MaxBetLimit,omitempty"`
	SportsMaxBetLimit *float64 `json:"SportsMaxBetLimit,omitempty"`
}
