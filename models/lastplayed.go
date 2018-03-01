package models

type LastPlayed struct {
	Id      *int64   `json:"Id,omitempty"`
	GameId  *int64   `json:"GameId,omitempty"`
	Created *IGCTime `json:"Created,omitempty"`
}
