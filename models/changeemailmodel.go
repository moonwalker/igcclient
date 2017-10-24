package models

type ChangeEmailModel struct {
	Password *string `json:"Password,omitempty"`
	Email    *string `json:"Email,omitempty"`
}
