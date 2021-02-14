package main

import (
	"fmt"
	"reflect"
)

func main() {
	const pi = 20
	const literal string = "bhagwat"
	value := int64(20)
	value = value + pi
	// const fname = 2
	// const sname string = "_sagute"

	fmt.Println("test", value, reflect.TypeOf(value), reflect.TypeOf(pi), reflect.TypeOf(literal))
}
