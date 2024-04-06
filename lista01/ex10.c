#include <stdio.h>

int main(void){
    float a=0, b=0, c=0, d=0;

    printf("Informe os valores de A, B, C e D abaixo:\n\n");
    printf("A: ");
    scanf("%f", &a);
    printf("B: ");
    scanf("%f", &b);
    printf("C: ");
    scanf("%f", &c);
    printf("D: ");
    scanf("%f", &d);

    float determin = a*d - b*c;

    printf("O VALOR DO DETERMINANTE E = %.2f\n", determin);

}
