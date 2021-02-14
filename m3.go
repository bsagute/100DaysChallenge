package main

import "fmt"

func main() {
	fmt.Println("Data")

	//IMPPPPPPPPP IMPOR
	var m = map[string]string{}
	m["fname"] = "bhagwat"
	m["sname"] = "sagute"
	fmt.Println(m)
	delete(m, "sname")
	fmt.Println(m)

	mm := map[string]int{}
	mm = map[string]int{"A": 11, "ss": 54}
	fmt.Println("", mm)

}
