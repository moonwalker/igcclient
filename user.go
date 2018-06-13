package igcclient

import (
	"net/url"
	"strconv"

	. "github.com/moonwalker/igcclient/models"
)

type UserService service

// Sends an email.
// Currently supports the following email template ids:
// 	1 (Contact Form template),
// 	2 (Manual Deposit Request template)
// 		Does not consider currently authenticated user.
// 		All information is to be included in the EmailBodyPlaceholderValues.
// 		Key "Email" and "FullName" is reserved and not to be used - will exit with error if used.
//
// Expected error codes: UNSUPPORTED_EMAIL_TEMPLATE_ID, REQUEST_DATA_INVALID, RESERVED_KEY_EMAIL_TEMPLATE_PLACEHOLDER_VALUES
func (s *UserService) SendEmail(body SendEmailModel) (response OperationResponseOfDictionaryOfStringAndString, err error) {
	err = s.client.apiPost("user/sendemail", &body, &response, nil, nil)
	return
}

// Get the turnover of a user for the current session
func (s *UserService) SessionTurnover(authToken string) (response OperationResponseOfDecimal, err error) {
	err = s.client.apiPost("/user/sessionturnover", nil, &response, nil, &authToken)
	return
}

// Adds a Note to the Logged In User
func (s *UserService) AddUserNote(note string, authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/user/addusernote?note="+url.QueryEscape(note), nil, &response, nil, &authToken)
	return
}

// Returns a list of users with no activity since the fromDate specified as parameter
// fromDate - should be in following format: 2017-01-01
// limit - the maximum number if users to return
func (s *UserService) GetUsersWithNoActivity(fromDate string, limit int64, xAPIKey string) (response OperationResponseOfIEnumerableOfUserWithNoActivity, err error) {
	l := strconv.FormatInt(limit, 10)
	err = s.client.apiPost("/user/getuserswithnoactivity?fromdate="+url.QueryEscape(fromDate)+"&limit="+url.QueryEscape(l), nil, &response, &xAPIKey, nil)
	return
}

// Gets the Count of Similar Users with the exact same First Name and Last Name
func (s *UserService) GetSimilarUserCount(body SimilarUser) (response OperationResponseOfInt32, err error) {
	err = s.client.apiPost("/user/getsimilarusercount", &body, &response, nil, nil)
	return
}

// Adds in a KYC requirement where needed for Photo ID, Proof of Address and Proof of Payment
func (s *UserService) AddKYCRequirements(authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/user/addkycrequirements", nil, &response, nil, &authToken)
	return
}

// Gets the KYC for the user with supplied user id ordered by status and then by KYC type
// X-Api-Key header is required
func (s *UserService) KYCByUserID(userID int64, xAPIKey string) (response OperationResponseOfListOfKYCLinkedObject, err error) {
	id := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/user/kyc?userid="+url.QueryEscape(id), nil, &response, &xAPIKey, nil)
	return
}

// Gets the KYC for the current user ordered by status and then by KYC type
// AuthenticationToken header is required
func (s *UserService) KYC(authToken string) (response OperationResponseOfListOfKYCLinkedObject, err error) {
	err = s.client.apiPost("/user/kyc", nil, &response, nil, &authToken)
	return
}

// Uploads a KYC file for the specified KYC item.
func (s *UserService) KYCUpload(body KYCUpload, authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/user/kyc/upload", &body, &response, nil, &authToken)
	return
}

// Get the logged in user using the Authentication Token
func (s *UserService) User(authToken string) (response OperationResponseOfSafeUserDetails, err error) {
	err = s.client.apiPost("/user", nil, &response, nil, &authToken)
	return
}

// Get user using the userID
func (s *UserService) UserByID(userID int64, xAPIKey string) (response OperationResponseOfSafeUserDetails, err error) {
	id := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/user?userId="+url.QueryEscape(id), nil, &response, &xAPIKey, nil)
	return
}

// Get user last logins
func (s *UserService) GetLoginHistory(authToken string) (response OperationResponseOfListOfDBTokenIP, err error) {
	err = s.client.apiPost("/user/getloginhistory", nil, &response, nil, &authToken)
	return
}

// Accept terms for the logged in user
func (s *UserService) AcceptTerms(authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/user/acceptterms", nil, &response, nil, &authToken)
	return
}

// Update user details
func (s *UserService) Update(body UpdateUserObject, authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/user/update", &body, &response, nil, &authToken)
	return
}

// Email validation
func (s *UserService) CheckEmail(body CheckUser) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/user/check/email", &body, &response, nil, nil)
	return
}

// Username validation
func (s *UserService) CheckUsername(body CheckUser) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/user/check/username", &body, &response, nil, nil)
	return
}

// Combination validation This check if the user firstname, lastname, address and phone match with an existing user
func (s *UserService) CheckCombination(body UserObject) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/user/check/combination", &body, &response, nil, nil)
	return
}

// Retrieves the list of available limit durations.
func (s *UserService) LimitsGetLimitDurations() (response OperationResponseOfIEnumerableOfLimitDurationObject, err error) {
	err = s.client.apiPost("/user/limits/getlimitdurations", nil, &response, nil, nil)
	return
}

// Retrieves the list of available limits.
func (s *UserService) LimitsGetLimits() (response OperationResponseOfIEnumerableOfLimit, err error) {
	err = s.client.apiPost("/user/limits/getlimits", nil, &response, nil, nil)
	return
}

// Get limits for the user with the user id provided
func (s *UserService) LimitsGetUserLimitsByUserID(userID int64, xAPIKey string) (response OperationResponseOfIEnumerableOfUserLimit, err error) {
	id := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/user/limits/getuserlimits?userId="+url.QueryEscape(id), nil, &response, &xAPIKey, nil)
	return
}

// Get limits for the current logged in user
func (s *UserService) LimitsGetUserLimits(authToken string) (response OperationResponseOfIEnumerableOfUserLimit, err error) {
	err = s.client.apiPost("/user/limits/getuserlimits", nil, &response, nil, &authToken)
	return
}

// Set limit for user logged in
// USER_LIMIT_INVALID error expected if limit is set with no duration and is not of type session,
// stake per session or max stake per bet.
func (s *UserService) LimitsSetUserLimit(body SetUserLimitModel, authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/user/limits/setuserlimit", &body, &response, nil, &authToken)
	return
}

// Set multiple limits for the logged in user
// USER_LIMIT_INVALID eror expected if limit is set with no duration and is not of type session,
// stake per session or max stake per bet.
func (s *UserService) LimitsSetUserLimits(body []SetUserLimitModel, authToken string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/user/limits/setuserlimits", &body, &response, nil, &authToken)
	return
}

// Associates a user with an affiliate reference.
// Errors:
// 	USER_NOT_FOUND if user does not exist.
// 	USER_ALREADY_AFFILIATED if user is already affiliated.
// 	INVALID_AFFILIATE_REFERENCE if affiliate reference is null or empty.
func (s *UserService) SetAffiliateReference(userID int64, affiliateReference string, xAPIKey string) (response OperationResponseOfBoolean, err error) {
	id := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/user/setaffiliatereference?userid=+"+url.QueryEscape(id)+"&affiliatereference="+url.QueryEscape(affiliateReference), nil, &response, &xAPIKey, nil)
	return
}
