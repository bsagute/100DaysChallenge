package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{4, 5, 8, 9, 7, 8, 9, 5, 5, 8, 959, 59, 56, 59}
	sort.Sort(sort.IntSlice(array))
	for i := 0; i < len(array)-1; i++ {
		if array[i] == array[i+1] {

			array = append(array, array[i+1])
		}
	}
	fmt.Println("S", array)
}
