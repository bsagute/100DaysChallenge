package main

import "fmt"

func main() {
	fmt.Println("Ss")
	mom := map[string]map[string]string{

		"A": map[string]string{
			"name":  "bhagwat",
			"lname": "sagute",
		},
		"B": map[string]string{
			"D": "DD",
			"E": "EE",
		}}
	fmt.Println("", mom["B"]["E"])
	// mmp := map[string]map[string]string{}
	//IMPPPPPPPPPPPPPPPPPPPPPP
	delete(mom["B"], "E")
	fmt.Println("MOM", mom["B"]["E"])
	fmt.Println("DELETES", mom)
}
