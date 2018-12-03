package igcclient

import (
	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type IPWhoisService service

// Get the country object from the IP passed in the IPWhoisRequest object
func (s *IPWhoisService) IPWhois(body models.IPWhoisRequest, headers map[string]string, log logger.Logger) (response models.OperationResponseOfCountriesObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/IPWhois", nil, &body, &response, &headers, log)
	return
}
