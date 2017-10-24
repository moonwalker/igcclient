package models

type LastPlayed struct {
	Id      *int     `json:"Id,omitempty"`
	GameId  *int     `json:"GameId,omitempty"`
	Created *IGCTime `json:"Created,omitempty"`
}
