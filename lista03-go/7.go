package main

import "fmt"

func main(){
	var N, M int = 0, 0

	fmt.Printf("Insira o valor de N: ")
	fmt.Scan(&N)
		for N<1 && N <= 10000{
			fmt.Printf("Valor incorreto. Tente novamente: ")
			fmt.Scan(&N)
		}
	
		array := make([]int, N)
		M = array[0]

		fmt.Printf("Agora informe %d elementos:\n", N)
		for i:=0; i<N; i++{
			fmt.Scan(&array[i])
			
			if array[i] > M {
				M = array[i]
			}
		}

		fmt.Printf("Maior elemento: %d\n\n", M)
}

