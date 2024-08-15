package global

import (
	"time"
)

var (
	ITERATION       *int
	HOST            *string
	PORT            *string
	ACT             *string
	MODE            *string
	test            = false
	Run             = &test
	Success, Failed int
)

var DataLogon = []byte{
	0, 80, 8, 0, 32, 32, 1, 0, 0, 192, 0, 4, 151, 4, 0, 0, 0, 144, 0, 51, 49, 48, 51, 54, 50, 53, 49, 54, 48, 48, 48, 48, 48, 49, 57,
	57, 57, 49, 49, 53, 57, 50, 49, 0, 55, 72, 84, 76, 69, 48, 52, 49, 48, 48, 49, 48, 48, 49, 49, 48, 51, 54, 50, 53, 49, 54, 48, 48,
	48, 48, 48, 48, 48, 52, 48, 48, 54, 52, 57, 56, 55, 49}

func IsRun() {
	var cur bool
	Run = &cur
	for !*Run {
		cur = time.Now().Second()%2 == 0
		if cur {
			Run = &cur
		}
		time.Sleep(250 * time.Microsecond)
	}
}
