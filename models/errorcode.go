package models

type ErrorCode struct {
	ErrorCodeId *int    `json:"ErrorCodeId"`
	Error       *string `json:"Error"`
	Message     *string `json:"Message"`
}
