package models

type Limit struct {
	LimitType *LimitType `json:"LimitType,omitempty"` // The Limit Type
	Name      *string    `json:"Name,omitempty"`      // The Name of the Limit
	Visible   *bool      `json:"Visible,omitempty"`   // Is the limit visible?
}
