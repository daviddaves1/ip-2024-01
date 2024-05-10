package main

import "fmt"

func main(){
	var n int = 0

	fmt.Printf("Informe o valor de N: ");
	fmt.Scan(&n)
	fmt.Printf("\n")

	array := make([]int, n)

	fmt.Printf("Agora informe %d valores inteiros:\n", n)
	for i:=0; i<n; i++{
		fmt.Scan(&array[i]);
	}

	for i:=n-1; i>=0; i--{
		fmt.Printf("Saída: %d\n", array[i])
	}
}