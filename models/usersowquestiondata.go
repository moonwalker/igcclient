package models

type UserSowQuestionData struct {
	SowQuestionID *int64   `json:"SowQuestionId,omitempty"`
	SowAnswerIDs  *[]int64 `json:"SowAnswerIds,omitempty"`
	FreeText      *string  `json:"FreeText,omitempty"`
}
