package models

type BankObject struct {
	Name         *string `json:"Name,omitempty"`
	BankId       *int64  `json:"BankId,omitempty"`
	DisplayOrder *int64  `json:"DisplayOrder,omitempty"`
}
