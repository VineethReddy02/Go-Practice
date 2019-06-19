package main

import (
	"encoding/json"
	"fmt"
)

// type Counter struct {
// 	name    string
// 	address string
// }

func main() {
	map1 := make(map[string]string)
	name1 := ""
	fmt.Scanln(&name1)
	add1 := ""
	fmt.Scanln(&add1)
	map1["name"] = name1
	map1["address"] = add1

	fmt.Println(map1)
	bar, _ := json.Marshal(map1)
	fmt.Println(string(bar))
}
