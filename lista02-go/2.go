package main

import "fmt"

func main() {
	var populA, populB float64 = 0, 0
	var anos int32 = 0

	for populA >= populB {
		fmt.Printf("Informe a população do país A: ")
		fmt.Scan(&populA)
		fmt.Printf("Informe a população do país B: ")
		fmt.Scan(&populB)
		if populB <= populA {
			fmt.Printf("A população do país B deve ser maior que a população do país A. Insira os dados novamente\n")
			continue
		}
	}

	for populA <= populB {
		populB = populB * (1 + 0.015)
		populA = populA * (1 + 0.03)
		anos++
	}

	fmt.Printf("ANOS = %d\n", anos)
}




/* for i:=int(populA); i<int(populB); i++{

		i = i - int(populA)
		populA = populA * (1 + 0.03)
		anos++
	} */