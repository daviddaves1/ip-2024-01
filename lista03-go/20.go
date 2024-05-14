package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanln(&n)

	pontos := make([][]float64, n)
	for i := 0; i < n; i++ {
		pontos[i] = make([]float64, 3)
		fmt.Scanln(&pontos[i][0], &pontos[i][1], &pontos[i][2])
	}

	maxMagnitude := calcularMaximaMagnitudeVetor(pontos)
	for _, mag := range maxMagnitude {
		fmt.Printf("%.2f\n", mag)
	}
}

func calcularMaximaMagnitudeVetor(pontos [][]float64) []float64 {
	maxMagnitude := make([]float64, len(pontos)-1)

	for i := 0; i < len(pontos)-1; i++ {
		max := 0.0
		for j := 0; j < 3; j++ {
			diff := pontos[i+1][j] - pontos[i][j]
			if math.Abs(diff) > max {
				max = math.Abs(diff)
			}
		}
		maxMagnitude[i] = max
	}

	return maxMagnitude
}
