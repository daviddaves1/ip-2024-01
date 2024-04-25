package main

import "fmt"

func main() {
	var num float64

	fmt.Printf("Informe um número real: ")
	fmt.Scan(&num)

	frac := fmt.Sprintf("%.0f/%d", num*10, 10)

	fmt.Printf("SAÍDA: %s\n", num, frac)
}
