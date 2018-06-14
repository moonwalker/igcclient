package models

type FrontUserFreeSpins struct {
	UserFreeSpinsID *int64   `json:"UserFreeSpinsID,omitempty"`
	UserID          *int64   `json:"UserId,omitempty"`
	BonusID         *int64   `json:"BonusId,omitempty"`
	BonusName       *string  `json:"BonusName,omitempty"`
	VendorID        *int64   `json:"VendorId,omitempty"`
	DateCreated     *IGCTime `json:"DateCreated,omitempty"`
	PromoCode       *string  `json:"PromoCode,omitempty"`
}
