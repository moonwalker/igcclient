package igcclient

import (
	"fmt"

	"github.com/moonwalker/igcclient/models"
)

type RealityCheckService service

// Gets the reality check settings for the authenticated user
func (s *RealityCheckService) GetUserRealityCheck(authToken string) (response models.OperationResponseOfRealityCheckUserSetting, err error) {
	err = s.client.apiPost("/realitycheck/getuserrealitycheck", nil, &response, nil, &authToken)
	return
}

// Saves the reality check interval for the authenticated user
func (s *RealityCheckService) SaveUserRealityCheck(interval int64, authToken string) (response models.OperationResponse, err error) {
	err = s.client.apiPost(fmt.Sprintf("/realitycheck/saveuserrealitycheck?interval=%d", interval), nil, &response, nil, &authToken)
	return
}

// Gets the reality check options for the current user's country of access
func (s *RealityCheckService) GetRealityCheckOptions(authToken string) (response models.OperationResponseOfListOfRealityCheckOption, err error) {
	err = s.client.apiPost("/realitycheck/getrealitycheckoptions", nil, &response, nil, &authToken)
	return
}
