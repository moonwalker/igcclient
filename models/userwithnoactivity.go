package models

type UserWithNoActivity struct {
	UserId           *int64   `json:"UserId,omitempty"`
	LastActivityDate *IGCTime `json:"LastActivityDate,omitempty"`
}
