package models

type DevicePlatformModel struct {
	Platform     *DevicePlatform `json:"Platform,omitempty"`
	Type         *DeviceType     `json:"Type,omitempty"`
	PlatformName *string         `json:"PlatformName,omitempty"`
	TypeName     *string         `json:"TypeName,omitempty"`
}
