package models

type FrontUserBonuses struct {
	UserBonusID    *int64       `json:"UserBonusId,omitempty"`
	UserID         *int64       `json:"UserId,omitempty"`
	BonusID        *int64       `json:"BonusId,omitempty"`
	LockedMoney    *float64     `json:"LockedMoney,omitempty"`
	BonusMoney     *float64     `json:"BonusMoney,omitempty"`
	Status         *int64       `json:"Status,omitempty"`
	DateCreated    *IGCTime     `json:"DateCreated,omitempty"`
	WageringAmount *float64     `json:"WageringAmount,omitempty"`
	ExpiresOn      *IGCTime     `json:"ExpiresOn,omitempty"`
	ProductID      *ProductType `json:"ProductId,omitempty"`
}
