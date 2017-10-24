package models

type BonusProductPublic struct {
	ProductName *string            `json:"ProductName,omitempty"` // Product Name
	Bonuses     *[]UserBonusPublic `json:"Bonuses,omitempty"`     // A List of Bonuses for this Bonus Product
}
