package models

type DeviceTypeObject struct {
	DeviceTypeId *int    `json:"DeviceTypeId,omitempty"`
	Name         *string `json:"Name,omitempty"`
	Referral     *string `json:"Referral,omitempty"`
}
