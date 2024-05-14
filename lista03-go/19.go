package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Entrada:\n")
	fmt.Scanln(&n)

	numeros := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&numeros[i])
	}

	elementosÚnicos := encontrar_elementos_unicos(numeros)
	for _, elemento := range elementosÚnicos {
		fmt.Printf("Saída:\n")
		fmt.Println(elemento)
	}
}

func encontrar_elementos_unicos(numeros []int) []int {
	elementos_unicos := make([]int, 0)
	elementoAtual := numeros[0]
	for i := 1; i < len(numeros); i++ {
		if numeros[i] != elementoAtual {
			elementos_unicos = append(elementos_unicos, elementoAtual)
			elementoAtual = numeros[i]
		}
	}

	elementos_unicos = append(elementos_unicos, elementoAtual)
	return elementos_unicos
}
