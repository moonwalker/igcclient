package models

type KYCLinkedObject struct {
	KycId            *int     `json:"KycId,omitempty"`
	UserId           *int     `json:"UserId,omitempty"`
	TypeId           *int     `json:"TypeId,omitempty"`
	Path             *string  `json:"Path,omitempty"` // Filename of the KYC object
	Status           *int     `json:"Status,omitempty"`
	ReasonText       *string  `json:"ReasonText,omitempty"`
	StatusMessage    *string  `json:"StatusMessage,omitempty"`
	Expires          *string  `json:"Expires,omitempty"`
	DateUploaded     *IGCTime `json:"DateUploaded,omitempty"`
	DateVerified     *IGCTime `json:"DateVerified,omitempty"`
	LastModifiedBy   *int     `json:"LastModifiedBy,omitempty"`
	DateLastModified *IGCTime `json:"DateLastModified,omitempty"`
	DateCreated      *IGCTime `json:"DateCreated,omitempty"`
	KycTypeName      *string  `json:"KycTypeName,omitempty"`
	HelpText         *string  `json:"HelpText,omitempty"`
	ShowToUser       *bool    `json:"ShowToUser,omitempty"`
	NeedUserAction   *bool    `json:"NeedUserAction,omitempty"`
	StatusText       *string  `json:"StatusText,omitempty"`
}
