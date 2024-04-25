package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Println("Digite um número inteiro N:")
    fmt.Scan(&n)

    for h:=1; h<=n; h++ {
        for cat1 :=1; cat1<h; cat1++ {
            cat2 := int(math.Sqrt(float64(h*h - cat1*cat1)))
            if cat2*cat2+cat1*cat1 == h*h {
                fmt.Printf("hipotenusa = %d, catetos %d e %d\n", h, cat1, cat2)
            }
        }
    }
}
