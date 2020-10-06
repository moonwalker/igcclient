package models

type ChangeSecurityQuestionModel struct {
	Password           *string `json:"Password,omitempty"`
	SecurityQuestionID *int64  `json:"SecurityQuestionID,omitempty"`
	SecurityAnswer     *string `json:"SecurityAnswer,omitempty"`
}
