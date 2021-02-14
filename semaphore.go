package main

import "fmt"

func main() {
	fmt.Println("test")
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		count := 10

		for i := 0; i < count; i++ {
			ch <- i
		}
		done <- true
	}()
	go func() {
		count := 10

		for i := 0; i < count; i++ {
			ch <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(ch)
	}()

	for v := range ch {
		fmt.Println("ss", v)

	}

}
