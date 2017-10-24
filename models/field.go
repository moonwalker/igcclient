package models

type Field struct {
	Name        *string       `json:"name,omitempty"`
	Placeholder *string       `json:"placeholder,omitempty"`
	DesignType  *string       `json:"design:type,omitempty"`
	Constraints *[]Constraint `json:"constraints,omitempty"`
	Meta        *Meta         `json:"meta,omitempty"`
	Values      *string       `json:"values,omitempty"`
}
