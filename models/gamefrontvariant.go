package models

type GameFrontVariant struct {
	GameVariantID     *int64   `json:"GameVariantId,omitempty"`
	GameID            *int64   `json:"GameId,omitempty"`
	MinBet            *float64 `json:"MinBet,omitempty"`
	MaxBet            *float64 `json:"MaxBet,omitempty"`
	DenominatorAmount *float64 `json:"DenominatorAmount,omitempty"`
	IsDefault         *bool    `json:"IsDefault,omitempty"`
}
