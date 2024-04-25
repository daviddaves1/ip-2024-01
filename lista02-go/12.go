package main

import (
	"fmt"
)

func main() {
	var vi, vIni, vFim float64

	fmt.Printf("-------------------------------------------------------\n")
	fmt.Printf("Valor do ingresso em R$: ")
	fmt.Scan(&vi)
	fmt.Printf("Valor inicial do intervalo em R$: ")
	fmt.Scan(&vIni)
	fmt.Printf("Valor final do intervalo em R$: ")
	fmt.Scan(&vFim)

	for vIni >= vFim {
		fmt.Printf("Intervalo inválido. Informe o valor inicial do intervalo novamente: ")
		fmt.Scan(&vIni)
		continue
	}
	fmt.Printf("-------------------------------------------------------\n\n\n")

	calcular(vi, vIni, vFim)
}

func calcular(vi, vIni, vFim float64) {
	for v := vIni; v <= vFim; v += 1.0 {
		df := 200.0 + 0.05*float64(v-vIni)
		vendas := calcularVendas(v, vi)
		lucro := float64(vendas)*(v-vi) - df

		fmt.Printf("================================================================\n")
		fmt.Printf("V: %.2f, N: %d, L: %.2f\n", v, vendas, lucro)
	}

	vFim, lMax, nIngressos := calcularMelhor(vi, vIni, vFim)
	fmt.Printf("Melhor valor final: %.2f\n", vFim)
	fmt.Printf("Lucro: %.2f\n", lMax)
	fmt.Printf("Numero de ingressos: %d\n", nIngressos)
}

func calcularVendas(v, vi float64) int {
	if vi == 0 {
		return 145
	}
	return int(120 + (25*(v-vi))/0.5)
}

func calcularMelhor(vi, vIni, vFim float64) (float64, float64, int) {
	vFimMax, lMax, nIngressos := vIni, 0.0, 0

	for v := vIni; v <= vFim; v += 1.0 {
		df := 200.0 + 0.05*float64(v-vIni)
		vendas := calcularVendas(v, vi)
		lucro := float64(vendas)*(v-vi) - df

		if lucro > lMax {
			vFimMax, lMax, nIngressos = v, lucro, vendas
		}
	}

	return vFimMax, lMax, nIngressos
}
