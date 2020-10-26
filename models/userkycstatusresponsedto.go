package models

type UserKycStatusResponseDTO struct {
	ApprovalTypeID *KycApprovalType   `json:"ApprovalTypeId,omitempty"`
	StatusID       *KycApprovalStatus `json:"StatusId,omitempty"`
}

type KycApprovalType int64

const (
	KAT_PhotoId                      KycApprovalType = 1
	KAT_ProofOfAddress               KycApprovalType = 2
	KAT_PaymentMethod                KycApprovalType = 3
	KAT_SourceOfWealthQuestionnaire  KycApprovalType = 4
	KAT_SourceOfWealthSupportingDocs KycApprovalType = 5
	KAT_SourceOfFunds                KycApprovalType = 6
	KAT_SSN                          KycApprovalType = 7
)

func KYCApprovalTypes() []KycApprovalType {
	return []KycApprovalType{
		KAT_PhotoId,
		KAT_ProofOfAddress,
		KAT_PaymentMethod,
		KAT_SourceOfWealthQuestionnaire,
		KAT_SourceOfWealthSupportingDocs,
		KAT_SourceOfFunds,
		KAT_SSN,
	}
}

var (
	kycApprovalTypeNames = [...]string{
		"Photo Id",
		"ProofOfAddress",
		"Payment Method",
		"Source Of Wealth Questionnaire",
		"Source Of Wealth Supporting Docs",
		"Source Of Funds",
		"SSN",
	}
)

func (kat KycApprovalType) String() string {
	if kat < KAT_PhotoId || kat > KAT_SSN {
		return "Unknown"
	}
	return kycApprovalTypeNames[kat-1]
}

type KycApprovalStatus int64

const (
	KAS_NotSet        KycApprovalStatus = 1
	KAS_Pending       KycApprovalStatus = 2
	KAS_ManualPending KycApprovalStatus = 3
	KAS_Approved      KycApprovalStatus = 4
	KAS_Rejected      KycApprovalStatus = 5
)

func KYCApprovalStatuses() []KycApprovalStatus {
	return []KycApprovalStatus{
		KAS_NotSet,
		KAS_Pending,
		KAS_ManualPending,
		KAS_Approved,
		KAS_Rejected,
	}
}

var (
	kycApprovalStatusNames = [...]string{
		"Not Set",
		"Pending",
		"Manual Pending",
		"Approved",
		"Rejected",
	}
)

func (kas KycApprovalStatus) String() string {
	if kas < KAS_NotSet || kas > KAS_Rejected {
		return "Unknown"
	}
	return kycApprovalStatusNames[kas-1]
}
