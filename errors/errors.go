package errors

const (
	USER_NOT_FOUND                                      = 0   // User was not found.
	GAME_NOT_FOUND                                      = 2   // Game was not found.
	WITHDRAW_AMOUNT_NEGATIVE                            = 3   // Withdrawal amount specified was negative.
	DEPOSIT_AMOUNT_NEGATIVE                             = 6   // Deposit amount specified was negative.
	INSUFFICIENT_FUNDS                                  = 7   // User does not have enough funds.
	WALLET_TRANSACTION_DESCRIPTION_REQUIRED             = 8   // A wallet transaction requires a description.
	UNKOWN_ERROR                                        = 9   // Unexpected error occurred.
	CURRENCY_NOT_SUPPORTED                              = 10  // The currency code supplied was not found or is not supported.
	SERVICE_AUTHENTICATION_ERROR                        = 11  // Service Authentication Failed.
	PLAYER_LIMIT_EXCEEDED                               = 12  // Play Limit has been reached.
	NO_GAMES_FOUND_FOR_FILTER                           = 13  // No games found for the specified filters.
	INVALID_GAME_CATEGORY                               = 14  // Invalid or no Game Category Specified.
	INVALID_TRANSACTION_TYPE                            = 15  // Invalid transaction type specified.
	ACCOUNT_BLOCKED                                     = 17  // Your account has been blocked.
	KYC_SEND_DOCUMENTS                                  = 19  // Please send KYC documents by email.
	INVALID_AUTHENTICATION_TOKEN                        = 20  // You have supplied an invalid Authentication Token.
	UNAUTHORIZED_AUTHENTICATION_TOKEN                   = 21  // The authentication token passed is not authorized.
	MISSING_AUTHENTICATION_TOKEN                        = 22  // No authentication token was found with the request.
	ACCOUNT_COOLDOWN                                    = 23  // Your account has been locked for some minutes because there are too many failed attempts.
	EMAIL_NOT_VERIFIED                                  = 24  // Your email was not verified.
	GAME_SESSION_INVALID                                = 25  // An invalid game session was supplied.
	GAME_SESSION_EXPIRED                                = 26  // The game session supplied is expired.
	USER_IS_SELF_EXCLUDED                               = 27  // The user is self-excluded.
	TNC_ACCEPTANCE_REQUIRED                             = 28  // The user is required to accept Terms & Conditions.
	USER_IP_RESTRICTED                                  = 29  // The user's IP address is restricted.
	USER_NEEDS_COOLING_OFF                              = 30  // Self-exclusion is over and the user must contact the operator to lift the exclusion.
	USER_IN_COOLING_OFF                                 = 31  // Self-exclusion is over but the user is in the cooling period.
	DAILY_PROTECTION_EXCEEDED                           = 32  // The user exceeded their daily protection limit.
	WEEKLY_PROTECTION_EXCEEDED                          = 35  // The user exceeded their weekly protection limit.
	MONTHLY_PROTECTION_EXCEEDED                         = 37  // The user exceeded their monthly protection limit.
	GAMEPLAY_DURATION_EXCEEDED                          = 38  // The user exceeded their game play duration.
	LOSS_LIMIT_EXCEEDED                                 = 39  // The user exceeded their loss limit.
	USER_BLOCKED_FROM_GAME                              = 40  // The user is not permitted to play this game.
	USERNAME_ALREADY_EXISTS                             = 41  // This username already exists.
	EMAIL_ALREADY_EXISTS                                = 42  // This email already exists.
	CANNOT_DELEGATE_DELEGATED                           = 43  // You cannot delegate and already delegated token.
	INTERNAL_SERVER_ERROR                               = 44  // An Internal Server error has occurred! Please check your request and try again. If you think your request is correct, please contact the person in charge.
	INVALID_VERIFICATION_CODE                           = 45  // The Verification Code you have entered is invalid or has already been used.
	TRANSACTION_ALREADY_PROCESSED_DIFF                  = 46  // The transaction has already been processed, with different details.
	TRANSACTION_ALREADY_PROCESSED                       = 47  // The transaction has already been processed.
	INVALID_PASSWORD                                    = 48  // You have supplied an invalid Password.
	INVALID_SECURITY_ANSWER                             = 49  // The Security Answer you have supplied is Invalid.
	EMAIL_NOT_FOUND                                     = 50  // No user was found for the supplied email address.
	INVALID_PASSWORD_RECOVERY_CODE                      = 51  // The Password Recovery Code you have supplied is invalid.
	USER_LOGGED_OUT                                     = 52  // The session token used has been logged out.
	CHALLENGE_INVALID                                   = 54  // Your challenge is invalid. This means that either it has been already used or it has timed out.
	DATE_RANGE_TOO_BIG                                  = 55  // The date range requested is too big (Maximum is 62 days).
	SYSTEM_WALLET_NOT_INITIALISED                       = 56  // The system wallet has not been initialised correctly.
	PARAMETER_NOT_SUPPLIED                              = 57  // A required parameter was not supplied.
	WAGER_NOT_FOUND                                     = 58  // The casino wager was not found.
	USER_EXCEEDED_LIMITS                                = 59  // The user exceeded his/her limit.
	GAME_ACTIVITY_NOT_FOUND                             = 60  // The game activity was not found.
	OPERATION_NOT_ALLOWED                               = 61  // The operation is not allowed.
	USER_SELF_EXCLUDED                                  = 62  // User has self-excluded himself and is not allowed to login.
	USERNAME_TOO_LONG                                   = 65  // The username you have entered exceeds 17 characters.
	RESPONSIBLE_GAMING_LIMIT_EXCEEDED                   = 66  // The user exceeded his/her limit.
	REQUEST_DATA_INVALID                                = 98  // Request data is missing entirely or has critical parts missing.
	MOBILE_PREFIX_NULL_OR_EMPTY                         = 99  // Mobile prefix is null or empty.
	MOBILE_NULL_OR_EMPTY                                = 100 // Mobile is null or empty.
	MOBILE_OR_PREFIX_INVALID                            = 101 // Mobile and prefix combination is invalid.
	MOBILE_MOBILE_VERIFICATION_CODE_COMBINATION_INVALID = 102 // Mobile & mobile verification code combination is invalid.
	MOBILE_VERIFICATION_CODE_NULL_OR_EMPTY              = 103 // Mobile verification code is null or empty.
	SMS_PROVIDER_ERROR                                  = 104 // Something went wrong with the SMS provider.
	SMS_NOT_DELIVERED                                   = 105 // SMS not delivered to recipient.
	X_FORWARDED_FOR_NOT_VALID_IPv4                      = 106 // X-Forwarded-For did not contain a valid IPv4 IP.
	VERIFICATION_TYPE_ID_NOT_VALID_FOR_IP               = 107 // Verification Type is not valid for IP.
	VALIDATION_ERROR                                    = 108 // Validation Error.
	LAST_SMS_NOT_DELIVERED                              = 109 // Last SMS was not delivered to recipient.
	PHONE_OR_PREFIX_INVALID                             = 110 // Phone and phone code combination is invalid.
	EMAIL_IS_INVALID                                    = 111 // Email is not a valid email.
	USERNAME_IS_INVALID                                 = 112 // Username is not a valid username.
	ERROR_NOT_DEFINED                                   = 113 // Error has not yet been defined.
	MOBILE_AND_PREFIX_COMBINATION_ALREADY_REGISTERED    = 114 // Mobile and prefix combination is already registered with us.
	PHONE_AND_PREFIX_COMBINATION_ALREADY_REGISTERED     = 115 // Phone and phone code combination is already registered with us.
	MOBILE_AND_PREFIX_INVALID                           = 116 // Mobile and prefix combination is invalid.
	PHONE_AND_PREFIX_INVALID                            = 117 // Phone and phone prefix combination is invalid.
	MOBILE_AND_PREFIX_NOT_TRUSTWORTHY                   = 118 // Mobile and prefix combination has been deemed untrustworthy.
	PHONE_AND_PREFIX_NOT_TRUSTWORTHY                    = 119 // Phone and phone prefix combination has been deemed untrustworthy.
	SESSION_SELF_EXCLUSION                              = 120 // Player is not allowed to have a session because of a session self-exclusion.
	STAKE_TOO_HIGH                                      = 121 // Stake too high.
	FPP_LESS_THAN_ZERO                                  = 122 // FPP Value is less than 0.
	MISSING_USER_AGENT                                  = 123 // No Client UserAgent was found with the request.
	USER_IS_NOT_OF_LEGAL_AGE                            = 124 // User is not of legal age.
	INPUT_ILLEGAL                                       = 125 // The odds you have chosen cannot be placed together.
	SESSION_STAKE_PROTECTION_EXCEEDED                   = 140 // Session stake protection exceeded.
	MAX_STAKE_PROTECTION_EXCEEDED                       = 141 // Max stake protection exceeded.
	GAME_IS_OFFLINE                                     = 142 // This Game is currently offline.
	DEMO_PLAY_NOT_SUPPORTED                             = 143 // This Game does not support demo play.
	SMS_HTTP_REQUEST_FAILED                             = 144 // Communication between server and SMS provider failed.
	SMS_TIMEOUT_HAS_NOT_EXPIRED                         = 145 // Timeout for mobile number has not yet expired.
	SMS_STATUS_CHECK_FAILED                             = 146 // Communication between server and SMS status provider failed.
	SMS_STATUS_CHECK_NO_REFID                           = 147 // No reference id was found for SMS record.
	MAXIMUM_BONUS_BET                                   = 148 // User bet over the maximum bonus bet limit.
)
