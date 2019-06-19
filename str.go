package main
import "bufio"
import "os"
import "fmt"

func main() {
in := bufio.NewReader(os.Stdin)
line, _ := in.ReadString('\n')
  for _, r := range line {
        c := string(r)
        fmt.Println(c)
    }
}

UTIB0000114