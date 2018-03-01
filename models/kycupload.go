package models

type KYCUpload struct {
	Data      *[]byte `json:"data,omitempty"`      // The data for the KYC item.
	KYCID     *int64  `json:"KYCID,omitempty"`     // The KYC ID. This needs to exist in the database for you to upload the data to it
	Extension *string `json:"extension,omitempty"` // The file extension of the KYC item. Examples are: ".png", ".pdf", ".jpg", ".jpeg", ".bmp", ".gif"
}
