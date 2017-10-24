package models

type WalletResponse struct {
	Id                    *string            `json:"Id,omitempty"`
	PlatformTransactionId *string            `json:"PlatformTransactionId,omitempty"`
	CreatedOn             *IGCTime           `json:"CreatedOn,omitempty"`
	Amount                *map[string]string `json:"Amount,omitempty"`
	TransactionType       *string            `json:"TransactionType,omitempty"`
	Reverts               *Reverts           `json:"Reverts,omitempty"`
}
