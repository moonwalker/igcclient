package models

type FrontEndUserBonusObject struct {
	Name           *string  `json:"Name,omitempty"`
	Description    *string  `json:"Description,omitempty"`
	TermsLinks     *string  `json:"TermsLinks,omitempty"`
	RealBonusMoney *float64 `json:"RealBonusMoney,omitempty"`
	BonusMoney     *float64 `json:"BonusMoney,omitempty"`
	Status         *string  `json:"Status,omitempty"`
	DateCreated    *IGCTime `json:"DateCreated,omitempty"`
	WageringAmount *float64 `json:"WageringAmount,omitempty"`
	BonusWagering  *float64 `json:"BonusWagering,omitempty"`
	WagerBy        *IGCTime `json:"WagerBy,omitempty"`
	BonusId        *int     `json:"BonusId,omitempty"`
	ProductName    *string  `json:"ProductName,omitempty"`
	UserBonusId    *int     `json:"UserBonusId,omitempty"`
	ForfeitAmount  *float64 `json:"ForfeitAmount,omitempty"`
	OpenBets       *int     `json:"OpenBets,omitempty"`
	PromoCode      *string  `json:"PromoCode,omitempty"`
	RedeemTypeId   *int     `json:"RedeemTypeId,omitempty"`
}
