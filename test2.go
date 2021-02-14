package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []float64{455, 54, 4}
	fmt.Println("S", a)
	s := a[:2]
	fmt.Println("S", s)
	fmt.Println("t", reflect.TypeOf(a))
	ss := make([]float64, 5)
	fmt.Println("ss", ss)
	ss = []float64{55, 55, 88, 55, 78, 45, 645, 45}
	fmt.Println("ss", ss)

}
