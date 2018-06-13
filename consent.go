package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"strconv"
)

type ConsentService service

// Get All Consents by Language X-Api-Key is required.
func (s *ConsentService) GetConsents(languageAlpha2Code string, xApiKey string) (response OperationResponseOfListOfPublicConsentModel, err error) {
	err = s.client.apiPost("/Consent/GetConsents?languageAlpha2Code="+languageAlpha2Code, nil, &response, &xApiKey, nil)
	return
}

// Checks if user has any pending mandatory consents
func (s *ConsentService) HasPendingConsents(authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/Consent/HasPendingConsents", nil, &response, nil, &authToken)
	return
}

// Checks if user has any migrated consents that require and update
func (s *ConsentService) RequiresLegacyUpdate(authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/Consent/RequiresLegacyUpdate", nil, &response, nil, &authToken)
	return
}

// Gets mandatory consents with content which user need to consent to by language
func (s *ConsentService) UserPendingConsents(languageAlpha2Code string, authToken string) (response OperationResponseOfPublicUserConsentsListModel, err error) {
	err = s.client.apiPost("/Consent/UserPendingConsents?languageAlpha2Code="+languageAlpha2Code, nil, &response, nil, &authToken)
	return
}

// Save user consents
func (s *ConsentService) SaveUserConsents(body []ConsentVersionSaveModel, authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/bonuses", &body, &response, nil, &authToken)
	return
}

// Gets list of user consents with content and user's acceptance
func (s *ConsentService) UserConsents(languageAlpha2Code string, authToken string) (response OperationResponseOfPublicUserConsentsListModel, err error) {
	err = s.client.apiPost("/Consent/UserConsents?languageAlpha2Code="+languageAlpha2Code, nil, &response, nil, &authToken)
	return
}

// Gets the user current bonuses AuthenticationToken is required.
func (s *ConsentService) Unsubscribe(triggerCode string, userId int64, xApiKey string) (response OperationResponseOfObject, err error) {
	id := strconv.FormatInt(userId, 10)
	err = s.client.apiPost("/Consent/Unsubscribe?triggerCode="+triggerCode+"&userId="+id, nil, &response, &xApiKey, nil)
	return
}
