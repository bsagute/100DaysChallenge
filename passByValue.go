package main

import "fmt"

func main() {
	fmt.Println("s")
	name := "Bhagwat"
	changeMe(&name)
	fmt.Println("CallS", name)

}

func changeMe(name *string) {

	fmt.Println("Changed ", name)
	*name = "Sagute"
	fmt.Println("LASS", name)
}
