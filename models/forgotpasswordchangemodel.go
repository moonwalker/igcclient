package models

type ForgotPasswordChangeModel struct {
	RecoveryCode *string `json:"RecoveryCode,omitempty"`
	Password     *string `json:"Password,omitempty"`
}
