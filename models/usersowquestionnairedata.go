package models

type UserSowQuestionnaireData struct {
	SowQuestions []*UserSowQuestionData `json:"SowQuestions,omitempty"`
}
