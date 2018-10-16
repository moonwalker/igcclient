package igcclient

import (
	"fmt"
	"net/url"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type WalletService service

// Gets the users OM wallet
func (s *WalletService) GetOMWallet(headers map[string]string, log logger.Logger) (response models.OperationResponseOfOMUserWalletPublic, err error) {
	err = s.client.apiPost("/wallet/getomwallet", nil, nil, &response, &headers, log)
	return
}

// Gets the current OM bonus
func (s *WalletService) GetOMCurrentBonus(headers map[string]string, log logger.Logger) (response models.OperationResponseOfBonus, err error) {
	err = s.client.apiPost("/wallet/getomcurrentbonus", nil, nil, &response, &headers, log)
	return
}

// Gets the User Wallet
func (s *WalletService) Wallet(headers map[string]string, log logger.Logger) (response models.OperationResponseOfUserWalletPublic, err error) {
	err = s.client.apiPost("/wallet", nil, nil, &response, &headers, log)
	return
}

// Gets the User Wallet
func (s *WalletService) WalletByUserID(userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfUserWalletPublic, err error) {
	q := url.Values{}
	q.Add("userid", fmt.Sprintf("%d", userID))
	err = s.client.apiPost("/wallet", &q, nil, &response, &headers, log)
	return
}

// No documentation available.
func (s *WalletService) Transactions(body models.WalletTransactionsSearchModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfWalletTransactionsModel, err error) {
	err = s.client.apiPost("/wallet/transactions", nil, &body, &response, &headers, log)
	return
}
