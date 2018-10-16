package igcclient

import (
	"fmt"
	"net/url"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type ConsentService service

// Get All Consents by Language X-Api-Key is required.
func (s *ConsentService) GetConsents(languageAlpha2Code string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfPublicConsentModel, err error) {
	q := url.Values{}
	q.Add("languageAlpha2Code", languageAlpha2Code)
	err = s.client.apiPost("/Consent/GetConsents", &q, nil, &response, &headers, log)
	return
}

// Checks if user has any pending mandatory consents
func (s *ConsentService) HasPendingConsents(headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/Consent/HasPendingConsents", nil, nil, &response, &headers, log)
	return
}

// Checks if user has any migrated consents that require and update
func (s *ConsentService) RequiresLegacyUpdate(headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/Consent/RequiresLegacyUpdate", nil, nil, &response, &headers, log)
	return
}

// Gets mandatory consents with content which user need to consent to by language
func (s *ConsentService) UserPendingConsents(languageAlpha2Code string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfPublicUserConsentsListModel, err error) {
	q := url.Values{}
	q.Add("languageAlpha2Code", languageAlpha2Code)
	err = s.client.apiPost("/Consent/UserPendingConsents", &q, nil, &response, &headers, log)
	return
}

// Save user consents
func (s *ConsentService) SaveUserConsents(body []models.ConsentVersionSaveModel, userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	q := url.Values{}
	q.Add("userId", fmt.Sprintf("%d", userID))
	err = s.client.apiPost("/Consent/SaveUserConsents", &q, &body, &response, &headers, log)
	return
}

// Gets list of user consents with content and user's acceptance
func (s *ConsentService) UserConsents(languageAlpha2Code string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfPublicUserConsentsListModel, err error) {
	q := url.Values{}
	q.Add("languageAlpha2Code", languageAlpha2Code)
	err = s.client.apiPost("/Consent/UserConsents", &q, nil, &response, &headers, log)
	return
}

// Unsubscribe all Consents by trigger code This will require CRM X-api key:
func (s *ConsentService) Unsubscribe(triggerCode string, userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfObject, err error) {
	q := url.Values{}
	q.Add("triggerCode", triggerCode)
	q.Add("userId", fmt.Sprintf("%d", userID))
	err = s.client.apiPost("/Consent/Unsubscribe", &q, nil, &response, &headers, log)
	return
}
