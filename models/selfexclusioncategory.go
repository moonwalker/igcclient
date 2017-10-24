package models

type SelfExclusionCategory struct {
	ID                       *int  `json:"ID,omitempty"`                       // This represents a unique ID for the SelfExclusionCategory.
	PeriodDurationID         *int  `json:"PeriodDurationID,omitempty"`         // PeriodDurationID references a pre-defined time quantity specifying for how long the account will be closed. For example: 24 hours or 7 days or 30 days etc. References a time quantity. For example: 24 hours or 7 days or 30 days etc. These can be fetched from ResponsibleGaming/SelfExclusion/Durations.
	IsVisible                *bool `json:"IsVisible,omitempty"`                // Toggles self exclusion category visibility from the front-end.
	GracePeriodDurationID    *int  `json:"GracePeriodDurationID,omitempty"`    // Grace period duration represents duration at which the end user can cancel the self-exclusion. References a time quantity. For example: 24 hours or 7 days or 30 days etc. These can be fetched from ResponsibleGaming/SelfExclusion/Durations.
	GraceEffectiveDurationID *int  `json:"GraceEffectiveDurationID,omitempty"` // When the user cancels, this is the amount of time that has to elapse before the self-exlusion is released. References a time quantity. For example: 24 hours or 7 days or 30 days etc. These can be fetched from ResponsibleGaming/SelfExclusion/Durations.
	MaxUsageTime             *int  `json:"MaxUsageTime,omitempty"`             // Represents the count of this type of self-exclusion you can have at one time.
}
