package main

import "fmt"

func main(){

	var a, b, c, d float64

	fmt.Printf("Informe os valores (A B C D): ")
	fmt.Scanf("%f%f%f%f", &a, &b, &c, &d)

	det := a*d - b*c

	fmt.Printf("VALOR DO DETERMINANTE: %.2f\n", det)
}