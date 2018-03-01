package models

type ErrorCode struct {
	ErrorCodeId *int64  `json:"ErrorCodeId"`
	Error       *string `json:"Error"`
	Message     *string `json:"Message"`
}
