package models

type LimitDuration int

const (
	LD_24Hours    LimitDuration = 1
	LD_1Week      LimitDuration = 2
	LD_1Month     LimitDuration = 3
	LD_3Months    LimitDuration = 4
	LD_NoDuration LimitDuration = 5
	LD_Session    LimitDuration = 6
)
