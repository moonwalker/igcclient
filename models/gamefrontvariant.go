package models

type GameFrontVariant struct {
	GameVariantId     *int     `json:"GameVariantId,omitempty"`
	GameId            *int     `json:"GameId,omitempty"`
	MinBet            *float64 `json:"MinBet,omitempty"`
	MaxBet            *float64 `json:"MaxBet,omitempty"`
	DenominatorAmount *float64 `json:"DenominatorAmount,omitempty"`
	IsDefault         *bool    `json:"IsDefault,omitempty"`
}
