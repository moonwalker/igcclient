package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"strconv"
	"net/url"
)

// countries is an private struct
// Use IGCClient.Countries() to access its methods
type CountriesService service

// GetCountries returns all the countries
func (s *CountriesService) Countries() (response OperationResponseOfListOfCountriesObject, err error) {
	err = s.client.apiPost("/countries", nil, &response, nil, nil)
	return
}

// GetCountriesTop returns the top countries
func (s *CountriesService) CountriesTop() (response OperationResponseOfListOfCountriesObject, err error) {
	err = s.client.apiPost("/countries/top", nil, &response, nil, nil)
	return
}

// GetCountryById returns country by country ID
func (s *CountriesService) CountryById(id int64) (response OperationResponseOfCountriesObject, err error) {
	strId := strconv.FormatInt(id, 10)
	err = s.client.apiPost("/countries/"+url.QueryEscape(strId), nil, &response, nil, nil)
	return
}

// GetCountryByAlphaCode2 returns country by the Alpha Code 2
func (s *CountriesService) CountryByAlphaCode2(alphaCode2 string) (response OperationResponseOfCountriesObject, err error) {
	err = s.client.apiPost("/countries/"+url.QueryEscape(alphaCode2), nil, &response, nil, nil)
	return
}

// GetCountryByAlphaCode3 returns country by the Alpha Code 3
func (s *CountriesService) CountryByAlphaCode3(alphaCode3 string) (response OperationResponseOfCountriesObject, err error) {
	err = s.client.apiPost("/countries/"+url.QueryEscape(alphaCode3), nil, &response, nil, nil)
	return
}
