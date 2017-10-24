package models

type PartialUserData struct {
	UserID               *int    `json:"UserID,omitempty"`
	IBAN                 *string `json:"IBAN,omitempty"`
	BankID               *int    `json:"BankID,omitempty"`
	SocialSecurityNumber *string `json:"SocialSecurityNumber,omitempty"`
}
