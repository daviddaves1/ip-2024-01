package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Digite un numero inteiro: ")
	fmt.Scan(&n)

	total := graos(n)
	fmt.Printf("SAÍDA: %d\n", total)
}

func graos(n int) int {
	total := 0
	graos := n

	for i := 1; i <= 64; i++ {
		total += graos
		graos *= 2
	}

	return total
}
