package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
)

type WalletService service

// Gets the users OM wallet
func (s *WalletService) GetOMWallet(authToken string) (response OperationResponseOfOMUserWalletPublic, err error) {
	err = s.client.apiPost("/wallet/getomwallet", nil, &response, nil, &authToken)
	return
}

// Gets the current OM bonus
func (s *WalletService) GetOMCurrentBonus(authToken string) (response OperationResponseOfBonus, err error) {
	err = s.client.apiPost("/wallet/getomcurrentbonus", nil, &response, nil, &authToken)
	return
}

// Gets the User Wallet
func (s *WalletService) Wallet(authToken string) (response OperationResponseOfUserWalletPublic, err error) {
	err = s.client.apiPost("/wallet", nil, &response, nil, &authToken)
	return
}

// No documentation available.
func (s *WalletService) Transactions(body WalletTransactionsSearchModel, authToken string) (response OperationResponseOfIEnumerableOfWalletTransactionsModel, err error) {
	err = s.client.apiPost("/wallet/transactions", &body, &response, nil, &authToken)
	return
}
