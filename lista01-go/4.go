package main

import "fmt"

func main(){
    
    var salario float64
    var kw float64
    
    fmt.Printf("Informe o valor do salario minimo: ")
    fmt.Scanf("%f", &salario)
    fmt.Printf("Informe quantos Kw a residencia gasta: ")
    fmt.Scanf("%f", &kw)
    
    valor_kw := (salario * 0.70) /100
    custo_consumo := valor_kw * kw
    custo_desconto := custo_consumo * (1 - 0.10) 
    
    fmt.Printf("Valor do Kw: %.2f   Custo do consumo: %.2f   Custo com o desconto: %.2f\n", valor_kw, custo_consumo, custo_desconto)
    
}
