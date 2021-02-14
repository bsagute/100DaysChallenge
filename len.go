package main

import "fmt"

func main() {
	fmt.Println("GO")
	name := "BHAGWAT"
	count := 0
	for index, _ := range name {
		fmt.Println(index)
		count++
	}
	a := "'a'"
	fmt.Println("sss", a)
	sa := []string{"A", "b"}
	test("A", "v", "c", "s", "ss", "sss")
	fmt.Println("kkk", 97*7, sa)
}

func test(a ...string) {
	fmt.Println("len", len(a))
}
