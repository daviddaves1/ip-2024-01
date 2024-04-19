package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Insira um numero inteiro maior que 1: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Numero invalido!\n")
		return
	}

	var sum float64
	for k := 1; k <= n; k++ {
		sum += 1.0 / float64(k)
	}

	fmt.Printf("Resultado = %.6f\n", sum)
}
