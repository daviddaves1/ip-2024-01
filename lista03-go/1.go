package main

import "fmt"

func main() {
	var N, M int
    
	fmt.Print("Insira um numero inteiro N: ")
	fmt.Scan(&N)
    
    numberN := make([]int, N)
	fmt.Printf("Insira %d numeros:\n", N)
    
	for i := 0; i < N; i++ {
		fmt.Scan(&numberN[i])
	}
    
	fmt.Print("Insira um numero inteiro M: ")
	fmt.Scan(&M)
    
    numberM := make([]int, M)
	fmt.Printf("Insira %d numeros:\n", M)
    
	for j := 0; j < M; j++ {
		fmt.Scan(&numberM[j])
	}
    
    found := make(map[int]bool)
    
    for _, num := range numberN {
        found[num] = true
    }
    
    for _, num := range numberM {
        if found[num] {
            fmt.Println("ACHEI")
        } else {
            fmt.Println("NÃO ACHEI")
        }
    }
}
