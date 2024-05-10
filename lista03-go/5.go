package main

import "fmt"

func main(){
	var n, maior, index int = -1, 0, 0

	for n != 0{

		fmt.Print("Informe o valor de N: ")
		fmt.Scan(&n)

		array := make([]int, n)
		maior = array[0]

		fmt.Printf("Agora informe %d elementos:\n", n)
		for i:=0; i<n; i++{
			fmt.Scan(&array[i])

			if array[i] > maior{
				maior = array[i]
				index = i
			}
		}
		fmt.Printf("Saída: %d	%d\n\n", index, maior)
	}
}