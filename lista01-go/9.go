// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func main() {
    
    var a, b, c float64
    
  fmt.Printf("Valor de (A B C): ")
  fmt.Scanf("%f%f%f", &a, &b, &c)
  
  delta := b*b - 4 * (a * c)
  
  fmt.Printf("Delta: %.2f\n", delta)
}
