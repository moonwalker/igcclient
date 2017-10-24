package models

type VerifyUserBySMSModel struct {
	MobileVerificationCode *string `json:"MobileVerificationCode,omitempty"`
	MobilePrefix           *string `json:"MobilePrefix,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty"`
}
