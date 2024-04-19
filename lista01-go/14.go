package main

import "fmt"
import "math"

func main(){

	var altura, aresta, area float64

	fmt.Printf("Informe a altura e aresta da piramide em metros (0 0): ")
	fmt.Scanf("%f %f", &altura, &aresta)

		area = 3 * (aresta*aresta) * math.Sqrt(3) /2 
		volume := (1/3) * area * altura

		fmt.Printf("O VOLUME DA PIRAMIDE E - %.2f METROS CUBICOS\n", volume)
}