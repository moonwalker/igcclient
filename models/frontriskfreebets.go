package models

type FrontRiskFreeBets struct {
	BonusId     *int64               `json:"BonusId,omitempty"`
	Status      *RiskFreeBetStatuses `json:"Status,omitempty"`
	DateClaimed *IGCTime             `json:"DateClaimed,omitempty"`
}
