package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func init() {
	fp, er := os.Create("log.txt")
	if er != nil {
		fmt.Println(er)

	}

	log.SetOutput(fp)
}
func main() {
	fmt.Println("ERRO")
	err := testError()

	if err != nil {
		fmt.Println(err)
		// log.Fatalln("EEEEE", err)
	}
	_, er := os.Open("finalssLog.txt")
	if er != nil {
		log.Println(er)
		log.Println("MANAMAMA", er)

	}

}
func testError() error {
	return errors.New("NODATA ERROR")
}
