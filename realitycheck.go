package igcclient

import (
	"fmt"

	. "github.com/moonwalker/igcclient/models"
)

type RealityCheckService service

// Gets the reality check settings for the authenticated user
func (s *RealityCheckService) GetUserRealityCheck(headers map[string]string) (response models.OperationResponseOfRealityCheckUserSetting, err error) {
	err = s.client.apiPost("/realitycheck/getuserrealitycheck", nil, &response, &headers)
	return
}

// Saves the reality check interval for the authenticated user
func (s *RealityCheckService) SaveUserRealityCheck(interval int64, headers map[string]string) (response models.OperationResponse, err error) {
	err = s.client.apiPost(fmt.Sprintf("/realitycheck/saveuserrealitycheck?interval=%d", interval), nil, &response, &headers)
	return
}

// Gets the reality check options for the current user's country of access
func (s *RealityCheckService) GetRealityCheckOptions(headers map[string]string) (response models.OperationResponseOfListOfRealityCheckOption, err error) {
	err = s.client.apiPost("/realitycheck/getrealitycheckoptions", nil, &response, &headers)
	return
}
