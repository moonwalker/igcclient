package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"strconv"
	"net/url"
)

type LanguagesService service

// Returns a list of cached Languages supported by the system. If parameter 'isLiveOnly' is not sent as part of the request, the default value is set
func (s *LanguagesService) Languages(isLiveOnly bool) (response OperationResponseOfListOfLanguageObject, err error) {
	b := strconv.FormatBool(isLiveOnly)
	err = s.client.apiPost("/languages?isLiveOnly="+url.QueryEscape(b), nil, &response, nil, nil)
	return
}

// Get language properties by language ID
func (s *LanguagesService) LanguagesById(id int) (response OperationResponseOfLanguageObject, err error) {
	i := strconv.Itoa(id)
	err = s.client.apiPost("/languages/"+url.QueryEscape(i), nil, &response, nil, nil)
	return
}

// Get language properties by alpha code 2
func (s *LanguagesService) LanguagesByAlphaCode2(alphaCode2 string) (response OperationResponseOfLanguageObject, err error) {
	err = s.client.apiPost("/languages/"+url.QueryEscape(alphaCode2), nil, &response, nil, nil)
	return
}

// Get language properties by alpha code 3
func (s *LanguagesService) LanguagesByAlphaCode3(alphaCode3 string) (response OperationResponseOfLanguageObject, err error) {
	err = s.client.apiPost("/languages/"+url.QueryEscape(alphaCode3), nil, &response, nil, nil)
	return
}

// Get language properties by name
func (s *LanguagesService) LanguagesByName(name string) (response OperationResponseOfLanguageObject, err error) {
	err = s.client.apiPost("/languages/"+url.QueryEscape(name), nil, &response, nil, nil)
	return
}
