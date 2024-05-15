package main

import "fmt"

func soma(a, b, c int) int {
	return a*a + b*b + c*c
}

func main() {
	resultado := soma(4, 6, 8)
	fmt.Println("Soma dos Quadrados:", resultado)
}
