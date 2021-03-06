package igcclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type ResponsibleGamingService service

//Fetches the responsible limit history for the past 6 months for the current logged in user.
func (s *ResponsibleGamingService) GetUserLimitHistory(headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfUserLimitResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/Limits/GetUserLimitHistory", nil, nil, &response, &headers, log)
	return
}

//Retrieves all self exclusion categories.
func (s *ResponsibleGamingService) SelfExclusionCategories(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfSelfExclusionCategory, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/SelfExclusion/Categories", nil, nil, &response, &headers, log)
	return
}

//Retrieves all self exclusion durations.
func (s *ResponsibleGamingService) SelfExclusionDurations(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfSelfExclusionCategoryDuration, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/SelfExclusion/Durations", nil, nil, &response, &headers, log)
	return
}

//Retrieves all timeout categories.
func (s *ResponsibleGamingService) TimeoutCategories(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfSelfExclusionCategory, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/Timeout/Categories", nil, nil, &response, &headers, log)
	return
}

//Retrieves all timeout durations.
func (s *ResponsibleGamingService) TimeoutDurations(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfSelfExclusionCategoryDuration, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/Timeout/Durations", nil, nil, &response, &headers, log)
	return
}

//Add a self-exclusion for the current authenticated user.
func (s *ResponsibleGamingService) AddSelfExclusion(selfExclusionCategoryID int64, blockReason int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	q := url.Values{}
	q.Add("selfExclusionCategoryID", fmt.Sprintf("%d", selfExclusionCategoryID))
	if blockReason != 0 {
		q.Add("blockReason", fmt.Sprintf("%d", blockReason))
	}
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/SelfExclusion/AddSelfExclusion", &q, nil, &response, &headers, log)
	return
}

//Add a timeout for the current authenticated user.
func (s *ResponsibleGamingService) AddTimeoutByEndDate(endDate string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	q := url.Values{}
	q.Add("endDate", endDate)
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/Timeout/AddTimeout", &q, nil, &response, &headers, log)
	return
}

//Add a timeout for the current authenticated user.
func (s *ResponsibleGamingService) AddTimeoutByTimeoutCategoryID(timeoutCategoryID int64, blockReason int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	q := url.Values{}
	q.Add("timeoutCategoryID", fmt.Sprintf("%d", timeoutCategoryID))
	if blockReason != 0 {
		q.Add("blockReason", fmt.Sprintf("%d", blockReason))
	}
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/Timeout/AddTimeout", &q, nil, &response, &headers, log)
	return
}

//Responsible for returning the active and future limits for a specific user (Headers supported AuthenticationToken The list of Limits will contains all important limit details and for session limits Amount = Amount in mins Amount_HR - amount in hour Amount_DAY - Amount in Days for other limits ignore Amount_HR and Amount_DAY are not applicable
func (s *ResponsibleGamingService) GetUserLimits(headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfUserLimitResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/Limits/GetUserLimits", nil, nil, &response, &headers, log)
	return
}

//No documentation available.
func (s *ResponsibleGamingService) SetUserLimits(body []models.SetUserLimitModelV2, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/Limits/SetUserLimits", nil, &body, &response, &headers, log)
	return
}

//No documentation available.
func (s *ResponsibleGamingService) GetLimits(headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfUserLimitResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/Limits/GetLimits", nil, nil, &response, &headers, log)
	return
}

//This api will provide the configured deposit limits information and the total win/loss during the past period of input months
func (s *ResponsibleGamingService) UserPlayStatsAndLimits(months int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfUserRgStatsResponseDTO, err error) {
	if months < 1 {
		months = 1
	} else if months > 12 {
		months = 12
	}
	q := url.Values{}
	q.Add("months", fmt.Sprintf("%d", months))
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/userPlayStatsAndLimits", &q, nil, &response, &headers, log)
	return
}

//Get the last to expire Active Responsible User Blocks end date of a user.
func (s *ResponsibleGamingService) GetEndDateOfActiveResponsibleUserBlocks(userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfUserBlockResponseDTO, err error) {
	q := url.Values{}
	q.Add("userId", fmt.Sprintf("%d", userID))
	err = s.client.apiReq(http.MethodPost, "/v2/ResponsibleGaming/GetEndDateOfActiveResponsibleUserBlocks", &q, nil, &response, &headers, log)
	return
}
