package main

import (
	"fmt"
	"time"
)

func main2() {

	fmt.Println("Welcome ")
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			fmt.Println("CHAN", <-ch)
		}
	}()
	fmt.Println("NAND", <-ch)
	time.Sleep(time.Millisecond * 2)

}
