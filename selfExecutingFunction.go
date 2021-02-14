package main

import "fmt"

func main() {

	fmt.Println("ds")
	func() {
		fmt.Println("SELF EXECUTING")
	}()
}
