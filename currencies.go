package igcclient

import (
	"net/url"
	"strconv"

	. "github.com/moonwalker/igcclient/models"
)

type CurrenciesService service

// Gets all the currencies
func (s *CurrenciesService) Currencies() (response OperationResponseOfListOfCurrenciesObject, err error) {
	err = s.client.apiPost("/currencies", nil, &response, nil, nil)
	return
}

// Gets a single currency by currency ID
func (s *CurrenciesService) CurrencyByID(id int64) (response OperationResponseOfCurrenciesObject, err error) {
	i := strconv.FormatInt(id, 10)
	err = s.client.apiPost("/currencies/"+url.QueryEscape(i), nil, &response, nil, nil)
	return
}

// Gets a single curency by currency code
func (s *CurrenciesService) CurrencyByCode(code string) (response OperationResponseOfCurrenciesObject, err error) {
	err = s.client.apiPost("/currencies/"+url.QueryEscape(code), nil, &response, nil, nil)
	return
}
