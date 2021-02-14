package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const pi = 3.14

func main() {

	fmt.Println("data", pi)
	res, _ := http.Get("https://www.golang-book.com/")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("body", string(body))

}
