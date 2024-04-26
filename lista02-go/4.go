package main

import (
	"fmt"
)

func main() {
	var n, i, s float64
	var k int

	fmt.Printf("_______________________________________________________________________\n")
	fmt.Printf("Insira os valores de N, I, e S: ")
	fmt.Scan(&n, &i, &s)

	for n < 0 || n > 9 {
		fmt.Printf("N deve estar entre 0 e 9. Digite o valor de N novamente: ")
		fmt.Scan(&n)
	}

	fmt.Printf("Insira o valor de K: ")
	fmt.Scan(&k)
	fmt.Printf("_______________________________________________________________________\n\n\n\n")

	fmt.Printf("==================================================================================================\n")

	if n >= 0 && n <= 9 {
		fmt.Println("Tabuada de soma:")
		fmt.Printf("%.2f + %.2f = %.2f\n", n, i, n+i)
		for index:=1; index<k; index++ {
			i += s
			fmt.Printf("%.2f + %.2f = %.2f\n", n, i, n+i)
		}
		fmt.Printf("\n")

		i = 0
		fmt.Printf("Tabuada de subtração:\n")
		fmt.Printf("%.2f - %.2f = %.2f\n", n, i, n-i)
		for index:=1; index<k; index++ {
			i += s
			fmt.Printf("%.2f - %.2f = %.2f\n", n, i, n-i)
		}
		fmt.Printf("\n")

		i = 1
		fmt.Printf("Tabuada de multiplicação:\n")
		fmt.Printf("%.2f x %.2f = %.2f\n", n, i, n*i)
		for index:=1; index<k; index++ {
			i += s
			fmt.Printf("%.2f x %.2f = %.2f\n", n, i, n*i)
		}
		fmt.Printf("\n")

		i = 1
		fmt.Printf("Tabuada de divisão:\n")
		fmt.Printf("%.2f / %.2f = %.2f\n", n, i, n/i)
		for index:=1; index<k; index++ {
			i += s
			if i != 0 {
				fmt.Printf("%.2f / %.2f = %.2f\n", n, i, n/i)
			} else {
				fmt.Printf("Impossível dividir por zero\n")
			}
		}
	}
	fmt.Println("==================================================================================================")
}
