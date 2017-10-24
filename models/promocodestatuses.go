package models

type PromoCodeStatuses int

const (
	PCS_Created   PromoCodeStatuses = 1
	PCS_Claimed   PromoCodeStatuses = 2
	PCS_Cancelled PromoCodeStatuses = 3
	PCS_Rejected  PromoCodeStatuses = 4
	PCS_Expired   PromoCodeStatuses = 5
)
