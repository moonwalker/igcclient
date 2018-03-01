package models

type FrontUserFreeSpins struct {
	UserFreeSpinsID *int64   `json:"UserFreeSpinsID,omitempty"`
	UserId          *int64   `json:"UserId,omitempty"`
	BonusId         *int64   `json:"BonusId,omitempty"`
	BonusName       *string  `json:"BonusName,omitempty"`
	VendorId        *int64   `json:"VendorId,omitempty"`
	DateCreated     *IGCTime `json:"DateCreated,omitempty"`
	PromoCode       *string  `json:"PromoCode,omitempty"`
}
