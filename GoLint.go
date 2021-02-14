package main

import (
	"fmt"
	"log"

	"os"
)

var isConnected = false

func init() {
	fmt.Println("INIT CALLED ")

	wr, err := os.Create("log.txt")
	if err != nil {
		log.Println("ERRRO IS")
	}
	log.SetOutput(wr)
}
func main() {
	fmt.Println("TEST  gOLINT")
	goLint()
	_, err := os.Open("GosLint.go")
	if err != nil {
		// log.Println("ERROE IS ", err)
		// log.Fatalln("ERROR NOT FOUND", err)
		// fmt.Fatalln(err)
		panic(err)
	}
	// fmt.Println("", filePtr.Name())
	fmt.Println("LAST LINE")

}

//THIS function is used to call
func goLint() string {
	if isConnected {
		return "WORKING FINE"
	}
	return "CHECK CONNECTION SPEED"
}
