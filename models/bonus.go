package models

type Bonus struct {
	AvailableDepositBonus *[]FirstDepositBonus `json:"AvailableDepositBonus,omitempty"`
	BlockingDepositBonus  *[]FirstDepositBonus `json:"BlockingDepositBonus,omitempty"`
	UsedDepositBonus      *[]FirstDepositBonus `json:"UsedDepositBonus,omitempty"`
	AvailableFreeBetBonus *[]FreeBetBonus      `json:"AvailableFreeBetBonus,omitempty"`
	BlockingFreeBetBonus  *[]FreeBetBonus      `json:"BlockingFreeBetBonus,omitempty"`
	UsedFreeBetBonus      *[]FreeBetBonus      `json:"UsedFreeBetBonus,omitempty"`
	Error                 *bool                `json:"Error,omitempty"`
}
