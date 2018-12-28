package models

type UserRgStatsResponseDTO struct {
	Limits    *[]UserLimitResponseDTO `json:"Limits,omitempty"`
	PlayStats *PlayerStatsResponseDTO `json:"Limits,omitempty"`
}
