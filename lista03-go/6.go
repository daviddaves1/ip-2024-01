package main

import "fmt"

func main(){
	var n, soma int = 0, 0

	fmt.Printf("Insira o valor de N: ")
	fmt.Scan(&n)
	for n >= 5000{
		fmt.Printf("N não pode ser maior ou igual a 5000. Insira de novo: ")
		fmt.Scan(&n)
	}
	fmt.Print("Agora insira %d elementos:\n", n)

	array := make([]int, n)
	
	for i:=0; i<n; i++{
		fmt.Scan(&array[i])
		soma += array[i]
	}
	fmt.Printf("Saída: %d\n", soma)
}