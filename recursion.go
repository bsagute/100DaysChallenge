package main

import "fmt"

func main() {
	fmt.Println("Enter factorial number:- ")
	number := 0
	fmt.Scan(&number)
	f := fact(number)
	fmt.Println("Factorial of number is :-", f)
}
func fact(number int) int {

	if number == 0 {
		return 1
	}

	return number * fact(number-1)

}
