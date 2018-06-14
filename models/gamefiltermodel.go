package models

type GameFilterModel struct {
	Language     *string `json:"Language"`               // Filter by language.
	Category     *string `json:"Category,omitempty"`     // (Optional) Filter by game category tag. Default = All Categories.
	Sort         *int64  `json:"Sort,omitempty"`         // (Optional) Result sort order. Default = Alphabetic. 0 = Alphabetic, 1 = Popularity, 2 = Vendor
	Device       *int64  `json:"Device,omitempty"`       // (Optional) Filter by device. Default = Device From User Agent. 1 = Desktop, 2 = iPad, 3 = iPhone, 4 = Android Phone, 5 = Windows Mobile 7, 6 = Windows Mobile 8, 7 = iPod, 8 = AndroidTablet.
	Country      *string `json:"Country,omitempty"`      // (Optional) Filter by country. Default = All Countries.
	SearchString *string `json:"SearchString,omitempty"` // (Optional) Game name search string. Default = No Search Filter.
	IsMinigame   *bool   `json:"IsMinigame,omitempty"`   // (Optional) Gets or sets a value indicating whether this instance is minigame. Default = False
	CurrencyCode *string `json:"CurrencyCode,omitempty"` // (Optional) Filter by currency. Default = Operator Base Currency
}
