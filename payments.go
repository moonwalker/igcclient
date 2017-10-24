package igcclient

// TODO Note that there is two url endpoints in the api and that we won't implement any of them at the moment

import (
	. "github.com/moonwalker/igcclient/models"
	"net/url"
	"strconv"
)

type PaymentsService service

// Get all the Deposit methods for a specific country
func (s *PaymentsService) GetAllDepositMethods(countryId int) (response OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	i := strconv.Itoa(countryId)
	err = s.client.apiPost("/payments/getalldepositmethods/"+url.QueryEscape(i), nil, &response, nil, nil)
	return
}

// Get all the Deposit methods for a specific country
func (s *PaymentsService) GetAllWithdrawMethods(countryId int) (response OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	i := strconv.Itoa(countryId)
	err = s.client.apiPost("/payments/getallwithdrawmethods/"+url.QueryEscape(i), nil, &response, nil, nil)
	return
}

// Retrieves all payment methods regardless of their availability or country.
func (s *PaymentsService) GetAllPaymentMethods() (response OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	err = s.client.apiPost("/payments/getallpaymentmethods", nil, &response, nil, nil)
	return
}

// Get the Deposit methods that the user can use. User Authentication (authentication token) is required.
func (s *PaymentsService) GetUserDepositMethods(countryId int, currencyId int, authToken string) (response OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	country := strconv.Itoa(countryId)
	currency := strconv.Itoa(currencyId)
	err = s.client.apiPost("/payments/getuserdepositmethods/"+url.QueryEscape(country)+"/"+url.QueryEscape(currency), nil, &response, nil, &authToken)
	return
}

// Get the Withdraw methods that the user can use. User Authentication (authentication token) is required.
func (s *PaymentsService) GetUserWithdrawMethods(countryId int, userId int, currencyId int, authToken string) (response OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	country := strconv.Itoa(countryId)
	user := strconv.Itoa(userId)
	currency := strconv.Itoa(currencyId)
	err = s.client.apiPost("/payments/getuserwithdrawmethods/"+url.QueryEscape(country)+"/"+url.QueryEscape(user)+"/"+url.QueryEscape(currency), nil, &response, nil, &authToken)
	return
}

// Make a Deposit request. User Authentication (authentication token) is required.
func (s *PaymentsService) GetDepositRequest(body PayRequestModel, authToken string) (response OperationResponseOfPaymentApiResponse, err error) {
	err = s.client.apiPost("/payments/getdepositrequest", &body, &response, nil, &authToken)
	return
}

// Make a Quick Deposit Deposit request. User Authentication (authentication token) is required.
func (s *PaymentsService) GetQuickDepositRequest(body QuickPayRequestModel, authToken string) (response OperationResponseOfPaymentApiResponse, err error) {
	err = s.client.apiPost("/payments/getquickdepositrequest", &body, &response, nil, &authToken)
	return
}

// Make a Withdraw request. User Authentication (authentication token) is required.
func (s *PaymentsService) GetWithdrawRequest(body PayRequestModel, authToken string) (response OperationResponseOfPaymentApiResponse, err error) {
	err = s.client.apiPost("/payments/getwithdrawrequest", &body, &response, nil, &authToken)
	return
}

// Returns the last successful deposit.
func (s *PaymentsService) LastSuccessfulDeposit(authToken string) (response OperationResponseOfTransactionModel, err error) {
	err = s.client.apiPost("/payments/lastsuccessfuldeposit", nil, &response, nil, &authToken)
	return
}

// Returns the transaction details of a given transaction.
func (s *PaymentsService) TransactionDetails(body TransactionSearchModel, authToken string) (response OperationResponseOfTransactionModel, err error) {
	err = s.client.apiPost("/payments/transactiondetails", nil, &response, nil, &authToken)
	return
}

// This method cancels and withdrawal transaction which is in Pending or WaitingApproval status. User Authentication (authentication token) is required.
func (s *PaymentsService) CancelWithdrawal(transactionId int, authToken string) (response OperationResponseOfBoolean, err error) {
	i := strconv.Itoa(transactionId)
	err = s.client.apiPost("/payments/cancelwithdrawal/"+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Get the Form Meta Data for Deposit. User Authentication (authentication token) is required.
func (s *PaymentsService) GetDepositMetaData(paymentTypeId int, authToken string) (response OperationResponseOfPaymentMetaDataModel, err error) {
	i := strconv.Itoa(paymentTypeId)
	err = s.client.apiPost("/payments/getdepositmetadata/"+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Get the Form Meta Data for Withdraw. User Authentication (authentication token) is required.
func (s *PaymentsService) GetWithdrawMetaData(paymentTypeId int, authToken string) (response OperationResponseOfPaymentMetaDataModel, err error) {
	i := strconv.Itoa(paymentTypeId)
	err = s.client.apiPost("/payments/getwithdrawmetadata/"+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Get the Form Meta Data for Quick Deposit. User Authentication (authentication token) is required.
func (s *PaymentsService) GetQuickDepositMetaData(paymentTypeId int, authToken string) (response OperationResponseOfPaymentMetaDataModel, err error) {
	i := strconv.Itoa(paymentTypeId)
	err = s.client.apiPost("/payments/getquickdepositmetadata/"+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Creates a wallet and a payment transaction (Headers required X-Api-Key and AuthenticationToken)
func (s *PaymentsService) OperatorWalletsTransactions(body WalletRequest, operatorTransactionId int, authToken string) (response OperationResponseOfWalletResponse, err error) {
	i := strconv.Itoa(operatorTransactionId)
	err = s.client.apiPost("/payments/operatorwallets/transactions/"+url.QueryEscape(i), &body, &response, nil, &authToken)
	return
}
