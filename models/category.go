package models

type Category struct {
	ID           *int               `json:"ID,omitempty"`
	DisplayName  *string            `json:"DisplayName,omitempty"`
	OrderNumber  *int               `json:"OrderNumber,omitempty"`
	IsSelected   *bool              `json:"IsSelected,omitempty"`
	Visible      *bool              `json:"Visible,omitempty"`
	LanguageID   *int               `json:"LanguageID,omitempty"`
	Tag          *string            `json:"Tag,omitempty"`
	LanguageCode *string            `json:"LanguageCode,omitempty"`
	SelectedID   *int               `json:"SelectedID,omitempty"`
	ImageURL     *string            `json:"ImageURL,omitempty"`
	VendorID     *int               `json:"VendorID,omitempty"`
	VendorName   *string            `json:"VendorName,omitempty"`
	ParentId     *int               `json:"ParentId,omitempty"`
	Devices      *map[string]string `json:"Devices,omitempty"`
}
