/* Dados dois números inteiros x e n, calcula o valor de x^n */
package main
import "fmt"

func main() {
  var x, n, result int = 1, 0, 0;
  
  fmt.Scan(&x)
  fmt.Scan(&n)

    result = potencia(x, y)
    fmt.Printf("%d\n", result)
  
}

func potencia(x int, n int) int{
    if n == 0{
        return 1;
    } else{
        return x * potencia(x, n-1)
    }
}



//x^4 = x*x*x*x