package main

import "fmt"

func main(){ 
	var N, contar, receive int = 0, 0, 0

	fmt.Printf("Informe o valor de N: ")
	fmt.Scan(&N)

	array := make([]int, N)

	fmt.Printf("Agora informe %d valores inteiros:\n", N)
	for i:=0; i<N; i++{
		fmt.Scan(&array[i])

		if array[i] == array[i-i]{
			receive = array[i]
			contar++
		}
	}
	fmt.Printf("Saída:\n")
	fmt.Println(receive)
	fmt.Println(contar)
}