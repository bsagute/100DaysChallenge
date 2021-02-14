package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Sss")
	m := map[int]string{
		1: "Bhagwat",
		2: "Sagute",
		3: "Ravsaheb",
		4: "TalentProIndia",
		5: "pune"}
	fmt.Println("", m)

	if value, exists := m[6]; exists {
		delete(m, 5)
		fmt.Println("DELETED ", m, value)

	} else {
		fmt.Println("VALUE NOT EXISTS", reflect.TypeOf(value), exists)
	}

}
