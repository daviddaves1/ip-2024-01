package main

import (
    "fmt"
    "math"
)

func main() {
    var x float64
    var N int

    fmt.Println("Digite o valor de x:")
    fmt.Scan(&x)
    fmt.Println("Digite a quantidade de termos N:")
    fmt.Scan(&N)

    eX := 0.0
    for n:=0; n<=N; n++ {
        eX += math.Pow(x, float64(n)) / float64(fatorial(n))
    }

    fmt.Printf("e^%.2f = %.6f\n", x, eX)
}

func fatorial(n int) int {
    if n == 0 {
        return 1
    }
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}
