package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go red("Bhagwat")
	go red("Sagute")
	wg.Wait()
	fmt.Println("AutoMacity")
}
func red(s string) {
	count := 20
	for i := 0; i < count; i++ {
		// x := counter
		// x++

		atomic.AddInt64(&counter, 1)
		time.Sleep(time.Millisecond * 10)
		// counter = x

		fmt.Println(s, "----", counter)
	}
	wg.Done()

}
