package models

type AllBonusTypes int

const (
	ABT_NoDeposit                    AllBonusTypes = 3
	ABT_Cashback                     AllBonusTypes = 4
	ABT_CasinoWelcome                AllBonusTypes = 5
	ABT_Reload                       AllBonusTypes = 6
	ABT_FreeSpins                    AllBonusTypes = 7
	ABT_SportsWelcome                AllBonusTypes = 10
	ABT_RiskFreeBetBonus             AllBonusTypes = 11
	ABT_SportsReload                 AllBonusTypes = 13
	ABT_PersonalizedRiskFreeBetBonus AllBonusTypes = 16
	ABT_SportsNoDeposit              AllBonusTypes = 18
	ABT_ManualCashback               AllBonusTypes = 19
	ABT_GenericFreeSpins             AllBonusTypes = 20
	ABT_SportsManualCashback         AllBonusTypes = 21
)
