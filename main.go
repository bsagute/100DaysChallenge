package main

import "fmt"

func main() {
	ar := []int{1, 1, 2, 2, 3, 3, 5, 6, 6, 7, 7, 8, 8}
	mp := map[int]int{}
	count := 0
	for _, el := range ar {
		mp[el] = el
		if value, isExists := mp[el]; isExists && count <= 2 {
			fmt.Printf("%d %d", value, count)
			count++
		}

	}
	// for ok, e := range mp {
	// 	fmt.Println("ssss", ok, e)
	// }
	// fmt.Println("Sample FOOD MORNIG ", ar, mp)
}
