package models

type GameDetailsRestrictedCountry struct {
	CountryId   *int64  `json:"CountryId,omitempty"`
	CountryCode *string `json:"CountryCode,omitempty"`
	Name        *string `json:"Name,omitempty"`
}