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
	LT_NetLoss         LimitType = 9
	LT_Wagering        LimitType = 10
	LT_NetDeposit      LimitType = 11
)

func LimitTypes() []LimitType {
	return []LimitType{
		LT_Deposit,
		LT_CasinoWagering,
		LT_CasinoNetLoss,
		LT_SportsWagering,
		LT_SportsNetLoss,
		LT_Session,
		LT_StakePerSession,
		LT_MaxStakePerBet,
		LT_NetLoss,
		LT_Wagering,
		LT_NetDeposit,
	}
}

var (
	limitTypeNames = [...]string{
		"Deposit",
		"Casino Wagering",
		"Casino Net Loss",
		"Sports Wagering",
		"Sports Net Loss",
		"Session",
		"Stake Per Session",
		"Max Stake Per Bet",
		"Net Loss",
		"Wagering",
		"Net Deposit",
	}
)

func (lt LimitType) String() string {
	if lt < LT_Deposit || lt > LT_NetDeposit {
		return "Unknown"
	}
	return limitTypeNames[lt-1]
}
