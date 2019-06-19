package main

import "fmt"

func inc(z int) int {

	fmt.Println("GOROUTINE-1", z)
	return z + 3
}

func dec(z int) int {
	fmt.Println("GOROUTINE-2", z)
	return z - -6
}

func main() {
	x := 3
	go func() {
		x = inc(x)
	}()
	x = x + 1
	go func() {
		x = dec(x)
	}()
	// fmt.Println(x)
	fmt.Scanln()
	fmt.Println(x)
}

/* Race conditions is when two or more routines have access to the same resource such as variable or data structure
this type code can create bugs as the change in flow occurs. So when you run this code press enter,
pressing it immediately results in different ouput compared to pressing enter after a time delay

So delay in milliseconds in processing results in different output.
*/
