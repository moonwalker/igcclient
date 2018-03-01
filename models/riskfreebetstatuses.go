package models

type RiskFreeBetStatuses int64

const (
	RFBS_Claimed   RiskFreeBetStatuses = 1
	RFBS_Pending   RiskFreeBetStatuses = 2
	RFBS_Cancelled RiskFreeBetStatuses = 3
	RFBS_Rejected  RiskFreeBetStatuses = 4
	RFBS_Expired   RiskFreeBetStatuses = 5
	RFBS_Used      RiskFreeBetStatuses = 6
)
