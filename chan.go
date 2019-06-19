package main

import (
	"fmt"
	"time"
)

func send(s string, messages chan string) {
	fmt.Println("HELLLOOOO")
	messages <- s
}
func rec(messages chan string) string {
	fmt.Println("Waiting")
	s := <-messages
	fmt.Println(s)
	fmt.Println("HEEEYEYYYYY")
	return s
}

func main() {
	messages := make(chan string)
	// fmt.Println(<-messages)

	// fmt.Println(<-messages)
	go rec(messages)
	// fmt.Println(<-messages)

	fmt.Println("VINEETH")
	fmt.Println("REDDY")

	go send("HII", messages)
	time.Sleep(100 * time.Millisecond)

}
