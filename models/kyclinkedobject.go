package models

type KYCLinkedObject struct {
	KYCID            *int64   `json:"KycId,omitempty"`
	UserID           *int64   `json:"UserId,omitempty"`
	TypeID           *int64   `json:"TypeId,omitempty"`
	Path             *string  `json:"Path,omitempty"` // Filename of the KYC object
	Status           *int64   `json:"Status,omitempty"`
	ReasonText       *string  `json:"ReasonText,omitempty"`
	StatusMessage    *string  `json:"StatusMessage,omitempty"`
	Expires          *string  `json:"Expires,omitempty"`
	DateUploaded     *IGCTime `json:"DateUploaded,omitempty"`
	DateVerified     *IGCTime `json:"DateVerified,omitempty"`
	LastModifiedBy   *int64   `json:"LastModifiedBy,omitempty"`
	DateLastModified *IGCTime `json:"DateLastModified,omitempty"`
	DateCreated      *IGCTime `json:"DateCreated,omitempty"`
	KycTypeName      *string  `json:"KycTypeName,omitempty"`
	HelpText         *string  `json:"HelpText,omitempty"`
	ShowToUser       *bool    `json:"ShowToUser,omitempty"`
	NeedUserAction   *bool    `json:"NeedUserAction,omitempty"`
	StatusText       *string  `json:"StatusText,omitempty"`
}

type KycDocumentStatus int64

const (
	KDS_NotSet         KycDocumentStatus = 0
	KDS_Approved       KycDocumentStatus = 1
	KDS_Requested      KycDocumentStatus = 2
	KDS_Rejected       KycDocumentStatus = 3
	KDS_Processing     KycDocumentStatus = 4
	KDS_UserRequested  KycDocumentStatus = 5
	KDS_PendingRequest KycDocumentStatus = 6
)

func KYCDocumentStatuses() []KycDocumentStatus {
	return []KycDocumentStatus{
		KDS_NotSet,
		KDS_Approved,
		KDS_Requested,
		KDS_Rejected,
		KDS_Processing,
		KDS_UserRequested,
		KDS_PendingRequest,
	}
}

var (
	kycDocumentStatusNames = [...]string{
		"Not Set",
		"Approved",
		"Requested",
		"Rejected",
		"Processing",
		"UserRequested",
		"PendingRequest",
	}
)

func (kds KycDocumentStatus) String() string {
	if kds < KDS_NotSet || kds > KDS_PendingRequest {
		return "Unknown"
	}
	return kycDocumentStatusNames[kds]
}
