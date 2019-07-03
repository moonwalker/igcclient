package models

type SowAnswer struct {
	ID       *int64  `json:"Id,omitempty"`
	Text     *string `json:"Text,omitempty"`
	FreeText *string `json:"FreeText,omitempty"`
}
