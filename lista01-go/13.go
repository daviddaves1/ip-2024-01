package main

import "fmt"

func main(){

	var nota float64
	
	fmt.Printf("Informe o a nota do aluno: ")
	fmt.Scanf("%f", &nota)

	if nota >= 0 && nota <= 6{
		fmt.Printf("NOTA - %.2f CONCEITO D\n", nota)
	}
	if nota > 6 && nota <= 7.5{
		fmt.Printf("NOTA - %.2f CONCEITO C\n", nota)
	}
	if nota > 7.5 && nota <= 9{
		fmt.Printf("NOTA - %.2f CONCEITO B\n", nota)
	}
	if nota > 9 && nota <= 10{
		fmt.Printf("NOTA %.2f CONCEITO A\n", nota)
	}

}