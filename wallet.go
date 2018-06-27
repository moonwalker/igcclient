package igcclient

import (
	"github.com/moonwalker/igcclient/models"
)

type WalletService service

// Gets the users OM wallet
func (s *WalletService) GetOMWallet(headers map[string]string) (response models.OperationResponseOfOMUserWalletPublic, err error) {
	err = s.client.apiPost("/wallet/getomwallet", nil, &response, &headers)
	return
}

// Gets the current OM bonus
func (s *WalletService) GetOMCurrentBonus(headers map[string]string) (response models.OperationResponseOfBonus, err error) {
	err = s.client.apiPost("/wallet/getomcurrentbonus", nil, &response, &headers)
	return
}

// Gets the User Wallet
func (s *WalletService) Wallet(headers map[string]string) (response models.OperationResponseOfUserWalletPublic, err error) {
	err = s.client.apiPost("/wallet", nil, &response, &headers)
	return
}

// No documentation available.
func (s *WalletService) Transactions(body models.WalletTransactionsSearchModel, headers map[string]string) (response models.OperationResponseOfIEnumerableOfWalletTransactionsModel, err error) {
	err = s.client.apiPost("/wallet/transactions", &body, &response, &headers)
	return
}
