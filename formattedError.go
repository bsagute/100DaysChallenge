package main

import (
	"fmt"
	"log"
	"os"
)

func init() {

	fp, ers := os.Create("FormattedError.txt")
	if ers != nil {
		log.Println("ERROR IS OCCURED", ers)
	}
	log.SetOutput(fp)

}

func main() {
	fmt.Println("FormatterError")
	count := 2563
	fmt.Errorf("Data is called %v", count)
}
