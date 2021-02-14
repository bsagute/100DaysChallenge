package main

import "fmt"

func main() {
	var list []int
	list = []int{45, 8, 4, 5, 8, 4, 887, 54, 8, 54, 8, 7}
	filterList := filter(list, func(a int) bool {
		return a < 10
	})

	fmt.Println("Filterred List is ", filterList)
}
func filter(list []int, callback func(a int) bool) []int {
	var l []int
	for _, v := range list {

		if callback(v) {

			l = append(l, v)
		}
	}
	return l
}
