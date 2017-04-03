package utmp

import (
	"encoding/binary"
	"io"
)

const (
	Empty        = 0x0
	RunLevel     = 0x1
	BootTime     = 0x2
	NewTime      = 0x3
	OldTime      = 0x4
	InitProcess  = 0x5
	LoginProcess = 0x6
	UserProcess  = 0x7
	DeadProcess  = 0x8
	Accounting   = 0x9
)

const (
	LineSize = 32
	NameSize = 32
	HostSize = 256
)

// utmp structures
// see man utmp
type ExitStatus struct {
	Termination int16
	Exit        int16
}

type TimeVal struct {
	Sec  int32
	Usec int32
}

type Utmp struct {
	Type int16
	// alignment
	_       [2]byte
	Pid     int32
	Device  [LineSize]byte
	Id      [4]byte
	User    [NameSize]byte
	Host    [HostSize]byte
	Exit    ExitStatus
	Session int32
	Time    TimeVal
	Addr    [4]int32
	// Reserved member
	Reserved [20]byte
}

// Read utmps
func Read(file io.Reader) ([]*Utmp, error) {
	var us []*Utmp

	for {
		u, readErr := readLine(file)
		if readErr != nil {
			if readErr == io.EOF {
				break
			}
			return nil, readErr
		}
		us = append(us, u)
	}

	return us, nil
}

// read utmp
func readLine(file io.Reader) (*Utmp, error) {
	u := new(Utmp)

	err := binary.Read(file, binary.LittleEndian, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
