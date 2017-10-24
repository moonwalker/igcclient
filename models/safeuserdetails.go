package models

type SafeUserDetails struct {
	PreferredLanguageID            *int             `json:"PreferredLanguageID,omitempty"`
	CountryCode                    *string          `json:"CountryCode,omitempty"`
	Language                       *string          `json:"Language,omitempty"`
	Currency                       *string          `json:"Currency,omitempty"`
	PrefLanguage                   *string          `json:"PrefLanguage,omitempty"`
	GamemodeVisible                *bool            `json:"GamemodeVisible,omitempty"`
	AcceptedTermsOn                *IGCTime         `json:"AcceptedTermsOn,omitempty"`
	AffiliateID                    *int             `json:"AffiliateID,omitempty"`
	LastDepositDate                *IGCTime         `json:"LastDepositDate,omitempty"`
	DepositCount                   *int             `json:"DepositCount,omitempty"`
	LastLoginDate                  *IGCTime         `json:"LastLoginDate,omitempty"`
	CreatedDate                    *IGCTime         `json:"CreatedDate,omitempty"`
	AdditionalFields               *PartialUserData `json:"AdditionalFields,omitempty"`
	UserID                         *int             `json:"UserID,omitempty"`
	Username                       *string          `json:"Username,omitempty"`
	Password                       *string          `json:"Password,omitempty"`
	Email                          *string          `json:"Email,omitempty"`
	SexID                          *int             `json:"SexID,omitempty"`
	Gender                         *string          `json:"Gender,omitempty"`
	Title                          *string          `json:"Title,omitempty"`
	FirstName                      *string          `json:"FirstName,omitempty"`
	LastName                       *string          `json:"LastName,omitempty"`
	BirthDate                      *IGCTime         `json:"BirthDate,omitempty"`
	MobilePrefix                   *string          `json:"MobilePrefix,omitempty"`
	Mobile                         *string          `json:"Mobile,omitempty"`
	PhonePrefix                    *string          `json:"PhonePrefix,omitempty"`
	Phone                          *string          `json:"Phone,omitempty"`
	Address1                       *string          `json:"Address1,omitempty"`
	Address2                       *string          `json:"Address2,omitempty"`
	City                           *string          `json:"City,omitempty"`
	PostalCode                     *string          `json:"PostalCode,omitempty"`
	CountryID                      *int             `json:"CountryID,omitempty"`
	LanguageID                     *int             `json:"LanguageID,omitempty"`
	SecurityQuestionID             *int             `json:"SecurityQuestionID,omitempty"`
	SecurityAnswer                 *string          `json:"SecurityAnswer,omitempty"`
	CurrencyID                     *int             `json:"CurrencyID,omitempty"`
	AllowsNewsAndOffers            *bool            `json:"AllowsNewsAndOffers,omitempty"`
	AllowsNewsAndOffersSMS         *bool            `json:"AllowsNewsAndOffersSMS,omitempty"`
	Reference                      *string          `json:"Reference,omitempty"`
	RegistrationVerificationTypeID *int             `json:"RegistrationVerificationTypeID,omitempty"`
}
