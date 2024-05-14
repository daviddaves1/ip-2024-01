package main

import (
    "fmt"
    "sort"
)

func main() {
    var n int
    fmt.Scanln(&n)

    numeros := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanln(&numeros[i])
    }

    count := contarElementosUnicos(numeros)
    fmt.Println(count)
}

func contarElementosUnicos(numeros []int) int {
    sort.Ints(numeros)

    count := 0
    for i := 0; i < len(numeros); {
        j := i + 1
        for j < len(numeros) && numeros[j] == numeros[i] {
            j++
        }
        if j-i == 1 {
            count++
        }
        i = j
    }

    return count
}
