package main

import (
    "fmt"
)

func inverterPalavras(frase string) string {
    palavras := []string{}
    palavra := ""
    for _, char := range frase {
        if char == ' ' {
            palavras = append(palavras, palavra)
            palavra = ""
        } else {
            palavra = string(char) + palavra
        }
    }
    palavras = append(palavras, palavra)
    return fmt.Sprint(palavras)
}

func main() {
    var linha string
    for {
        _, err := fmt.Scanln(&linha)
        if err != nil {
            break
        }
        fraseCorrigida := inverterPalavras(linha)
        fmt.Println(fraseCorrigida)
    }
}
