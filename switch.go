package main

import "fmt"

func main() {

	fmt.Println("Enter your name :=")
	name := ""
	fmt.Scan(&name)
	switch name {
	case "bhagwat":
		fmt.Println("Hello mr " + name)

	case "sagute":
		fmt.Println("Sagute Bhagwat Ravsaheb ")
	case "gullu", "mullu", "soni":
		fmt.Println("GODDTY")

	default:
		fmt.Println("No Name Found")
	}
}
