package main

import "fmt"

func main() {
	length := 0
	width := 0
	fmt.Println("Enter the length : -")
	fmt.Scan(&length)
	fmt.Println("Enter the width  : -")
	fmt.Scan(&width)
	area := length * width
	flatCost := 0
	fmt.Println("Enter the flat Cost :-")
	fmt.Scan(&flatCost)
	areaPerSqFt := flatCost / area
	fmt.Println("Flat Cost  >>>", flatCost)
	fmt.Println("Total Area >>>", area)
	fmt.Println("SqFt Value >>>", areaPerSqFt)

}
