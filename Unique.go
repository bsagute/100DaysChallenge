package main

import "fmt"

func main() {
	ar := []int{1, 5, 8, 6, 9, 11, 9, 7, 8, 9, 2, 5, 6}
	list := GetUnique(ar)
	fmt.Println("", list)
	fmt.Println("", ar)
}
func GetUnique(ar []int) []int {
	uni := []int{}
	mp := map[int]bool{}
	for _, el := range ar {
		if _, isExist := mp[el]; !isExist {
			mp[el] = isExist
			uni = append(uni, el)
		}
	}
	return uni
}
