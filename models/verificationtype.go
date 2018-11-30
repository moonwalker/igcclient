package models

type VerificationType int64

const (
	VerificationTypeSms         VerificationType = 1
	VerificationTypeEmail       VerificationType = 2
	VerificationTypeManual      VerificationType = 3
	VerificationTypeSmsAndEmail VerificationType = 4
)
