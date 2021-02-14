package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("GOod, Binary Tree")
	ar := []int{1, 33, 342, 234, 234, 234, 2342, 34234, 234}
	sort.Sort(sort.IntSlice(ar))
	//	sort.Sort(sort.Reverse(ar))
	for _, v := range ar {
		fmt.Println("", v)
	}
	for i := len(ar); i > 0; i-- {
		fmt.Printf("%d \t", ar[i-1])
	}
	fmt.Println(ar)

}
