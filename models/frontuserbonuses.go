package models

type FrontUserBonuses struct {
	UserBonusId    *int64       `json:"UserBonusId,omitempty"`
	UserId         *int64       `json:"UserId,omitempty"`
	BonusId        *int64       `json:"BonusId,omitempty"`
	LockedMoney    *float64     `json:"LockedMoney,omitempty"`
	BonusMoney     *float64     `json:"BonusMoney,omitempty"`
	Status         *int64       `json:"Status,omitempty"`
	DateCreated    *IGCTime     `json:"DateCreated,omitempty"`
	WageringAmount *float64     `json:"WageringAmount,omitempty"`
	ExpiresOn      *IGCTime     `json:"ExpiresOn,omitempty"`
	ProductId      *ProductType `json:"ProductId,omitempty"`
}
