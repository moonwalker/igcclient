package models

type Jackpot struct {
	GameId    int64  `json:"GameId,omitempty"`    // The game identifier.
	GameName  string `json:"GameName,omitempty"`  // The name of the game.
	Vendor    string `json:"Vendor,omitempty"`    // The Game vendor.
	Currency  string `json:"Currency,omitempty"`  // The currency. (ISO Alpha 3)
	Amount    string `json:"Amount,omitempty"`    // The amount.
	JackpotId string `json:"JackpotId,omitempty"` // The jackpot identifier.
}
