package igcclient

import (
	"net/url"

	. "github.com/moonwalker/igcclient/models"
)

type AuthenticationService service

// Used to log-in an end user to the system
func (s *AuthenticationService) Login(loginModel LoginModel, headers map[string]string) (response OperationResponseOfNullableIfGuid, err error) {
	err = s.client.apiPost("/authentication/login", &loginModel, &response, &headers)
	return
}

// Checks if the User is logged in
func (s *AuthenticationService) IsLoggedIn(headers map[string]string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/authentication/isloggedin", nil, &response, &headers)
	return
}

// Logout the User
func (s *AuthenticationService) Logout(headers map[string]string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/authentication/logout", nil, &response, &headers)
	return
}

// Authenticate a user using a one time challenge token.
// param challengeToken: challenge token aquired from Poker/GetChallengeToken
func (s *AuthenticationService) AuthenticateWithChallenge(challengeToken string, headers map[string]string) (response OperationResponseOfGuid, err error) {
	err = s.client.apiPost("/v2/authentication/authenticatewithchallenge?challengeToken="+url.QueryEscape(challengeToken), nil, &response, &headers)
	return
}

// Retrieves password validation Regex for operator depending on Core.Settings. Default if not found: ^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{6,25}$
func (s *AuthenticationService) GetPasswordRegex(headers map[string]string) (response OperationResponseOfString, err error) {
	err = s.client.apiPost("/v2/authentication/getpasswordregex", nil, &response, &headers)
	return
}

// Get user account activation token
func (s *AuthenticationService) VerifyActivationTokenGet(body GetVerificationTokenModel, headers map[string]string) (response OperationResponseOfGuid, err error) {
	err = s.client.apiPost("/v2/authentication/verify/activationtoken/get", &body, &response, &headers)
	return
}

// Send a verification SMS to an unverified user using mobile prefix and mobile combination.
// The SMS language is selected using the user's registered LanguageID property. Sends SMS even though user is blocked.
func (s *AuthenticationService) VerifySMSSend(body ShortMessageServiceModel, headers map[string]string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/v2/authentication/verify/sms/send", &body, &response, &headers)
	return
}

// Activate a user's account using the correct mobile prefix, mobile, and mobile verification code combination.
// We allow blocked users to be verified. We return ACCOUNT_BLOCKED in the list of errors even though successfully verified.
func (s *AuthenticationService) VerifySMS(body VerifyUserBySMSModel, headers map[string]string) (response OperationResponseOfString, err error) {
	err = s.client.apiPost("/v2/authentication/verify/sms", &body, &response, &headers)
	return
}

// Verifies an email address with the GUID recieved in the email. Upon successful verificaton, the user will be logged in.
// Allows user to verify even though they are blocked. If blocked we will return ACCOUNT_BLOCKED error code.
func (s *AuthenticationService) VerifyEmail(verificationCode string, headers map[string]string) (response OperationResponseOfGuid, err error) {
	err = s.client.apiPost("/Authentication/Verify/Email/"+url.QueryEscape(verificationCode), nil, &response, &headers)
	return
}

// Used to send the verify email. Returns false if email is null or empty.
// USER_NOT_FOUND if user with the given email does not exist or exists but is already verified.
func (s *AuthenticationService) VerifyEmailSend(email string, headers map[string]string) (response OperationResponseOfBoolean, err error) {
	err = s.client.apiPost("/authentication/verify/email/send?email="+url.QueryEscape(email), nil, &response, &headers)
	return
}

// Allows to bind UserAgent and IP to the authentication token
// param: New IP and User Agent to bind to
func (s *AuthenticationService) AllowIP(body AllowIPModel, headers map[string]string) (response OperationResponse, err error) {
	err = s.client.apiPost("/authentication/allow/ip", &body, &response, &headers)
	return
}

// Get security question for email
func (s *AuthenticationService) SecurityQuestion(body EmailModel, headers map[string]string) (response OperationResponseOfSecurityQuestion, err error) {
	err = s.client.apiPost("/authentication/securityquestion", &body, &response, &headers)
	return
}

// Change User's password
func (s *AuthenticationService) ChangePassword(body ChangePasswordModel, headers map[string]string) (response OperationResponse, err error) {
	err = s.client.apiPost("/authentication/change/password", &body, &response, &headers)
	return
}

// Change User's Email Address
func (s *AuthenticationService) ChangeEmail(body ChangeEmailModel, headers map[string]string) (response OperationResponse, err error) {
	err = s.client.apiPost("/authentication/change/email", &body, &response, &headers)
	return
}

// Change User's Security Question
func (s *AuthenticationService) ChangeSecurityQuestion(body ChangeSecurityQuestionModel, headers map[string]string) (response OperationResponse, err error) {
	err = s.client.apiPost("/authentication/change/securityquestion", &body, &response, &headers)
	return
}

// Start Forgot Password Process
func (s *AuthenticationService) ForgotPassword(body ForgotPasswordModel, headers map[string]string) (response OperationResponse, err error) {
	err = s.client.apiPost("/authentication/forgotpassword", &body, &response, &headers)
	return
}

// Attempts to send a mobile verification code to the mobile number specified in the arguments for reset password purposes
func (s *AuthenticationService) ForgotPasswordSMS(body ShortMessageServiceModel, headers map[string]string) (response OperationResponse, err error) {
	err = s.client.apiPost("/authentication/forgotpassword/sms", &body, &response, &headers)
	return
}

// Change User's Forgotten Password by using a Mobile Verification Code
func (s *AuthenticationService) ForgotPasswordChangeSMS(body ForgotPasswordChangeBySMSModel, headers map[string]string) (response OperationResponse, err error) {
	err = s.client.apiPost("/authentication/forgotpassword/change/sms", &body, &response, &headers)
	return
}

// Change User's Forgotten Password
func (s *AuthenticationService) ForgotPasswordChange(body ForgotPasswordChangeModel, headers map[string]string) (response OperationResponse, err error) {
	err = s.client.apiPost("/authentication/forgotpassword/change", &body, &response, &headers)
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
func (s *AuthenticationService) Register(body RegistrationData) (response OperationResponseOfDictionaryOfStringAndString, err error) {
	err = s.client.apiPost("/v2/authentication/register", &body, &response, &headers)
	return
}
