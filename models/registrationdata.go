package models

type RegistrationData struct {
	VerificationTypeID     *int64                      `json:"VerificationTypeID"`               // VerificationTypeID from UserVerifyRegistrationTypes()
	Username               *string                     `json:"Username"`                         // Username - tested against ^[a-zA-Z0-9._-]{4,20}$
	Password               *string                     `json:"Password"`                         // Password regex is to be validated against regex found in Core.Settings Innovation.PublicAPI.Authentication.PasswordComplexityRegex
	Email                  *string                     `json:"Email,omitempty"`                  // email address
	SexID                  *int64                      `json:"SexID,omitempty"`                  // sex identifier
	Title                  *string                     `json:"Title,omitempty"`                  // title.
	FirstName              *string                     `json:"FirstName"`                        // first name. - Required
	LastName               *string                     `json:"LastName"`                         // last name. - Required
	BirthDate              *IGCTime                    `json:"BirthDate"`                        // birth date - Required
	Address1               *string                     `json:"Address1"`                         // address1 - Required
	Address2               *string                     `json:"Address2,omitempty"`               // address2.
	PostalCode             *string                     `json:"PostalCode,omitempty"`             // postal code
	City                   *string                     `json:"City,omitempty"`                   // city.
	CountryID              *int64                      `json:"CountryID,omitempty"`              // country identifier.
	LanguageID             *int64                      `json:"LanguageID,omitempty"`             // language identifier.
	MobilePrefix           *string                     `json:"MobilePrefix"`                     // mobile prefix. (+356, +21 etc)
	Mobile                 *string                     `json:"Mobile"`                           // mobile - Required
	MobileVerificationCode *string                     `json:"MobileVerificationCode,omitempty"` // mobile verification code.
	PhonePrefix            *string                     `json:"PhonePrefix,omitempty"`            // phone prefix.
	Phone                  *string                     `json:"Phone,omitempty"`                  // phone
	SecurityQuestionID     *int64                      `json:"SecurityQuestionID,omitempty"`     // security question identifier. set as 0 for not set.
	SecurityAnswer         *string                     `json:"SecurityAnswer,omitempty"`         // security answer
	CurrencyID             *int64                      `json:"CurrencyID,omitempty"`             // currency identifier
	AllowsNewsAndOffers    *bool                       `json:"AllowsNewsAndOffers,omitempty"`    // value indicating whether client allows news and offers by Email
	AllowsNewsAndOffersSMS *bool                       `json:"AllowsNewsAndOffersSMS,omitempty"` // value indicating whether client allows news and offers by SMS
	Reference              *string                     `json:"Reference,omitempty"`              // btag value reference
	UrlReferrer            *string                     `json:"UrlReferrer,omitempty"`            // Request URL referrer
	UserAgent              *string                     `json:"UserAgent,omitempty"`              // client's user agent
	AffiliateSourceID      *string                     `json:"AffiliateSourceID,omitempty"`      // Optional fields represents unique identifier for a particular affiliate. Usually retrieved from the query string or cookie.
	AffiliateClientUUID    *string                     `json:"AffiliateClientUUID,omitempty"`    // Optional field represents unique identifier for user in association with the affiliate. Usually retrieved from the query string or cookie.
	Tags                   *[]string                   `json:"Tags,omitempty"`                   // Optional: A list of tags to be associated to the user. Supported tags: Sports, Casino, Poker
	UserAdditionalFields   *RegistrationAdditionalData `json:"UserAdditionalFields,omitempty"`   // Contains additional optional fields in relation to user registration
}
