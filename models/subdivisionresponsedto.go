package models

type SubdivisionResponseDTO struct {
	SubdivisionId       *int64  `json:"SubdivisionId,omitempty"`
	IsoCodeSubdivision1 *string `json:"IsoCodeSubdivision1,omitempty"`
	IsoNameSubdivision1 *string `json:"IsoNameSubdivision1,omitempty"`
	IsoCodeSubdivision2 *string `json:"IsoCodeSubdivision2,omitempty"`
	IsoNameSubdivision2 *string `json:"IsoNameSubdivision2,omitempty"`
}
