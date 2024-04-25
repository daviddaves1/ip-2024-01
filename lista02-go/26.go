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

    sinX := 0.0
    sign := 1.0
    factorial := 1.0
    power := x
    for n := 0; n <= N; n++ {
        if n != 0 {
            factorial *= float64(2*n) * float64(2*n + 1)
            power *= x * x
        }
        sinX += sign * power /factorial
        sign *= -1
    }

    fmt.Printf("seno(%.2f) = %.6f\n", x, sinX)
}