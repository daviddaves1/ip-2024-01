#include <stdio.h>

int main(void){

    float nota1=0, nota2=0, nota3=0;

    printf("Insira as 3 notas do aluno: ");
    scanf("%f%f%f", &nota1, &nota2, &nota3);

    float media = (nota1 + nota2 + nota3)/3;

    printf("Media do aluno = %.2f\n\n", media);

    if(media >= 6){
        printf("APROVADO\n");
    } else{
        printf("REPROVADO\n");
    }
}

