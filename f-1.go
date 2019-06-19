package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// /**
//   Helper Function
//   To read from keyboard strings containing
//   blancks or white spaces
//   based on the work presented in
//   https://stackoverflow.com/questions/43843477/scanln-in-golang-doesnt-accept-whitespace
// **/
func inputAnyString() string {
	var line string = ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line = scanner.Text()
		fmt.Printf("             The input was:      %q\n", line)
		fmt.Println()
		return line
	}
	return line
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Define a Struct

type name struct {
	fname [20]string
	lname [20]string
}

func main() {

	s := make([]name, 0)

	var fileName string
	// Use of a helper function to enter name and address
	// including the case of white space separated strings
	fmt.Print("\n\nEnter the name of the file:      ")
	fileName = inputAnyString()
	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		// n := new(name)

		var n name
		v := strings.Fields(scanner.Text())

		if len(v[0]) <= 20 {
			w := 20 - len(v[0])
			// fmt.Println("w len :", w)
			b := make([]string, w)

			for i := range v[0] {
				b = append(b, string(v[0][i]))
			}

			// fmt.Println("b2 len :", len(b))
			// fmt.Println("b :", b)

			for i := 0; i < 20; i++ {
				n.fname[i] = string(b[i])

			}

		} else {
			b := v[0][:20]

			for i := 0; i < 20; i++ {
				n.fname[i] = string(b[i])
			}
		}

		if len(v[1]) <= 20 {
			w := 20 - len(v[1])
			// fmt.Println("w len :", w)
			b := make([]string, w)

			for i := range v[1] {
				b = append(b, string(v[1][i]))
			}

			// fmt.Println("b2 len :", len(b))
			// fmt.Println("b :", b)

			for i := 0; i < 20; i++ {
				n.lname[i] = string(b[i])

			}

		} else {
			b := v[1][:20]

			for i := 0; i < 20; i++ {
				n.lname[i] = string(b[i])
			}
		}

		//fmt.Println("n.fname :", n.fname)
		//fmt.Println("n.lname :", n.lname)

		s = append(s, n)
	}
	//fmt.Println(s)

	for i := range s {
		fname := ""
		lname := ""
		for j := range s[i].fname {
			//fname = string.add()   fname.add(j)
			fname = fname + s[i].fname[j]
		}

		for j := range s[i].lname {
			//fname = string.add()   fname.add(j)
			lname = lname + s[i].lname[j]
		}

		//fmt.Print("Hola", s[i].fname, s[i].lname)
		fmt.Println("First Name:         ", fname)
		fmt.Println("Last Name:          ", lname, "\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
