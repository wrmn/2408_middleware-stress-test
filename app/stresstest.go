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

func hit(wg *sync.WaitGroup, c chan struct{}, i int) {
	defer wg.Done()
	if *global.MODE == "time" {
		for !*global.Run {
			time.Sleep(1 * time.Millisecond)
		}
	} else {
		<-c
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

func HitIteration(wg *sync.WaitGroup, c chan struct{}) {
	for i := 0; i < *global.ITERATION; i++ {
		wg.Add(1)
		go hit(wg, c, i)
	}
}
