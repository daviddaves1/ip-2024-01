package main

import "fmt"

func main() {
	var N int = 0

	fmt.Printf("Informe o valor de N: ")
	fmt.Scan(&N)

	for N < 1 || N > 1000 {
		fmt.Printf("Valor incorreto. Tente novamente: ")
		fmt.Scan(&N)
	}

	array := make([]int, N)
	fmt.Printf("Agora informe %d valores inteiros:\n", N)
	for i:=0; i<N; i++ {
		fmt.Scan(&array[i])
	}

	orderBy(array, N)

	fmt.Printf("Array ordenado:\n")
	for i:=0; i<N; i++ {
		fmt.Printf("%d\t", array[i])
	}
}

func orderBy(array []int, size int) {
	var aux int = 0

	for c:=0; c<size-1; c++ {
		for i:=0; i<size-1-c; i++ {
			if array[i] > array[i+1] {
				aux = array[i]
				array[i] = array[i+1]
				array[i+1] = aux
			}
		}
	}
}
/* Usando o algoritmo de ordenação bubble sorte através de uma função, onde o loop externo iterage sobre todos os elementos, menos sobre o último, já que o mesmo será o maior depois de toda a iteração, havendo uma condição com esse propósito no loop interno. Depois no próximo loop iteramos sobre os valores, mas não percorremos o último elemento que foi iterado no loop externo, pois ele já foi percorrido, e assim colocamos uma condição que verifica se o índice atual do array é maior do que o próximo índice, onde caso a condição seja verdadeira, fazemos a troca dos elementos para ficarem em ordem crescente. */
