package main

import "fmt"

func main() {
    var m int
    fmt.Println("Digite um número inteiro maior que zero:")
    fmt.Scan(&m)

    for n:=1; n<=m; n++ {
        cubo := n*n*n
        soma := 0
        impar := 1

        for soma < cubo {
            soma += impar
            fmt.Printf("%d", impar)
            if soma < cubo {
                fmt.Printf("+")
            }
            impar += 2
        }
		
        fmt.Printf(" = %d\n", cubo)
    }
}
