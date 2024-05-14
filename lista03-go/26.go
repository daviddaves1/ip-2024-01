package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scanln(&T)

	for t := 0; t < T; t++ {
		numbers := make([]int, 9)
		for i := 0; i < 9; i++ {
			fmt.Scanln(&numbers[i])
		}

		resultado := encontrarCombinação(numbers)
		for _, num := range resultado {
			fmt.Println(num)
		}
	}
}

func encontrarCombinação(numbers []int) []int {
	sort.Ints(numbers)

	for i := 0; i < len(numbers)-6; i++ {
		for j := i + 1; j < len(numbers)-5; j++ {
			for k := j + 1; k < len(numbers)-4; k++ {
				for l := k + 1; l < len(numbers)-3; l++ {
					for m := l + 1; m < len(numbers)-2; m++ {
						for n := m + 1; n < len(numbers)-1; n++ {
							for o := n + 1; o < len(numbers); o++ {
								soma := numbers[i] + numbers[j] + numbers[k] + numbers[l] + numbers[m] + numbers[n] + numbers[o]
								if soma == 100 {
									return []int{numbers[i], numbers[j], numbers[k], numbers[l], numbers[m], numbers[n], numbers[o]}
								}
							}
						}
					}
				}
			}
		}
	}

	return nil
}
