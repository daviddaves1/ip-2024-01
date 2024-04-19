package main
import "fmt"

func main(){

	var n1, n2, n3, media float64

	fmt.Print("Informe as 3 notas do aluno (0 0 0): ")
	fmt.Scan(&n1, &n2, &n3)

	media = (n1+n2+n3)/3

	fmt.Printf("Media = %.2f\n", media)

	if(media < 6){
		fmt.Print("REPROVADO(A)\n")
	} else{
		fmt.Print("APROVADO(A)\n")
	}
}