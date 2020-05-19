package goixc

import (
	"bytes"
	"time"
)

type DateYMD struct{ time.Time }

func (d *DateYMD) UnmarshalJSON(b []byte) (err error) {
	s := string(bytes.TrimSpace(bytes.Trim(b, `"`)))
	d.Time, err = time.Parse("2006-01-02", s)
	return
}
