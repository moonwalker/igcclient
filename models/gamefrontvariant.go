package models

type GameFrontVariant struct {
	GameVariantId     *int64   `json:"GameVariantId,omitempty"`
	GameId            *int64   `json:"GameId,omitempty"`
	MinBet            *float64 `json:"MinBet,omitempty"`
	MaxBet            *float64 `json:"MaxBet,omitempty"`
	DenominatorAmount *float64 `json:"DenominatorAmount,omitempty"`
	IsDefault         *bool    `json:"IsDefault,omitempty"`
}
