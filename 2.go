package main

import "fmt"

var X = 25

func Increment() int {
	x++

	return x

}
func main2() {

	fmt.Println("Data", 25, x, Increment(), Increment(), Increment(), x)
	fmt.Println("Data", Increment())

}
