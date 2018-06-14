package models

type Vendor struct {
	ID         *int64      `json:"Id,omitempty"`         // The vendor identifier.
	Name       *string     `json:"Name,omitempty"`       // The vendor name.
	TypeName   *string     `json:"TypeName,omitempty"`   // The vendor type name.
	VendorType *VendorType `json:"VendorType,omitempty"` // The vendor type value.
}
