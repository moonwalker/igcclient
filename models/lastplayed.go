package models

type LastPlayed struct {
	ID      *int64   `json:"Id,omitempty"`
	GameID  *int64   `json:"GameId,omitempty"`
	Created *IGCTime `json:"Created,omitempty"`
}
