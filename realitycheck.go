package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"strconv"
	"net/url"
)

type RealityCheckService service

// Gets the reality check settings for the authenticated user
func (s *RealityCheckService) GetUserRealityCheck(authToken string) (response OperationResponseOfRealityCheckUserSetting, err error) {
	err = s.client.apiPost("realitycheck/getuserrealitycheck", nil, &response, nil, &authToken)
	return
}

// Saves the reality check interval for the authenticated user
func (s *RealityCheckService) SaveUserRealityCheck(interval int, authToken string) (response OperationResponse, err error) {
	i := strconv.Itoa(interval)
	err = s.client.apiPost("realitycheck/saveuserrealitycheck?interval="+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Gets the reality check options for the current user's country of access
func (s *RealityCheckService) GetRealityCheckOptions(authToken string) (response OperationResponseOfListOfRealityCheckOption, err error) {
	err = s.client.apiPost("realitycheck/getrealitycheckoptions", nil, &response, nil, &authToken)
	return
}