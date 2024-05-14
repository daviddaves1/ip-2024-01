package main

import (
	"fmt"
	"strings"
)

func maxValor(n, d int, number string) int {
	stack := []byte{}
	for i := 0; i < len(number); i++ {
		for len(stack) > 0 && stack[len(stack)-1] < number[i] && len(stack)+len(number)-i > n-d {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, number[i])
	}
	result := strings.TrimLeft(string(stack), "0")
	if result == "" {
		result = "0"
	}
	var max int
	fmt.Sscanf(result, "%d", &max)
	return max
}

func main() {
	var n, d int
	for {
		fmt.Scan(&n, &d)
		if n == 0 && d == 0 {
			break
		}
		var number string
		fmt.Scan(&number)
		fmt.Println(maxValor(n, d, number))
	}
}
