package main

import "fmt"

const name = "BHAGWAT"

var number = 20

func main() {
	x := 10
	x += 1
	fmt.Println("----", x)
	fmt.Println("s", name)

	//
	max := max(2)
	fmt.Println("LLL", max)

	scope()
	fmt.Println("", number)
}

func scope() {
	number++
	fmt.Println("", name)
}

func max(n int) int {
	return n + 20
}
