package igcclient

import (
	"net/url"
	"strconv"

	. "github.com/moonwalker/igcclient/models"
)

type RolesService service

// Gets list of roles associated to a user
func (s *RolesService) RolesByUserID(userID int64, xAPIKey string) (response OperationResponseOfIEnumerableOfRoleResponse, err error) {
	id := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/roles/user?userId="+url.QueryEscape(id), nil, &response, &xAPIKey, nil)
	return
}

// Fetches all existing roles.
func (s *RolesService) Roles(xAPIKey string) (response OperationResponseOfIEnumerableOfRoleResponse, err error) {
	err = s.client.apiPost("/roles", nil, &response, &xAPIKey, nil)
	return
}
