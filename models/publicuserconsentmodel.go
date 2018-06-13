package models

type PublicUserConsentModel struct {
	Consented       *bool
	ConsentId       *int64
	Name            *string
	Version         *float64
	IsMandatory     *bool
	ConsentContents *[]PublicConsentContentModel
}
