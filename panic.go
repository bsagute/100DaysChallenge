package main

import (
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("PANIC")
	wg.Add(1)
	go testPanic()
	wg.Wait()
}

func testPanic() {

	fmt.Println("FIRST")
	fmt.Println("SECOND")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("RECOVER IS CALLED ---->", r)

		}
	}()
	fmt.Println("AFTER RECOVERY")

	gota()

}

func gota() {

	func() {
		fmt.Println("SELF EXUECUTE")
		_, err := os.Open("Sample.txt")
		if err != nil {
			fmt.Println("ERROR OCCURED ")
			panic(err)

		}
		fmt.Println("AFTER PANIC")

	}()
	wg.Done()
}
