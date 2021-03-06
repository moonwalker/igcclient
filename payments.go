package igcclient

// TODO Note that there is two url endpoints in the api and that we won't implement any of them at the moment

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type PaymentsService service

// Get all the Deposit methods for a specific country
func (s *PaymentsService) GetAllDepositMethods(countryID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	i := strconv.FormatInt(countryID, 10)
	err = s.client.apiReq(http.MethodPost, "/payments/getalldepositmethods/"+url.QueryEscape(i), nil, nil, &response, &headers, log)
	return
}

// Get all the Deposit methods for a specific country
func (s *PaymentsService) GetAllWithdrawMethods(countryID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	i := strconv.FormatInt(countryID, 10)
	err = s.client.apiReq(http.MethodPost, "/payments/getallwithdrawmethods/"+url.QueryEscape(i), nil, nil, &response, &headers, log)
	return
}

// Retrieves all payment methods regardless of their availability or country.
func (s *PaymentsService) GetAllPaymentMethods(headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	err = s.client.apiReq(http.MethodPost, "/payments/getallpaymentmethods", nil, nil, &response, &headers, log)
	return
}

// Get the Deposit methods that the user can use. User Authentication (authentication token) is required.
func (s *PaymentsService) GetUserDepositMethods(countryID int64, currencyID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	country := strconv.FormatInt(countryID, 10)
	currency := strconv.FormatInt(currencyID, 10)
	err = s.client.apiReq(http.MethodPost, "/payments/getuserdepositmethods/"+url.QueryEscape(country)+"/"+url.QueryEscape(currency), nil, nil, &response, &headers, log)
	return
}

// Get the Withdraw methods that the user can use. User Authentication (authentication token) is required.
func (s *PaymentsService) GetUserWithdrawMethods(countryID int64, userID int64, currencyID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfPaymentMethod, err error) {
	country := strconv.FormatInt(countryID, 10)
	user := strconv.FormatInt(userID, 10)
	currency := strconv.FormatInt(currencyID, 10)
	err = s.client.apiReq(http.MethodPost, "/payments/getuserwithdrawmethods/"+url.QueryEscape(country)+"/"+url.QueryEscape(user)+"/"+url.QueryEscape(currency), nil, nil, &response, &headers, log)
	return
}

// Make a Deposit request. User Authentication (authentication token) is required.
func (s *PaymentsService) GetDepositRequest(body models.PayRequestModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfPaymentApiResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/payments/getdepositrequest", nil, &body, &response, &headers, log)
	return
}

// Make a Quick Deposit Deposit request. User Authentication (authentication token) is required.
func (s *PaymentsService) GetQuickDepositRequest(body models.QuickPayRequestModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfPaymentApiResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/payments/getquickdepositrequest", nil, &body, &response, &headers, log)
	return
}

// Make a Withdraw request. User Authentication (authentication token) is required.
func (s *PaymentsService) GetWithdrawRequest(body models.PayRequestModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfPaymentApiResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/payments/getwithdrawrequest", nil, &body, &response, &headers, log)
	return
}

// Returns the last successful deposit.
func (s *PaymentsService) LastSuccessfulDeposit(headers map[string]string, log logger.Logger) (response models.OperationResponseOfTransactionModel, err error) {
	err = s.client.apiReq(http.MethodPost, "/payments/lastsuccessfuldeposit", nil, nil, &response, &headers, log)
	return
}

// Returns the transaction details of a given transaction.
func (s *PaymentsService) TransactionDetails(body models.TransactionSearchModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfTransactionModel, err error) {
	err = s.client.apiReq(http.MethodPost, "/payments/transactiondetails", nil, &body, &response, &headers, log)
	return
}

// This method cancels and withdrawal transaction which is in Pending or WaitingApproval status. User Authentication (authentication token) is required.
func (s *PaymentsService) CancelWithdrawal(transactionId int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	i := strconv.FormatInt(transactionId, 10)
	err = s.client.apiReq(http.MethodPost, "/payments/cancelwithdrawal/"+url.QueryEscape(i), nil, nil, &response, &headers, log)
	return
}

// Get the Form Meta Data for Deposit. User Authentication (authentication token) is required.
func (s *PaymentsService) GetDepositMetaData(paymentTypeID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfPaymentMetaDataModel, err error) {
	i := strconv.FormatInt(paymentTypeID, 10)
	err = s.client.apiReq(http.MethodPost, "/payments/getdepositmetadata/"+url.QueryEscape(i), nil, nil, &response, &headers, log)
	return
}

// Get the Form Meta Data for Withdraw. User Authentication (authentication token) is required.
func (s *PaymentsService) GetWithdrawMetaData(paymentTypeID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfPaymentMetaDataModel, err error) {
	i := strconv.FormatInt(paymentTypeID, 10)
	err = s.client.apiReq(http.MethodPost, "/payments/getwithdrawmetadata/"+url.QueryEscape(i), nil, nil, &response, &headers, log)
	return
}

// Get the Form Meta Data for Quick Deposit. User Authentication (authentication token) is required.
func (s *PaymentsService) GetQuickDepositMetaData(paymentTypeID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfPaymentMetaDataModel, err error) {
	i := strconv.FormatInt(paymentTypeID, 10)
	err = s.client.apiReq(http.MethodPost, "/payments/getquickdepositmetadata/"+url.QueryEscape(i), nil, nil, &response, &headers, log)
	return
}

// Creates a wallet and a payment transaction (Headers required X-Api-Key and AuthenticationToken)
func (s *PaymentsService) OperatorWalletsTransactions(body models.WalletRequest, operatorTransactionId int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfWalletResponse, err error) {
	i := strconv.FormatInt(operatorTransactionId, 10)
	err = s.client.apiReq(http.MethodPost, "/payments/operatorwallets/transactions/"+url.QueryEscape(i), nil, &body, &response, &headers, log)
	return
}
