package main

import "fmt"

func swap(m *int, n *int) {
	t := *m
	*m = *n
	*n = t
}
func bubblesort(l int, sli []int) {
	end := l
	for {
		if end == 0 {
			break
		}
		for i := 0; i < l-1; i++ {
			if sli[i] > sli[i+1] {
				swap(&sli[i], &sli[i+1])
			}

		}
		end = end - 1
	}
}
func main() {
	sli := []int{}
	//i := 0
	for i := 0; i < 10; i++ {
		k := 0
		fmt.Scan(&k)
		sli = append(sli, k)
	}
	bubblesort(len(sli), sli)
	fmt.Println(sli)
}
