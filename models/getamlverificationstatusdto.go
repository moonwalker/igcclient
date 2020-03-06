package models

// AML Verification Status response:
// 0 - AML verification not in progress, user is not restricted.
// 1 - AML verification in progress, user is restricted by pending Source of Wealth Questionnaire request.
// 2 - AML verification in progress, user is restricted by pending Source of Wealth Supporting Documents request.
// 3 - AML verification in progress, user is restricted by pending Source of Funds request.
type AmlVerificationStatus int64

const (
	AVS_NotInProgress                AmlVerificationStatus = 0
	AVS_SourceOfWealthQuestionnaire  AmlVerificationStatus = 1
	AVS_SourceOfWealthSupportingDocs AmlVerificationStatus = 2
	AVS_SourceOfFunds                AmlVerificationStatus = 3
)

func AmlVerificationStatuss() []AmlVerificationStatus {
	return []AmlVerificationStatus{
		AVS_NotInProgress,
		AVS_SourceOfWealthQuestionnaire,
		AVS_SourceOfWealthSupportingDocs,
		AVS_SourceOfFunds,
	}
}

var (
	AmlVerificationStatusNames = [...]string{
		"Not In Progress",
		"Source Of Wealth Questionnaire",
		"Source Of Wealth Supporting Docs",
		"Source Of Funds",
	}
)

func (avs AmlVerificationStatus) String() string {
	if avs < AVS_NotInProgress || avs > AVS_SourceOfFunds {
		return "Unknown"
	}
	return AmlVerificationStatusNames[avs]
}

type GetAmlVerificationStatusDTO struct {
	AmlVerificationStatus *AmlVerificationStatus `json:"amlVerificationStatus"`
}
