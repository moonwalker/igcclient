package igcclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
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
func (s *UserService) SendEmail(body models.SendEmailModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfDictionaryOfStringAndString, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/sendemail", nil, &body, &response, &headers, log)
	return
}

// Get the turnover of a user for the current session
func (s *UserService) SessionTurnover(headers map[string]string, log logger.Logger) (response models.OperationResponseOfDecimal, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/sessionturnover", nil, nil, &response, &headers, log)
	return
}

// Adds a Note to the Logged In User
func (s *UserService) AddUserNote(note string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	q := url.Values{}
	q.Add("note", note)
	err = s.client.apiReq(http.MethodPost, "/user/addusernote", &q, nil, &response, &headers, log)
	return
}

// Returns a list of users with no activity since the fromDate specified as parameter
// fromDate - should be in following format: 2017-01-01
// limit - the maximum number if users to return
func (s *UserService) GetUsersWithNoActivity(fromDate string, limit int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfUserWithNoActivity, err error) {
	q := url.Values{}
	q.Add("fromdate", fromDate)
	q.Add("limit", fmt.Sprintf("%d", limit))
	err = s.client.apiReq(http.MethodPost, "/user/getuserswithnoactivity", &q, nil, &response, &headers, log)
	return
}

// Gets the Count of Similar Users with the exact same First Name and Last Name
func (s *UserService) GetSimilarUserCount(body models.SimilarUser, headers map[string]string, log logger.Logger) (response models.OperationResponseOfInt32, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/getsimilarusercount", nil, &body, &response, &headers, log)
	return
}

// Adds in a KYC requirement where needed for Photo ID, Proof of Address and Proof of Payment
func (s *UserService) AddKYCRequirements(headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/addkycrequirements", nil, nil, &response, &headers, log)
	return
}

// Gets the KYC for the user with supplied user id ordered by status and then by KYC type
// X-Api-Key header is required
func (s *UserService) KYCByUserID(userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfKYCLinkedObject, err error) {
	q := url.Values{}
	q.Add("userid", fmt.Sprintf("%d", userID))
	err = s.client.apiReq(http.MethodPost, "/user/kyc", &q, nil, &response, &headers, log)
	return
}

// Gets the KYC for the current user ordered by status and then by KYC type
// AuthenticationToken header is required
func (s *UserService) KYC(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfKYCLinkedObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/kyc", nil, nil, &response, &headers, log)
	return
}

// Uploads a KYC file for the specified KYC item.
func (s *UserService) KYCUpload(body models.KYCUpload, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/kyc/upload", nil, &body, &response, &headers, log)
	return
}

// Get the logged in user using the Authentication Token
func (s *UserService) User(headers map[string]string, log logger.Logger) (response models.OperationResponseOfSafeUserDetails, err error) {
	err = s.client.apiReq(http.MethodPost, "/user", nil, nil, &response, &headers, log)
	return
}

// Get the logged in user using the Authentication Token
func (s *UserService) UserV2(headers map[string]string, log logger.Logger) (response models.OperationResponseOfUserResponseDTO, err error) {
	err = s.client.apiReq(http.MethodGet, "/v2/user", nil, nil, &response, &headers, log)
	return
}

// Get user using the userID
func (s *UserService) UserByID(userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfSafeUserDetails, err error) {
	q := url.Values{}
	q.Add("userid", fmt.Sprintf("%d", userID))
	err = s.client.apiReq(http.MethodPost, "/user", &q, nil, &response, &headers, log)
	return
}

// Get user last logins
func (s *UserService) GetLoginHistory(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfDBTokenIP, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/getloginhistory", nil, nil, &response, &headers, log)
	return
}

// Accept terms for the logged in user
func (s *UserService) AcceptTerms(headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/acceptterms", nil, nil, &response, &headers, log)
	return
}

// Update user details
func (s *UserService) Update(body models.UpdateUserObject, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/update", nil, &body, &response, &headers, log)
	return
}

// Email validation
func (s *UserService) CheckEmail(body models.CheckUser, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/check/email", nil, &body, &response, &headers, log)
	return
}

// Username validation
func (s *UserService) CheckUsername(body models.CheckUser, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/check/username", nil, &body, &response, &headers, log)
	return
}

// Combination validation This check if the user firstname, lastname, address and phone match with an existing user
func (s *UserService) CheckCombination(body models.UserObject, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/check/combination", nil, &body, &response, &headers, log)
	return
}

// Retrieves the list of available limit durations.
func (s *UserService) LimitsGetLimitDurations(headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfLimitDurationObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/limits/getlimitdurations", nil, nil, &response, &headers, log)
	return
}

// Retrieves the list of available limits.
func (s *UserService) LimitsGetLimits(headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfLimit, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/limits/getlimits", nil, nil, &response, &headers, log)
	return
}

// Get limits for the user with the user id provided
func (s *UserService) LimitsGetUserLimitsByUserID(userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfUserLimit, err error) {
	q := url.Values{}
	q.Add("userid", fmt.Sprintf("%d", userID))
	err = s.client.apiReq(http.MethodPost, "/user/limits/getuserlimits", &q, nil, &response, &headers, log)
	return
}

// Get limits for the current logged in user
func (s *UserService) LimitsGetUserLimits(headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfUserLimit, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/limits/getuserlimits", nil, nil, &response, &headers, log)
	return
}

// Set limit for user logged in
// USER_LIMIT_INVALID error expected if limit is set with no duration and is not of type session,
// stake per session or max stake per bet.
func (s *UserService) LimitsSetUserLimit(body models.SetUserLimitModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/limits/setuserlimit", nil, &body, &response, &headers, log)
	return
}

// Set multiple limits for the logged in user
// USER_LIMIT_INVALID eror expected if limit is set with no duration and is not of type session,
// stake per session or max stake per bet.
func (s *UserService) LimitsSetUserLimits(body []models.SetUserLimitModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/limits/setuserlimits", nil, &body, &response, &headers, log)
	return
}

// Associates a user with an affiliate reference.
// Errors:
// 	USER_NOT_FOUND if user does not exist.
// 	USER_ALREADY_AFFILIATED if user is already affiliated.
// 	INVALID_AFFILIATE_REFERENCE if affiliate reference is null or empty.
func (s *UserService) SetAffiliateReference(userID int64, affiliateReference string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	q := url.Values{}
	q.Add("userid", fmt.Sprintf("%d", userID))
	q.Add("affiliatereference", affiliateReference)
	err = s.client.apiReq(http.MethodPost, "/user/setaffiliatereference", &q, nil, &response, &headers, log)
	return
}

// Close user's account at their request. Requires user's date of birth to proceed Errors: USER_NOT_FOUND if user does not exist. INVALID_DATE_RANGE if date of birth is Empty/Incorrect. INVALID_DOB if date of birth is incorrect.
func (s *UserService) CloseAccount(body models.CloseAccountDOBRequestDto, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/user/closeaccount", nil, &body, &response, &headers, log)
	return
}

func (s *UserService) PasswordSetOnce(body models.PasswordRequestDto, headers map[string]string, log logger.Logger) (response models.OperationResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/user/password/setonce", nil, &body, &response, &headers, log)
	return
}
