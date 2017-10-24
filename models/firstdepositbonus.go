package models

type FirstDepositBonus struct {
	BonusID            *int         `json:"BonusID,omitempty"`
	BonusName          *string      `json:"BonusName,omitempty"`
	Status             *string      `json:"Status,omitempty"`
	ManualBonus        *bool        `json:"ManualBonus,omitempty"`
	StartDate          *IGCTime     `json:"StartDate,omitempty"`
	EndDate            *IGCTime     `json:"EndDate,omitempty"`
	ExpireDate         *IGCTime     `json:"ExpireDate,omitempty"`
	LiveBettingOpt     *string      `json:"LiveBettingOpt,omitempty"`
	IncludeHorseRacing *bool        `json:"IncludeHorseRacing,omitempty"`
	TermsAndCondition  *string      `json:"TermsAndCondition,omitempty"`
	Currency           *string      `json:"Currency,omitempty"`
	ExRateToEUR        *float64     `json:"ExRateToEUR,omitempty"`
	Amount             *float64     `json:"Amount,omitempty"`
	CanBeRefused       *bool        `json:"CanBeRefused,omitempty"`
	CurrentTurnover    *float64     `json:"CurrentTurnover,omitempty"`
	TurnoverLeft       *float64     `json:"TurnoverLeft,omitempty"`
	RequiredOdds       *float64     `json:"RequiredOdds,omitempty"`
	Percentage         *string      `json:"Percentage,omitempty"`
	MinimumDeposit     *float64     `json:"MinimumDeposit,omitempty"`
	MaximumDeposit     *float64     `json:"MaximumDeposit,omitempty"`
	BetType            *string      `json:"BetType,omitempty"`
	Withdrawable       *bool        `json:"Withdrawable,omitempty"`
	PromoCode          *string      `json:"PromoCode,omitempty"`
	ReceiveDate        *IGCTime     `json:"ReceiveDate,omitempty"`
	RedeemedAmount     *float64     `json:"RedeemedAmount,omitempty"`
	RedeemedTime       *IGCTime     `json:"RedeemedTime,omitempty"`
	BonusStatus        *BonusStatus `json:"BonusStatus,omitempty"`
}
