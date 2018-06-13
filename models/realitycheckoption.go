package models

type RealityCheckOption struct {
	CountryID *int64 `json:"CountryId,omitempty"` // Country internal Id
	Interval  *int64 `json:"Interval,omitempty"`  // Interval in seconds
}
