package main

import (
	"fmt"
)

func main() {
	var size int

	for {
		fmt.Printf("Informe quantos valores deseja ordenar: ")
		fmt.Scan(&size)

		if size == 0 {
			fmt.Printf("Programa encerrado.\n")
			break
		}

		sequencia := make([]float64, size) //fazendo o array alocar os dados dinamicamente ao lê-los
		fmt.Printf("Agora innforme os valores que serão ordenados: ")
		for i:=0; i<size; i++ {
			fmt.Scan(&sequencia[i])
		}

		ordenar := true
		for i:=1; i<size; i++ {
			if sequencia[i-1] > sequencia[i] { //ordenando crescentemente a sequência numérica
				ordenar = false
				break
			}
		}

		if ordenar {
			fmt.Println("ORDENADA")
		} else {
			fmt.Println("DESORDENADA")
		}
	}
}
