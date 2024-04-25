package main

import "fmt"

func primo(primo int) bool {
    if primo <= 1 {
        return false
    }
    for i:=2; i*i <= primo; i++ {
        if primo%i == 0 {
            return false
        }
    }
    return true   
}

func main() {
    var num int

    fmt.Printf("Informe o número: ")
    fmt.Scan(&num)

    if primo(num) {
        fmt.Printf("PRIMO\n", num)
    } else {
        fmt.Printf("NÃO PRIMO\n", num)
    }
}
