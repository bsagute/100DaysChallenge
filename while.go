package main

import "fmt"

func main() {
	fmt.Println("while")

	i := 0
	for i < 10 {
		i++
		fmt.Println("Print:::", i)
	}

	// For with no condition
	j := 0
	for {
		fmt.Println("::::Data", j)
		j++
		if j == 5 {
			break
		}
	}

}
