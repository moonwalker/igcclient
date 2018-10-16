package igcclient

import (
	"fmt"
	"net/url"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type RolesService service

// Gets list of roles associated to a user
func (s *RolesService) RolesByUserID(userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfRoleResponse, err error) {
	q := url.Values{}
	q.Add("userId", fmt.Sprintf("%d", userID))
	err = s.client.apiPost("/roles/user", &q, nil, &response, &headers, log)
	return
}

// Fetches all existing roles.
func (s *RolesService) Roles(headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfRoleResponse, err error) {
	err = s.client.apiPost("/roles", nil, nil, &response, &headers, log)
	return
}
