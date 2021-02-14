package main

import "fmt"

var x = 10

func main() {
	fmt.Println("Data", incP())
	fmt.Println("Data", incP())

}
func incP() int {
	x++
	return x
}
