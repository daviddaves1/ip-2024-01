package main

import (
	"fmt"
	"strings"
	"math"
)

func main() {
	var entrada string
	fmt.Scanln(&entrada)

	if !éEntradaVálida(entrada) {
		fmt.Println("FORMATO INVÁLIDO!")
		return
	}

	frases := dividirETratarFrases(entrada)
	frequências := calcularFrequencias(frases)
	distância := calcularDistancia(frequências)

	fmt.Println(frequências[0])
	fmt.Println(frequências[1])
	fmt.Printf("%.2f\n", distância)
}

func éEntradaVálida(entrada string) bool {
	return strings.Count(entrada, ";") == 1
}

func dividirETratarFrases(entrada string) []string {
	frases := strings.Split(entrada, ";")
	for i := range frases {
		frases[i] = strings.ToLower(strings.ReplaceAll(frases[i], " ", ""))
	}
	return frases
}

func calcularFrequencias(frases []string) [][]int {
	frequências := make([][]int, len(frases))
	for i, frase := range frases {
		frequências[i] = calcularFrequencia(frase)
	}
	return frequências
}

func calcularFrequencia(frase string) []int {
	frequência := make([]int, 5)
	for _, char := range frase {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			frequência[int(char-'a')]++
		}
	}
	return frequência
}

func calcularDistancia(frequências [][]int) float64 {
	distância := 0.0
	for i := 0; i < len(frequências[0]); i++ {
		diferença := frequências[0][i] - frequências[1][i]
		distância += math.Pow(float64(diferença), 2)
	}
	return math.Sqrt(distância)
}
