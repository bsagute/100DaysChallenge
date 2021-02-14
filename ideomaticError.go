package main

import (
	"errors"
	"log"
	"os"
)

func init() {
	fp, err := os.Create("ServerLog.txt")
	if err != nil {
		log.Println("Errror While creating log File ", err)
	}
	log.SetOutput(fp)
}

var NOT_FOUND = errors.New("CANDIDATE DETAILS NOT FOUND!!!")

func main() {

	err := getError()
	if err != nil {
		log.Println("ERROR IS L::", err)
	}

}

func getError() error {
	return NOT_FOUND

}
