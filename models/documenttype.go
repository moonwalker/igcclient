package models

type DocumentType int64

const (
	DT_Id                    DocumentType = 1
	DT_ProofOfAddress        DocumentType = 2
	DT_PaymentMethod         DocumentType = 3
	DT_KycApproval           DocumentType = 4
	DT_SOWQuestionnaire      DocumentType = 5
	DT_SOWSupportingDocument DocumentType = 6
	DT_SourceOfFunds         DocumentType = 7
	DT_SocialSecurityNumber  DocumentType = 8
)

func DocumentTypes() []DocumentType {
	return []DocumentType{
		DT_Id,
		DT_ProofOfAddress,
		DT_PaymentMethod,
		DT_KycApproval,
		DT_SOWQuestionnaire,
		DT_SOWSupportingDocument,
		DT_SourceOfFunds,
		DT_SocialSecurityNumber,
	}
}

var (
	documentTypeNames = [...]string{
		"ID",
		"Proof Of Address",
		"Payment Method",
		"Kyc Approval",
		"Source Of Wealth Questionnaire",
		"Source Of Wealth Supporting Document",
		"Source Of Funds",
		"Social Security Number",
	}
)

func (dt DocumentType) String() string {
	if dt < DT_Id || dt > DT_SocialSecurityNumber {
		return "Unknown"
	}
	return documentTypeNames[dt-1]
}
