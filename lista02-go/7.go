package main

import (
    "fmt"
)

func main() {
    var qtde, qtde_par, qtde_imp int = 0, 0, 0
    var soma_par, soma_imp float32 = 0, 0
    
    fmt.Printf("Informe a quantidade de numeros, incluindo o 0 para quando quiser sair. Insira a quantidade   : ")
    fmt.Scan(&qtde)
    fmt.Printf("Agora informe %d numeros inteiros:\n", qtde)

    numeros := make([]int, qtde)

    for i := 0; i < qtde; i++ {
        fmt.Scan(&numeros[i])
        if numeros[i] == 0{
            fmt.Printf("Leitura terminada!\n")
            break
        }
        if numeros[i]%2==0{
            soma_par += float32(numeros[i])
            qtde_par++
        } else{
            soma_imp += float32(numeros[i])
            qtde_imp++
        }
    }
    if qtde_par > 0 {
        fmt.Printf("MEDIA PAR = %.2f\n", soma_par/float32(qtde_par))
    }
    if qtde_imp > 0 {
        fmt.Printf("MEDIA IMPAR = %.2f\n", soma_imp/float32(qtde_imp))
    }
}
