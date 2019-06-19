package main

import (
	"fmt"
	"sort"
)

func main() {
	var v int
	var i int
	s := make([]int, 3)
	for {
		fmt.Println("Enter a number: ")
		fmt.Scanln(&v)
		if i < len(s) {
			for n, m := range s {
				if m == 0 {
					s[n] = v
					break
				}
			}
		} else {
			s = append(s, v)
		}
		i = i + 1
		sort.Sort(sort.IntSlice(s))
		fmt.Println(s)
	}
}
