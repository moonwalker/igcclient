package models

type SendEmailModel struct {
	EmailTemplateID            *int64             `json:"EmailTemplateID,omitempty"`
	SendEmailTo                *string            `json:"SendEmailTo,omitempty"`
	EmailBodyPlaceholderValues *map[string]string `json:"EmailBodyPlaceholderValues,omitempty"`
	LanguageCode               *string            `json:"LanguageCode,omitempty"`
}
