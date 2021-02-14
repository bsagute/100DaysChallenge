package main

import "fmt"

func main() {
	sl := new([100]int)[:50]
	fmt.Println("s", len(sl), cap(sl))
	slice := make([]int, 2, 3)
	fmt.Println("s", len(slice), cap(slice))
	slice = append(slice, 5)
	slice = append(slice, 155, 255, 535, 448, 668, 68)
	fmt.Println("s", len(slice), cap(slice), slice)
	slice = append(slice[:2], slice[3:len(slice)-1]...)
	fmt.Println("", slice)
}
