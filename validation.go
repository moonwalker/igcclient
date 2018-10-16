package igcclient

import (
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type ValidationService service

// Determines whether the mobile and prefix combination isvalid.
// If ignoreExisting is set to true [ignore existing].
func (s *ValidationService) Mobile(body models.ValidationMobileModel, ignoreExisting bool, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	ignore := strconv.FormatBool(ignoreExisting)
	q := url.Values{}
	q.Add("ignoreExisting", ignore)
	err = s.client.apiPost("/validate/mobile", &q, &body, &response, &headers, log)
	return
}

// Determines whether the is username available and valid.
// If ignoreExisting is set to true [ignore existing].
func (s *ValidationService) Username(body models.ValidationUsernameModel, ignoreExisting bool, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	ignore := strconv.FormatBool(ignoreExisting)
	q := url.Values{}
	q.Add("ignoreExisting", ignore)
	err = s.client.apiPost("/validate/username", &q, &body, &response, &headers, log)
	return
}

// Determines whether the mobile and prefix combination isvalid.
// If ignoreExisting is set to true [ignore existing].
func (s *ValidationService) Email(body models.ValidationEmailModel, ignoreExisting bool, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	ignore := strconv.FormatBool(ignoreExisting)
	q := url.Values{}
	q.Add("ignoreExisting", ignore)
	err = s.client.apiPost("/validate/email", &q, &body, &response, &headers, log)
	return
}
