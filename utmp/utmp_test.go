package utmp_test

import (
	"os"
	"testing"

	"github.com/neko-neko/utmpdump/utmp"
)

func TestReadIPv6(t *testing.T) {
	f, err := os.Open("testdata/utmp")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	records, err := utmp.Read(f)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(records), 1; got != want {
		t.Fatalf("unexpected number of records: got %d, want %d", got, want)
	}
	gr := utmp.NewGoUtmp(records[0])
	if got, want := gr.Addr, "2a02:168:4a00:19:2e7d:2730:69d4:372a"; got != want {
		t.Errorf("Addr: got %s, want %s", got, want)
	}
}
