package main

import "fmt"

func main() {
	fmt.Println("s")
	list := []float64{1, 8, 98, 98, 7, 54, 2, 6, 8, 5}
	visit(list, func(n float64) {
		fmt.Println("::::", n)
	})
	fmt.Println("e")
}

func visit(list []float64, test func(n float64)) {
	for _, value := range list {
		test(value)
	}

}
