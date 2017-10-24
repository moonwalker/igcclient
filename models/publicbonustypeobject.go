package models

type PublicBonusTypeObject struct {
	TypeId    *int    `json:"TypeId,omitempty"`
	ProductId *int    `json:"ProductId,omitempty"`
	Type      *string `json:"Type,omitempty"`
}
