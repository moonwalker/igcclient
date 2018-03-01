package models

type DBTokenIP struct {
	ForwardedFor  *string  `json:"ForwardedFor,omitempty"`
	RemoteIP      *string  `json:"RemoteIP,omitempty"`
	ClientIP      *string  `json:"ClientIP,omitempty"`
	Country       *string  `json:"Country,omitempty"`
	Token         *string  `json:"Token,omitempty"`
	DateCreated   *IGCTime `json:"DateCreated,omitempty"`
	LastAccess    *IGCTime `json:"LastAccess,omitempty"`
	IsInvalidated *bool    `json:"IsInvalidated,omitempty"`
	AdminUserID   *int64   `json:"AdminUserID,omitempty"`
	AdminUserName *string  `json:"AdminUserName,omitempty"`
}
