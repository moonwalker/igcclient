package models

type GameJackpot struct {
	GameID         *int64   `json:"GameId,omitempty"`
	GameName       *string  `json:"GameName,omitempty"`
	Vendor         *string  `json:"Vendor,omitempty"`
	CurrencyFormat *string  `json:"CurrencyFormat,omitempty"`
	CurrencyCode   *string  `json:"CurrencyCode,omitempty"`
	Amount         *float64 `json:"Amount,omitempty"`
	JackpotID      *string  `json:"JackpotId,omitempty"`
}
