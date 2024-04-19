package main

import "fmt"

func main(){
    
    var fahrenheit float64
    var celsius float64
    var num_fahr int
    
    for i:=0; i<num_fahr; i++{
        
        fmt.Printf("Informe quantos fahrenheit serão calculados: ")
        fmt.Scanf("%d", &num_fahr)
        
        fmt.Printf("Informe o valor do %d fahrenheit: ", i)
        fmt.Scanf("%f", &fahrenheit)
        
        celsius = (5 * (fahrenheit - 32)) /9
    }
    
    for i:=0; i<num_fahr; i++ {
        fmt.Printf("%.2f FAHRENHEIT equivale a %.2f CELSIUS\n", fahrenheit, celsius)
    }
    
}
