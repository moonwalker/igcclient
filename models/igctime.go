package models

import (
	"fmt"
	"strings"
	"time"
)

type IGCTime struct {
	time.Time
}

const igcTimeFormat = "2006-01-02T15:04:05.999" // RFC3339Nano without offset
var nilTime = (time.Time{}).UnixNano()

func (t *IGCTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\n\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(igcTimeFormat, s)
	return
}

func (t *IGCTime) MarshalJSON() ([]byte, error) {
	if t == nil || t.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(igcTimeFormat))), nil
}
