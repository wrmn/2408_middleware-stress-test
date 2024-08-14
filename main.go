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
	global.PORT = flag.String("p", "443", "tagret port")

	flag.Parse()

	switch *global.ACT {
	case "tcp":
		go global.IsRun()
		for i := 0; i < *global.ITERATION; i++ {
			wg.Add(1)
			go app.Hit(&wg, i)
		}
		for !*global.Run {
			time.Sleep(1 * time.Millisecond)
		}
		cur := time.Now()
		wg.Wait()
		fmt.Printf("Success %d, Failed %d \n", global.Success, global.Failed)
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
