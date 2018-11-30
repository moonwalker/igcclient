package models

type DocumentResponseDTO struct {
	LastFourDigitsSsnNumber *string `json:"LastFourDigitsSsnNumber,omitempty"`
	PersonID                *string `json:"PersonId,omitempty"`
}
