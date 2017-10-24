package models

type FrontUserBonuses struct {
	UserBonusId    *int         `json:"UserBonusId,omitempty"`
	UserId         *int         `json:"UserId,omitempty"`
	BonusId        *int         `json:"BonusId,omitempty"`
	LockedMoney    *float64     `json:"LockedMoney,omitempty"`
	BonusMoney     *float64     `json:"BonusMoney,omitempty"`
	Status         *int         `json:"Status,omitempty"`
	DateCreated    *IGCTime     `json:"DateCreated,omitempty"`
	WageringAmount *float64     `json:"WageringAmount,omitempty"`
	ExpiresOn      *IGCTime     `json:"ExpiresOn,omitempty"`
	ProductId      *ProductType `json:"ProductId,omitempty"`
}
