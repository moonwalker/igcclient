package models

type UserWithNoActivity struct {
	UserID           *int64   `json:"UserId,omitempty"`
	LastActivityDate *IGCTime `json:"LastActivityDate,omitempty"`
}
