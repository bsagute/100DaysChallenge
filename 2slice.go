package main

import "fmt"

func main() {
	s := make(map[string]int)
	changeMe(s)
	fmt.Println("LLLLL", s["A"])
}

func changeMe(sl map[string]int) {
	fmt.Println("ss", sl)
	sl["A"] = 11

}
