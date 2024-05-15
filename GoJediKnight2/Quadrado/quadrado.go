package main

import "fmt"

func quadradoDaSoma(a, b, c int) int {
	soma := a + b + c
	return soma * soma
}

func main() {
	resultado := quadradoDaSoma(5, 3, 6)
	fmt.Println("Quadrado da Soma:", resultado)
}
