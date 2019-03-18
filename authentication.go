package igcclient

import (
	"net/http"
	"net/url"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type AuthenticationService service

// Used to log-in an end user to the system
func (s *AuthenticationService) Login(loginModel models.LoginModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfAuthResponseDTO, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/authentication/login", nil, &loginModel, &response, &headers, log)
	return
}

// Checks if the User is logged in
func (s *AuthenticationService) IsLoggedIn(headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/authentication/isloggedin", nil, nil, &response, &headers, log)
	return
}

// Logout the User
func (s *AuthenticationService) Logout(headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/authentication/logout", nil, nil, &response, &headers, log)
	return
}

// Authenticate a user using a one time challenge token.
// param challengeToken: challenge token aquired from Poker/GetChallengeToken
func (s *AuthenticationService) AuthenticateWithChallenge(challengeToken string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfGuid, err error) {
	q := url.Values{}
	q.Add("challengeToken", challengeToken)
	err = s.client.apiReq(http.MethodPost, "/v2/authentication/authenticatewithchallenge", &q, nil, &response, &headers, log)
	return
}

// Retrieves password validation Regex for operator depending on Core.Settings. Default if not found: ^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{6,25}$
func (s *AuthenticationService) GetPasswordRegex(headers map[string]string, log logger.Logger) (response models.OperationResponseOfString, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/authentication/getpasswordregex", nil, nil, &response, &headers, log)
	return
}

// Get user account activation token
func (s *AuthenticationService) VerifyActivationTokenGet(body models.GetVerificationTokenModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfGuid, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/authentication/verify/activationtoken/get", nil, &body, &response, &headers, log)
	return
}

// Send a verification SMS to an unverified user using mobile prefix and mobile combination.
// The SMS language is selected using the user's registered LanguageID property. Sends SMS even though user is blocked.
func (s *AuthenticationService) VerifySMSSend(body models.ShortMessageServiceModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/authentication/verify/sms/send", nil, &body, &response, &headers, log)
	return
}

// Activate a user's account using the correct mobile prefix, mobile, and mobile verification code combination.
// We allow blocked users to be verified. We return ACCOUNT_BLOCKED in the list of errors even though successfully verified.
func (s *AuthenticationService) VerifySMS(body models.VerifyUserBySMSModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfString, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/authentication/verify/sms", nil, &body, &response, &headers, log)
	return
}

// Verifies an email address with the GUID recieved in the email. Upon successful verificaton, the user will be logged in.
// Allows user to verify even though they are blocked. If blocked we will return ACCOUNT_BLOCKED error code.
func (s *AuthenticationService) VerifyEmail(verificationCode string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfGuid, err error) {
	err = s.client.apiReq(http.MethodPost, "/Authentication/Verify/Email/"+url.QueryEscape(verificationCode), nil, nil, &response, &headers, log)
	return
}

// Used to send the verify email. Returns false if email is null or empty.
// USER_NOT_FOUND if user with the given email does not exist or exists but is already verified.
func (s *AuthenticationService) VerifyEmailSend(email string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	q := url.Values{}
	q.Add("email", email)
	err = s.client.apiReq(http.MethodPost, "/authentication/verify/email/send", &q, nil, &response, &headers, log)
	return
}

// Allows to bind UserAgent and IP to the authentication token
// param: New IP and User Agent to bind to
func (s *AuthenticationService) AllowIP(body models.AllowIPModel, headers map[string]string, log logger.Logger) (response models.OperationResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/authentication/allow/ip", nil, &body, &response, &headers, log)
	return
}

// Get security question for email
func (s *AuthenticationService) SecurityQuestion(body models.EmailModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfSecurityQuestion, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/authentication/securityquestion", nil, &body, &response, &headers, log)
	return
}

// Change User's password
func (s *AuthenticationService) ChangePassword(body models.ChangePasswordModel, headers map[string]string, log logger.Logger) (response models.OperationResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/authentication/change/password", nil, &body, &response, &headers, log)
	return
}

// Change User's Email Address
func (s *AuthenticationService) ChangeEmail(body models.ChangeEmailModel, headers map[string]string, log logger.Logger) (response models.OperationResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/authentication/change/email", nil, &body, &response, &headers, log)
	return
}

// Change User's Security Question
func (s *AuthenticationService) ChangeSecurityQuestion(body models.ChangeSecurityQuestionModel, headers map[string]string, log logger.Logger) (response models.OperationResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/authentication/change/securityquestion", nil, &body, &response, &headers, log)
	return
}

// Start Forgot Password Process
func (s *AuthenticationService) ForgotPassword(body models.ResetPasswordModel, headers map[string]string, log logger.Logger) (response models.OperationResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/authentication/forgotpassword", nil, &body, &response, &headers, log)
	return
}

// Attempts to send a mobile verification code to the mobile number specified in the arguments for reset password purposes
func (s *AuthenticationService) ForgotPasswordSMS(body models.ShortMessageServiceModel, headers map[string]string, log logger.Logger) (response models.OperationResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/authentication/forgotpassword/sms", nil, &body, &response, &headers, log)
	return
}

// Change User's Forgotten Password by using a Mobile Verification Code
func (s *AuthenticationService) ForgotPasswordChangeSMS(body models.ForgotPasswordChangeBySMSModel, headers map[string]string, log logger.Logger) (response models.OperationResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/authentication/forgotpassword/change/sms", nil, &body, &response, &headers, log)
	return
}

// Change User's Forgotten Password
func (s *AuthenticationService) ForgotPasswordChange(body models.ForgotPasswordChangeModel, headers map[string]string, log logger.Logger) (response models.OperationResponse, err error) {
	err = s.client.apiReq(http.MethodPost, "/authentication/forgotpassword/change", nil, &body, &response, &headers, log)
	return
}

// Before calling this method the method "GetUserVerifyRegistrationTypes()" should be called to get a list of allowed verification types for a particular IP.
// This is of course is only valid if the country-verification type whitelist is going to be used
//
// Note that verification types can be whitelisted for each country from the Backoffice.
// Note: Verification Type ID property is always required when posting to this endpoint
//
// Verification Type ID 1 (SMS): SMS verification is required during registration
// 	#1 POST to this endpoint with the properties UserName, Email, Mobile Prefix and Mobile
// 	#2 If all properties are valid the error code "MOBILE_MOBILE_VERIFICATION_CODE_COMBINATION_INVALID" will be returned. This indicates that an SMS has been sent to the provided Mobile.
// 	#3 POST to the endpoint once more when the user enters the Mobile Verification Code and the rest of his details
//
// Verification Type ID 2 (Email): Email verification is required after registration
// 	#1 POST to this endpoint with all the required properties. This sends out a verification Email once a user record is created.
//
// Verification Type ID 3 (SMS and Email): This is a combination of Type 1 and Type 2
// 	#1 POST to this endpoint with the properties UserName, Email, Mobile Prefix and Mobile
// 	#2 If all properties are valid the error code "MOBILE_MOBILE_VERIFICATION_CODE_COMBINATION_INVALID" will be returned. This indicates that an SMS has been sent to the provided Mobile.
// 	#3 POST to the endpoint once more when the user enters the Mobile Verification Code and the rest of his details. At this point we will send out a verification Email.
//
// Verification Type ID 4 (MANUAL): No verification methods are activated automatically
// 	#1 POST to this endpoint with all the required properties. THIS WILL NOT SEND OUT ANY VERIFICATION EMAILS OR SMS's ** At this point you are able to call "Verify/SMS/Send" or/and "Verify/Email/Send" to verify the user at any point desired.
//
func (s *AuthenticationService) Register(body models.RegistrationData, headers map[string]string, log logger.Logger) (response models.OperationResponseOfDictionaryOfStringAndString, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/authentication/register", nil, &body, &response, &headers, log)
	return
}
