package models

type UserSecurityQuestionRequestDTO struct {
	SecurityQuestionID *int64  `json:"SecurityQuestionId"`
	SecurityAnswer     *string `json:"SecurityAnswer"`
}
