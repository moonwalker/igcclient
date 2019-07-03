package models

type SowQuestionType int64

const (
	SQT_SingleAnswer       SowQuestionType = 1
	SQT_MultipleAnswer     SowQuestionType = 2
	SQT_MultipleOrNoAnswer SowQuestionType = 3
	SQT_Country            SowQuestionType = 4
)
