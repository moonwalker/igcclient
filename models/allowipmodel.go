package models

type AllowIPModel struct {
	IP        *string `json:"IP,omitempty"`        // The IP of the End User
	UserAgent *string `json:"UserAgent,omitempty"` // The User Agent of the end user
}
