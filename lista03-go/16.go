package main

import (
    "fmt"
    "sort"
)

func main() {
    var n, k int
    fmt.Scanln(&n, &k)

    temposDeChegada := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanln(&temposDeChegada[i])
    }

    if aulaCancelada(temposDeChegada, k) {
        fmt.Println("SIM")
    } else {
        alunosPresentes := getAlunosPresentes(temposDeChegada)
        sort.Slice(alunosPresentes, func(i, j int) bool {
            return alunosPresentes[i] > alunosPresentes[j]
        })

        fmt.Println("NAO")
        for _, aluno := range alunosPresentes {
            fmt.Println(aluno)
        }
    }
}

func aulaCancelada(temposDeChegada []int, k int) bool {
    alunosPresentes := 0
    for _, tempoDeChegada := range temposDeChegada {
    	if tempoDeChegada <= 0 {
            alunosPresentes++
        }
    }

    return alunosPresentes < k
}

func getAlunosPresentes(temposDeChegada []int) []int {
    alunosPresentes := make([]int, 0)
    for _, tempoDeChegada := range temposDeChegada {
        if tempoDeChegada <= 0 {
            alunosPresentes = append(alunosPresentes, -tempoDeChegada)
        }
    }

    return alunosPresentes
}
