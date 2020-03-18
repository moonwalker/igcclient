package models

type KycRequestType int64

const (
	KT_PhotoId                      KycRequestType = 1
	KT_ProofOfAddress               KycRequestType = 2
	KT_PaymentMethod                KycRequestType = 3
	KT_KYCApproval                  KycRequestType = 4
	KT_SourceOfWealthQuestionnaire  KycRequestType = 5
	KT_SourceOfWealthSupportingDocs KycRequestType = 6
	KT_SourceOfFunds                KycRequestType = 7
	KT_SSN                          KycRequestType = 8
)

func KycRequestTypes() []KycRequestType {
	return []KycRequestType{
		KT_PhotoId,
		KT_ProofOfAddress,
		KT_PaymentMethod,
		KT_KYCApproval,
		KT_SourceOfWealthQuestionnaire,
		KT_SourceOfWealthSupportingDocs,
		KT_SourceOfFunds,
		KT_SSN,
	}
}

var (
	kycRequestTypeNames = [...]string{
		"Photo Id",
		"ProofOfAddress",
		"Payment Method",
		"KYC Approval",
		"Source Of Wealth Questionnaire",
		"Source Of Wealth Supporting Docs",
		"Source Of Funds",
		"SSN",
	}
)

func (kt KycRequestType) String() string {
	if kt < KT_PhotoId || kt > KT_SSN {
		return "Unknown"
	}
	return kycRequestTypeNames[kt-1]
}
