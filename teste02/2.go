package main

import (
	"fmt"
)

func troca(vetor []int, i, j int) {
	vetor[i], vetor[j] = vetor[j], vetor[i]
}

func trocaOpostosSeMenor(vetor []int, n int) {
	for i := 0; i < n/2; i++ {
		oposto := n - 1 - i
		if vetor[i] < vetor[oposto] {
			troca(vetor, i, oposto)
		}
	}
}

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		t--

		var n int
		fmt.Scan(&n)

		vetor := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&vetor[i])
		}

		trocaOpostosSeMenor(vetor, n)

		for _, v := range vetor {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
