package models

type RealityCheckUserSetting struct {
	UserId   *int `json:"UserId,omitempty"`   // User identifier
	Interval *int `json:"Interval,omitempty"` // Interval in seconds
}
