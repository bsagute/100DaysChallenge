package main

import (
	"fmt"
	"sync"
	"time"
)

var mx sync.Mutex

var counter = 0

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go add("Bhagwat")
	go add("Sagute")
	wg.Wait()

	fmt.Println("sasas", mx)
}

func add(s string) {
	count := 100
	for i := 0; i < count; i++ {

		mx.Lock()
		x := counter
		x++
		time.Sleep(time.Millisecond * 5)
		counter = x
		fmt.Println(s, "---", "COUNTER", counter)
		mx.Unlock()

	}
	wg.Done()
}
