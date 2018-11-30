package models

type AddressResponseDTO struct {
	Address1    *string                 `json:"Address1,omitempty"`
	Address2    *string                 `json:"Address2,omitempty"`
	City        *string                 `json:"City,omitempty"`
	PostalCode  *string                 `json:"PostalCode,omitempty"`
	Subdivision *SubdivisionResponseDTO `json:"Subdivision,omitempty"`
}
