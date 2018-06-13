package models

type PayRequestModel struct {
	PaymentMethodID *int64   `json:"PaymentMethodId"`           // Always required. The PaymentMethodId should be retrieved by calling GetUserDepositMethods or GetUserWithdrawMethods
	PaymentTypeID   *int64   `json:"PaymentTypeId"`             // Always required. The PaymentTypeId should be retrieved by calling GetUserDepositMethods or GetUserWithdrawMethods
	Amount          *float64 `json:"Amount"`                    // Always required. The amount the users has requested
	BonusCode       *string  `json:"BonusCode,omitempty"`       // Optional. The Bonus Code, if any
	LanguageAlpha2  *string  `json:"LanguageAlpha2,omitempty"`  // Optional. The Language Alpha2 the user has choosen. If none passed, the default user language will be used
	UserIP          *string  `json:"UserIP"`                    // Always required. The user IP Address
	UserAgent       *string  `json:"UserAgent"`                 // Always required. The user User Agent
	AccountID       *string  `json:"AccountId,omitempty"`       // Required only for Wallet Deposit or Withdrawal. The Wallet Account Id. This is required for payment types like Neteller
	SecureID        *string  `json:"SecureId,omitempty"`        // Required only for Wallet Deposit or Withdrawal. The Wallet Secure Id. This is required for payment types like Neteller
	CreditCard      *string  `json:"CreditCard,omitempty"`      // Required only for IGC Credit Card Deposit. This is the credit card number. Not required for EPG
	ExpiryDate      *string  `json:"ExpiryDate,omitempty"`      // Required only for IGC Credit Card Deposit. This is the credit card expiry date. Not required for EPG
	CV2             *string  `json:"CV2,omitempty"`             // Required only for IGC Credit Card Deposit. This is the credit card cv2. Not required for EPG
	UserCardID      *int64   `json:"UserCardId,omitempty"`      // Required only for IGC or EPG Credit Card Withdrawals. This is mandatory for withdrawals and should be available from the meta data api call.
	PayeeName       *string  `json:"PayeeName,omitempty"`       // Required only for Bank Withdrawl. This is the Customer name to be shown in the bank withdrawal transaction
	CountryCode     *string  `json:"CountryCode,omitempty"`     // Required only for Bank Withdrawl. This is the Country code to be shown in the bank withdrawal transaction
	AccountNumber   *string  `json:"AccountNumber,omitempty"`   // Required only for Bank Withdrawl. This is the Account Number to be shown in the bank withdrawal transaction
	BankName        *string  `json:"BankName,omitempty"`        // Required only for Bank Withdrawl. This is the Bank Name to be shown in the bank withdrawal transaction
	BankCode        *string  `json:"BankCode,omitempty"`        // Required only for Bank Withdrawl. This is the Bank Code to be shown in the bank withdrawal transaction
	BranchCode      *string  `json:"BranchCode,omitempty"`      // Required only for Bank Withdrawl. This is the Branch Code to be shown in the bank withdrawal transaction
	BranchAddress   *string  `json:"BranchAddress,omitempty"`   // Required only for Bank Withdrawl. This is the Branch Address to be shown in the bank withdrawal transaction
	Iban            *string  `json:"Iban,omitempty"`            // Required only for Bank Withdrawl. This is the IBAN to be shown in the bank withdrawal transaction
	Swift           *string  `json:"Swift,omitempty"`           // Required only for Bank Withdrawl. This is the SWIFT/BIC to be shown in the bank withdrawal transaction
	AdditionalInfo1 *string  `json:"AdditionalInfo1,omitempty"` // Required only for Bank Withdrawl. This is the Additional Info 1 to be shown in the bank withdrawal transaction
	AdditionalInfo2 *string  `json:"AdditionalInfo2,omitempty"` // Required only for Bank Withdrawl. This is the Additional Info 2 to be shown in the bank withdrawal transaction
	AdditionalInfo3 *string  `json:"AdditionalInfo3,omitempty"` // Required only for Bank Withdrawl. This is the Additional Info 3 to be shown in the bank withdrawal transaction
}
