package models

type OMUserWalletPublic struct {
	Amount              *float64 `json:"Amount,omitempty"`              // Real cash amount that you can bet with
	FirstDepositBalance *float64 `json:"FirstDepositBalance,omitempty"` // Amount of money available from you deposit bonus
	FreeBetBalance      *float64 `json:"FreeBetBalance,omitempty"`      // Amount on money available for your free bet
	Currency            *string  `json:"Currency,omitempty"`            // Currency
	BalanceTime         *IGCTime `json:"BalanceTime,omitempty"`         // When the balance was checked
	Error               *bool    `json:"Error,omitempty"`               // If there was an error in the call
	ErrorMessage        *string  `json:"ErrorMessage,omitempty"`        // Error message from OM
}
