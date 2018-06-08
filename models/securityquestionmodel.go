package models

type SecurityQuestionModel struct {
	ID       *int64  `json:"ID,omitempty"`
	Question *string `json:"Question,omitempty"`
}
