package models

type UserWalletPublic struct {
	UserId      *int                  `json:"UserId,omitempty"`      // The User ID
	Balance     *float64              `json:"Balance,omitempty"`     // The User Balance
	Currency    *string               `json:"Currency,omitempty"`    // The User Currency
	UserBonuses *[]BonusProductPublic `json:"UserBonuses,omitempty"` // A List of User Bonuses Available
}
