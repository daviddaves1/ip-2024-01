package main

import (
	"fmt"
	"math"
)

type Ponto struct {
	x, y, z float64
}

func distancia(p1, p2 Ponto) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func main() {
	var N int
	fmt.Scanln(&N)

	pontos := make([]Ponto, N)
	for i := range pontos {
		fmt.Scanln(&pontos[i].x, &pontos[i].y, &pontos[i].z)
	}

	for i := 1; i < N; i++ {
		fmt.Printf("%.2f\n", distancia(pontos[i-1], pontos[i]))
	}
}
