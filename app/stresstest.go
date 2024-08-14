package app

import (
	"2408_middleware-stress-test/connections"
	"2408_middleware-stress-test/global"
	"2408_middleware-stress-test/utility"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Hit(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	for !*global.Run {
		time.Sleep(1 * time.Millisecond)
	}
	msg, tid := utility.GetSampleLogon()
	_, err := connections.SendTCPRequest(msg)
	if err != nil {
		fmt.Printf("[x] iteration %s : Failed  [tid : %s] Error : %s \n", utility.PadLeft(3, ' ', strconv.Itoa(i)), tid, err.Error())
		global.Failed++
		return

	}
	fmt.Printf("[v] iteration %s : Success [tid : %s]\n", utility.PadLeft(3, ' ', strconv.Itoa(i)), tid)
	global.Success++
}
