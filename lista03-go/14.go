package main

import "fmt"

func main(){ 
	var N, contar, soma, mediaImpar, mediaPar int

	fmt.Printf("Informe o valor de N: ")
	fmt.Scan(&N)

	num := make([]int, N)

	fmt.Printf("Digite os %d numeros inteiros correspondentes a mediana:\n", N)
	for i:=0; i<N; i++{
		fmt.Scan(&num[i])
		soma += num[i]
		contar++
	}

		if soma % 2==0{
			mediaPar = soma/contar
			fmt.Printf("Mediana: %.2f\n", float64(mediaPar))
		} 
		if contar % 2==0{
				mediaImpar = num[N/2 + 1]
				fmt.Printf("Mediana: %.2f\n", float64(mediaImpar))
		}
}