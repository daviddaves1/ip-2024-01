package main

import "fmt"

func main() {
	var num_casos int

	fmt.Printf("Informe o numero de casos: ")
	fmt.Scan(&num_casos)

	for i:=0; i<num_casos; i++ {
		
		var total_pessoas int
		var popular, geral, arquib, cadeiras float64

		fmt.Printf("Informe os dados abaixo sobre o jogo %d:\n\n", i+1)

		fmt.Printf("Total de pessoas que compraram ingressos para o caso de teste: ")
		fmt.Scan(&total_pessoas)
		fmt.Printf("Percentual que compraram na categoria popular: ")
		fmt.Scan(&popular)
		fmt.Printf("Percentual que compraram na categoria geral: ")
		fmt.Scan(&geral)
		fmt.Printf("Percentual que compraram na categoria arquibancada: ")
		fmt.Scan(&arquib)
		fmt.Printf("Percentual que compraram na categoria cadeiras: ")
		fmt.Scan(&cadeiras)

		//obs: convertendo a variavel "total_pessoas" para float
		renda_total := float64(total_pessoas) * (popular / 100) * 1
		renda_total += float64(total_pessoas) * (geral / 100) * 5
		renda_total += float64(total_pessoas) * (arquib / 100) * 10
		renda_total += float64(total_pessoas) * (cadeiras / 100) * 20

		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i+1, renda_total)
	}
}
