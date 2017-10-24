package models

type GameCategoryFront struct {
	GameId       *int    `json:"GameId,omitempty"`
	CategoryId   *int    `json:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty"`
	OrderNumber  *int    `json:"OrderNumber,omitempty"`
}
