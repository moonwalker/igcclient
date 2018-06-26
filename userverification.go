package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
)

type UserVerificationService service

// Gets the available verification types for the user's ip. User Verification types can be whitelisted for each country.
func (s *UserVerificationService) RegistrationTypes(headers map[string]string) (response OperationResponseOfListOfCountryVerificationTypeWhitelist, err error) {
	err = s.client.apiPost("/user/verify/registration/types", nil, &response, &headers)
	return
}
