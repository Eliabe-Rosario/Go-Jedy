package main

import "fmt"

func mediaDoAluno(nota1, nota2, nota3, nota4 float64) (float64, string) {
	media := (nota1 + nota2 + nota3 + nota4) / 4
	if media >= 5 {
		return media, "Aprovado"
	} else {
		return media, "Reprovado"
	}
}

func main() {
	media, situacao := mediaDoAluno(10, 5, 5, 9)
	fmt.Println("Situacao:", situacao)
	fmt.Println("Média:", media)
}
