package igcclient

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type CurrenciesService service

// Gets all the currencies
func (s *CurrenciesService) Currencies(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfCurrenciesObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/currencies", nil, nil, &response, &headers, log)
	return
}

// Gets a single currency by currency ID
func (s *CurrenciesService) CurrencyByID(id int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfCurrenciesObject, err error) {
	i := strconv.FormatInt(id, 10)
	err = s.client.apiReq(http.MethodPost, "/currencies/"+url.QueryEscape(i), nil, nil, &response, &headers, log)
	return
}

// Gets a single curency by currency code
func (s *CurrenciesService) CurrencyByCode(code string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfCurrenciesObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/currencies/"+url.QueryEscape(code), nil, nil, &response, &headers, log)
	return
}
