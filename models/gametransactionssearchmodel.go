package models

type GameTransactionsSearchModel struct {
	From *IGCTime `json:"From,omitempty"`
	To   *IGCTime `json:"To,omitempty"`
}
