package models

type Constraint struct {
	Name        *string `json:"name,omitempty"`
	Operator    *string `json:"operator,omitempty"`
	Constraints *string `json:"constraints,omitempty"`
}
