package models

type BankObject struct {
	Name         *string `json:"Name,omitempty"`
	BankID       *int64  `json:"BankId,omitempty"`
	DisplayOrder *int64  `json:"DisplayOrder,omitempty"`
}
