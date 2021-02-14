package main

import "fmt"

func main() {
	fmt.Println("DeadLock")
	ch := make(chan int)
	go func() {

		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	// for  {

	// 	fmt.Println("VALUE", <-ch)
	// }

	for n := range ch {
		fmt.Println("", n)

	}
}
