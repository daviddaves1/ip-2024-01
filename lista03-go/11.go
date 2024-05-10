package main

import "fmt"

func main() {
	var N, maior, menor int = 0, 0, 0

	fmt.Printf("Valor de N: ")
	fmt.Scan(&N)

	V := make([]int, N)
	W := make([]int, N)

	fmt.Printf("Insira %d valores:\n", N)
	for i:=0; i<N; i++{
		fmt.Scan(&V[i])
		W[N-1-i] = V[i]
	}

	maior = V[0]
	for i:=0; i<N; i++{
		if maior < V[i] {
			maior = V[i]
		}
	}

	menor = W[0]
	for i:=0; i<N; i++{
		menor = W[i]
	}

	fmt.Printf("\n\n")
	for i:=0; i<N; i++{
		fmt.Printf("%d	", V[i])
	}

	fmt.Println()
	for i:=0; i<N; i++{
		fmt.Printf("%d	", W[i])
	}

	fmt.Printf("\n")
	fmt.Println(maior)
	fmt.Println(menor)
}
