package models

type LimitType int64

const (
	LT_Deposit         LimitType = 1
	LT_CasinoWagering  LimitType = 2
	LT_CasinoNetLoss   LimitType = 3
	LT_SportsWagering  LimitType = 4
	LT_SportsNetLoss   LimitType = 5
	LT_Session         LimitType = 6
	LT_StakePerSession LimitType = 7
	LT_MaxStakePerBet  LimitType = 8
)
