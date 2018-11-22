package models

type CloseAccountDOBRequestDto struct {
	DOB             SplitDateOfBirth `json:"DOB"`
	UserBlockReason string           `json:"UserBlockReason"`
}
