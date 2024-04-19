package main

import "fmt" //Novo valor=100+(100 × 0,50)=100+50=150

func main(){

	var salario, reajust float64

	fmt.Printf("Informe o salario atual do funcionario: ")
	fmt.Scanf("%f", &salario)

	if salario <= 300 {
		reajust = salario + (salario * 0.50)
	} else{
		reajust = salario + (salario * 0.30)
	}

	fmt.Print("Salario com reajuste:  R$ %.2f\n", reajust)
}