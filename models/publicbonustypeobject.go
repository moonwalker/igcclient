package models

type PublicBonusTypeObject struct {
	TypeId    *int64  `json:"TypeId,omitempty"`
	ProductId *int64  `json:"ProductId,omitempty"`
	Type      *string `json:"Type,omitempty"`
}
