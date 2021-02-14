package main

import "fmt"

func main() {

	fmt.Println("Enter the max number :- ")
	max := 0
	fmt.Scan(&max)
	fmt.Println("Enter the total table number :- ")
	tmax := 0
	fmt.Scan(&tmax)
	for j := 1; j <= tmax; j++ {

		for i := 1; i <= max; i++ {
			fmt.Println("", j, "*", i, "=", j*i)
		}
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	}

}
