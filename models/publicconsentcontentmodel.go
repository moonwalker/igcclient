package models

type PublicConsentContentModel struct {
	LanguageAlpha2Code *string `json:"LanguageAlpha2Code,omitempty"`
	Header             *string `json:"Header,omitempty"`
	Content            *string `json:"Content,omitempty"`
}
