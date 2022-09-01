package rfctime

import (
	"fmt"
	"strings"
	"time"
)

type RFC3339Time struct {
	time.Time
}

const ctLayout = time.RFC3339

func (ct *RFC3339Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *RFC3339Time) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *RFC3339Time) IsSet() bool {
	return ct.UnixNano() != nilTime
}
