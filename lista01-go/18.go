package main

import (
	"fmt"
)

func main() {
	var a1, r, n int

	fmt.Print("Insira o valor inicial da progressão aritmética, a razão e o número de elementos: ")
	fmt.Scan(&a1, &r, &n)

	sum := 0
	for i:=0; i<n; i++ {
		sum += a1 + i*r
	}

	fmt.Print("valor da soma dos n primeiro elementos da progressão = %d\n", sum)
}
