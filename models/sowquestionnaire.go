package models

type SowQuestionnaire struct {
	SowQuestions []*SowQuestion `json:"SowQuestions,omitempty"`
}
