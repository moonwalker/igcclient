package models

type BonusDetails struct {
	UserBonuses  *FrontUserBonuses   `json:"UserBonuses,omitempty"`
	FreeSpins    *FrontUserFreeSpins `json:"FreeSpins,omitempty"`
	RiskFreeBets *FrontRiskFreeBets  `json:"RiskFreeBets,omitempty"`
}
