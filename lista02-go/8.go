package main

import "fmt"

func main() {
    var qtde_times int
    fmt.Printf("Informe a quantidade de times: ")
    fmt.Scan(&qtde_times)

    if qtde_times < 2 {
        fmt.Println("Campeonato invalido!")
        return
    }

    num_finais := qtde_times * (qtde_times - 1) /32

    k := 1
	fmt.Printf("Numero de finais: %d\n", num_finais)
    for i:=0; i<qtde_times; i++ {
        for j:=i+1; j<qtde_times; j++ {
            fmt.Printf("Final %d: Time%d X Time%d\n", k, i+1, j+1)
            k++
        }
    }
}