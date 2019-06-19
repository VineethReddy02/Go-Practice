package main

import "fmt"
import "time"

func main() {
	go fmt.Println("HII HEY")
	fmt.Printf("HEELLLO")
	time.Sleep(100 * time.Millisecond)
}
