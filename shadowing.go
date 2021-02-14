package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("s")
	welcomeMessage := func() {
		fmt.Println("GOOD MORNING")
	}

	welcomeMessage()
	wrapperFunction := functionCall()
	fmt.Println("TYPE:::", reflect.TypeOf(wrapperFunction))
	fmt.Println("wrapperFunction", wrapperFunction())

}
func functionCall() func() string {

	testVariable := "BHAGWAT_SHADOWING"
	return func() string {
		return testVariable
	}

}
