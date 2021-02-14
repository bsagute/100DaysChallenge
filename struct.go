package main

import (
	"fmt"
	"reflect"
)

type foo int
type person struct {
	fname string
	lname string
	age   int
}

func main() {
	var age foo
	age = 25

	firstPerson := person{"Bhagwat", "sagute", 24}
	fmt.Println("ss", age, reflect.TypeOf(age), firstPerson.lname)

}
