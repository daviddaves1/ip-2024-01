package main

import (
	"fmt"
	"sort"
)

func main() {
	var tamA, tamB int

	for {
		fmt.Println("Digite o tamanho do conjunto A:")
		fmt.Scanln(&tamA)
		if tamA >= 1 && tamA <= 100 {
			break
		}
		fmt.Println("Tamanho do conjunto A inválido. Digite novamente:")
	}

	for {
		fmt.Println("Digite o tamanho do conjunto B:")
		fmt.Scanln(&tamB)
		if tamB >= 1 && tamB <= 100 {
			break
		}
		fmt.Println("Tamanho do conjunto B inválido. Digite novamente:")
	}

	conjuntoA := make([]int, tamA)
	conjuntoB := make([]int, tamB)

	fmt.Println("Digite os elementos do conjunto A:")
	for i := 0; i < tamA; i++ {
		for {
			var elem int
			fmt.Scanln(&elem)
			if !contem(conjuntoA[:i], elem) {
				conjuntoA[i] = elem
				break
			}
			fmt.Println("Elemento já existente no conjunto A. Digite novamente:")
		}
	}

	fmt.Println("Digite os elementos do conjunto B:")
	for i := 0; i < tamB; i++ {
		for {
			var elem int
			fmt.Scanln(&elem)
			if !contem(conjuntoB[:i], elem) {
				conjuntoB[i] = elem
				break
			}
			fmt.Println("Elemento já existente no conjunto B. Digite novamente:")
		}
	}

	uniao := uniaoConjuntos(conjuntoA, conjuntoB)
	interseccao := interseccaoConjuntos(conjuntoA, conjuntoB)

	fmt.Println("União:", formatarConjunto(uniao))
	fmt.Println("Intersecção:", formatarConjunto(interseccao))
}

func contem(conjunto []int, elemento int) bool {
	for _, item := range conjunto {
		if item == elemento {
			return true
		}
	}
	return false
}

func uniaoConjuntos(conjuntoA, conjuntoB []int) []int {
	uniao := append(conjuntoA[:], conjuntoB[:]...)
	sort.Ints(uniao)
	return removerDuplicatas(uniao)
}

func interseccaoConjuntos(conjuntoA, conjuntoB []int) []int {
	interseccao := []int{}
	for _, itemA := range conjuntoA {
		for _, itemB := range conjuntoB {
			if itemA == itemB {
				interseccao = append(interseccao, itemA)
				break
			}
		}
	}
	sort.Ints(interseccao)
	return removerDuplicatas(interseccao)
}

func removerDuplicatas(slice []int) []int {
	deduped := []int{}
	for i, v := range slice {
		if i == 0 || v != deduped[len(deduped)-1] {
			deduped = append(deduped, v)
		}
	}
	return deduped
}

func formatarConjunto(slice []int) string {
	if len(slice) == 0 {
		return "()"
	}

	str := "(" + fmt.Sprint(slice[0])
	for i := 1; i < len(slice); i++ {
		str += "," + fmt.Sprint(slice[i])
	}
	str += ")"
	return str
}
