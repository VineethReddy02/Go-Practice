package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Enter a floating point number to truncate into integer")
	fmt.Scanln(&x)
	var y int = int(x)
	fmt.Println(y)
}
