package models

type GameFront struct {
	GameID               *int64               `json:"GameId,omitempty"`
	VendorID             *int64               `json:"VendorId,omitempty"`
	VendorGameID         *string              `json:"VendorGameId,omitempty"`
	VendorGameIDAlterate *string              `json:"VendorGameIdAlterate,omitempty"`
	Name                 *string              `json:"Name,omitempty"`
	ShortDescription     *string              `json:"ShortDescription,omitempty"`
	Slug                 *string              `json:"Slug,omitempty"`
	SuggestedHeight      *int64               `json:"SuggestedHeight,omitempty"`
	SuggestedWidth       *int64               `json:"SuggestedWidth,omitempty"`
	IsEnabled            *bool                `json:"IsEnabled,omitempty"`
	IsNew                *bool                `json:"IsNew,omitempty"`
	IsLive               *bool                `json:"IsLive,omitempty"`
	IsHot                *bool                `json:"IsHot,omitempty"`
	IsRecommended        *bool                `json:"IsRecommended,omitempty"`
	Description          *string              `json:"Description,omitempty"`
	AllowRealPlay        *bool                `json:"AllowRealPlay,omitempty"`
	AllowDemoPlay        *bool                `json:"AllowDemoPlay,omitempty"`
	Category             *string              `json:"Category,omitempty"`
	GameType             *string              `json:"GameType,omitempty"`
	Vendor               *string              `json:"Vendor,omitempty"`
	IsMinigame           *bool                `json:"IsMinigame,omitempty"`
	IsParent             *bool                `json:"IsParent,omitempty"`
	ParentGameID         *int64               `json:"ParentGameId,omitempty"`
	ChildSortOrder       *int64               `json:"ChildSortOrder,omitempty"`
	ChildGames           *[]GameFront         `json:"ChildGames,omitempty"`
	ChildGameValue       *string              `json:"ChildGameValue,omitempty"`
	InMaintenance        *bool                `json:"InMaintenance,omitempty"`
	IsOffline            *bool                `json:"IsOffline,omitempty"`
	Categories           *[]GameCategoryFront `json:"Categories,omitempty"`
	Variants             *[]GameFrontVariant  `json:"Variants,omitempty"`
	Jackpot              *GameJackpot         `json:"Jackpot,omitempty"`
	OrderNumber          *int64               `json:"OrderNumber,omitempty"`
	AdminCategoryID      *int64               `json:"AdminCategoryId,omitempty"`
	IsLiveCasinoLobby    *bool                `json:"IsLiveCasinoLobby,omitempty"`
}
