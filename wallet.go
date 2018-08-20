package igcclient

import (
	"fmt"
	"github.com/moonwalker/igcclient/models"
	"net/url"
)

type WalletService service

// Gets the users OM wallet
func (s *WalletService) GetOMWallet(headers map[string]string) (response models.OperationResponseOfOMUserWalletPublic, err error) {
	err = s.client.apiPost("/wallet/getomwallet", nil, nil, &response, &headers)
	return
}

// Gets the current OM bonus
func (s *WalletService) GetOMCurrentBonus(headers map[string]string) (response models.OperationResponseOfBonus, err error) {
	err = s.client.apiPost("/wallet/getomcurrentbonus", nil, nil, &response, &headers)
	return
}

// Gets the User Wallet
func (s *WalletService) Wallet(headers map[string]string) (response models.OperationResponseOfUserWalletPublic, err error) {
	err = s.client.apiPost("/wallet", nil, nil, &response, &headers)
	return
}

// Gets the User Wallet
func (s *WalletService) WalletByUserID(userID int64, headers map[string]string) (response models.OperationResponseOfUserWalletPublic, err error) {
	q := url.Values{}
	q.Add("userid", fmt.Sprintf("%d", userID))
	err = s.client.apiPost("/wallet", &q, nil, &response, &headers)
	return
}

// No documentation available.
func (s *WalletService) Transactions(body models.WalletTransactionsSearchModel, headers map[string]string) (response models.OperationResponseOfIEnumerableOfWalletTransactionsModel, err error) {
	err = s.client.apiPost("/wallet/transactions", nil, &body, &response, &headers)
	return
}
