package models

type PublicUserConsentsListModel struct {
	UserID               *int64                    `json:"UserId,omitempty"`
	RequiresLegacyUpdate *bool                     `json:"RequiresLegacyUpdate,omitempty"`
	UserConsents         *[]PublicUserConsentModel `json:"UserConsents,omitempty"`
}
