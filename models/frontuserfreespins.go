package models

type FrontUserFreeSpins struct {
	UserFreeSpinsID *int     `json:"UserFreeSpinsID,omitempty"`
	UserId          *int     `json:"UserId,omitempty"`
	BonusId         *int     `json:"BonusId,omitempty"`
	BonusName       *string  `json:"BonusName,omitempty"`
	VendorId        *int     `json:"VendorId,omitempty"`
	DateCreated     *IGCTime `json:"DateCreated,omitempty"`
	PromoCode       *string  `json:"PromoCode,omitempty"`
}
