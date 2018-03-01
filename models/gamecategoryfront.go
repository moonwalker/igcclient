package models

type GameCategoryFront struct {
	GameId       *int64  `json:"GameId,omitempty"`
	CategoryId   *int64  `json:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty"`
	OrderNumber  *int64  `json:"OrderNumber,omitempty"`
}
