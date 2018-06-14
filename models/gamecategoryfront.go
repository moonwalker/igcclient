package models

type GameCategoryFront struct {
	GameID       *int64  `json:"GameId,omitempty"`
	CategoryID   *int64  `json:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty"`
	OrderNumber  *int64  `json:"OrderNumber,omitempty"`
}
