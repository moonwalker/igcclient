package igcclient

import (
	"fmt"

	"github.com/moonwalker/igcclient/models"
)

type ResponsibleGamingService service

//Fetches the responsible limit history for the past 6 months for the current logged in user.
func (s *ResponsibleGamingService) GetUserLimitHistory(headers map[string]string) (response models.OperationResponseOfIEnumerableOfUserLimitResponse, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Limits/GetUserLimitHistory", nil, &response, &headers)
	return
}

//Retrieves all self exclusion categories.
func (s *ResponsibleGamingService) SelfExclusionCategories(headers map[string]string) (response models.OperationResponseOfListOfSelfExclusionCategory, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/SelfExclusion/Categories", nil, &response, &headers)
	return
}

//Retrieves all self exclusion durations.
func (s *ResponsibleGamingService) SelfExclusionDurations(headers map[string]string) (response models.OperationResponseOfListOfSelfExclusionCategoryDuration, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/SelfExclusion/Durations", nil, &response, &headers)
	return
}

//Retrieves all timeout categories.
func (s *ResponsibleGamingService) TimeoutCategories(headers map[string]string) (response models.OperationResponseOfListOfSelfExclusionCategory, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Timeout/Categories", nil, &response, &headers)
	return
}

//Retrieves all timeout durations.
func (s *ResponsibleGamingService) TimeoutDurations(headers map[string]string) (response models.OperationResponseOfListOfSelfExclusionCategoryDuration, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Timeout/Durations", nil, &response, &headers)
	return
}

//Add a self-exclusion for the current authenticated user.
func (s *ResponsibleGamingService) AddSelfExclusion(selfExclusionCategoryID int64, headers map[string]string) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiPost(fmt.Sprintf("/v2/ResponsibleGaming/SelfExclusion/AddSelfExclusion?selfExclusionCategoryID=%d", selfExclusionCategoryID), nil, &response, &headers)
	return
}

//Add a timeout for the current authenticated user.
func (s *ResponsibleGamingService) AddTimeoutByEndDate(endDate string, headers map[string]string) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiPost(fmt.Sprintf("/v2/ResponsibleGaming/Timeout/AddTimeout?endDate=%s", endDate), nil, &response, &headers)
	return
}

//Add a timeout for the current authenticated user.
func (s *ResponsibleGamingService) AddTimeoutByTimeoutCategoryID(timeoutCategoryID int64, headers map[string]string) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiPost(fmt.Sprintf("/v2/ResponsibleGaming/Timeout/AddTimeout?timeoutCategoryID=%d", timeoutCategoryID), nil, &response, &headers)
	return
}

//Responsible for returning the active and future limits for a specific user (Headers supported AuthenticationToken The list of Limits will contains all important limit details and for session limits Amount = Amount in mins Amount_HR - amount in hour Amount_DAY - Amount in Days for other limits ignore Amount_HR and Amount_DAY are not applicable
func (s *ResponsibleGamingService) GetUserLimits(headers map[string]string) (response models.OperationResponseOfIEnumerableOfUserLimitResponse, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Limits/GetUserLimits", nil, &response, &headers)
	return
}

//No documentation available.
func (s *ResponsibleGamingService) SetUserLimits(body []models.SetUserLimitModelV2, headers map[string]string) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Limits/SetUserLimits", &body, &response, &headers)
	return
}

//No documentation available.
func (s *ResponsibleGamingService) GetLimits(headers map[string]string) (response models.OperationResponseOfIEnumerableOfUserLimitResponse, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Limits/GetLimits", nil, &response, &headers)
	return
}
