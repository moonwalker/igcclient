package models

type RealityCheckUserSetting struct {
	UserID   *int64 `json:"UserId,omitempty"`   // User identifier
	Interval *int64 `json:"Interval,omitempty"` // Interval in seconds
}
