package models

type SelfExclusionCategoryDuration struct {
	ID     *int `json:"ID,omitempty"`
	Days   *int `json:"Days,omitempty"`   // Duration in day units.
	Months *int `json:"Months,omitempty"` // Duration in month units.
}
