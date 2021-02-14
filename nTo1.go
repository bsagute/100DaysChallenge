package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Semaphpre")
	c := make(chan int)
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(c)
	}()
	count := 0
	for v := range c {
		count++
		fmt.Println("Cale", v)
	}
	fmt.Println("COUNT", count)

}
