package filter

import "github.com/neko-neko/utmpdump/utmp"

type Filter interface {
	Filter(utmps []*utmp.GoUtmp) []*utmp.GoUtmp
}
