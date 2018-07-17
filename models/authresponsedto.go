package models

type AuthResponseDTO struct {
	Token        *string                   `json:"Token"`
	BlockReasons *[]BlockReasonResponseDTO `json:"BlockReasons"`
	Result       *string                   `json:"Result"`
}
