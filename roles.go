package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"strconv"
	"net/url"
)

type RolesService service

// Gets list of roles associated to a user
func (s *RolesService) RolesByUserId(userId int, xApiKey string) (response OperationResponseOfIEnumerableOfRoleResponse, err error) {
	id := strconv.Itoa(userId)
	err = s.client.apiPost("/roles/user?userId="+url.QueryEscape(id), nil, &response, &xApiKey, nil)
	return
}

// Fetches all existing roles.
func (s *RolesService) Roles(xApiKey string) (response OperationResponseOfIEnumerableOfRoleResponse, err error) {
	err = s.client.apiPost("/roles", nil, &response, &xApiKey, nil)
	return
}