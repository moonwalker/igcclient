package models

type UserResponseDTO struct {
	UserID            *int64                      `json:"UserId,omitempty"`
	Username          *string                     `json:"Username,omitempty"`
	Email             *string                     `json:"Email,omitempty"`
	FirstName         *string                     `json:"FirstName,omitempty"`
	LastName          *string                     `json:"LastName,omitempty"`
	BirthDate         *IGCTime                    `json:"BirthDate,omitempty"`
	GenderID          *int64                      `json:"GenderId,omitempty"`
	AccountVerifiedOn *IGCTime                    `json:"AccountVerifiedOn,omitempty"`
	VerificationType  *VerificationType           `json:"VerificationType,omitempty"`
	AuthMethod        *AuthMethod                 `json:"AuthMethod,omitempty"`
	Phone             *PhoneResponseDTO           `json:"Phone,omitempty"`
	AffiliateID       *int64                      `json:"AffiliateId,omitempty"`
	CreatedDate       *IGCTime                    `json:"CreatedDate,omitempty"`
	Language          *LanguageResponseDTO        `json:"Language,omitempty"`
	Country           *CountryResponseDTO         `json:"Country,omitempty"`
	Currency          *CurrencyResponseDTO        `json:"Currency,omitempty"`
	Address           *AddressResponseDTO         `json:"Address,omitempty"`
	Document          *DocumentResponseDTO        `json:"Document,omitempty"`
	ExternalDetails   *ExternalDetailsResponseDTO `json:"ExternalDetails,omitempty"`
	DisplayName       *string                     `json:"DisplayName,omitempty"`
}
