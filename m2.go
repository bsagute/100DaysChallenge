package main

import "fmt"

func main() {
	fmt.Println("DD")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("aaa", m)
	mm := make(map[string]int)
	fmt.Println("ss", mm)
	mm = map[string]int{"B": 5, "C": 57, "DD": 55}
	fmt.Println("ss", mm)
	value, ok := mm["D"]
	fmt.Println("::::", value, ok)

}
