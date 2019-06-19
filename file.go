package main

import (
	"bufio"
	"fmt"
	"os"
)

type lin struct {
	fname string
	lname string
}

func main() {
	filename := ""
	fmt.Scanln(&filename)
	lines := []*lin{}
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	//i := 0
	for scanner.Scan() {
		line := scanner.Text()
		x := new(lin)
		x.fname = line
		scanner.Scan()
		line1 := scanner.Text()
		x.lname = line1
		lines = append(lines, x)
	}

	//fmt.Println(lines)
	for _, x := range lines {
		fmt.Println()
		fmt.Print("First Name: ", x.fname)
		fmt.Print("||")
		fmt.Print("Last Name: ", x.lname)
	}
}
