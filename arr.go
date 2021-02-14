package main

import "fmt"

func main() {
	// fmt.Println("Arr")
	// var arr [100]int
	letter := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println("SLICE", letter)
	fmt.Println("SLICE", letter[0:2])
	fmt.Println("SLICE", len(letter[:]))
	fmt.Println("SLICE", letter[2:])
	fmt.Println("ABC"[0])
	fmt.Println("ABC"[1])
	fmt.Println(string("ABC"[2]))
	// fmt.Println("", arr, letter)
	// for i, v := range arr {
	// 	fmt.Println("", string(i), i, v)

	// }
}
