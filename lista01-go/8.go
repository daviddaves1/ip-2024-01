package main

import "fmt"

func main(){
    
    var raio float64
    var altura float64
    const pi = 3.14
    
    fmt.Printf("Informe o raio em m²: ")
    fmt.Scanf("%f", &raio)
    
    fmt.Printf("Informe a altura em m²: ")
    fmt.Scanf("%f", &altura)
    
    area_circulo := pi * (raio * raio)
    area_lateral := 2 * pi * raio * altura
    
    valor_total := 2 * area_circulo + area_lateral
    
    fmt.Printf("Valor do custo da lata = %.2f\n", valor_total*100)
}
