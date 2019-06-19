package main

import "fmt"

func main() {
	var n, k int
	fmt.Println("Enter the number of people: ")
	fmt.Scanln(&n)
	fmt.Println("Enter the position to be killed: ")
	fmt.Scanln(&k)
	fmt.Println(n, k)
	//const f := n
	var arr [10]int
	for i := 1; i <= n; i++ {
		//	if i == k {
		arr[i] = 0
		//	}
	}
	fmt.Println(arr)
}
