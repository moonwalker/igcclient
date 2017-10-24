package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"net/url"
	"strconv"
)

type ResponsibleGamingService service

// Retrieves all self exclusion categories.
func (s *ResponsibleGamingService) SelfExclusionCategories() (response OperationResponseOfListOfSelfExclusionCategory, err error) {
	err = s.client.apiPost("/responsiblegaming/selfexclusion/categories", nil, &response, nil, nil)
	return
}

// Retrieves all self exclusion durations.
func (s *ResponsibleGamingService) SelfExclusionDurations() (response OperationResponseOfListOfSelfExclusionCategoryDuration, err error) {
	err = s.client.apiPost("/responsiblegaming/selfexclusion/durations", nil, &response, nil, nil)
	return
}

// Add an exclusion for the current authenticated user.
func (s *ResponsibleGamingService) SelfExclusionAdd(selfExclusionCategoryID int, authToken string) (response OperationResponseOfBoolean, err error) {
	i := strconv.Itoa(selfExclusionCategoryID)
	err = s.client.apiPost("/v2/responsiblegaming/selfexclusion/add?selfexclusioncategoryid="+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Add a timeout for the current authenticated user. Contrary to SelfExclusions, Timeouts are not propagated to other brands running on the same license.
func (s *ResponsibleGamingService) TimeoutAdd(timeoutCategoryID int, authToken string) (response OperationResponseOfBoolean, err error) {
	i := strconv.Itoa(timeoutCategoryID)
	err = s.client.apiPost("/v2/responsiblegaming/timeout/add?timeoutcategoryid="+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Responsible for returning the active and future limits for a specific user. Requires the AuthenticationToken header to be set
func (s *ResponsibleGamingService) LimitsGetUserLimits(authToken string) (response OperationResponseOfObject, err error) {
	err = s.client.apiPost("/v2/responsiblegaming/limits/getuserlimits", nil, &response, nil, &authToken)
	return
}

// Responsible for returning the active and future limits for a specific user. Requires the X-Api-Key header to be set
func (s *ResponsibleGamingService) LimitsGetUserLimitsByUserId(userId int, xApiKey string) (response OperationResponseOfObject, err error) {
	i := strconv.Itoa(userId)
	err = s.client.apiPost("/v2/responsiblegaming/limits/getuserlimits?userid="+url.QueryEscape(i), nil, &response, &xApiKey, nil)
	return
}
