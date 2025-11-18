package main

import (
	"flag"
	"fmt"
)

func CliADD(a, b int) int {
	return a + b

}

func SUB(a, b int) int {
	return a - b
}

func MUL(a, b int) int {
	return a * b
}

func DIV(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func main() {
	a := flag.Int("a", 0, "number 1")
	b := flag.Int("b", 0, "number 2")
	fmt.Println(CliADD(*a, *b))

}
