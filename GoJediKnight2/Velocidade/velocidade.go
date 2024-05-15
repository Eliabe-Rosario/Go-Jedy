package main

import "fmt"

func Velocidade(distancia float64, tempo float64) float64 {
	return (distancia * 1000) / (tempo * 60)
}

func main() {
	velocidade := Velocidade(12, 6)
	fmt.Println("Velocidade do Projétil (m/s):", velocidade)
}
