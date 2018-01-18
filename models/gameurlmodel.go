package models

type GameURLModel struct {
	Language   *string `json:"Language,omitempty"`
	UserIP     *string `json:"UserIP,omitempty"`
	UserAgent  *string `json:"UserAgent,omitempty"`
	PlayForFun *bool   `json:"PlayForFun,omitempty"`
	VariantID  *int64  `json:"VariantID,omitempty"`
}
