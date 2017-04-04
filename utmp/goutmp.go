package utmp

import (
	"bytes"
	"fmt"
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
		Addr:    addrToString(u.Addr),
	}
}

// Integer ip address to string
func addrToString(addr [4]int32) string {
	if addr[1] == 0 && addr[2] == 0 && addr[3] == 0 {
		return fmt.Sprintf(
			"%d.%d.%d.%d",
			addr[0]&0xFF,
			(addr[0]>>8)&0xFF,
			(addr[0]>>16)&0xFF,
			(addr[0]>>24)&0xFF,
		)
	} else {
		return fmt.Sprintf(
			"%x:%x:%x:%x:%x:%x:%x:%x",
			addr[0]&0xffff,
			(addr[0]>>16)&0xffff,
			addr[1]&0xffff,
			(addr[1]>>16)&0xffff,
			addr[2]&0xffff,
			(addr[2]>>16)&0xffff,
			addr[3]&0xffff,
			(addr[3]>>16)&0xffff,
		)
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
