package models

type WalletTransactionsSearchModel struct {
	From *IGCTime `json:"From,omitempty"`
	To   *IGCTime `json:"To,omitempty"`
}
