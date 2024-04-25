package main

import "fmt"

func main() {
    var n, soma int
    fmt.Println("Digite um número inteiro maior que 1:")
    fmt.Scan(&n)

    soma = 1

    for i:=2; i<=n/2; i++ {
        if n%i == 0 {
            soma += i
        }
    }

    fmt.Printf("%d = 1", n)

    for i:=2; i<=n/2; i++ {
        if n%i == 0 {
            fmt.Printf(" + %d", i)
        }
    }

    if soma == n {
        fmt.Printf(" = %d (NUMERO PERFEITO)\n", soma)
    } else {
        fmt.Printf(" = %d (NUMERO NAO E PERFEITO)\n", soma)
    }
}
