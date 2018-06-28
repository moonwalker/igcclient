package igcclient

import (
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
)

// countries is an private struct
// Use IGCClient.Countries() to access its methods
type CountriesService service

// GetCountries returns all the countries
func (s *CountriesService) Countries(headers map[string]string) (response models.OperationResponseOfListOfCountriesObject, err error) {
	err = s.client.apiPost("/countries", nil, &response, &headers)
	return
}

// GetCountriesTop returns the top countries
func (s *CountriesService) CountriesTop(headers map[string]string) (response models.OperationResponseOfListOfCountriesObject, err error) {
	err = s.client.apiPost("/countries/top", nil, &response, &headers)
	return
}

// GetCountryByID returns country by country ID
func (s *CountriesService) CountryByID(id int64, headers map[string]string) (response models.OperationResponseOfCountriesObject, err error) {
	strID := strconv.FormatInt(id, 10)
	err = s.client.apiPost("/countries/"+url.QueryEscape(strID), nil, &response, &headers)
	return
}

// GetCountryByAlphaCode2 returns country by the Alpha Code 2
func (s *CountriesService) CountryByAlphaCode2(alphaCode2 string, headers map[string]string) (response models.OperationResponseOfCountriesObject, err error) {
	err = s.client.apiPost("/countries/"+url.QueryEscape(alphaCode2), nil, &response, &headers)
	return
}

// GetCountryByAlphaCode3 returns country by the Alpha Code 3
func (s *CountriesService) CountryByAlphaCode3(alphaCode3 string, headers map[string]string) (response models.OperationResponseOfCountriesObject, err error) {
	err = s.client.apiPost("/countries/"+url.QueryEscape(alphaCode3), nil, &response, &headers)
	return
}
