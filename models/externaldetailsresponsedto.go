package models

type ExternalDetailsResponseDTO struct {
	PlayerID     *string `json:"PlayerId,omitempty"`
	Status       *string `json:"Status,omitempty"`
	GovernmentID *string `json:"GovernmentId,omitempty"`
	Source       *string `json:"Source,omitempty"`
}
