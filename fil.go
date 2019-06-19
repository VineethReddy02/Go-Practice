package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	word, _ := os.Open("s.txt")
	words := bufio.NewScanner(word)
	words.Split(bufio.ScanWords)
	in := ""
	flag := 0
	a := ""
	fmt.Scanln(&in)
	//z:=0
	for words.Scan() {
		a = words.Text()
		//z++
		if in == a {

			fmt.Println("Welcome to memory game ", a)
			flag = 1
			//words.Scan()
			//b := words.Text()
			//fmt.Println(b)
			break
		}
	}
	if flag == 0 {
		f, err := os.OpenFile("s.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		f.WriteString("\n")
		if _, err = f.WriteString(in); err != nil {
			panic(err)
		}
		fmt.Println("New User create and Login Successful!! Hi", in)
	}
	i := 1
	j := 10
	p := 1
Loop:
	k := random(i, j)
	fmt.Println(k)
	duration := time.Duration(10) * time.Second
	time.Sleep(duration)
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	t := 0
	fmt.Scanln(&t)

	if k == t {
		fmt.Println("Correct, You are Proceeding to round ", p)
		p++
		i = i * 10
		j = j * 10
		goto Loop
	} else {
		fmt.Println("You Lost in round ", p)
		return
	}

}
