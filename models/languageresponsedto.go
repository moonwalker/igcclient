package models

type LanguageResponseDTO struct {
	LanguageID *int64   `json:"LanguageId,omitempty"`
	Alpha2Code *string  `json:"Alpha2Code,omitempty"`
	Alpha3Code *string  `json:"Alpha3Code,omitempty"`
	Name       *string  `json:"Name,omitempty"`
	IsLive     *bool    `json:"IsLive,omitempty"`
	LCID       *string  `json:"LCID,omitempty"`
	Created    *IGCTime `json:"Created,omitempty"`
}
