package models

type RecentWinnersV2RequestModel struct {
	MaximumWin *float64 `json:"MaximunWin,omitempty"` // (Optional) Filter by maximum win value
	Device     *int64   `json:"Device,omitempty"`     // (Optional) Filter by device. Default = Device From User Agent. 1 = Desktop, 2 = iPad, 3 = iPhone, 4 = Android Phone, 5 = Windows Mobile 7, 6 = Windows Mobile 8, 7 = iPod, 8 = AndroidTablet.
	Count      *int64   `json:"Count"`                // The number of results to return.
	MinimumWin *float64 `json:"MinimumWin"`           // Filter by minimum win value.
	GameID     *int64   `json:"GameId,omitempty"`     // (Optional) Filter by game.
}
