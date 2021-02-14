package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 1215848

	fmt.Println("ENTER NUMBER  ")
	// fmt.Scanf("%d", &number)
	fmt.Println("number", reflect.TypeOf(number))
	reverse_no := 0
	fmt.Println("NUMBER", number)
	for {
		remainder := number % 10

		reverse_no = reverse_no*10 + remainder

		number = number / 10

		if number == 0 {

			break

		}

	}
	fmt.Println("REVERSE", reverse_no)
}
