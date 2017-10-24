package models

type Meta struct {
	MetaType      *string `json:"meta:type,omitempty"`
	ReadOnly      *bool   `json:"read-only,omitempty"`
	Value         *string `json:"value,omitempty"`
	Step          *int    `json:"step,omitempty"`
	Url           *string `json:"url,omitempty"`
	DictionaryKey *string `json:"dictionary-key,omitempty"`
}
