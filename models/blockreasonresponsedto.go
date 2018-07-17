package models

type BlockReasonResponseDTO struct {
	BlockReasonType int64  `json:"BlockReasonType"`
	Name            string `json:"Name"`
	IsSystemReason  bool   `json:"IsSystemReason"`
	BlockType       int64  `json:"BlockType"`
}
