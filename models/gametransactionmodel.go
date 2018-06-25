package models

type GameTransactionModel struct {
	ActivityID    *string  `json:"ActivityId,omitempty"`    // The transaction identifier.
	Date          *IGCTime `json:"Date,omitempty"`          // The date / time of the transaction.
	Game          *string  `json:"Game,omitempty"`          // The name of the game.
	Vendor        *string  `json:"Vendor,omitempty"`        // The game vendors name.
	Currency      *string  `json:"Currency,omitempty"`      // The currency of the transaction.
	ActivityType  *string  `json:"ActivityType,omitempty"`  // The type of the activity.
	Bet           *float64 `json:"Bet,omitempty"`           // The bet value.
	Win           *float64 `json:"Win,omitempty"`           // The win value.
	BalanceBefore *float64 `json:"BalanceBefore,omitempty"` // The balance before value.
	BalanceAfter  *float64 `json:"BalanceAfter,omitempty"`  // The balance after value.
	WalletAmount  *float64 `json:"WalletAmount,omitempty"`  // The Wallet amount.
	BonusAmount   *float64 `json:"BonusAmount,omitempty"`   // The Bonus Wallet amount.
	Bonus         *string  `json:"Bonus,omitempty"`         // The Bonus Wallet description.
}
