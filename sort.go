package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := make([]int, 3)
	i := 0
Loop:
	var x int
	fmt.Scanln(&x)
	// fmt.Println(x)
	//y := string(x)
	// z := "X"
	if x == 0 {
		return
	}
	if i < 3 {
		arr[i] = x
		i++
	} else {
		arr = append(arr, x)
	}
	sort.Sort(sort.IntSlice(arr))
	fmt.Println(arr)
	goto Loop
}
