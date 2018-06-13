package models

type PublicUserConsentsListModel struct {
	UserId               *int64                    `json:"UserId,omitempty"`
	RequiresLegacyUpdate *bool                     `json:"RequiresLegacyUpdate,omitempty"`
	UserConsents         *[]PublicUserConsentModel `json:"UserConsents,omitempty"`
}
