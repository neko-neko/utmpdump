package printer

import (
	"fmt"

	"github.com/neko-neko/utmpdump/utmp"
)

// TSV printer
type CsvPrinter struct{}

// Format TSV
func (t *CsvPrinter) Print(u *utmp.GoUtmp) {
	fmt.Printf(
		"%d,%d,%s,%s,%s,%s,%d,%d,%d,%s,%s\n",
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
