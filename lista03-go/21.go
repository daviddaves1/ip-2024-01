// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message
//21


package main
import (
	"fmt"
	"math"
)

func main() {
    var N int = 0
    
  fmt.Println("Insira o valor de N: ")
  fmt.Scan(&N)

  pot := math.Pow(10, 5)
  for N < 1 || N > int(pot) {
	fmt.Printf("Valor incorreto. Tente novamente: ")
	fmt.Scan(&N)
  }
  
  array := make([]int, N)
  
  fmt.Printf("Agora insira %d valores inteiros iguais ou maior que 0:\n", N)
  for i:=0; i<N; i++{
    fmt.Scan(&array[i])
    
    for array[i] < 0{
        fmt.Printf("Valor incorreto. Tente novamente: ")
        fmt.Scan(&array[i])
    }
}
    
    fmt.Printf("PARES:\n")
    orderBypar(array, N)
    for i:=0; i<N; i++{
        if array[i] % 2==0{
         fmt.Printf("%d  ", array[i])   
        }
    }
    
    fmt.Printf("\n\n")
    fmt.Printf("IMPARES:\n")
    orderByimpar(array, N)
    for i:=0; i<N; i++{
        if array[i] % 2!=0{
         fmt.Printf("%d  ", array[i])   
        }
    }
}

func orderBypar(array []int, N int){
    var aux int = 0
    for c:=0; c<N-1; c++{
        for i:=0; i<N-1-c; i++{
                if array[i] > array[i+1]{
                    aux = array[i]
                    array[i] = array[i+1]
                    array[i+1] = aux
            }           
        }
    }
}

func orderByimpar(array []int, N int){
    var aux int = 0
    for c:=0; c<N-1; c++{
        for i:=0; i<N-1-c; i++{
                if array[i] < array[i+1]{
                    aux = array[i]
                    array[i] = array[i+1]
                    array[i+1] = aux
            }
        }
    }
}
