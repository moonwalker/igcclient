package models

type FrontAllUserBonuses struct {
	UserBonuses      *[]FrontEndUserBonusObject `json:"UserBonuses,omitempty"`
	ClaimedFreeSpins *[]FrontUserFreeSpins      `json:"ClaimedFreeSpins,omitempty"`
}
