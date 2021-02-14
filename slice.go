package main

import "fmt"

func main() {
	fmt.Println("start")
	sl := make([]int, 2, 5)
	fmt.Println("Before", len(sl), cap(sl))
	sl = append(sl, 10)
	sl = append(sl, 20)
	fmt.Println("POINTER", &sl[0])
	sl[0] = 11
	fmt.Println("POINTER", &sl[0])
	sl[1] = 15
	sl = append(sl, 30)
	fmt.Println("POINTER", &sl[0])
	sl = append(sl, 60)
	fmt.Println("SAfter", len(sl), cap(sl))
	fmt.Println("POINTER", &sl[0])
	fmt.Println("SL", sl)

	var aar [2]int
	fmt.Println("ARR", aar)
	aar[0] = 01
	aar[1] = 55
	aar[1] += 5
	aar[0]++
	fmt.Println("ARR", aar)

}
