package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("ss")
	res, _ := http.Get("http://example.com/")
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("", string(data))
}
