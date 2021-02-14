package main

import "fmt"

func main() {

	var data = []float64{4, 65, 5, 8, 8, 2, 78.8, 56, 56}
	res := Average(data...)
	fmt.Println("RESULT ", res)
	fmt.Printf("%.2f \n", res)
}

func Average(n ...float64) float64 {

	total := 0.0
	for _, v := range n {
		total += v

	}
	return total / float64(len(n))
}
