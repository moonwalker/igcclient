package igcclient

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type BanksService service

// Get all the banks sorted in display order, ascending, bank name decrypted
func (s *BanksService) Banks(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfBankObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/banks", nil, nil, &response, &headers, log)
	return
}

// Get a single bank by Bank ID
func (s *BanksService) BankByID(id int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBankObject, err error) {
	i := strconv.FormatInt(id, 10)
	err = s.client.apiReq(http.MethodPost, "/banks/"+url.QueryEscape(i), nil, nil, &response, &headers, log)
	return
}
