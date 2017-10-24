package models

type SecurityQuestion struct {
	ID       *int    `json:"ID,omitempty"`
	Question *string `json:"Question,omitempty"`
}
