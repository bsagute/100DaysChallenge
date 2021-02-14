package main

import (
	"fmt"
	"reflect"
)

const data = "SOL"

func main() {
	fmt.Println("Variadic")
	avg := Average(4, 5, 66, 87, 9, 878, 45)
	fmt.Println("AVG", avg)
}

func Average(n ...float64) float64 {

	fmt.Println("LIST", n)
	fmt.Println("TYPE", reflect.TypeOf(n))
	vars := 0.0
	for _, s := range n {
		vars += s

	}
	fmt.Println("TOTAL :- ", vars)
	return vars / float64(len(n))
}
