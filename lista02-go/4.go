package main

import (
	"fmt"
)

func main() {
	var n, i, s float64 = 0, 0, 0
	var k int = 0

		fmt.Printf("\n\n")
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
		
		const value, zero = 1, 0
		temp := s
		
		fmt.Printf("==================================================================================================\n")
		if n >= 0 && n <= 9 {
			fmt.Printf("Tabuada de soma:\n")
			fmt.Printf("%.2f + %.2f = %.2f\n", n, i, n+i)
			for index:=value; index<k; index++ {
				if index >= 1 {
					fmt.Printf("%.2f + %.2f = %.2f\n", n, i+s, n+(i+s))
					s += temp
				}
			}
			fmt.Printf("\n")
			fmt.Printf("Tabuada de subtração:\n")
			fmt.Printf("%.2f - %.2f = %.2f\n", n, i, n-i)
			s = zero
			for index:=value; index<k; index++ {
				if index >= 1 {
					fmt.Printf("%.2f - %.2f = %.2f\n", n, i+s, n-(i+s))
					s += temp
				}
			}
			fmt.Printf("\n")
			fmt.Printf("Tabuada de multiplicação:\n")
			fmt.Printf("%.2f x %.2f = %.2f\n", n, i, n*i)
			s = zero
			for index:=value; index<k; index++ {
				if index >= 1 {
					fmt.Printf("%.2f x %.2f = %.2f\n", n, i+s, n*(i+s))
					s += temp
				}      
			}
			fmt.Printf("\n")
			fmt.Printf("Tabuada de divisão:\n")
			fmt.Printf("%.2f / %.2f = %.2f\n", n, i, n/i)
			s = zero
			for index:=value; index<k; index++ {
				if index >= 1 {
					fmt.Printf("%.2f / %.2f = %.2f\n", n, i+s, n/(i+s))
					s += temp
				}
			}
		}
		fmt.Printf("==================================================================================================\n")
}
