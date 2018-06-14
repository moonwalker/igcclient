package models

type PublicUserConsentModel struct {
	Consented       *bool
	ConsentID       *int64
	Name            *string
	Version         *float64
	IsMandatory     *bool
	ConsentContents *[]PublicConsentContentModel
}
