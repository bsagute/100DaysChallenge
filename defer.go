package main

import "fmt"

func main() {
	greet1("Bhagwat")
	defer greet2("Ravsaheb")
	greet3("Sagute")
}

func greet1(name string) {
	fmt.Println("Welcom Mr. ", name)
}

func greet2(name string) {
	fmt.Println("Welcom Mr. ", name)
}

func greet3(name string) {
	fmt.Println("Welcom Mr. ", name)
}
