package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DTime struct {
	Layout string
	Zone   string
}

func NewTime() *DTime {
	return &DTime{
		Layout: "2006-01-02 15:04:05",
		Zone:   "-07:00",
	}
}

func (t *DTime) UnixToStr(n int64, layouts ...string) string {
	layout := t.Layout
	if len(layouts) > 0 && layouts[0] != "" {
		layout = layouts[0]
	}

	var ss, ns string
	s := strconv.FormatInt(n, 10)
	l := len(s)
	if l > 10 {
		ss = s[:10]
		ns = s[10:]
		//000000000
		fillLen := 9 - len(ns)
		ns = ns + strings.Repeat("0", fillLen)
	} else {
		ss = s[:10]
	}

	si, _ := strconv.ParseInt(ss, 10, 64)
	ni, _ := strconv.ParseInt(ns, 10, 64)
	tm := time.Unix(si, ni)

	return tm.Format(layout)
}

func (t *DTime) StrToUnix(s string) int64 {
	ti, _ := time.Parse(t.Layout+" "+t.Zone, s)
	return ti.Unix()
}

func (t *DTime) GetTimeToStr(tm int64, fm ...string) string {
	ti := time.Unix(tm, 0)
	var format string

	if len(fm) < 1 {
		format = "%04d-%02d-%02d %02d:%02d:%02d"
	} else {
		format = fm[0]
	}

	return fmt.Sprintf(format, ti.Year(), ti.Month(), ti.Day(), ti.Hour(), ti.Minute(), ti.Second())
}

func (t *DTime) GetStrToTime(s string) int64 {
	s += " " + t.Layout
	ti, _ := time.Parse(t.Layout, s)
	return ti.Unix()
}
