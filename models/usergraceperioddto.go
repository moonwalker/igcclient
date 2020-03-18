package models

type UserGracePeriodDTO struct {
	DocumentTypeID *KycRequestType `json:"DocumentTypeId,omitempty"`
	DocumentType   *string         `json:"DocumentType,omitempty"`
	StartDate      *IGCTime        `json:"StartDate,omitempty"`
	EndDate        *IGCTime        `json:"EndDate,omitempty"`
}
