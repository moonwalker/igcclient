package models

type LimitDuration int64

const (
	LD_24Hours    LimitDuration = 1
	LD_1Week      LimitDuration = 2
	LD_1Month     LimitDuration = 3
	LD_3Months    LimitDuration = 4
	LD_NoDuration LimitDuration = 5
	LD_Session    LimitDuration = 6
)

func LimitDurations() []LimitDuration {
	return []LimitDuration{
		LD_24Hours,
		LD_1Week,
		LD_1Month,
		LD_3Months,
		LD_NoDuration,
		LD_Session,
	}
}

var (
	limitDurationNames = [...]string{
		"1 Day",
		"1 Week",
		"1 Month",
		"3 Months",
		"No Duration",
		"Session",
	}
)

func (ld LimitDuration) String() string {
	if ld < LD_24Hours || ld > LD_Session {
		return "Unknown"
	}
	return limitDurationNames[ld-1]
}
