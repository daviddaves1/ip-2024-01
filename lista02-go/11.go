package main

import (
	"fmt"
	"math"
)

func main() {
	var number, e, r float64 = 0, 0, 1

	fmt.Println("Digite o número:")
	fmt.Scan(&number)
	fmt.Println("Digite o erro:")
	fmt.Scan(&e)
	
	for math.Abs(number-r*r) > e {
		erro := math.Abs(number-r*r)
		fmt.Printf("r: %.9f, erro: %.9f\n", r, erro)
		r = (r + number/r) /2
	}

	fmt.Printf("r: %.9f, erro: %.9f\n", r, math.Abs(number-r*r))
}
