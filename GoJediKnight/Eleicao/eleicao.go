package main

import "fmt"

func eleicaoSindical(votosA, votosB, votosC, nulos, emBranco, totalEleitores int) {
	fmt.Println("Total de eleitores:", totalEleitores)
	fmt.Printf("Votos validos: %.2f%%\n", float64(votosA+votosB+votosC)/float64(totalEleitores)*100)
	fmt.Printf("Votos válidos do candidato A : %.2f%%\n", float64(votosA)/float64(totalEleitores)*100)
	fmt.Printf("Votos válidos do candidato B: %.2f%%\n", float64(votosB)/float64(totalEleitores)*100)
	fmt.Printf("Votos válidos do candidato C: %.2f%%\n", float64(votosC)/float64(totalEleitores)*100)
	fmt.Printf("Votos Nulos: %.2f%%\n", float64(nulos)/float64(totalEleitores)*100)
	fmt.Printf("Votos Brancos: %.2f%%\n", float64(emBranco)/float64(totalEleitores)*100)
}

func main() {
	eleicaoSindical(400, 150, 120, 50, 30, 725)
}
