package models

type DeviceTypeObject struct {
	DeviceTypeID *int64  `json:"DeviceTypeId,omitempty"`
	Name         *string `json:"Name,omitempty"`
	Referral     *string `json:"Referral,omitempty"`
}
