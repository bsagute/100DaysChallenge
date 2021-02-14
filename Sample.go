package main

import "fmt"

func main() {
	fmt.Println("WELCOME")

	chp := getReturnChan(boring("bhagwat"), boring("sagute"))
	fmt.Println(chp)
	for i := 0; i <= 10; i++ {
		fmt.Println("", <-chp)
	}

}

func boring(s string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", s, i)
		}
	}()
	return ch
}

func getReturnChan(ch, ch2 <-chan string) <-chan string {
	ck := make(chan string)

	go func() {
		for {
			ck <- <-ch
		}
	}()
	go func() {
		for {
			ck <- <-ch2
		}
	}()

	return ck
}
