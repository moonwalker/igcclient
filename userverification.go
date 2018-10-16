package igcclient

import (
	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type UserVerificationService service

// Gets the available verification types for the user's ip. User Verification types can be whitelisted for each country.
func (s *UserVerificationService) RegistrationTypes(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfCountryVerificationTypeWhitelist, err error) {
	err = s.client.apiPost("/user/verify/registration/types", nil, nil, &response, &headers, log)
	return
}
