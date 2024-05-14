package main

import (
    "fmt"
    "strconv"
)

func main() {
    var testCases int
    fmt.Scanln(&testCases)

    for i := 0; i < testCases; i++ {
        var cpf string
        fmt.Scanln(&cpf)

        isValid := validarCPF(cpf)
        if isValid {
            fmt.Println("CPF valido")
        } else {
            fmt.Println("CPF invalido")
        }
    }
}

func validarCPF(cpf string) bool {
    cpf = limparCPF(cpf)

    if len(cpf) != 11 {
        return false
    }

    primeiroDigito := calcularDigito(cpf, 10)
    if cpf[9] != byte(primeiroDigito+'0') {
        return false
    }

    segundoDigito := calcularDigito(cpf, 11)
    if cpf[10] != byte(segundoDigito+'0') {
        return false
    }

    return true
}

func limparCPF(cpf string) string {
    cpfLimpo := ""
    for _, char := range cpf {
        if char >= '0' && char <= '9' {
            cpfLimpo += string(char)
        }
    }
    return cpfLimpo
}

func calcularDigito(cpf string, multiplicador int) int {
    soma := 0
    for i, char := range cpf {
        digit, _ := strconv.Atoi(string(char))
        soma += digit * (multiplicador - i)
    }

    resto := soma % 11
    if resto < 2 {
        return 0
    }
    return 11 - resto
}
