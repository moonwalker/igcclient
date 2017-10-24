package models

type LimitDurationObject struct {
	LimitDurationType *LimitDuration `json:"LimitDurationType,omitempty"` // The limit duration type.
	Name              *string        `json:"Name,omitempty"`              // A descriptive name of the limit duration.
	Visible           *bool          `json:"Visible,omitempty"`           // Shows whether duration is to be made visible from the front-end
}
