package main

import "fmt"

func main(){
	var n int = 0

	fmt.Printf("Informe o valor de N: ")
	fmt.Scan(&n)

	array := make([]int, n)

	fmt.Printf("Agora informe %d valores:\n", n)
	for i:=0; i<n; i++{
		fmt.Scan(&array[i])
	}

	fmt.Println("Saída:\n")
	for i:=0; i<n; i++{
		fmt.Printf("%d	", array[i])
	}
}