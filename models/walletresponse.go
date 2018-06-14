package models

type WalletResponse struct {
	ID                    *string            `json:"Id,omitempty"`
	PlatformTransactionID *string            `json:"PlatformTransactionId,omitempty"`
	CreatedOn             *IGCTime           `json:"CreatedOn,omitempty"`
	Amount                *map[string]string `json:"Amount,omitempty"`
	TransactionType       *string            `json:"TransactionType,omitempty"`
	Reverts               *Reverts           `json:"Reverts,omitempty"`
}
