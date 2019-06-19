package main

import "fmt"

type Interf interface{ Eat() }

type typed struct{ name string }

func (d *typed) Eat() {
	if d == nil {
		fmt.Println("HIIIII")
	} else {
		fmt.Println(d.name)
	}

}
func main() {
	var i1 Interf
	var d1 *typed
	d1.name = "VINEETH"
	i1 = d1
	i1.Eat()

}
