package models

type ResetPasswordModel struct {
	BirthDate         *IGCTime                          `json:"BirthDate,omitempty"`
	SecurityQuestions *[]UserSecurityQuestionRequestDTO `json:"SecurityQuestions,omitempty"`
	Email             *string                           `json:"email"`
}
