package main

import (
	"fmt"
	"sort"
)

func main() {

	s := "zacsd"
	a := "csadz"
	Check(s, a)
	fmt.Println("OK", s, a)
}

type SortRune []rune

func (a SortRune) Len() int           { return len(a) }
func (a SortRune) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortRune) Less(i, j int) bool { return a[i] < a[j] }

func Check(s, a string) bool {
	t := []rune(s)
	sort.Sort(SortRune(t))
	fmt.Println("sssss", t)
	if len(a) != len(s) {
		return false
	}

	fmt.Println("")
	return true
}
