package printer

import (
	"github.com/neko-neko/utmpdump/utmp"
)

const (
	PrintTypeJson = "json"
	PrintTypeCsv  = "csv"
	PrintTypeTsv  = "tsv"
)

var PrintTypes = map[string]struct{}{
	PrintTypeJson: struct{}{},
	PrintTypeCsv:  struct{}{},
	PrintTypeTsv:  struct{}{},
}

// Printer interface
type Printer interface {
	Print(utmp *utmp.GoUtmp)
}

// Get formatter
func GetPrinter(formatType string) Printer {
	switch formatType {
	case PrintTypeJson:
		return &JsonPrinter{}
	case PrintTypeCsv:
		return &CsvPrinter{}
	case PrintTypeTsv:
		return &TsvPrinter{}
	default:
		return &JsonPrinter{}
	}
}
