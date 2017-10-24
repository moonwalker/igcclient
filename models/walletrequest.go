package models

type WalletRequest struct {
	Amount          *map[string]string `json:"Amount"`              // Always required. The transaction amount and currency code
	TransactionType *string            `json:"TransactionType"`     // Always Required. The Transaction Type. Possible values are Deposit and Withdrawal
	BonusCode       *string            `json:"BonusCode,omitempty"` // Optional. The Bonus code, if any
}
