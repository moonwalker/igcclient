package models

type ChangePasswordModel struct {
	OldPassword *string `json:"OldPassword,omitempty"`
	NewPassword *string `json:"NewPassword,omitempty"`
}
