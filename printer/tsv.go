package printer

import (
	"fmt"

	"github.com/neko-neko/utmpdump/utmp"
)

// TSV printer
type TsvPrinter struct{}

// Format TSV
func (t *TsvPrinter) Print(u *utmp.GoUtmp) {
	fmt.Printf(
		"%d\t%d\t%s\t%s\t%s\t%s\t%d\t%d\t%d\t%s\t%s\n",
		u.Type,
		u.Pid,
		u.Device,
		u.Id,
		u.User,
		u.Host,
		u.Exit.Exit,
		u.Exit.Termination,
		u.Session,
		u.Time,
		u.Addr,
	)
}
