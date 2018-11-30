package models

type AuthMethod int64

const (
	AuthMethodOneFactor        AuthMethod = 1
	AuthMethodTwoFactorBySms   AuthMethod = 2
	AuthMethodTwoFactorByEmail AuthMethod = 4
	AuthMethodTwoFactorByKba   AuthMethod = 8
)
