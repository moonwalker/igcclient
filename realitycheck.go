package igcclient

import (
	"net/url"
	"strconv"

	. "github.com/moonwalker/igcclient/models"
)

type RealityCheckService service

// Gets the reality check settings for the authenticated user
func (s *RealityCheckService) GetUserRealityCheck(authToken string) (response OperationResponseOfRealityCheckUserSetting, err error) {
	err = s.client.apiPost("/realitycheck/getuserrealitycheck", nil, &response, nil, &authToken)
	return
}

// Saves the reality check interval for the authenticated user
func (s *RealityCheckService) SaveUserRealityCheck(interval int64, authToken string) (response OperationResponse, err error) {
	i := strconv.FormatInt(interval, 10)
	err = s.client.apiPost("/realitycheck/saveuserrealitycheck?interval="+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Gets the reality check options for the current user's country of access
func (s *RealityCheckService) GetRealityCheckOptions(authToken string) (response OperationResponseOfListOfRealityCheckOption, err error) {
	err = s.client.apiPost("/realitycheck/getrealitycheckoptions", nil, &response, nil, &authToken)
	return
}
