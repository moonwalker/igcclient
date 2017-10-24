package models

type GameDetailsAffiliates struct {
	GameId      *int    `json:"GameId,omitempty"`      // The game identifier.
	Name        *string `json:"Name,omitempty"`        // The game name.
	Vendor      *string `json:"Vendor,omitempty"`      // The game vendor name.
	IsEnabled   *bool   `json:"IsEnabled,omitempty"`   // A value indicating whether this game is enabled.
	Description *string `json:"Description,omitempty"` // The game description.
	Devices     *string `json:"Devices,omitempty"`     // A list of compatible devices.
}
