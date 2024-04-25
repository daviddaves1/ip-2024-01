package main

import "fmt"

func main() {
	var (
		matric   []int
		salarioH []float32
		num_horas []int
		contar   int
	)

	fmt.Print("Quantos funcionários deseja registrar? ")
	fmt.Scan(&contar)

	matric = make([]int, contar)
	salarioH = make([]float32, contar)
	num_horas = make([]int, contar)

	for i:=0; i<contar; i++ {
		fmt.Printf("Informe os seguintes dados do funcionário %d (matrícula, horas trabalhadas, salário por hora): ", i+1)
		fmt.Scan(&matric[i], &num_horas[i], &salarioH[i])

		salarioTotal := salarioH[i] * float32(num_horas[i])
		salarioH[i] = salarioTotal
	}
	fmt.Printf("\n")
	fmt.Println("Saída:\n")
	for i:=0; i<contar; i++ {
		fmt.Printf("Matrícula: %d	Salário: %.2f\n", matric[i], salarioH[i])
	}
}
