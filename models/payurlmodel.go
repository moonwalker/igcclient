package models

type PayURLModel struct {
	UserIP         *string `json:"UserIP,omitempty"`
	UserAgent      *string `json:"UserAgent,omitempty"`
	LanguageAlpha2 *string `json:"LanguageAlpha2,omitempty"`
}
