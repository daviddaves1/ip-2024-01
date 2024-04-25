package main

import (
	"fmt"
)

func main() {
	var num int

	for {
		fmt.Println("Digite um número inteiro maior que 1:")
		fmt.Scan(&num)

		if num <= 1 {
			fmt.Printf("Fatoracao nao e possivel para o numero %d!\n", num)
		} else {
			fmt.Printf("%d = ", num)
			fatorar(num)
			fmt.Printf("\n")
		}
	}
}

func fatorar(num int) {
	for i := 2; i <= num; i++ {
		for num%i == 0 {
			fmt.Printf("%d", i)
			num /= i
			if num > 1 {
				fmt.Printf(" x ")
			}
		}
	}
}
