package models

type RegistrationAdditionalData struct {
	IBAN                 *string `json:"IBAN,omitempty"`                 // Bank account IBAN number as per ISO 13616-1:2007. Since IBAN check varies per country, this check should be done at operator frontend level.
	BankID               *int64  `json:"BankID,omitempty"`               // Reference to Core.Banks table. Call Banks/GetBanks method to fetch all existing banks.
	SocialSecurityNumber *string `json:"SocialSecurityNumber,omitempty"` // User's Social Security number
}
