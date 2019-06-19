package main

import "fmt"

type line struct {
	fname string
	lname string
}

func main() {
	sli := []line{}
	sli = append(sli, "vineeth", "vin")
	sli = append(sli, "vin")
	fmt.Println(sli)
}
