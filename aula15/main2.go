// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func main() {
    var elements int = 0
    
  fmt.Scan(&elements);
  
  var array = make([]float32, elements)
  var soma float32 = 0
  
  for i:=0; i<elements; i++{
      fmt.Scan(&array[i])
      soma += array[i]
  }
  fmt.Printf("Soma: %.2f\n", soma);
      
}