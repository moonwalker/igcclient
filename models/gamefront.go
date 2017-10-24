package models

type GameFront struct {
	GameId               *int                 `json:"GameId,omitempty"`
	VendorId             *int                 `json:"VendorId,omitempty"`
	VendorGameId         *string              `json:"VendorGameId,omitempty"`
	VendorGameIdAlterate *string              `json:"VendorGameIdAlterate,omitempty"`
	Name                 *string              `json:"Name,omitempty"`
	ShortDescription     *string              `json:"ShortDescription,omitempty"`
	Slug                 *string              `json:"Slug,omitempty"`
	SuggestedHeight      *int                 `json:"SuggestedHeight,omitempty"`
	SuggestedWidth       *int                 `json:"SuggestedWidth,omitempty"`
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
	ParentGameId         *int                 `json:"ParentGameId,omitempty"`
	ChildSortOrder       *int                 `json:"ChildSortOrder,omitempty"`
	ChildGames           *[]GameFront         `json:"ChildGames,omitempty"`
	ChildGameValue       *string              `json:"ChildGameValue,omitempty"`
	InMaintenance        *bool                `json:"InMaintenance,omitempty"`
	IsOffline            *bool                `json:"IsOffline,omitempty"`
	Categories           *[]GameCategoryFront `json:"Categories,omitempty"`
	Variants             *[]GameFrontVariant  `json:"Variants,omitempty"`
	Jackpot              *GameJackpot         `json:"Jackpot,omitempty"`
	OrderNumber          *int                 `json:"OrderNumber,omitempty"`
	AdminCategoryId      *int                 `json:"AdminCategoryId,omitempty"`
	IsLiveCasinoLobby    *bool                `json:"IsLiveCasinoLobby,omitempty"`
}
