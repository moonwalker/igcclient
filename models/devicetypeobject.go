package models

type DeviceTypeObject struct {
	DeviceTypeId *int64  `json:"DeviceTypeId,omitempty"`
	Name         *string `json:"Name,omitempty"`
	Referral     *string `json:"Referral,omitempty"`
}
