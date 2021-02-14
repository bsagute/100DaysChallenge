package main

import (
	"fmt"
	"reflect"
	"sort"
)

var a = "bhagwat"
var b = 20
var c float32 = 20
var f, m, n string = "datas", "datdsfa", "datfdfda11"
var d = false
var aa, bb, cc, dd int
var list = []int{1, 1, 55, 25, 5, 6, 8, 8, 8, 78, 78, 7, 7, 8, -25, 0, 25, -202}

func main3() {
	e := 4.15
	fmt.Println("test", list)
	sort.Ints(list)
	fmt.Println(f, m, n)
	fmt.Printf("%d,%d,%d ,%T \n", aa, bb, cc, dd)
	fmt.Println("TYPE", reflect.TypeOf(aa))
	fmt.Printf("%T", e)
	fmt.Printf("%d \n", len(list))
	fmt.Printf("%d \n", list[len(list)-3])
	fmt.Println("test", list)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", &d)
	fmt.Println("test")
}
