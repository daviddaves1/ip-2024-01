package main

import "fmt"

func main(){

	var h, m, s int

	fmt.Printf("Insira os valores em horas, minutos e segundos (00 00 00): ")
	fmt.Scanf("%d%d%d", &h, &m, &s)

	result := h*3600 + s + m*60

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", result)
}