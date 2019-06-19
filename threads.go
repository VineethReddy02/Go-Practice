package main

import "fmt"
import "sync"

func even(i int, c chan bool) {
	fmt.Println(i, "Thread-1")
	// c <- true
}

func odd(i int, c chan bool) {
	fmt.Println(i, "Thread-2")
	// c <- true
}
func main() {
	var wg sync.WaitGroup
	c := make(chan bool, 1)
	c <- true
	wg.Add(1)
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if i == n {
			wg.Done()
		}
		if i%2 == 0 {
			//<-c
			go even(i, c)
		}
		if i%2 != 0 {
			// <-c
			go odd(i, c)
		}

	}
	wg.Wait()
}
