package printer

import (
	"encoding/json"
	"fmt"

	"github.com/neko-neko/utmpdump/utmp"
)

// TSV printer
type JsonPrinter struct{}

// Format TSV
func (t *JsonPrinter) Print(u *utmp.GoUtmp) {
	json, _ := json.Marshal(u)
	fmt.Println(string(json))
}
