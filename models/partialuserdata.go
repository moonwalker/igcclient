package models

type PartialUserData struct {
	UserID               *int64  `json:"UserID,omitempty"`
	IBAN                 *string `json:"IBAN,omitempty"`
	BankID               *int64  `json:"BankID,omitempty"`
	SocialSecurityNumber *string `json:"SocialSecurityNumber,omitempty"`
}
