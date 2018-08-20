package igcclient

import (
	"github.com/moonwalker/igcclient/models"
)

type SecurityQuestionsService service

func (s *SecurityQuestionsService) SecurityQuestions(headers map[string]string) (response models.OperationResponseOfListOfSecurityQuestionModel, err error) {
	err = s.client.apiPost("/securityquestions", nil, nil, &response, &headers)
	return
}
