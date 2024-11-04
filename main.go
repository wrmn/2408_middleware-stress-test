package main

import (
	"2408_middleware-stress-test/app"
	"2408_middleware-stress-test/global"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	global.ACT = flag.String("test", "", "test case")
	global.ITERATION = flag.Int("i", 0, "iteration")
	global.HOST = flag.String("h", "", "target host")
	global.MODE = flag.String("m", "chan", "hit mode")
	global.PORT = flag.String("p", "443", "tagret port")

	flag.Parse()
	fmt.Println(*global.ACT)
	switch *global.ACT {
	case "tcp":
		c := make(chan struct{})
		if *global.MODE == "time" {
			go global.IsRun()
		}

		app.HitIteration(&wg, c)

		switch *global.MODE {
		case "time":
			for !*global.Run {
				time.Sleep(1 * time.Nanosecond)
			}
		case "chan":
			time.Sleep(1 * time.Second)
			close(c)
		default:
			log.Fatal("invalid command 2")
		}

		cur := time.Now()
		wg.Wait()
		fmt.Printf("Success %d, Failed %d, Invalid response %d \n", global.Success, global.Failed, global.Invalid)
		fmt.Printf("Time spend : %d ms ", time.Since(cur).Milliseconds())
	case "itr":
		go global.IsRun()
		for i := 0; i < *global.ITERATION; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for !*global.Run {
					time.Sleep(1 * time.Microsecond)
				}
				fmt.Println(time.Now().UnixNano())
			}()
		}
		wg.Wait()
		fmt.Printf("Success %d, Failed %d", global.Success, global.Failed)
	case "test":
		go global.IsRun()
		for {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(time.Now().Second())
			fmt.Println(time.Now().Second() % 3)
			fmt.Println(*global.Run)

		}
	default:
		log.Fatal("invalid command")
	}
}
