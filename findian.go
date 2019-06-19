package main

import "fmt"
import "strings"
import "os"
import "bufio"

func main() {
	fmt.Println("Enter the string :)")
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.ToLower(str)
	a := 0
	b := 0
	k := len(str) - 2
	//fmt.Println(k)
	if str[0] == 'i' && str[k-1] == 'n' {
		a++
		for i := 1; i <= k-2; i++ {
			if str[i] == 'a' {
				b++
				fmt.Println("Found!")
				return
			}
		}
		if b == 0 {
			fmt.Println("Not Found!")
			return
		}
	} else {
		fmt.Println("Not Found!")
	}
}
