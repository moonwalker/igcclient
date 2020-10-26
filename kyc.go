package igcclient

import (
	"fmt"
	"net/http"
	"net/url"

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

// Gets the AML verification status. Headers required: AuthenticationToken
func (s *KYCService) GetAMLVerificationStatus(headers map[string]string, log logger.Logger) (response models.OperationResponseOfGetAmlVerificationStatusDTO, err error) {
	err = s.client.apiReq(http.MethodGet, "/kyc/getamlverificationstatus", nil, nil, &response, &headers, log)
	return
}

// Gets the list of KYC types and statuses. Headers required: AuthenticationToken
func (s *KYCService) GetKYCVerificationStatus(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfUserKycStatusResponseDTO, err error) {
	err = s.client.apiReq(http.MethodGet, "/user/getkycstatus", nil, nil, &response, &headers, log)
	return
}

// Get the kyc Status for the specified userId. Available only with X-Api-Key!
func (s *KYCService) GetKYCVerificationStatusByUserID(userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfUserKycStatusResponseDTO, err error) {
	q := url.Values{}
	q.Add("userid", fmt.Sprintf("%d", userID))
	err = s.client.apiReq(http.MethodGet, "/user/getkycstatus", &q, nil, &response, &headers, log)
	return
}
