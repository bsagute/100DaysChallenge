package main

import (
	"fmt"
	"reflect"
)

func main() {

	f := 'a'
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))
	fmt.Printf("%T \n", rune(f))
}
