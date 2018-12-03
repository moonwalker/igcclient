package igcclient

import (
	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
	"net/http"
)

type SecurityQuestionsService service

func (s *SecurityQuestionsService) SecurityQuestions(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfSecurityQuestionModel, err error) {
	err = s.client.apiReq(http.MethodPost, "/securityquestions", nil, nil, &response, &headers, log)
	return
}
