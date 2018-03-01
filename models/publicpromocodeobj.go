package models

type PublicPromoCodeObj struct {
	PromoId     *int64             `json:"PromoId,omitempty"`
	BonusId     *int64             `json:"BonusId,omitempty"`
	PromoCode   *string            `json:"PromoCode,omitempty"`
	UserId      *int64             `json:"UserId,omitempty"`
	DateCreated *IGCTime           `json:"DateCreated,omitempty"`
	ExpiresOn   *IGCTime           `json:"ExpiresOn,omitempty"`
	Status      *PromoCodeStatuses `json:"Status,omitempty"`
}
