package models

type PublicPromoCodeObj struct {
	PromoID     *int64             `json:"PromoId,omitempty"`
	BonusID     *int64             `json:"BonusId,omitempty"`
	PromoCode   *string            `json:"PromoCode,omitempty"`
	UserID      *int64             `json:"UserId,omitempty"`
	DateCreated *IGCTime           `json:"DateCreated,omitempty"`
	ExpiresOn   *IGCTime           `json:"ExpiresOn,omitempty"`
	Status      *PromoCodeStatuses `json:"Status,omitempty"`
}
