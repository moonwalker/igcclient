package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"strconv"
	"net/url"
)

type CurrenciesService service

// Gets all the currencies
func (s *CurrenciesService) Currencies() (response OperationResponseOfListOfCurrenciesObject, err error) {
	err = s.client.apiPost("/currencies", nil, &response, nil, nil)
	return
}

// Gets a single currency by currency ID
func (s *CurrenciesService) CurrencyById(id int) (response OperationResponseOfCurrenciesObject, err error) {
	i := strconv.Itoa(id)
	err = s.client.apiPost("/currencies/"+url.QueryEscape(i), nil, &response, nil, nil)
	return
}

// Gets a single curency by currency code
func (s *CurrenciesService) CurrencyByCode(code string) (response OperationResponseOfCurrenciesObject, err error) {
	err = s.client.apiPost("/currencies/"+url.QueryEscape(code), nil, &response, nil, nil)
	return
}
