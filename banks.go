package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"strconv"
	"net/url"
)

type BanksService service

// Get all the banks sorted in display order, ascending, bank name decrypted
func (s *BanksService) Banks() (response OperationResponseOfListOfBankObject, err error) {
	err = s.client.apiPost("/banks", nil, &response, nil, nil)
	return
}

// Get a single bank by Bank ID
func (s *BanksService) BankById(id int64) (response OperationResponseOfBankObject, err error) {
	i := strconv.FormatInt(id, 10)
	err = s.client.apiPost("/banks/"+url.QueryEscape(i), nil, &response, nil, nil)
	return
}
