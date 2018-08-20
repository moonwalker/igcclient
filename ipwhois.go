package igcclient

import (
	"github.com/moonwalker/igcclient/models"
)

type IPWhoisService service

// Get the country object from the IP passed in the IPWhoisRequest object
func (s *IPWhoisService) IPWhois(body models.IPWhoisRequest, headers map[string]string) (response models.OperationResponseOfCountriesObject, err error) {
	err = s.client.apiPost("/IPWhois", nil, &body, &response, &headers)
	return
}
