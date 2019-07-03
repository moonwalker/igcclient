package models

type SowQuestion struct {
	ID              *int64           `json:"Id,omitempty"`
	Text            *string          `json:"Text,omitempty"`
	SowQuestionType *SowQuestionType `json:"SowQuestionType,omitempty"`
	SowAnswers      []*SowAnswer     `json:"SowAnswers,omitempty"`
}
