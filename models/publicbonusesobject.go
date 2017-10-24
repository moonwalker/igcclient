package models

type PublicBonusesObject struct {
	BonusId *int
	TypeId *AllBonusTypes
	Name *string
	Description *string
	DateFrom *IGCTime
	DateTo *IGCTime
	Period *string
	Enabled *bool
	TermsLink *string
	ManualBonus *bool
	MaxBetOn *bool
	IncludeInList *bool
	PromoCodes *[]string
	ProductId *ProductType
	RedeemTypeId *int
}
