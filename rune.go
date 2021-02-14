package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	list := []int{55, 8, 9, 787, 6354, 24}
	sort.Ints(list)
	fmt.Println("rune", list[len(list)-1])
	fmt.Println("rune", list)
	a := `"'a'
	 name"`
	fmt.Println(reflect.TypeOf(a), a)

}
