package main

import "fmt"

func main() {

	f := sqrt(2, 3, 4)
	m := find(f)
	// fmt.Println("::LLL", <-m)
	// fmt.Println("::LLL", <-m)
	// fmt.Println("::LLL", <-m)
	for v := range m {
		fmt.Println("PP", v)
	}

	fmt.Println("Bhagwat")
}
func sqrt(a ...int) chan int {
	ch := make(chan int)

	go func() {
		for _, num := range a {
			ch <- num * num
		}
		close(ch)
	}()

	return ch

}
func find(ch chan int) chan int {
	c := make(chan int)

	go func() {
		for n := range ch {
			c <- n
		}
		close(c)
	}()
	return c
}
