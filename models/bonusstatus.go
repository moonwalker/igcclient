package models

type BonusStatus int64

const (
	BS_Available BonusStatus = 1 // The bonus which is currently available to player
	BS_Used      BonusStatus = 2 // The bonuses which had been used by player
	BS_Blocking  BonusStatus = 3 // The bonus which has been used by player, but, for it’s turnover has never been reached currently, it block player from withdrawing his money out and from using other available bonuses
)
