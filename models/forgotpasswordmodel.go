package models

type ForgotPasswordModel struct {
	SecurityAnswer *string `json:"SecurityAnswer,omitempty"` // Security answer to the security question. Ommiting the security answer parameter to avoid security question check.
	Email          *string `json:"Email"`
}
