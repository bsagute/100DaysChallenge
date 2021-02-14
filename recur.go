package main

import (
	"fmt"
)

func main() {
	fmt.Println("GOE")

	c := getFacts(50)

	d := calc(c)

	for e := range d {
		fmt.Println("KKKK", e)
	}

}

func getFacts(n int) (ch chan int) {
	ch = make(chan int)

	go func() {
		for i := 2; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func calc(ch chan int) (k chan int) {
	k = make(chan int)

	go func() {
		for f := range ch {
			k <- cFact(f)
		}
		close(k)
	}()

	return k
}

func cFact(n int) int {
	// fmt.Println("FF", n)
	total := 1
	for i := 1; i < n; i++ {
		total *= i

	}
	return total

}
