package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, _ := ioutil.ReadFile("trunc.go")
	fmt.Println(string(dat))
	//fmt.Println(k)
}
