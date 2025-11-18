package main

import (
	"flag"
	"fmt"
)

func CliADD(a, b int) int {
	return a + b

}

func CliSUB(a, b int) int {
	return a - b

}

func CliMUL(a, b int) int {
	return a * b

}

func CliDIV(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b

}

func main() {
	a := flag.Int("a", 0, "number 1")
	b := flag.Int("a", 0, "number 2")
	fmt.Println(CliADD(*a, *b))
	fmt.Println(CliSUB(*a, *b))
	fmt.Println(CliMUL(*a, *b))
	fmt.Println(CliDIV(*a, *b))

}
