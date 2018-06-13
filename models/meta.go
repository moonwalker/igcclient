package models

type Meta struct {
	MetaType      *string `json:"meta:type,omitempty"`
	ReadOnly      *bool   `json:"read-only,omitempty"`
	Value         *string `json:"value,omitempty"`
	Step          *int64  `json:"step,omitempty"`
	URL           *string `json:"url,omitempty"`
	DictionaryKey *string `json:"dictionary-key,omitempty"`
}
