package main

import (
	"fmt"
)

func main() {
	var codigo uint64
	var precoCompra, precoVenda float64
	var numVendas int

	var menor10, entre10e20, maior20, codMaiorLucro, codMaisVendida uint64
	var maiorLucro, totalCompra, totalVenda float64

	for {
	    fmt.Printf("Digite o código, preço de compra, preço de venda e numero de vendas: ")
		fmt.Scan(&codigo, &precoCompra, &precoVenda, &numVendas)
		
		lucro := (precoVenda - precoCompra) * float64(numVendas)
		lucroPercentual := (lucro / (precoCompra * float64(numVendas))) * 100

		// Atualização das estatísticas
		if lucroPercentual < 10 {
			menor10++
		} else if lucroPercentual <= 20 {
			entre10e20++
		} else {
			maior20++
		}

		if lucro > maiorLucro {
			maiorLucro = lucro
			codMaiorLucro = codigo
		}

		if uint64(numVendas) > codMaisVendida {
			codMaisVendida = codigo
		}

		totalCompra += precoCompra * float64(numVendas)
		totalVenda += precoVenda * float64(numVendas)
	}

	// Impressão dos resultados
	fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", menor10)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% e menor ou igual a 20%%: %d\n", entre10e20)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %d\n", maior20)
	fmt.Printf("Codigo da mercadoria que gerou maior lucro: %d\n", codMaiorLucro)
	fmt.Printf("Codigo da mercadoria mais vendida: %d\n", codMaisVendida)
	fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%\n", totalCompra, totalVenda, ((totalVenda-totalCompra)/totalCompra)*100)
}
