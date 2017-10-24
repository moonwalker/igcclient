package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"strconv"
	"net/url"
)

type ValidationService service

// Determines whether the mobile and prefix combination isvalid.
// If ignoreExisting is set to true [ignore existing].
func (s *ValidationService) Mobile(body ValidationMobileModel, ignoreExisting bool) (response OperationResponseOfBoolean, err error) {
	ignore := strconv.FormatBool(ignoreExisting)
	err = s.client.apiPost("/validate/mobile?ignoreExisting="+url.QueryEscape(ignore), &body, &response, nil, nil)
	return
}

// Determines whether the is username available and valid.
// If ignoreExisting is set to true [ignore existing].
func (s *ValidationService) Username(body ValidationUsernameModel, ignoreExisting bool) (response OperationResponseOfBoolean, err error) {
	ignore := strconv.FormatBool(ignoreExisting)
	err = s.client.apiPost("/validate/username?ignoreExisting="+url.QueryEscape(ignore), &body, &response, nil, nil)
	return
}

// Determines whether the mobile and prefix combination isvalid.
// If ignoreExisting is set to true [ignore existing].
func (s *ValidationService) Email(body ValidationEmailModel, ignoreExisting bool) (response OperationResponseOfBoolean, err error) {
	ignore := strconv.FormatBool(ignoreExisting)
	err = s.client.apiPost("/validate/email?ignoreExisting="+url.QueryEscape(ignore), &body, &response, nil, nil)
	return
}
