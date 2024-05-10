/* Dada uma sequência de N notas entre 0 e 10, escreva um programa que exiba o valor da última nota
informada e quantas vezes ela apareceu no conjunto. O programa deve exibir ainda a maior nota informada
e a posição (índice do vetor) da sua primeira ocorrência.
Entrada
Na primeira linha há um inteiro N, sendo 1  N  10000 representando a quantidade de notas da
sequência. Não é necessário validar o valor de N na entrada. Nas N linhas seguintes haverá um número
inteiro entre 0 e 10, inclusive, em cada linha.
Saída
O programa gera 2 linhas de saída. A primeira linha exibirá a frequência da última nota informada e
a segunda linha exibirá a maior nota e a posição (índice do vetor) da sua primeira ocorrência, seguindo o
formato da saída apresentado a seguir. Não se esqueça de quebrar uma linha após a última impressão.
Exemplo
Entrada
11
5
6
3
4
3
8
7
4
8
6
4
Saída
Nota 4, 3 vezes
Nota 8, indice 5
13 */

package main

import "fmt"

func main() {
	var N, contar, index int = 0, 0, 0
	var nota []float32

	fmt.Printf("Informe quantas notas quer colocar em sequência: ")
	fmt.Scan(&N)

	for N < 1 || N > 10000 {
		fmt.Printf("Valor incorreto. Tente novamente: ")
		fmt.Scan(&N)
	}

	nota = make([]float32, N)

	var last, maior float32 = 0, 0
	maior = nota[0]

	fmt.Printf("Informe as %d notas:\n", N)

	for i := 0; i < N; i++ {
		fmt.Scan(&nota[i])

		if nota[i] == nota[N-1] {
			last = nota[i]
			contar++
		}

		if nota[i] > maior {
			maior = nota[i]
			index = i
		}
	}

	fmt.Printf("Nota %.2f,		%d vezes\n", last, contar)
	fmt.Printf("Nota %.2f,		indice %d\n", maior, index)
}