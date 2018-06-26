package igcclient

import "github.com/moonwalker/igcclient/models"

type SecurityQuestionsService service

func (s *SecurityQuestionsService) SecurityQuestions() (response models.OperationResponseOfListOfSecurityQuestionModel, err error) {
	err = s.client.apiPost("/securityquestions", nil, &response, nil)
	return
}
