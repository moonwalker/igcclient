package models

type UserObject struct {
	PasswordSalt                   *string  `json:"PasswordSalt,omitempty"`
	SecurityAnswerSalt             *string  `json:"SecurityAnswerSalt,omitempty"` // The Security Answer Salt
	SignupIP                       *string  `json:"SignupIP,omitempty"`           // The IP used during signup
	LanguageCode                   *string  `json:"LanguageCode,omitempty"`       // The ISO 639-1 Language Code
	SecurityQuestion               *string  `json:"SecurityQuestion,omitempty"`   // The Security Question
	GamMatrixUserID                *int64   `json:"GamMatrixUserID,omitempty"`    // The GamMatrix User ID. The User ID of GamMatrix
	Created                        *IGCTime `json:"Created,omitempty"`            // The Date on which the user was created
	EmailGUID                      *string  `json:"EmailGUID,omitempty"`          // The GUID which is used to verify the email
	LoginTries                     *int64   `json:"LoginTries,omitempty"`         // The number of login tries
	LoginTryCoolDown               *IGCTime `json:"LoginTryCoolDown,omitempty"`   // The cool down needed before next login
	PreferredLanguage              *string  `json:"PreferredLanguage,omitempty"`  // The Preferred Language
	KYC                            *bool    `json:"KYC,omitempty"`                // Needs KYC
	IsBlocked                      *bool    `json:"IsBlocked,omitempty"`          // This is true if the user is blocked.
	GamemodeVisible                *bool    `json:"GamemodeVisible,omitempty"`    // Should you show the game mode?
	TrackFirstDeposit              *bool    `json:"TrackFirstDeposit,omitempty"`  // Track the first deposit?
	AccountVerifiedOn              *IGCTime `json:"AccountVerifiedOn,omitempty"`  // When was the account verified on?
	KYCIDApproved                  *bool    `json:"KYCIDApproved,omitempty"`
	KYCAddressApproved             *bool    `json:"KYCAddressApproved,omitempty"`
	KYCPaymentApproved             *bool    `json:"KYCPaymentApproved,omitempty"`
	IsSuperAdmin                   *bool    `json:"IsSuperAdmin,omitempty"` // Is the user a super admin?
	KYCVerifiedBefore              *bool    `json:"KYCVerifiedBefore,omitempty"`
	ChatGUID                       *string  `json:"ChatGUID,omitempty"`        // The Chat GUID
	DepositCount                   *int64   `json:"DepositCount,omitempty"`    // The Deposit Count
	LastDepositDate                *IGCTime `json:"LastDepositDate,omitempty"` // The last deposit Date
	UserID                         *int64   `json:"UserID,omitempty"`
	Username                       *string  `json:"Username,omitempty"`
	Password                       *string  `json:"Password,omitempty"`
	Email                          *string  `json:"Email,omitempty"`
	SexID                          *int64   `json:"SexID,omitempty"`
	Gender                         *string  `json:"Gender,omitempty"`
	Title                          *string  `json:"Title,omitempty"`
	FirstName                      *string  `json:"FirstName,omitempty"`
	LastName                       *string  `json:"LastName,omitempty"`
	BirthDate                      *IGCTime `json:"BirthDate,omitempty"`
	MobilePrefix                   *string  `json:"MobilePrefix,omitempty"`
	Mobile                         *string  `json:"Mobile,omitempty"`
	PhonePrefix                    *string  `json:"PhonePrefix,omitempty"`
	Phone                          *string  `json:"Phone,omitempty"`
	Address1                       *string  `json:"Address1,omitempty"`
	Address2                       *string  `json:"Address2,omitempty"`
	City                           *string  `json:"City,omitempty"`
	PostalCode                     *string  `json:"PostalCode,omitempty"`
	CountryID                      *int64   `json:"CountryID,omitempty"`
	LanguageID                     *int64   `json:"LanguageID,omitempty"`
	SecurityQuestionID             *int64   `json:"SecurityQuestionID,omitempty"`
	SecurityAnswer                 *string  `json:"SecurityAnswer,omitempty"`
	CurrencyID                     *int64   `json:"CurrencyID,omitempty"`
	AllowsNewsAndOffers            *bool    `json:"AllowsNewsAndOffers,omitempty"`
	AllowsNewsAndOffersSMS         *bool    `json:"AllowsNewsAndOffersSMS,omitempty"`
	Reference                      *string  `json:"Reference,omitempty"`
	RegistrationVerificationTypeID *int64   `json:"RegistrationVerificationTypeID,omitempty"`
}
