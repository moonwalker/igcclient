package models

type ForgotPasswordChangeBySMSModel struct {
	MobilePrefix *string `json:"MobilePrefix,omitempty"`
	Mobile       *string `json:"Mobile,omitempty"`
	RecoveryCode *string `json:"RecoveryCode,omitempty"`
	Password     *string `json:"Password,omitempty"`
}
