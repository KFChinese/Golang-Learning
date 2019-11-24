package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	var i2 uint64
	var d2 float64
	var s2 string //makes s2 string

	fmt.Scan(&i2)  // Gets i2 input
	fmt.Scan(&d2)  // Gets d2 input
	scanner.Scan() // Initalizes string for "is the best... "
	s2 = scanner.Text()

	fmt.Println(i + i2) // Prints i & i2
	fmt.Printf("%.1f\n", d+d2)
	fmt.Printf("%s%s", s, s2)

}
