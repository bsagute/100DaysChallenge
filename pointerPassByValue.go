package main

import "fmt"

func main() {

	number := 10
	fmt.Println("Before", number)
	s := "BHAGWAT"
	change(&number, s)
	fmt.Println("After", number, s)
	fmt.Println(&number)

}
func change(p *int, s string) {
	fmt.Println(p)
	fmt.Println(s)
	s = "SAGUTE"

	*p = 11
	return
}
