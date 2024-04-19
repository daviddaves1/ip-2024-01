package main

import "fmt"

func main(){
    
    var conta_cliente int
    var consumo_mc float64
    var consumidor string
    var preco float64
    
    fmt.Printf("Conta do Cliente: ")
    fmt.Scanf("%d", &conta_cliente)
    
    fmt.Printf("Consumo por M³: ")
    fmt.Scanf("%f", &consumo_mc)
    
    fmt.Printf("Tipo de Consumidor: ")
    fmt.Scanf("%s", &consumidor)
    
    if consumidor == "R" {
        preco = 5 + (consumo_mc * 0.05)  
    }
    if consumidor == "C" {
        if consumo_mc <= 80 {
            preco = 500
        } else{
            preco = 500 + ((consumo_mc - 80) * 0.25)
        }
    }
    if consumidor == "I" {
        if consumo_mc <= 100 {
            preco = 800
        } else{
            preco = 800 + ((consumo_mc - 100) * 0.04)
        }
    }
    
    fmt.Printf("Conta %d    Valor da Conta: %.2f\n", conta_cliente, preco)
}
