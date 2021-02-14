package main

import "fmt"

func main() {
	fmt.Println(greet("AA", "BB", "CC"))
}

func greet(AA, BB, CC string) string {
	return fmt.Sprint(AA, "__", BB, "__", CC)
}

func returnData(BB, CC, AA string) (concat string) {

	concat = fmt.Sprint(AA, "__", BB, "__", CC)
	fmt.Println(concat)
	return
}
