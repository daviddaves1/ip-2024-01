package main

import "fmt"

func main() {
    const qtde_provas, qtde_list = 8, 5
    var (
        matric int
        nota_provas [qtde_provas]float32
        nota_lists [qtde_list]float32
        notaEndWork, soma_notas, soma_lists float32
        frequencia int
    )

    fmt.Printf("\n\n\n")
    fmt.Printf("--------------------------------------------------------------------------\n")
    fmt.Printf("Informe os dados do aluno abaixo:\n\n")
    for {
        soma_notas, soma_lists = 0, 0

        fmt.Print("Matricula: ")
        fmt.Scan(&matric)
        if matric == -1 {
            fmt.Printf("Voce saiu com sucesso!\n")
            break
        }

        fmt.Printf("Notas das 8 provas (0 0 0 0 0 0 0 0): ")
        for i:=0; i<qtde_provas; i++ {
            fmt.Scan(&nota_provas[i])
            soma_notas += nota_provas[i]
        }
        soma_notas /= qtde_provas

        fmt.Printf("Notas das 5 listas de exercicios (0 0 0 0 0): ")
        for i:=0; i<qtde_list; i++ {
            fmt.Scan(&nota_lists[i])
            soma_lists += nota_lists[i] 
        }
        soma_lists /= qtde_list

        fmt.Print("Nota do trabalho final: ")
        fmt.Scan(&notaEndWork)
        fmt.Print("Frequencia: ")
        fmt.Scan(&frequencia)
        fmt.Printf("--------------------------------------------------------------------------\n\n\n")

        nota_final := (0.7 * soma_notas) + (0.15 * soma_lists) + (0.15 * notaEndWork)

        fmt.Printf("==========================================================================\n")
        fmt.Printf("Matricula: %d, Nota Final: %.2f, Situacao Final: ", matric, nota_final)

        if nota_final >= 6 && frequencia >= 96 {
            fmt.Println("APROVADO")
        } else if frequencia < 75 {
            fmt.Println("REPROVADO POR FREQUENCIA")
        } else if nota_final >= 6 {
            fmt.Println("REPROVADO POR NOTA")
        } else {
            fmt.Println("REPROVADO POR NOTA E POR FREQUENCIA")
        }
        fmt.Printf("==========================================================================\n\n")
    }
}
