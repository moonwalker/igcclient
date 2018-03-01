package models

type CountryVerificationTypeWhitelist struct {
	CountryID          *int64            `json:"CountryID,omitempty"`
	VerificationTypeID *int64            `json:"VerificationTypeID,omitempty"`
	VerificationType   *VerificationType `json:"VerificationType,omitempty"`
}
