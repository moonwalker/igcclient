package models

type PublicConsentModel struct {
	ConsentID       *int64                       `json:"ConsentId,omitempty"`
	Name            *string                      `json:"Name,omitempty"`
	Version         *float64                     `json:"Version,omitempty"`
	IsMandatory     *bool                        `json:"IsMandatory,omitempty"`
	ConsentContents *[]PublicConsentContentModel `json:"ConsentContents,omitempty"`
}
