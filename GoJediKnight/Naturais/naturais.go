package main

import "fmt"

func somandoNaturais() int {
	soma := 0
	for i := 1; i <= 100; i++ {
		soma += i
	}
	return soma
}

func main() {
	fmt.Println("Soma dos naturais:", somandoNaturais())
}
