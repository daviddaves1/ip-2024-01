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

    x = x - 2*math.Pi*math.Floor(x/(2*math.Pi))

    cosX := 0.0
    sign := 1.0
    fatorial := 1.0
    power := 1.0

    for n := 0; n <= N; n++ {
        if n != 0 {
            fatorial *= float64(n)
            power *= x * x
        }
        cosX += sign * power / fatorial
        sign *= -1
    }

    fmt.Printf("cos(%.2f) = %.6f\n", x, cosX)
}
