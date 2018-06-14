package models

type QuickPayRequestModel struct {
	PaymentTypeID  *int64   `json:"PaymentTypeId"`            // Always required. The PaymentTypeId
	Amount         *float64 `json:"Amount"`                   // Always required. The amount the users has requested
	LanguageAlpha2 *string  `json:"LanguageAlpha2,omitempty"` // Optional. The Language Alpha2 the user has choosen. If none passed, the default user language will be used
	UserIP         *string  `json:"UserIP"`                   // Always required. The user IP Address
	UserAgent      *string  `json:"UserAgent"`                // Always required. The user User Agent
	WalletCardID   *int64   `json:"WalletCardId"`             // Always required. The Wallet Account or Card Id.
	SecureID       *string  `json:"SecureId,omitempty"`       // Required only for Wallet Deposit. The Wallet Secure Id. This is required for payment types like Neteller
	CV2            *string  `json:"CV2,omitempty"`            // Required only for Credit Card Deposit. This is the credit card cv2.
	BonusCode      *string  `json:"BonusCode,omitempty"`      // Optional. The Bonus Code, if any
}
