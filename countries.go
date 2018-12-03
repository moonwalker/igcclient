package igcclient

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

// countries is an private struct
// Use IGCClient.Countries() to access its methods
type CountriesService service

// GetCountries returns all the countries
func (s *CountriesService) Countries(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfCountriesObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/countries", nil, nil, &response, &headers, log)
	return
}

// GetCountriesTop returns the top countries
func (s *CountriesService) CountriesTop(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfCountriesObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/countries/top", nil, nil, &response, &headers, log)
	return
}

// GetCountryByID returns country by country ID
func (s *CountriesService) CountryByID(id int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfCountriesObject, err error) {
	strID := strconv.FormatInt(id, 10)
	err = s.client.apiReq(http.MethodPost, "/countries/"+url.QueryEscape(strID), nil, nil, &response, &headers, log)
	return
}

// GetCountryByAlphaCode2 returns country by the Alpha Code 2
func (s *CountriesService) CountryByAlphaCode2(alphaCode2 string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfCountriesObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/countries/"+url.QueryEscape(alphaCode2), nil, nil, &response, &headers, log)
	return
}

// GetCountryByAlphaCode3 returns country by the Alpha Code 3
func (s *CountriesService) CountryByAlphaCode3(alphaCode3 string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfCountriesObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/countries/"+url.QueryEscape(alphaCode3), nil, nil, &response, &headers, log)
	return
}
