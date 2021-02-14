package main

import "fmt"

func main() {

	toatl := getFact(4)
	for r := range toatl {
		fmt.Println("", r)
	}
	fmt.Println("GOTA")

}

func getFact(n int) chan int {
	ch := make(chan int)
	go func() {
		t := 1
		for i := n; i > 1; i-- {
			t *= i
		}
		ch <- t
		close(ch)
	}()

	return ch
}
