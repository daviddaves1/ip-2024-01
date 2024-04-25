package main

import (
    "fmt"
)

func main() {
    var n int32
    var num[n] int32

    fmt.Printf("Informe um numero inteiro N: ")
    fmt.Scan(&n)
    fmt.Printf("Agora informe %d numeros inteiros: ", n)
    for i:=0; i<n; i++{
        fmt.Scan(&num[i])
    }

    
}
