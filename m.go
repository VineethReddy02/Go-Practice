package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var myname string
	var myaddr string

	m := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a Name: ")
	scanner.Scan()
	myname = scanner.Text()
	fmt.Print("Enter an Address: ")
	scanner.Scan()
	myaddr = scanner.Text()

	m["name"] = myname
	m["addr"] = myaddr

	fmt.Println(m)
	personjson, _ := json.Marshal(m)
	fmt.Println(string(personjson))

}
