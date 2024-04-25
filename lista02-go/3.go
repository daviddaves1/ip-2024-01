package main

import "fmt"

func main() {
	var num, fat int = 0, 1

	fmt.Println("Insira o numero: ")
	fmt.Scan(&num)
	
	for i:=num; i>=1; i-- {
		fat *= i
	}

	fmt.Printf("%d! = %d\n", num, fat)
}