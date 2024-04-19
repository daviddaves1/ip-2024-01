package main

import "fmt"

func main(){
    
    var f float64
    var poleg float64
    var c float64
    
    fmt.Printf("Informe a temperatura em F: ")
    fmt.Scanf("%f", &f)
    
    fmt.Printf("Informe a chuva em Polegadas: ")
    fmt.Scanf("%f", &poleg)
    
    c = (5 * (f - 32)) /9
    mili := poleg * 25.4
    
    fmt.Printf("%.2f F é igual a = %.2f celsius\n", f, c)
    fmt.Printf("%.2f polegadas de chuva é igual a %.2f milímetros de chuva\n", poleg, mili)
    
}
