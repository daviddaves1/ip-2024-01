package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		var a, b int = 0, 0
		fmt.Scanln(&a, &b)

		a = inverta_num(a)
		b = inverta_num(a)

		if a > b {
			fmt.Println(a)
		} else {
			fmt.Println(b)
		}
	}
}

func inverta_num(n int) int {
	inverte_num := 0
	for n > 0 {
		ult_dig := n % 10
		inverte_num = inverte_num*10 + ult_dig
		n /= 10
	}
	return inverte_num
}
