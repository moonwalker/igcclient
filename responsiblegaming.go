package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"net/url"
	"strconv"
)

type ResponsibleGamingService service

//Fetches the responsible limit history for the past 6 months for the current logged in user.
func (s *ResponsibleGamingService) GetUserLimitHistory(authToken string) (response OperationResponseOfIEnumerableOfUserLimitResponse, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Limits/GetUserLimitHistory", nil, &response, nil, &authToken)
	return
}

//Retrieves all self exclusion categories.
func (s *ResponsibleGamingService) SelfExclusionCategories() (response OperationResponseOfListOfSelfExclusionCategory, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/SelfExclusion/Categories", nil, &response, nil, nil)
	return
}

//Retrieves all self exclusion durations.
func (s *ResponsibleGamingService) SelfExclusionDurations() (response OperationResponseOfListOfSelfExclusionCategoryDuration, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/SelfExclusion/Durations", nil, &response, nil, nil)
	return
}

//Retrieves all timeout categories.
func (s *ResponsibleGamingService) TimeoutCategories() (response OperationResponseOfListOfSelfExclusionCategory, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Timeout/Categories", nil, &response, nil, nil)
	return
}

//Retrieves all timeout durations.
func (s *ResponsibleGamingService) TimeoutDurations() (response OperationResponseOfListOfSelfExclusionCategoryDuration, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Timeout/Durations", nil, &response, nil, nil)
	return
}

//Add a self-exclusion for the current authenticated user.
func (s *ResponsibleGamingService) AddSelfExclusion(selfExclusionCategoryID int64, authToken string) (response OperationResponseOfBoolean, err error) {
	i := strconv.FormatInt(selfExclusionCategoryID, 10)
	err = s.client.apiPost("/v2/ResponsibleGaming/SelfExclusion/AddSelfExclusion?selfExclusionCategoryID="+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

//Add a timeout for the current authenticated user.
func (s *ResponsibleGamingService) AddTimeoutByEndDate(endDate string, authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Timeout/AddTimeout?endDate="+endDate, nil, &response, nil, &authToken)
	return
}

//Add a timeout for the current authenticated user.
func (s *ResponsibleGamingService) AddTimeoutByTimeoutCategoryID(timeoutCategoryID int64, authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Timeout/AddTimeout?timeoutCategoryID={timeoutCategoryID}", nil, &response, nil, &authToken)
	return
}

//Responsible for returning the active and future limits for a specific user (Headers supported AuthenticationToken The list of Limits will contains all important limit details and for session limits Amount = Amount in mins Amount_HR - amount in hour Amount_DAY - Amount in Days for other limits ignore Amount_HR and Amount_DAY are not applicable
func (s *ResponsibleGamingService) GetUserLimits(authToken string) (response OperationResponseOfIEnumerableOfUserLimitResponse, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Limits/GetUserLimits", nil, &response, nil, &authToken)
	return
}

//No documentation available.
func (s *ResponsibleGamingService) SetUserLimits(body []SetUserLimitModelV2, authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Limits/SetUserLimits", &body, &response, nil, &authToken)
	return
}

//No documentation available.
func (s *ResponsibleGamingService) GetLimits(authToken string) (response OperationResponseOfIEnumerableOfUserLimitResponse, err error) {
	err = s.client.apiPost("/v2/ResponsibleGaming/Limits/GetLimits", nil, &response, nil, &authToken)
	return
}
