package igcclient

import (
	"fmt"

	"github.com/moonwalker/igcclient/models"
)

type ConsentService service

// Get All Consents by Language X-Api-Key is required.
func (s *ConsentService) GetConsents(languageAlpha2Code string, headers map[string]string) (response models.OperationResponseOfListOfPublicConsentModel, err error) {
	err = s.client.apiPost("/Consent/GetConsents?languageAlpha2Code="+languageAlpha2Code, nil, &response, &headers)
	return
}

// Checks if user has any pending mandatory consents
func (s *ConsentService) HasPendingConsents(headers map[string]string) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/Consent/HasPendingConsents", nil, &response, &headers)
	return
}

// Checks if user has any migrated consents that require and update
func (s *ConsentService) RequiresLegacyUpdate(headers map[string]string) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/Consent/RequiresLegacyUpdate", nil, &response, &headers)
	return
}

// Gets mandatory consents with content which user need to consent to by language
func (s *ConsentService) UserPendingConsents(languageAlpha2Code string, headers map[string]string) (response models.OperationResponseOfPublicUserConsentsListModel, err error) {
	err = s.client.apiPost("/Consent/UserPendingConsents?languageAlpha2Code="+languageAlpha2Code, nil, &response, &headers)
	return
}

// Save user consents
func (s *ConsentService) SaveUserConsents(body []models.ConsentVersionSaveModel, userID int64, headers map[string]string) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiPost(fmt.Sprintf("/Consent/SaveUserConsents?userId=%d", userID), &body, &response, &headers)
	return
}

// Gets list of user consents with content and user's acceptance
func (s *ConsentService) UserConsents(languageAlpha2Code string, headers map[string]string) (response models.OperationResponseOfPublicUserConsentsListModel, err error) {
	err = s.client.apiPost("/Consent/UserConsents?languageAlpha2Code="+languageAlpha2Code, nil, &response, &headers)
	return
}

// Unsubscribe all Consents by trigger code This will require CRM X-api key:
func (s *ConsentService) Unsubscribe(triggerCode string, userID int64, headers map[string]string) (response models.OperationResponseOfObject, err error) {
	err = s.client.apiPost(fmt.Sprintf("/Consent/Unsubscribe?triggerCode=%s&userId=%d", triggerCode, userID), nil, &response, &headers)
	return
}
