package models

type PublicBonusTypeObject struct {
	TypeID    *int64  `json:"TypeId,omitempty"`
	ProductID *int64  `json:"ProductId,omitempty"`
	Type      *string `json:"Type,omitempty"`
}
