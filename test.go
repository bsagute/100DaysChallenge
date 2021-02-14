package main

import "fmt"

func main() {
	fmt.Println("Install")
	m := map[string]string{
		"A": "AA",
		"B": "BB"}

	delete(m, "A")
	if value, isExists := m["A"]; isExists {
		fmt.Println("m", value, isExists, m)

	} else {
		fmt.Println("m", value, isExists, m)

	}
}
