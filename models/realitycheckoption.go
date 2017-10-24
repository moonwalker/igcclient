package models

type RealityCheckOption struct {
	CountryId *int `json:"CountryId,omitempty"` // Country internal Id
	Interval  *int `json:"Interval,omitempty"`  // Interval in seconds
}
