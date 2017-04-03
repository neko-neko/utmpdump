package filter

import (
	"time"

	"github.com/neko-neko/utmpdump/utmp"
)

type TimeFilter struct{}

// filter time
func (f *TimeFilter) Filter(utmps []*utmp.GoUtmp, since time.Time, until time.Time) []*utmp.GoUtmp {
	var (
		result      = []*utmp.GoUtmp{}
		sinceIsZero = since.IsZero()
		untilIsZero = until.IsZero()
	)

	if sinceIsZero && untilIsZero {
		return utmps
	}

	if !sinceIsZero {
		for _, u := range utmps {
			utime, _ := time.Parse(utmp.TimeFormat, u.Time)
			if !utime.Before(since) {
				result = append(result, u)
			}
		}
	}

	if !untilIsZero {
		for _, u := range utmps {
			utime, _ := time.Parse(utmp.TimeFormat, u.Time)
			if !utime.After(until) {
				result = append(result, u)
			}
		}
	}

	return result
}
