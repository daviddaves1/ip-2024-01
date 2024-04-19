package main

import "fmt"

func main(){
    
    var num1 int
    var num2 int
    var num3 int
    
    fmt.Printf("Informe 3 numeros inteiros: ")
    fmt.Scanf("%d%d%d", &num1, &num2, &num3)
    
    qd := (num1 * 100 + num2 * 10 + num3) * (num1 * 100 + num2 * 10 + num3)
    
    fmt.Printf("Contatenação: %d%d%d\n", num1, num2, num3)
    fmt.Printf("Quadrado = %d\n", qd)
    
   
}
