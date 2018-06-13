package models

type ErrorCode struct {
	ErrorCodeID *int64  `json:"ErrorCodeId"`
	Error       *string `json:"Error"`
	Message     *string `json:"Message"`
}
