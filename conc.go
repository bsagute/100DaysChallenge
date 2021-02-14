package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 10
		time.Sleep(time.Second * 3)
	}()
	fmt.Println("Main", <-ch)

}
