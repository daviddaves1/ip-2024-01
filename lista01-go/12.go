package main

import "fmt"

func main(){

	var num_horas float64

	fmt.Printf("Informe o numero que horas que o lacatario usou a charrete: ")
	fmt.Scanf("%f", &num_horas)

	var valor_pagar float64

	if(num_horas >= 3){
		valor_pagar = (num_horas/3) * 10  //cada numero de horas é divido por 3h onde elas serão precisamente contabilizadas e depois prontas para serem multiplicadas pelo preço de R$ 10,00   
	} else{
		valor_pagar = num_horas * 5
	}

	fmt.Printf("Valor a pagar:	R$ %.2f\n", valor_pagar)
}