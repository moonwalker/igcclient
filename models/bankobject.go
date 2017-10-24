package models

type BankObject struct {
	Name         *string `json:"Name,omitempty"`
	BankId       *int    `json:"BankId,omitempty"`
	DisplayOrder *int    `json:"DisplayOrder,omitempty"`
}
