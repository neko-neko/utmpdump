package filter

import (
	"time"

	"github.com/neko-neko/utmpdump/utmp"
)

type TimeFilter struct {
	Since time.Time
	Until time.Time
}

// filter time
func (f *TimeFilter) Filter(utmps []*utmp.GoUtmp) []*utmp.GoUtmp {
	var (
		result      = []*utmp.GoUtmp{}
		sinceIsZero = f.Since.IsZero()
		untilIsZero = f.Until.IsZero()
	)

	if sinceIsZero && untilIsZero {
		return utmps
	}

	if !sinceIsZero {
		for _, u := range utmps {
			utime, _ := time.Parse(utmp.TimeFormat, u.Time)
			if !utime.Before(f.Since) {
				result = append(result, u)
			}
		}
	}

	if !untilIsZero {
		for _, u := range utmps {
			utime, _ := time.Parse(utmp.TimeFormat, u.Time)
			if !utime.After(f.Until) {
				result = append(result, u)
			}
		}
	}

	return result
}
