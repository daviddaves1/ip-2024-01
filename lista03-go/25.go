package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Informe os números sorteados:")
	sortedNumbers := readNumbers()
	sort.Ints(sortedNumbers)

	fmt.Println("Informe a quantidade de apostas:")
	n := readSingleInput()

	bets := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Informe as seis dezenas da aposta %d:\n", i+1)
		betNumbers := readNumbers()

		if len(betNumbers) != 6 {
			fmt.Println("Erro: Você deve inserir seis números para cada aposta.")
			return
		}

		if hasDuplicates(betNumbers) {
			fmt.Println("Erro: Não é permitido repetir números na mesma aposta.")
			return
		}

		bets[i] = betNumbers
	}

	countSena, countQuina, countQuadra := countWinners(sortedNumbers, bets)

	printResults(countSena, countQuina, countQuadra)
}

func readNumbers() []int {
	numbers := make([]int, 0)
	for i := 0; i < 6; i++ {
		var num int
		fmt.Scanln(&num)
		numbers = append(numbers, num)
	}
	return numbers
}

func readSingleInput() int {
	var input int
	fmt.Scanln(&input)
	return input
}

func hasDuplicates(numbers []int) bool {
	seen := make(map[int]bool)
	for _, num := range numbers {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func countWinners(sortedNumbers []int, bets [][]int) (int, int, int) {
	countSena := 0
	countQuina := 0
	countQuadra := 0

	for _, bet := range bets {
		countMatch := 0
		for _, num := range bet {
			if contains(sortedNumbers, num) {
				countMatch++
			}
		}

		switch countMatch {
		case 6:
			countSena++
		case 5:
			countQuina++
		case 4:
			countQuadra++
		}
	}

	return countSena, countQuina, countQuadra
}

func contains(numbers []int, target int) bool {
	for _, number := range numbers {
		if number == target {
			return true
		}
	}
	return false
}

func printResults(countSena, countQuina, countQuadra int) {
	if countSena > 0 {
		fmt.Printf("Houve %d acertadores para sena\n", countSena)
	} else {
		fmt.Println("Nao houve acertador para sena")
	}

	if countQuina > 0 {
		fmt.Printf("Houve %d acertador(es) da quina\n", countQuina)
	} else {
		fmt.Println("Nao houve acertador para quina")
	}

	if countQuadra > 0 {
		fmt.Printf("Houve %d acertador(es) da quadra\n", countQuadra)
	} else {
		fmt.Println("Nao houve acertador para quadra")
	}
}
