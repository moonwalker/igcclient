package models

type PublicBonusesObject struct {
	BonusID       *int64
	TypeID        *AllBonusTypes
	Name          *string
	Description   *string
	DateFrom      *IGCTime
	DateTo        *IGCTime
	Period        *string
	Enabled       *bool
	TermsLink     *string
	ManualBonus   *bool
	MaxBetOn      *bool
	IncludeInList *bool
	PromoCodes    *[]string
	ProductID     *ProductType
	RedeemTypeID  *int64
}
