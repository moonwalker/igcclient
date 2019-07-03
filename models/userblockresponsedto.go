package models

type UserBlockResponseDTO struct {
	BlockReasonID *int64   `json:"BlockReasonId,omitempty"`
	EndDate       *IGCTime `json:"EndDate"`
}
