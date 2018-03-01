package models

type UpdateUserObject struct {
	Address1               *string          `json:"Address1,omitempty"`
	Address2               *string          `json:"Address2,omitempty"`
	PostalCode             *string          `json:"PostalCode,omitempty"`
	City                   *string          `json:"City,omitempty"`
	MobilePrefix           *string          `json:"MobilePrefix,omitempty"`
	Mobile                 *string          `json:"Mobile,omitempty"`
	PhonePrefix            *string          `json:"PhonePrefix,omitempty"`
	Phone                  *string          `json:"Phone,omitempty"`
	AllowsNewsAndOffers    *bool            `json:"AllowsNewsAndOffers,omitempty"`
	AllowsNewsAndOffersSMS *bool            `json:"AllowsNewsAndOffersSMS,omitempty"`
	AdditionalFields       *PartialUserData `json:"AdditionalFields,omitempty"`
	LanguageID             *int64           `json:"LanguageID,omitempty"`
}
