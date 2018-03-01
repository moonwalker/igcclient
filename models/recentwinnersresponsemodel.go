package models

type RecentWinnersResponseModel struct {
	WinnerName  *string  `json:"WinnerName,omitempty"`  // First name and first letter of last name of user.
	CountryCode *string  `json:"CountryCode,omitempty"` // User's registered country code.
	UserName    *string  `json:"UserName,omitempty"`    // The username.
	Currency    *string  `json:"Currency,omitempty"`    // The currency of the win.
	GameID      *int64   `json:"GameID,omitempty"`      // The game identifier.
	GameName    *string  `json:"GameName,omitempty"`    // The name of the game.
	Win         *float64 `json:"Win,omitempty"`         // The amount of the win.
	BaseWin     *float64 `json:"BaseWin,omitempty"`     // The base currency amount of the win.
	GameSlug    *string  `json:"GameSlug,omitempty"`    // The game slug.
	Date        *IGCTime `json:"Date,omitempty"`        // The date / time of the win.
	FirstName   *string  `json:"FirstName,omitempty"`   // The user's first name
	LastName    *string  `json:"LastName,omitempty"`    // The user's last name
}
