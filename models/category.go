package models

type Category struct {
	ID           *int64             `json:"ID,omitempty"`
	DisplayName  *string            `json:"DisplayName,omitempty"`
	OrderNumber  *int64             `json:"OrderNumber,omitempty"`
	IsSelected   *bool              `json:"IsSelected,omitempty"`
	Visible      *bool              `json:"Visible,omitempty"`
	LanguageID   *int64             `json:"LanguageID,omitempty"`
	Tag          *string            `json:"Tag,omitempty"`
	LanguageCode *string            `json:"LanguageCode,omitempty"`
	SelectedID   *int64             `json:"SelectedID,omitempty"`
	ImageURL     *string            `json:"ImageURL,omitempty"`
	VendorID     *int64             `json:"VendorID,omitempty"`
	VendorName   *string            `json:"VendorName,omitempty"`
	ParentId     *int64             `json:"ParentId,omitempty"`
	Devices      *map[string]string `json:"Devices,omitempty"`
}
