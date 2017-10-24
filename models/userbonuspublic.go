package models

type UserBonusPublic struct {
	BonusName      *string  `json:"BonusName,omitempty"`      // Bonus Name
	BonusMoney     *float64 `json:"BonusMoney,omitempty"`     // Bonus Money
	RealBonusMoney *float64 `json:"RealBonusMoney,omitempty"` // Locked Bonus Money
}
