package models

type CountryVerificationTypeWhitelist struct {
	CountryID          *int              `json:"CountryID,omitempty"`
	VerificationTypeID *int              `json:"VerificationTypeID,omitempty"`
	VerificationType   *VerificationType `json:"VerificationType,omitempty"`
}
