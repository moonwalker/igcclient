package models

type UserWithNoActivity struct {
	UserId           *int     `json:"UserId,omitempty"`
	LastActivityDate *IGCTime `json:"LastActivityDate,omitempty"`
}
