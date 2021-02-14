package main

import "fmt"

func main() {
	reminder := 13 % 5
	division := 15 / 5
	dubleDivision := 15 % 5
	fmt.Println("s", reminder)
	fmt.Println("s", division)
	fmt.Println("s", dubleDivision)
	number := 0
	fmt.Println("Enter Number :=")
	fmt.Scan(&number)
	if (number % 2) == 1 {
		fmt.Println("ODD")
	} else {
		fmt.Println("EVEN")
	}
	list := []int{10, 54, 8, 78, 7, 8, 8, 7, 6, 4, 5, 454, 6, 46, 4, 5, 46478, 865, 98, 7}
	for _, number := range list {
		if number%2 == 1 {
			fmt.Printf("ODD")
			fmt.Println("", number)

		} else {
			fmt.Printf("EVEN")
			fmt.Println("", number)

		}
	}
	for i := 0; i < 10; i++ {
		fmt.Println("", i)
	}
}
