package main

import (
	"fmt"
	"math"
)

func main() {
	var t, n int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&n)
		numbers := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanln(&numbers[j])
		}

		minDist, comparisons := findMinDistance(numbers)
		fmt.Printf("%d %d\n", minDist, comparisons)
	}
}

func findMinDistance(numbers []int) (int, int) {
	minDist := math.MaxInt32
	comparisons := 0

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			currentDist := abs(numbers[i] - numbers[j])
			if currentDist < minDist {
				minDist = currentDist
			}
			comparisons++
		}
	}

	return minDist, comparisons
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
