package models

type PasswordRequestDto struct {
	NewPassword *string `json:"newPassword,omitempty"`
}
