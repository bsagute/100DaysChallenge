package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {

	name := "BSDA"
	same := "ASDB"
	ar := []int{}
	for _, v := range same {
		fmt.Println("", reflect.TypeOf(v))
		ar = append(ar, int(v))
	}
	sort.Sort(sort.IntSlice(ar))
	fmt.Println("", ar)
	fmt.Println("name", name)
	fmt.Println("Adn")
}
