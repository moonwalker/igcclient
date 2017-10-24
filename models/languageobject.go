package models

type LanguageObject struct {
	LanguageID *int     `json:"LanguageID,omitempty"`
	Alpha2Code *string  `json:"Alpha2Code,omitempty"`
	Alpha3Code *string  `json:"Alpha3Code,omitempty"`
	Language   *string  `json:"Language,omitempty"`
	IsLive     *bool    `json:"IsLive,omitempty"`
	LCID       *string  `json:"LCID,omitempty"`
	Created    *IGCTime `json:"Created,omitempty"`
	FlagURL    *string  `json:"FlagURL,omitempty"`
}
