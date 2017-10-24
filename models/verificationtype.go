package models

type VerificationType struct {
	ID        *int    `json:"ID,omitempty"`
	Type      *string `json:"Type,omitempty"`
	Notes     *string `json:"Notes,omitempty"`
	IsLive    *bool   `json:"IsLive,omitempty"`
	UsesSMS   *bool   `json:"UsesSMS,omitempty"`
	UsesEmail *bool   `json:"UsesEmail,omitempty"`
}
