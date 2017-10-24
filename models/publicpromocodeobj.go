package models

type PublicPromoCodeObj struct {
	PromoId     *int               `json:"PromoId,omitempty"`
	BonusId     *int               `json:"BonusId,omitempty"`
	PromoCode   *string            `json:"PromoCode,omitempty"`
	UserId      *int               `json:"UserId,omitempty"`
	DateCreated *IGCTime           `json:"DateCreated,omitempty"`
	ExpiresOn   *IGCTime           `json:"ExpiresOn,omitempty"`
	Status      *PromoCodeStatuses `json:"Status,omitempty"`
}
