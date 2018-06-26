package igcclient

type WalletService service

// Gets the users OM wallet
func (s *WalletService) GetOMWallet(headers map[string]string) (response OperationResponseOfOMUserWalletPublic, err error) {
	err = s.client.apiPost("/wallet/getomwallet", nil, &response, &headers)
	return
}

// Gets the current OM bonus
func (s *WalletService) GetOMCurrentBonus(headers map[string]string) (response OperationResponseOfBonus, err error) {
	err = s.client.apiPost("/wallet/getomcurrentbonus", nil, &response, &headers)
	return
}

// Gets the User Wallet
func (s *WalletService) Wallet(headers map[string]string) (response OperationResponseOfUserWalletPublic, err error) {
	err = s.client.apiPost("/wallet", nil, &response, &headers)
	return
}

// No documentation available.
func (s *WalletService) Transactions(body WalletTransactionsSearchModel, headers map[string]string) (response OperationResponseOfIEnumerableOfWalletTransactionsModel, err error) {
	err = s.client.apiPost("/wallet/transactions", &body, &response, &headers)
	return
}
