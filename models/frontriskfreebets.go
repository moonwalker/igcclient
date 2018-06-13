package models

type FrontRiskFreeBets struct {
	BonusID     *int64               `json:"BonusId,omitempty"`
	Status      *RiskFreeBetStatuses `json:"Status,omitempty"`
	DateClaimed *IGCTime             `json:"DateClaimed,omitempty"`
}
