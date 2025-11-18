package main

import "fmt"

func ADD(a, b int) int {
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
	result := ADD(5, 6)
	fmt.Println(result)
	fmt.Println(SUB(5, 6))
	fmt.Println(MUL(5, 6))
	fmt.Println(DIV(5, 6))

}
