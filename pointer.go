package main

import "fmt"

func main() {
	number := 20
	var pointer *int = &number

	*pointer = 50
	zero(number)
	fmt.Println(":::", &number)

}

func zero(y int) {
	y = 0
	fmt.Println("function value  address", &y)
}
