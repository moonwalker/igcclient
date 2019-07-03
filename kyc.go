package igcclient

import (
	"net/http"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type KYCService service

// Gets the Source of Wealth Questionnaire data - questions along with question types and possible answers. Headers required: AuthenticationToken
func (s *KYCService) GetSowQuestionnaire(headers map[string]string, log logger.Logger) (response models.OperationResponseOfSowQuestionnaire, err error) {
	err = s.client.apiReq(http.MethodGet, "/kyc/getsowquestionnaire", nil, nil, &response, &headers, log)
	return
}

// Saves the Source of Wealth Questionnaire user answers. Headers required: AuthenticationToken
func (s *KYCService) PostSowQuestionnaire(body models.UserSowQuestionnaireData, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/kyc/postsowquestionnaire", nil, &body, &response, &headers, log)
	return
}
