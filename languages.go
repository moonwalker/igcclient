package igcclient

import (
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
)

type LanguagesService service

// Returns a list of cached Languages supported by the system. If parameter 'isLiveOnly' is not sent as part of the request, the default value is set
func (s *LanguagesService) Languages(isLiveOnly bool, headers map[string]string) (response models.OperationResponseOfListOfLanguageObject, err error) {
	b := strconv.FormatBool(isLiveOnly)
	q := url.Values{}
	q.Add("isLiveOnly", b)
	err = s.client.apiPost("/languages", &q, nil, &response, &headers)
	return
}

// Get language properties by language ID
func (s *LanguagesService) LanguagesByID(id int64, headers map[string]string) (response models.OperationResponseOfLanguageObject, err error) {
	i := strconv.FormatInt(id, 10)
	err = s.client.apiPost("/languages/"+url.QueryEscape(i), nil, nil, &response, &headers)
	return
}

// Get language properties by alpha code 2
func (s *LanguagesService) LanguagesByAlphaCode2(alphaCode2 string, headers map[string]string) (response models.OperationResponseOfLanguageObject, err error) {
	err = s.client.apiPost("/languages/"+url.QueryEscape(alphaCode2), nil, nil, &response, &headers)
	return
}

// Get language properties by alpha code 3
func (s *LanguagesService) LanguagesByAlphaCode3(alphaCode3 string, headers map[string]string) (response models.OperationResponseOfLanguageObject, err error) {
	err = s.client.apiPost("/languages/"+url.QueryEscape(alphaCode3), nil, nil, &response, &headers)
	return
}

// Get language properties by name
func (s *LanguagesService) LanguagesByName(name string, headers map[string]string) (response models.OperationResponseOfLanguageObject, err error) {
	err = s.client.apiPost("/languages/"+url.QueryEscape(name), nil, nil, &response, &headers)
	return
}
