package models

type CheckUser struct {
	Username *string `json:"Username,omitempty"` // The Username
	Email    *string `json:"Email,omitempty"`    // The Email
}
