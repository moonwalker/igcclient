package igcclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type RealityCheckService service

// Gets the reality check settings for the authenticated user
func (s *RealityCheckService) GetUserRealityCheck(headers map[string]string, log logger.Logger) (response models.OperationResponseOfRealityCheckUserSetting, err error) {
	err = s.client.apiReq(http.MethodPost, "/realitycheck/getuserrealitycheck", nil, nil, &response, &headers, log)
	return
}

// Saves the reality check interval for the authenticated user
func (s *RealityCheckService) SaveUserRealityCheck(interval int64, headers map[string]string, log logger.Logger) (response models.OperationResponse, err error) {
	q := url.Values{}
	q.Add("interval", fmt.Sprintf("%d", interval))
	err = s.client.apiReq(http.MethodPost, "/realitycheck/saveuserrealitycheck", &q, nil, &response, &headers, log)
	return
}

// Gets the reality check options for the current user's country of access
func (s *RealityCheckService) GetRealityCheckOptions(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfRealityCheckOption, err error) {
	err = s.client.apiReq(http.MethodPost, "/realitycheck/getrealitycheckoptions", nil, nil, &response, &headers, log)
	return
}
