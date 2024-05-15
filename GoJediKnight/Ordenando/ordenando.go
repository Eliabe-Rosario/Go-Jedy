package main

import "fmt"

func ordenandoInteiros(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{5, 3, 7, 2, 8, 1, 9, 4, 6, 12, 11, 10}
	fmt.Println("Primeiro Array:", arr)
	fmt.Println("Ordenado:", ordenandoInteiros(arr))
}
