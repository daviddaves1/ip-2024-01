package main
import "fmt"

func main(){

	var x, y, i int

	fmt.Printf("Digite dois numeros: ")
	fmt.Scanf("%d%d", &x, &y)
	fmt.Printf("\n")

	//o código simplesmente não mostra nenhum resultado, não faço ideia do que possa ser

	if x%2==0 {
		for i=x; i<=10; i+=2 {
			fmt.Printf("%d   ", i)
		}
	} else{
		fmt.Printf("O primeiro numero não é par\n")
	}
}