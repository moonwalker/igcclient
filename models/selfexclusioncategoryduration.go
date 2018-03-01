package models

type SelfExclusionCategoryDuration struct {
	ID     *int64 `json:"ID,omitempty"`
	Days   *int64 `json:"Days,omitempty"`   // Duration in day units.
	Months *int64 `json:"Months,omitempty"` // Duration in month units.
}
