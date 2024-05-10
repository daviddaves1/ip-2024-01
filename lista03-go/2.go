package main

import "fmt"

func main(){
	var N, K, contar int = 0, 0, 0

	fmt.Printf("Informe o valor de N: ")
	fmt.Scan(&N)
		for N <= 0 && N <= 1000{
			fmt.Printf("Valor de N incorreto. Tente de novo: ")
			fmt.Scan(&N)
		}

	V := make([]int, N)

	fmt.Printf("Agora informe abaixo %d valores inteiros:\n", N)
	for i:=0; i<N; i++{
		fmt.Scan(&V[i])
	}

	fmt.Printf("Informe o valor de K: ")
	fmt.Scan(&K)

	for i:=0; i<N; i++{
		if V[i] >= K{
			contar++
		}
	}

	fmt.Printf("%d", contar)

}