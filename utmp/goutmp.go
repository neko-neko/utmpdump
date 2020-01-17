package utmp

import (
	"bytes"
	"time"
)

// Display time format
const TimeFormat = time.RFC1123

type GoExitStatus struct {
	Termination int
	Exit        int
}

type GoUtmp struct {
	Type    int
	Pid     int
	Device  string
	Id      string
	User    string
	Host    string
	Exit    GoExitStatus
	Session int
	Time    string
	Addr    string
}

// Convert Utmp to GoUtmp
func NewGoUtmp(u *Utmp) *GoUtmp {
	return &GoUtmp{
		Type:   int(u.Type),
		Pid:    int(u.Pid),
		Device: string(u.Device[:getByteLen(u.Device[:])]),
		Id:     string(u.Id[:getByteLen(u.Id[:])]),
		User:   string(u.User[:getByteLen(u.User[:])]),
		Host:   string(u.Host[:getByteLen(u.Host[:])]),
		Exit: GoExitStatus{
			Termination: int(u.Exit.Termination),
			Exit:        int(u.Exit.Exit),
		},
		Session: int(u.Session),
		Time:    time.Unix(int64(u.Time.Sec), 0).Format(TimeFormat),
		Addr:    u.Addr().String(),
	}
}

// get byte \0 index
func getByteLen(byteArray []byte) int {
	n := bytes.IndexByte(byteArray[:], 0)
	if n == -1 {
		return 0
	}

	return n
}
