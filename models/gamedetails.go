package models

type GameDetails struct {
	GameID              *int64                          `json:"GameId,omitempty"`              // The game identifier.
	Name                *string                         `json:"Name,omitempty"`                // The game name.
	VendorID            *int64                          `json:"VendorId,omitempty"`            // The game vendor id.
	Vendor              *string                         `json:"Vendor,omitempty"`              // The game vendor name.
	GameType            *string                         `json:"GameType,omitempty"`            // The type of game
	IsEnabled           *bool                           `json:"IsEnabled,omitempty"`           // A value indicating whether this game is enabled.
	Description         *string                         `json:"Description,omitempty"`         // The game description.
	Devices             *string                         `json:"Devices,omitempty"`             // A list of compatible devices.
	IsLive              *bool                           `json:"IsLive,omitempty"`              // Is a live casino game
	RestrictedCountries *[]GameDetailsRestrictedCountry `json:"RestrictedCountries,omitempty"` // List of countries that this game is restricted in
	ParentGameID        *int64                          `json:"ParentGameId,omitempty"`        // If the game is a child, will contain the parent game's id
	ChildSortOrder      *int64                          `json:"ChildSortOrder,omitempty"`      // If the game is a child, this will contain its order position
	ChildGameValue      *string                         `json:"ChildGameValue,omitempty"`      // If the game is a child, this will contain a description to show its difference from the other children
	ChildGames          *[]GameDetails                  `json:"ChildGames,omitempty"`          // List of child games
	SuggestedWidth      *int64                          `json:"SuggestedWidth,omitempty"`      // Suggested game width
	SuggestedHeight     *int64                          `json:"SuggestedHeight,omitempty"`     // Suggested game height
}
