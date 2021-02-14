package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var counter = 0

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("runtime.NumCPU()", runtime.GOMAXPROCS(1))
}
func main() {
	fmt.Println("RaceCondition")
	wg.Add(2)
	go inc("Bhagwat")
	go inc("Sagute")
	wg.Wait()
}

func inc(s string) {

	for i := 0; i < 20; i++ {
		x := counter
		x++
		//rand.Shuffle(6, swap(10, 57))
		fmt.Println("", rand.Intn(20))
		time.Sleep(time.Millisecond * 15)
		counter = x
		fmt.Println(s, i, counter)
	}
	wg.Done()
}
