package main

import "fmt"

func ADD(a, b int) int {
	return a + b

}

func SUB(a, b int) int {
	return a - b
}

func main() {
	result := ADD(5, 6)
	fmt.Println(result)
	fmt.Println(SUB(5, 6))

}
