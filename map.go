package main

import "fmt"

func main() {

	//LOOP on MAp
	m := map[int]string{
		1: "A",
		2: "AA",
		3: "AAA",
		4: "AAAA"}
	fmt.Println("s", m)
	sl := make([]int, 5, 10)
	sl = []int{45, 4, 5, 45, 4, 8, 48, 4, 5, 4, 8, 48}
	for key, value := range m {
		fmt.Println(key, "___", value)
	}
	for val, v := range sl {
		fmt.Println("KKKKK", val, v)

	}
}
