package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 15
	m["k2"] = 145
	m["k3"] = 168
	fmt.Println("MAP:::::>  ", m, len(m))
	fmt.Println("", m["k2"])
	delete(m, "k2")
	fmt.Println("MAP:::::>  ", m, len(m))
	//If the key is exist map return the TRUE else false
	value, k := m["k1"]
	fmt.Println("Ss", k, value)

}
