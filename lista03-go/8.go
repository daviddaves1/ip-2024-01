package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		fmt.Println("Entrade:\n")
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}

		fmt.Println("Saída:\n")
		fmt.Println(decimalParaBinario(n))
	}
}

func decimalParaBinario(n int) string {
	if n == 0 {
		return "0"
	}

	binario := make([]byte, 0)
	for n > 0 {
		resto := n % 2
		binario = append(binario, byte('0'+resto))
		n /= 2
	}

	for i := 0; i < len(binario)/2; i++ {
		binario[i], binario[len(binario)-1-i] = binario[len(binario)-1-i], binario[i]
	}

	return string(binario)
}
