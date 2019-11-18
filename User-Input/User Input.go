package main
import (
        "fmt"
        "os"
        "bufio"
)

func main() {

fmt.Println("Hello, World.")
reader := bufio.NewReader(os.Stdin)
 text, _ := reader.ReadString('\n')
fmt.Println(text)

}

