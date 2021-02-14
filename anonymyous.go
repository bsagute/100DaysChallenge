package main

import "fmt"

func main() {
	fmt.Println("test")

	sampel := wrapperFunction()
	fmt.Println(":::1 ", sampel())
	fmt.Println(":::2 ", sampel())
}

func wrapperFunction() func() int {
	return func() int {
		var x = 20
		x++
		return x
	}
}

func wrapFunction() func() string {
	b := "sss"
	return func() string {
		return b
	}

}
