package main

import "fmt"

func main() {
    var m, n int
    fmt.Println("Digite o número de linhas:")
    fmt.Scan(&m)
    fmt.Println("Digite o número de colunas:")
    fmt.Scan(&n)

    for l:= 2; l<=m; l++ {
        fmt.Printf("(%d,1)", l)
        
        for i:=2; i<=n; i++ {
            if i<=l {
                fmt.Printf("-(%d,%d)", l, i)
            }
        }
        fmt.Printf("\n")
    }
}
