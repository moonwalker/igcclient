package models

type PhoneResponseDTO struct {
	MobilePrefix *string `json:"MobilePrefix,omitempty"`
	MobileNumber *string `json:"MobileNumber,omitempty"`
	PhonePrefix  *string `json:"PhonePrefix,omitempty"`
	PhoneNumber  *string `json:"PhoneNumber,omitempty"`
}
