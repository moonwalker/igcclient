package igcclient

// TODO Note that there is two url endpoints in the api and that we won't implement any of them at the moment

import (
	"net/url"
	"strconv"

	. "github.com/moonwalker/igcclient/models"
)

type PaymentsService service

// Get all the Deposit methods for a specific country
func (s *PaymentsService) GetAllDepositMethods(countryID int64) (response OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	i := strconv.FormatInt(countryID, 10)
	err = s.client.apiPost("/payments/getalldepositmethods/"+url.QueryEscape(i), nil, &response, nil)
	return
}

// Get all the Deposit methods for a specific country
func (s *PaymentsService) GetAllWithdrawMethods(countryID int64) (response OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	i := strconv.FormatInt(countryID, 10)
	err = s.client.apiPost("/payments/getallwithdrawmethods/"+url.QueryEscape(i), nil, &response, nil)
	return
}

// Retrieves all payment methods regardless of their availability or country.
func (s *PaymentsService) GetAllPaymentMethods() (response OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	err = s.client.apiPost("/payments/getallpaymentmethods", nil, &response, nil)
	return
}

// Get the Deposit methods that the user can use. User Authentication (authentication token) is required.
func (s *PaymentsService) GetUserDepositMethods(countryID int64, currencyID int64, headers map[string]string) (response OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	country := strconv.FormatInt(countryID, 10)
	currency := strconv.FormatInt(currencyID, 10)
	err = s.client.apiPost("/payments/getuserdepositmethods/"+url.QueryEscape(country)+"/"+url.QueryEscape(currency), nil, &response, &headers)
	return
}

// Get the Withdraw methods that the user can use. User Authentication (authentication token) is required.
func (s *PaymentsService) GetUserWithdrawMethods(countryID int64, userID int64, currencyID int64, headers map[string]string) (response OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	country := strconv.FormatInt(countryID, 10)
	user := strconv.FormatInt(userID, 10)
	currency := strconv.FormatInt(currencyID, 10)
	err = s.client.apiPost("/payments/getuserwithdrawmethods/"+url.QueryEscape(country)+"/"+url.QueryEscape(user)+"/"+url.QueryEscape(currency), nil, &response, &headers)
	return
}

// Make a Deposit request. User Authentication (authentication token) is required.
func (s *PaymentsService) GetDepositRequest(body PayRequestModel, headers map[string]string) (response OperationResponseOfPaymentApiResponse, err error) {
	err = s.client.apiPost("/payments/getdepositrequest", &body, &response, &headers)
	return
}

// Make a Quick Deposit Deposit request. User Authentication (authentication token) is required.
func (s *PaymentsService) GetQuickDepositRequest(body QuickPayRequestModel, headers map[string]string) (response OperationResponseOfPaymentApiResponse, err error) {
	err = s.client.apiPost("/payments/getquickdepositrequest", &body, &response, &headers)
	return
}

// Make a Withdraw request. User Authentication (authentication token) is required.
func (s *PaymentsService) GetWithdrawRequest(body PayRequestModel, headers map[string]string) (response OperationResponseOfPaymentApiResponse, err error) {
	err = s.client.apiPost("/payments/getwithdrawrequest", &body, &response, &headers)
	return
}

// Returns the last successful deposit.
func (s *PaymentsService) LastSuccessfulDeposit(headers map[string]string) (response OperationResponseOfTransactionModel, err error) {
	err = s.client.apiPost("/payments/lastsuccessfuldeposit", nil, &response, &headers)
	return
}

// Returns the transaction details of a given transaction.
func (s *PaymentsService) TransactionDetails(body TransactionSearchModel, headers map[string]string) (response OperationResponseOfTransactionModel, err error) {
	err = s.client.apiPost("/payments/transactiondetails", nil, &response, &headers)
	return
}

// This method cancels and withdrawal transaction which is in Pending or WaitingApproval status. User Authentication (authentication token) is required.
func (s *PaymentsService) CancelWithdrawal(transactionId int64, headers map[string]string) (response OperationResponseOfBoolean, err error) {
	i := strconv.FormatInt(transactionId, 10)
	err = s.client.apiPost("/payments/cancelwithdrawal/"+url.QueryEscape(i), nil, &response, &headers)
	return
}

// Get the Form Meta Data for Deposit. User Authentication (authentication token) is required.
func (s *PaymentsService) GetDepositMetaData(paymentTypeID int64, headers map[string]string) (response OperationResponseOfPaymentMetaDataModel, err error) {
	i := strconv.FormatInt(paymentTypeID, 10)
	err = s.client.apiPost("/payments/getdepositmetadata/"+url.QueryEscape(i), nil, &response, &headers)
	return
}

// Get the Form Meta Data for Withdraw. User Authentication (authentication token) is required.
func (s *PaymentsService) GetWithdrawMetaData(paymentTypeID int64, headers map[string]string) (response OperationResponseOfPaymentMetaDataModel, err error) {
	i := strconv.FormatInt(paymentTypeID, 10)
	err = s.client.apiPost("/payments/getwithdrawmetadata/"+url.QueryEscape(i), nil, &response, &headers)
	return
}

// Get the Form Meta Data for Quick Deposit. User Authentication (authentication token) is required.
func (s *PaymentsService) GetQuickDepositMetaData(paymentTypeID int64, headers map[string]string) (response OperationResponseOfPaymentMetaDataModel, err error) {
	i := strconv.FormatInt(paymentTypeID, 10)
	err = s.client.apiPost("/payments/getquickdepositmetadata/"+url.QueryEscape(i), nil, &response, &headers)
	return
}

// Creates a wallet and a payment transaction (Headers required X-Api-Key and AuthenticationToken)
func (s *PaymentsService) OperatorWalletsTransactions(body WalletRequest, operatorTransactionId int64, headers map[string]string) (response OperationResponseOfWalletResponse, err error) {
	i := strconv.FormatInt(operatorTransactionId, 10)
	err = s.client.apiPost("/payments/operatorwallets/transactions/"+url.QueryEscape(i), &body, &response, &headers)
	return
}
