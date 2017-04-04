package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/neko-neko/utmpdump/filter"
	"github.com/neko-neko/utmpdump/printer"
	"github.com/neko-neko/utmpdump/utmp"
)

func main() {
	var (
		filePath   string
		until      string
		since      string
		outputType string
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s: %s [options]
-f, --file <file> load a specific file instead of utmp
-t, --until <YYYYMMDDHHMMSS> display the lines until the specified time
-s, --since <YYYYMMDDHHMMSS> display the lines since the specified time
-o, --output <json/tsv/csv> display the lines at specific format. Default format is json.
`, os.Args[0], os.Args[0])
	}

	// parse arguments
	flag.StringVar(&filePath, "f", "", "file path")
	flag.StringVar(&filePath, "file", "", "file path")
	flag.StringVar(&until, "t", "", "until")
	flag.StringVar(&until, "until", "", "until")
	flag.StringVar(&since, "s", "", "since")
	flag.StringVar(&since, "since", "", "since")
	flag.StringVar(&outputType, "o", "", "output type")
	flag.StringVar(&outputType, "output", "", "output type")
	flag.Parse()
	if filePath == "" {
		flag.Usage()
		os.Exit(1)
	}
	outputType = strings.ToLower(outputType)
	if _, exist := printer.PrintTypes[outputType]; outputType != "" && !exist {
		flag.Usage()
		os.Exit(1)
	}
	sinceTime, _ := time.Parse("20060102150405", since)
	untilTime, _ := time.Parse("20060102150405", until)

	// read file
	file, openErr := os.Open(filePath)
	defer file.Close()
	if openErr != nil {
		fmt.Fprintln(os.Stderr, openErr)
		os.Exit(1)
	}
	utmps, readErr := utmp.Read(file)
	if readErr != nil {
		fmt.Fprintln(os.Stderr, readErr)
		os.Exit(1)
	}

	// to go utmp
	var goUtmps []*utmp.GoUtmp
	for _, gu := range utmps {
		goUtmps = append(goUtmps, utmp.NewGoUtmp(gu))
	}

	// filter
	filters := []filter.Filter{
		&filter.TimeFilter{Since: sinceTime, Until: untilTime},
	}
	for _, filter := range filters {
		goUtmps = filter.Filter(goUtmps)
	}

	// print
	printer := printer.GetPrinter(outputType)
	for _, u := range goUtmps {
		printer.Print(u)
	}
}
