package models

type PlayerStatsResponseDTO struct {
	NetLoss       *float64 `json:"NetLoss"`
	NetGain       *float64 `json:"NetGain"`
	TotalRealBets *float64 `json:"TotalRealBets"`
	TotalRealWins *float64 `json:"TotalRealWins"`
}
