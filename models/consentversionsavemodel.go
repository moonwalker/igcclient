package models

type ConsentVersionSaveModel struct {
	ConsentID *int64   `json:"ConsentId,omitempty"`
	Version   *float64 `json:"Version,omitempty"`
	Consented *bool    `json:"Consented,omitempty"`
}
