package models

type SendEmailModel struct {
	EmailTemplateID            *int               `json:"EmailTemplateID,omitempty"`
	SendEmailTo                *string            `json:"SendEmailTo,omitempty"`
	EmailBodyPlaceholderValues *map[string]string `json:"EmailBodyPlaceholderValues,omitempty"`
	LanguageCode               *string            `json:"LanguageCode,omitempty"`
}
